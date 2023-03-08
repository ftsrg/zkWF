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

package eu.toldi.`zokrates-wrapper`

import eu.toldi.`zokrates-wrapper`.exceptions.*
import java.io.File
import java.net.URL

object Zokrates {
    private var isInitDone = false
    private lateinit var zokratesPath: String

    private fun checkInit() {
        if (isInitDone.not()) {
            throw ZokratesNotInititialised()
        }
    }

    private fun installZokrates() {
        val tmp = File(System.getProperty("user.dir") + "/tmp").also { file ->
            file.mkdir()
        }
        val dlFile = File(tmp, "zokrates.sh").also {
            it.writeText(URL("https://raw.githubusercontent.com/ZoKrates/ZoKrates/master/scripts/one_liner.sh").readText())
            it.setExecutable(true)
        }

        val pb = ProcessBuilder(dlFile.absoluteFile.absolutePath.toString(), "--to", ".zokrates")
        val process = pb.start()

        val exitCode = process.waitFor()
        if (exitCode != 0) {
            throw ZokratesInstallFailed()
        }

        dlFile.delete()
        tmp.delete()
    }

    fun initZokrates() {
        if (File(System.getProperty("user.dir") + "/.zokrates/bin/zokrates").exists().not()) {
            installZokrates()
        }

        zokratesPath = System.getProperty("user.dir") + "/.zokrates/bin/zokrates"
        isInitDone = true
    }

    fun compile(path: String, outPut: String = "out") {
        checkInit()
        compile(File(path), outPut)
    }

    fun compile(file: File, outPut: String = "out") {
        checkInit()
        val pb = ProcessBuilder(zokratesPath, "compile", "-i", file.absoluteFile.absolutePath.toString(), "-o", outPut)

        val process = pb.start()

        val exitCode = process.waitFor()
        if (exitCode != 0) {
            val output = process.inputStream.bufferedReader().use { it.readText() }
            throw ZokratesCompileFailed(output)
        }
    }

    fun computeWithness(args: List<String>, input: String = "out", output: String = "witness"): List<String> {
        checkInit()
        val sb = StringBuilder()
        var prefix = ""
        args.forEach {
            sb.append(prefix)
            prefix = " "
            sb.append(it)
        }
        println(sb.toString())
        val command = mutableListOf(zokratesPath, "compute-witness", "-i", input, "-o", output, "-a").apply {
            addAll(args)
        }
        val processBuilder = ProcessBuilder(command)

        val process = processBuilder.start()

        val exitCode = process.waitFor()
        if (exitCode != 0) {
            val output1 = process.inputStream.bufferedReader().use { it.readText() }
            println("Error $output1")
            throw ZokratesComputeWitnessFailed(output1)
        }
        return readWitness(output)
    }

    fun readWitness(witness: String = "witness"): List<String> {
        return File(witness).readText().split("\n").filter { it.startsWith("~out_") }.sorted().map { it.split(' ')[1] }
    }

    fun setup(input: String = "out") {
        val pb = ProcessBuilder(zokratesPath, "setup", "-i", input)
        val process = pb.start()
        val exitCode = process.waitFor()
        if (exitCode != 0) {
            val output = process.inputStream.bufferedReader().use { it.readText() }
            println("Error $output")
            throw ZokratesSetupFailed(output)
        }
    }

    fun generateProof(input: String = "out", witness: String = "witness", proof: String = "proof.json") {
        val pb = ProcessBuilder(zokratesPath, "generate-proof", "-i", input, "-w", witness, "-j", proof)
        val process = pb.start()
        val exitCode = process.waitFor()
        if (exitCode != 0) {
            val output = process.inputStream.bufferedReader().use { it.readText() }
            println("Error $output")
            throw ZokratesSetupFailed(output)
        }
    }

    fun exportVerifier(input: String = "out", output: String = "verifier.sol") {
        val pb = ProcessBuilder(zokratesPath, "export-verifier")
        val process = pb.start()
        val exitCode = process.waitFor()
        if (exitCode != 0) {
            val output = process.inputStream.bufferedReader().use { it.readText() }
            println("Error $output")
            throw ZokratesSetupFailed(output)
        }
    }
}