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

import com.beust.klaxon.Klaxon
import eu.toldi.bpmn_zkp.model.testing.State
import eu.toldi.`zokrates-wrapper`.model.ProofObject

data class StateProof(val state: State, val proof: ProofObject) {
    companion object {
        val klaxonx = Klaxon()
        fun fromJson(json: String): StateProof {
            return klaxonx.parse<StateProof>(json)!!
        }
    }

    fun toJson(): String {
        return klaxonx.toJsonString(this)
    }
}
