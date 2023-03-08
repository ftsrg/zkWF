/*
 * Copyright 2023 Contributors of the zkWF project
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package eu.toldi.bpmn_zkp

import com.andreapivetta.kolor.green
import com.andreapivetta.kolor.red
import eu.toldi.bpmn_zkp.helper.SolidityHelper
import eu.toldi.bpmn_zkp.model.Model
import eu.toldi.bpmn_zkp.model.helper.MeasuredTimes
import eu.toldi.bpmn_zkp.model.testing.State
import eu.toldi.bpmn_zkp.model.testing.TestCase
import eu.toldi.bpmn_zkp.model.testing.TestCases
import eu.toldi.bpmn_zkp.web3.CredentialStore
import eu.toldi.`zokrates-wrapper`.Zokrates
import org.web3j.abi.FunctionEncoder
import org.web3j.abi.datatypes.Type
import org.web3j.abi.datatypes.Utf8String
import org.web3j.abi.datatypes.generated.Uint256
import org.web3j.protocol.Web3j
import org.web3j.protocol.core.methods.response.TransactionReceipt
import org.web3j.protocol.http.HttpService
import java.io.File
import java.math.BigInteger
import kotlin.system.exitProcess
import kotlin.system.measureTimeMillis

fun error(message: String): Nothing {
    println(message.red())
    exitProcess(1)
}

fun fixLen(message: Double): String {
    if (message >= 100)
        return "$message"
    if (message >= 10)
        return " $message"
    return "  $message"
}

fun deploySC(random: Int): String {
    Zokrates.exportVerifier()

    //solidityHelper.compileContract(File("./model.sol"))
    solidityHelper.compileContract(File("./model.sol"))
    val initialState = State(
        model.getInitialStateVector().split(" ").filter { it != "" },
        (random).toString(),
        Array(model.variables.size) { "0" }.asList(),
        Array(model.messages.size) { Array(8) { "0" }.asList() }.asList()
    )
    val hashValue = Zokrates.computeWithness(initialState.toArgs(), "hash", "hash_deploy.result")
    val hash = web3.Model.Hash(
        Uint256(BigInteger(hashValue[0])),
        Uint256(BigInteger(hashValue[1])),
        Uint256(BigInteger(hashValue[2])),
        Uint256(BigInteger(hashValue[3])),
        Uint256(BigInteger(hashValue[4])),
        Uint256(BigInteger(hashValue[5])),
        Uint256(BigInteger(hashValue[6])),
        Uint256(BigInteger(hashValue[7]))
    )

    val ciphertext = Utf8String(initialState.toJson())
    val dataParams = listOf(hash, ciphertext)

    val parameter = FunctionEncoder.encodeConstructor(
        dataParams
    )

    return solidityHelper.deployContract("build/Model.bin", parameter)
}

var testCount = 0
lateinit var model: Model
var web3j: Web3j? = null
val gasUsage = mutableListOf<BigInteger>()

lateinit var solidityHelper: SolidityHelper
fun main(args: Array<String>) {
    if (args.isEmpty()) {
        error("No arguments specified")
    }
    if (args.contains("--help") or args.contains("-h")) {
        println("Usage: bpmn_tester [OPTIONS] <bpmnFile> <testCases>")
        println("Available options:")
        println("--deploy\tDeploy smart contract (default: false)")
        println("--skip-setup\tSkips the setup phase (default: false)")
        println("--skip-tests\tSkips all test cases, only does the setup and compilation phase (default: false)")
        exitProcess(0)
    }
    if (args.size < 2) {
        error("Too few arguments! See --help for details.")
    }
    web3j = if (args.contains("--deploy")) {
        Web3j.build(HttpService("http://127.0.0.1:8545"))
    } else null

    if (web3j != null) {
        solidityHelper = SolidityHelper(web3j!!, CredentialStore.credentialsList[0])
        val index = args.indexOf("--contract-address")
        if (index != -1 && args.size >= index + 1 + 2 + 1) {
            deployedContract = args[index + 1]
        }
    }

    if (args.contains("--skip-setup")) {
        setupDone = true
        compiled = true
    }
    val skipTest = args.contains("--skip-tests")
    val from = if (args.contains("--from")) {
        val index = args.indexOf("--from")
        if (args.size >= index + 1 + 2 + 1) { // index + value + file1 + file2 + 1 for index correction
            args[index + 1].toInt()
        } else 0
    } else 0
    val bpmnFile = File(args[args.size - 2])
    val testCasesFile = File(args[args.size - 1])
    model = Model(bpmnFile)
    model.generateZokratesCode()
    println("Code generated! ✔️".green())

    Zokrates.initZokrates()

    val measured = mutableListOf<MeasuredTimes>()
    var testCases: TestCases = try {
        TestCases.fromJson(testCasesFile.readText())
    } catch (e: java.lang.Exception) {
        println(e.stackTraceToString())
        error("Error parsing TestCase json")
    }
    var failCount = 0

    if (!skipTest) {
        println("${testCases.size} testcases loaded ✔".green())
        if (from > 0) {
            testCases = TestCases(testCases.drop(from - 1))
            println("Test case set reduced to ${testCases.size}")
        }
        testCases.forEach {
            val time = runTestCase(it)
            if (time.witnessTime > 0 && time.proofTime > 0) {
                println("Test #${it.ID} finished in ${time.witnessTime + time.proofTime}s ✔".green())
                measured.add(time)
            } else {
                failCount++
                println("❌ Error in test case #${it.ID}".red())
            }
        }
    } else {
        println("Compiling ZoKrates code")
        val compileTime = measureTimeMillis {
            Zokrates.compile("root.zok")
            compiled = true
        }

        println("Running setup")
        val setupTime = measureTimeMillis {
            Zokrates.setup()
            setupDone = true
        }
        measured.add(
            MeasuredTimes(
                ID = 1,
                compileTime = compileTime.toDouble() / 1000,
                setUpTime = setupTime.toDouble() / 1000,
                witnessTime = 0.0,
                proofTime = 0.0
            )
        )
    }
    println()
    println("Results:")
    println()
    if (args.contains("--skip-setup").not()) {
        println("Compile Time:: ${measured.first { it.compileTime != 0.0 }.compileTime}".green())
        println("Setup Time:: ${measured.first { it.setUpTime != 0.0 }.setUpTime}".green())
    }
    var average = 0.0;
    if (measured.isNotEmpty() && !skipTest) {
        println("+----+--------------+------------+")
        println("| ID | Witness Time | Proof time |")
        println("+----+--------------+------------+")
        testCases.forEach { test ->
            val time = measured.firstOrNull { it.ID == test.ID }
            if (time != null) {
                println("| ${test.ID} |      ${fixLen(time.witnessTime)}s |   ${fixLen(time.proofTime)}s |".green())
                average += time.witnessTime + time.proofTime
            } else
                println("| ${test.ID} |       -       |      -      |".red())
            println("+----+--------------+------------+")
        }
        println("Proofs were generated in ${average / (testCases.size - failCount)} s on average")
        println()
    }

    if (web3j != null) {
        var gasAverage = BigInteger.valueOf(0L)
        gasUsage.forEach { gasAverage = gasAverage.add(it) }
        println("Average gas usage: ${gasAverage.divide(BigInteger.valueOf(gasUsage.size.toLong()))}")
        println()
    }
    if (failCount == 0) {
        println("All tests passed ✔".green())
    } else {
        println("$failCount/${testCases.size} tests failed! ❌".red())
        exitProcess(1)
    }
}

fun callSC(stateProof: StateProof): TransactionReceipt {
    val proofObj = stateProof.proof
    val newHash = web3.Model.Hash(
        Uint256(BigInteger(proofObj.inputs[11].drop(2), 16)),
        Uint256(BigInteger(proofObj.inputs[12].drop(2), 16)),
        Uint256(BigInteger(proofObj.inputs[13].drop(2), 16)),
        Uint256(BigInteger(proofObj.inputs[14].drop(2), 16)),
        Uint256(BigInteger(proofObj.inputs[15].drop(2), 16)),
        Uint256(BigInteger(proofObj.inputs[16].drop(2), 16)),
        Uint256(BigInteger(proofObj.inputs[17].drop(2), 16)),
        Uint256(BigInteger(proofObj.inputs[18].drop(2), 16)),
    )
    val signature = web3.Model.Signature(
        listOf(
            Uint256(BigInteger(proofObj.inputs[8].drop(2), 16)),
            Uint256(BigInteger(proofObj.inputs[9].drop(2), 16))
        ),
        Uint256(BigInteger(proofObj.inputs[10].drop(2), 16))
    )
    val ciphertext = Utf8String(stateProof.state.toJson())
    val proofParam = web3.Model.Proof(
        web3.Model.G1Point(
            Uint256(BigInteger(proofObj.proof.a[0].drop(2), 16)),
            Uint256(BigInteger(proofObj.proof.a[1].drop(2), 16))
        ),
        web3.Model.G2Point(
            proofObj.proof.b[0].map { Uint256(BigInteger(it.drop(2), 16)) },
            proofObj.proof.b[1].map { Uint256(BigInteger(it.drop(2), 16)) }),
        web3.Model.G1Point(
            Uint256(BigInteger(proofObj.proof.c[0].drop(2), 16)),
            Uint256(BigInteger(proofObj.proof.c[1].drop(2), 16))
        )
    )

    return solidityHelper.callContract(
        deployedContract,
        "stepModel", listOf(newHash, ciphertext, signature, proofParam) as List<Type<Any>>
        //"stepModel", listOf(newHash, signature, proofParam) as List<Type<Any>>
    )
}

lateinit var deployedContract: String
var hashCompiled = false
var compiled = false
var setupDone = false

fun runTestCase(t: TestCase): MeasuredTimes {
    testCount++
    println("Setting up Test#${t.ID}")
    try {
        if (!hashCompiled) {
            Zokrates.compile("hash.zok", "hash")
            hashCompiled = true
        }

        val s_curr = t.initialState.toArgs()
        val s_next = t.newState.toArgs()
        println(s_curr)
        val hash = Zokrates.computeWithness(s_curr, "hash", "hash${t.ID}.result")
        val hash_next = Zokrates.computeWithness(s_next, "hash", "hash${t.ID}.result")

        val keys = getKeys(hashToHexFormat(hash), hashToHexFormat(hash_next), t.keyIndex)
        println(keys)
        val args = mutableListOf<String>().apply {
            addAll(hash)
            addAll(s_curr)
            addAll(s_next)
            addAll(keys)
        }
        println(args)
        println("Compiling Test#${t.ID}")
        val compileTime = if (!compiled) measureTimeMillis {
            Zokrates.compile("root.zok")
            compiled = true
        } else 0.0
        println("Computing witness Test#${t.ID}")
        val witnessTime = measureTimeMillis {
            val outHash = Zokrates.computeWithness(args, output = "test${t.ID}.result")
        }
        println("Running setup Test#${t.ID}")
        val setupTime = if (!setupDone) measureTimeMillis {
            Zokrates.setup()
            setupDone = true
        } else 0.0
        if (setupTime != 0.0 && web3j != null) {
            deployedContract = deploySC(t.initialState.randomness.toInt())
        }
        println("Proving Test#${t.ID}")
        val proofTime = measureTimeMillis {
            Zokrates.generateProof(witness = "test${t.ID}.result", proof = "proof${t.ID}.json")
        }
        val mesured = MeasuredTimes(
            ID = t.ID,
            compileTime = compileTime.toDouble() / 1000,
            setUpTime = setupTime.toDouble() / 1000,
            witnessTime = witnessTime.toDouble() / 1000,
            proofTime = proofTime.toDouble() / 1000
        )

        val proof =
            eu.toldi.`zokrates-wrapper`.model.ProofObject.fromJson(File("proof${t.ID}.json").readText())
        val stateProof = StateProof(t.newState, proof!!)
        File("stateProof${t.ID}.json").writeText(stateProof.toJson())
        if (web3j != null) {
            if (t.requireRedeploy) {
                deployedContract = deploySC(t.initialState.randomness.toInt())
            }
            val rel = callSC(stateProof)
            gasUsage.add(rel.cumulativeGasUsed)
            if (!rel.isStatusOK) {
                println(rel.revertReason.red())
            }
        }
        return mesured
    } catch (e: Exception) {
        println("Test#${t.ID} failed! Reason: ${e.stackTraceToString()}".red())
        println("Test#${t.ID} failed! See logs for more details...".red())
    }
    return MeasuredTimes(t.ID, -1.0, -1.0, -1.0, -1.0)
}

fun getKeys(hash0: String, hash1: String, keyIndex: Int): List<String> {
    println("python ../pycrypto/demo.py $hash0 $hash1")
    val pb = ProcessBuilder("python", "../pycrypto/demo.py", hash0, hash1, keyIndex.toString())


    val process = pb.start()
    process.waitFor()
    val output = process.inputStream.bufferedReader().use { it.readText() }
    return output.split(' ')
}


fun longToUInt32ByteArray(value: Long): ByteArray {
    val bytes = ByteArray(4)
    bytes[3] = (value and 0xFFFF).toByte()
    bytes[2] = ((value ushr 8) and 0xFFFF).toByte()
    bytes[1] = ((value ushr 16) and 0xFFFF).toByte()
    bytes[0] = ((value ushr 24) and 0xFFFF).toByte()
    return bytes
}

fun hashToHexFormat(hash: List<String>): String {
    return buildString {
        hash.map { longToUInt32ByteArray(it.toLong()) }.forEach {
            for (b in it) {
                val st = String.format("%02X", b)
                append(st.lowercase())
            }
        }
    }
}
