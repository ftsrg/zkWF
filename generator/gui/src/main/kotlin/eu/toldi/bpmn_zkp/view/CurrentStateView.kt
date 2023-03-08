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
import eu.toldi.bpmn_zkp.event.EventListener
import eu.toldi.bpmn_zkp.event.EventType
import eu.toldi.bpmn_zkp.model.Model
import eu.toldi.bpmn_zkp.model.testing.State

class CurrentStateView(m: Model) : StateView(m, "Current State", false), EventListener {


    init {
        EventBus.registerEventListener(this)
    }

    override fun onEvent(event: Event) {
        if (event.type == EventType.CURRENT_STATE_CHANGED) {
            val state = State.fromJson(event.data)
            stateVector.text = buildString {
                state.stateVector.forEach {
                    append("$it ")
                }
                append(state.randomness)
            }
            if(variables.isNotEmpty()) {
                variables.forEachIndexed { index, variable ->
                    variable.text = state.variables[index]
                }
            }
            if(messageList.isNotEmpty()) {
                messageList.forEachIndexed { index, message ->
                    message.text = buildString {
                        state.messages[index].forEach {
                            append("$it ")
                        }
                    }.trim()
                }
            }
        }
    }
}