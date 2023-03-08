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

class Counter(val length: Int) {
    private val counter = IntArray(length) { 0 }

    fun isDone(): Boolean {
        counter.forEach {
            if (it != 1)
                return false
        }
        return true
    }

    operator fun inc(): eu.toldi.bpmn_zkp.Counter {
        for (i in counter.indices) {
            counter[i]++

            if (counter[i] != 2) break else {
                if (i + 1 != counter.size) {
                    if (counter[i + 1] == 1) continue
                    counter[i + 1]++
                } else break
                for (j in i downTo 0) {
                    counter[j] = 0
                }
                break
            }
        }
        return this
    }

    override fun toString(): String {
        val stringBuilder = StringBuilder()
        counter.forEach {
            stringBuilder.append(it)
        }
        return stringBuilder.toString()
    }
}