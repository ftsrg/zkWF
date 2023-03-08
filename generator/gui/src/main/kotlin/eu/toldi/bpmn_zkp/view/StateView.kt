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
import eu.toldi.bpmn_zkp.model.Model
import javafx.scene.Parent
import javafx.scene.control.TextField
import tornadofx.*
import java.util.*
import eu.toldi.bpmn_zkp.event.EventBus
import eu.toldi.bpmn_zkp.event.EventType


abstract class StateView(val m: Model,val headingTitle:String,val editable: Boolean) : View() {

    var stateVector: TextField by singleAssign()
    val variables: MutableList<TextField> = mutableListOf()
    val messageList: MutableList<TextField> = mutableListOf()
    val random = Random()

    override val root = vbox {
        paddingAll = 2.0
        label(headingTitle) {
            //TODO: make this a heading
        }
        form {

            paddingAll = 2.0
            hbox {
                paddingAll = 2.0
                label("State Vector")
                stateVector = textfield(buildString {
                    append(m.getInitialStateVector())
                    append(kotlin.math.abs(random.nextInt() * 10000))
                }).apply {
                    isEditable = editable
                }
            }
            if (m.variables.isNotEmpty()) {
                vbox {
                    paddingAll = 2.0
                    label("Variables")
                    m.variables.forEach {
                        hbox {
                            paddingAll = 2.0
                            label(it.value.type.toString() + " " + it.value.name)
                            val vr = textfield("0").apply {
                                isEditable = editable
                            }
                            variables.add(vr)
                        }

                    }
                }
            }
            if (m.messages.isNotEmpty()) {
                vbox {
                    paddingAll = 2.0
                    label("Messages")
                    m.messages.forEach {
                        hbox {
                            paddingAll = 2.0
                            label(it.id)
                            val vr = textfield(buildString {
                                for (i in 1..8)
                                    append("0 ")
                            }).apply {
                                isEditable = editable
                            }
                            messageList.add(vr)
                        }

                    }
                }
            }
        }

    }
}