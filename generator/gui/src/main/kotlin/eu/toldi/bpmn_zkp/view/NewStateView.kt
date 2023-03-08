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
import eu.toldi.bpmn_zkp.event.EventListener
import eu.toldi.bpmn_zkp.event.EventType
import eu.toldi.bpmn_zkp.model.Model

class NewStateView(m: Model) : StateView(m,"New State",true), EventListener{
    override fun onEvent(event: Event) {
        when(event.type){
            EventType.STEP_INPROGRESS -> makeEditable(false)
            EventType.STEP_COMPLETE -> makeEditable(true)
        }
    }

    fun makeEditable(editable: Boolean) {
        stateVector.isEditable = editable
        variables.forEach {
            it.isEditable = editable
        }
        messageList.forEach {
            it.isEditable = editable
        }
    }

}