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

import eu.toldi.`zokrates-wrapper`.Zokrates


fun getKeys(hash0: String, hash1: String): List<String> {
    println("python ../pycrypto/demo.py $hash0 $hash1")
    val pb = ProcessBuilder("python", "../pycrypto/demo.py", hash0, hash1)


    val process = pb.start()
    process.waitFor()
    val output = process.inputStream.bufferedReader().use { it.readText() }
    return output.split(' ')
}

val s_curr = mutableListOf(
    "1",
    "0",
    "0",
    "0",
    "0",
    "1",
    "0",
    "0",
    "0",
    "0",
    "15",
    "0",
    "0",
    "0",
    "0",
    "0",
    "0",
    "0",
    "0",
    "0",
    "0",
    "0",
    "0",
    "0",
    "0",
    "0",
    "0",
    "0",
    "0",
    "0",
    "0",
    "0",
    "0",
    "0",
    "0",
    "0",
    "0"
)
val s_next = mutableListOf(
    "2",
    "1",
    "0",
    "0",
    "0",
    "1",
    "0",
    "0",
    "0",
    "0",
    "99",
    "1",
    "0",
    "0",
    "0",
    "0",
    "0",
    "0",
    "0",
    "0",
    "0",
    "0",
    "0",
    "0",
    "0",
    "0",
    "0",
    "0",
    "0",
    "0",
    "0",
    "0",
    "0",
    "0",
    "0",
    "0",
    "0"
)

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

fun main() {
    Zokrates.initZokrates()
    Zokrates.compile("hash.zok")


    val hash = Zokrates.computeWithness(s_curr)
    val hash_next = Zokrates.computeWithness(s_next)
    println(hash_next)
    println(
        buildString {
            append('{')
            var prefix = ""
            hash.forEachIndexed { index, s ->
                append(prefix)
                prefix = ","
                append(" '${'a' + index}': $s")
            }
            append('}')
        }
    )
    println(
        buildString {
            append('{')
            var prefix = ""
            hash_next.forEachIndexed { index, s ->
                append(prefix)
                prefix = ","
                append(" '${'a' + index}': $s")
            }
            append('}')
        }
    )
    val keys = getKeys(hashToHexFormat(hash), hashToHexFormat(hash_next))
    println(keys)
    val args = mutableListOf<String>().apply {
        addAll(hash)
        addAll(s_curr)
        addAll(s_next)
        addAll(keys)
    }
    println(args)

    Zokrates.compile("root.zok")
    val outHash = Zokrates.computeWithness(args)
    println(outHash)
}