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

package eu.toldi.bpmn_zkp.view

import eu.toldi.bpmn_zkp.event.Event
import eu.toldi.bpmn_zkp.event.EventBus
import eu.toldi.bpmn_zkp.event.EventType
import eu.toldi.bpmn_zkp.model.Model
import eu.toldi.bpmn_zkp.model.StateProof
import eu.toldi.bpmn_zkp.model.helper.MeasuredTimes
import eu.toldi.bpmn_zkp.model.testing.State
import eu.toldi.bpmn_zkp.model.testing.TestCase
import eu.toldi.bpmn_zkp.model.testing.TestCases
import eu.toldi.`zokrates-wrapper`.Zokrates
import javafx.scene.control.ComboBox
import javafx.scene.control.ProgressIndicator
import javafx.scene.control.TextArea
import javafx.scene.layout.HBox
import javafx.stage.FileChooser
import kotlinx.coroutines.*
import tornadofx.*
import java.io.File
import java.time.ZonedDateTime
import java.time.format.DateTimeFormatter
import java.util.*
import kotlin.system.measureTimeMillis

class ComputeView(private val m: Model, private val measurementTable: HBox) : View("Compute") {

    var testCases: TestCases? = null
    var compiled = false
    var setupDone = false
    var testCount = 0
    val measurements = mutableListOf<MeasuredTimes>()

    private lateinit var loading: HBox
    private lateinit var status: TextArea

    private lateinit var currentStateView: StateView
    lateinit var newStateView: StateView
    private lateinit var keyIndex: ComboBox<String>
    private val random = Random()
    private var contract: web3.Model? = null

    override val root = vbox {

        currentStateView = CurrentStateView(m)
        val hbox1 = hbox {

        }
        hbox1 += currentStateView
        newStateView = NewStateView(m)
        hbox {}.apply { this += newStateView }

        hbox {
            paddingAll = 2.0
            label("Participant")
            keyIndex = combobox<String> {
                items = m.participants.map { it.name }.toObservable()
            }

        }
        hbox {
            paddingAll = 4.0
            checkbox("Skip compilation and setup") {
                action {
                    setupDone = isSelected
                    compiled = isSelected

                }
            }
        }
        button("Generate Proof") {
            useMaxWidth = true
            isDisable = true
            EventBus.registerEventListener { event ->
                if (event.type == EventType.CONTRACT_DEPLOYED) {
                    isDisable = false
                }
            }
            action {
                try {
                    loading += ProgressIndicator()

                    GlobalScope.launch(Dispatchers.IO) {
                        /// Zokrates.compile("hash.zok", "hash")


                        val initialState = State(
                            currentStateView.stateVector.text.split(" ").dropLast(1),
                            currentStateView.stateVector.text.split(" ").last(),
                            currentStateView.variables.map { it.text.trim() },
                            currentStateView.messageList.map { it.text.split(" ").filter { it != "" } }
                        )
                        val s_curr = initialState

                        val newState = State(
                            newStateView.stateVector.text.split(" ").dropLast(1),
                            newStateView.stateVector.text.split(" ").last(),
                            newStateView.variables.map { it.text.trim() },
                            newStateView.messageList.map { it.text.split(" ").filter { it != "" } }
                        )
                        val keysMap = getKeys()
                        println(keysMap[keyIndex.selectedItem!!]!!)

                        val testCase = TestCase(initialState, newState, keysMap[keyIndex.selectedItem!!]!!)

                        val measured = runTestCase(testCase)
                        if (measured.witnessTime != -1.0) {
                            measurements.add(measured)
                            EventBus.emitEvent(Event(EventType.PROOF_GENERATED, "stateProof${testCase.ID}.json"))
                            val proof =
                                eu.toldi.`zokrates-wrapper`.model.ProofObject.fromJson(File("proof${testCase.ID}.json").readText())
                            val stateProof = StateProof(newState, proof!!)
                            File("stateProof${testCase.ID}.json").writeText(stateProof.toJson())
                            withContext(Dispatchers.Main) {
                                loading.clear()
                                measurementTable.clear()
                                measurementTable += MeasurementTableView(measurements)
                                information(
                                    "Success!",
                                    "ZoKrates setup done in s !"
                                )
                            }
                        } else {
                            withContext(Dispatchers.Main) {
                                loading.clear()
                            }
                        }
                    }
                } catch (e: java.lang.Exception) {
                    loading.clear()
                    error("Error occurred", "Error! Computing witness failed")
                    e.printStackTrace()
                }
            }
        }

        borderpane {
            top {
                vbox {
                    label("Load testcase")
                    var testCasesCount = hbox {
                        label("No testcases loaded")
                    }
                    button("select...").apply {
                        action {
                            val files = chooseFile(
                                "Select a Testcase file",
                                listOf<FileChooser.ExtensionFilter>(
                                    FileChooser.ExtensionFilter(
                                        "Json TestCases (*.json)",
                                        "*.json"
                                    )
                                ).toTypedArray(),
                                File(".").absoluteFile
                            )
                            if (files.isNotEmpty()) {
                                try {
                                    testCases = TestCases.fromJson(files[0].readText())
                                    GlobalScope.launch(Dispatchers.Main) {
                                        setStatusMessage("${testCases!!.size} testcases loaded")
                                    }

                                    testCasesCount.clear()
                                    //testCasesCount += label("${testCases!!.size} testcases loaded")
                                } catch (e: java.lang.Exception) {
                                    println(e.stackTraceToString())
                                    error("Error occurred", "Error parsing TestCase json")
                                }
                            }
                        }
                    }
                    button("Run testcases").apply {
                        action {
                            loading += ProgressIndicator()
                            if (testCases != null) {
                                try {
                                    val fixedThreadPool = newFixedThreadPoolContext(2, "Tester Threads")
                                    GlobalScope.launch {
                                        launch {
                                            println("Testcase #1")
                                            launch(Dispatchers.Default) {
                                                addMeasurement(runTestCase(testCases!!.first()))
                                                val proof =
                                                    eu.toldi.`zokrates-wrapper`.model.ProofObject.fromJson(File("proof${testCases!!.first().ID}.json").readText())
                                                val stateProof = StateProof(testCases!!.first().newState, proof!!)
                                                File("stateProof${testCases!!.first().ID}.json").writeText(stateProof.toJson())
                                            }.join()
                                            println("Running other tests")

                                            testCases!!.drop(1).forEach {
                                                launch(fixedThreadPool) {
                                                    addMeasurement(runTestCase(it))
                                                    val proof =
                                                        eu.toldi.`zokrates-wrapper`.model.ProofObject.fromJson(File("proof${it.ID}.json").readText())
                                                    val stateProof = StateProof(it.newState, proof!!)
                                                    File("stateProof${it.ID}.json").writeText(stateProof.toJson())
                                                }
                                            }
                                        }.join()
                                        setStatusMessage("Done")
                                        withContext(Dispatchers.Main) {
                                            loading.clear()
                                        }
                                    }
                                } catch (e: java.lang.Exception) {
                                    loading.clear()
                                }
                            }
                        }
                    }
                }
            }
            center {
                loading = hbox {
                    paddingAll = 5
                }
            }
            bottom {
                status = textarea("") {
                    isEditable = false
                }

            }
        }
    }

    var hashCompiled = false
    suspend fun runTestCase(t: TestCase): MeasuredTimes {
        testCount++
        setStatusMessage("Setting up Test#${t.ID}")
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
            setStatusMessage("Compiling Test#${t.ID}")
            val compileTime = if (!compiled) measureTimeMillis {
                Zokrates.compile("root.zok")
                compiled = true
            } else 0.0
            setStatusMessage("Computing witness Test#${t.ID}")
            val witnessTime = measureTimeMillis {
                val outHash = Zokrates.computeWithness(args, output = "test${t.ID}.result")
            }
            setStatusMessage("Running setup Test#${t.ID}")
            val setupTime = if (!setupDone) measureTimeMillis {
                Zokrates.setup()
                EventBus.emitEvent(Event(EventType.SETUP_DONE, ""))
                setupDone = true
            } else 0.0
            setStatusMessage("Proving Test#${t.ID}")
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
            return mesured
        } catch (e: Exception) {
            println("Test#${t.ID} failed! Reason: ${e.stackTraceToString()}")
            setStatusMessage("Test#${t.ID} failed! See logs for more details...")
        }
        return MeasuredTimes(t.ID, -1.0, -1.0, -1.0, -1.0)
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

    fun getKeys(hash0: String, hash1: String, keyIndex: Int): List<String> {
        println("python ../pycrypto/demo.py $hash0 $hash1")
        val pb = ProcessBuilder("python", "../pycrypto/demo.py", hash0, hash1, keyIndex.toString())


        val process = pb.start()
        process.waitFor()
        val output = process.inputStream.bufferedReader().use { it.readText() }
        return output.split(' ')
    }

    suspend fun setStatusMessage(msg: String) {
        withContext(Dispatchers.Main) {
            status.text += "[${ZonedDateTime.now().format(DateTimeFormatter.RFC_1123_DATE_TIME)}] $msg\n"
        }
    }

    suspend fun addMeasurement(measuredTimes: MeasuredTimes) {
        measurements.add(measuredTimes)
        withContext(Dispatchers.Main) {
            measurementTable.clear()
            measurementTable += MeasurementTableView(measurements)
        }
    }


    fun getKeyByIndex(index: Int): String {
        val pb = ProcessBuilder("python", "../pycrypto/getKey.py", index.toString())

        val process = pb.start()
        process.waitFor()
        val output = process.inputStream.bufferedReader().use { it.readText() }.split(" ").take(2)
        return "${output[0]} ${output[1]}"
    }

    fun getKeys(): Map<String, Int> {

        val result = mutableMapOf<String, Int>()
        val keys = Array<String>(m.participants.size) {
            getKeyByIndex(it)
        }.toList()
        m.participants.forEach {
            result[it.name] = keys.indexOf(it.publicKey!!.replace(",", ""))
        }
        return result
    }
}