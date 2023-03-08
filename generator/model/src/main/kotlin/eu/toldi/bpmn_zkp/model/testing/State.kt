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

package eu.toldi.bpmn_zkp.model.testing

import com.beust.klaxon.Klaxon

data class State(
    val stateVector: List<String>,
    val randomness: String,
    val variables: List<String>,
    val messages: List<List<String>>,
) {

    companion object {
        val klaxonx = Klaxon()
        fun fromJson(json: String): State {
            return klaxonx.parse<State>(json)!!
        }
    }

    fun toArgs(): List<String> {
        val result = mutableListOf<String>()
        stateVector.forEach {
            result.add(it.trim())
        }
        result.add(randomness.trim())
        variables.forEach {
            result.add(it.trim())
        }
        messages.forEach {
            result.addAll(it)
        }
        return result
    }

    fun toJson(): String {
        return klaxonx.toJsonString(this)
    }
}