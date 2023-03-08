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

package eu.toldi.bpmn_zkp.model.helper.parse

import eu.toldi.bpmn_zkp.model.Model
import eu.toldi.bpmn_zkp.model.bpmn.Event
import eu.toldi.bpmn_zkp.model.bpmn.Task
import eu.toldi.bpmn_zkp.model.helper.TagNames
import eu.toldi.bpmn_zkp.model.helper.VariableTypeNames
import eu.toldi.bpmn_zkp.model.state.Variable
import eu.toldi.bpmn_zkp.model.state.VariableType
import org.w3c.dom.Node


class TaskParser : ElementParser {
    override fun parse(model: Model, element: Node): Event {
        val name = element.attributes.getNamedItem("id").textContent
        //val key = element.attributes.getNamedItem("zkp:publicKey").textContent
        val variablesNode = element.attributes.getNamedItem("zkp:variables")
        val properties = element.childNodes
        var toId = ""
        var tiId = ""
        for (i in 0 until properties.length) {
            val node = properties.item(i)
            when (node.nodeName) {
                TagNames.INCOMING_TRANSITION, TagNames.INCOMING_TRANSITION2 -> tiId = node.textContent
                TagNames.OUTGOING_TRANSITION, TagNames.OUTGOING_TRANSITION2 -> toId = node.textContent
            }
        }
        val transition_in = model.getTransition(tiId)
        val transition_out = model.getTransition(toId)

        val nodeVars = mutableSetOf<Variable>()
        if (variablesNode != null) {
            val vars = variablesNode.textContent.split(',')
            vars.forEach {
                val nameAndType = it.trim().split(' ')
                assert(nameAndType.size > 2)
                val type: VariableType = when (nameAndType[0]) {
                    VariableTypeNames.FIELD -> VariableType.FIELD
                    VariableTypeNames.U32 -> VariableType.U32
                    VariableTypeNames.BOOL -> VariableType.BOOL
                    else -> throw java.lang.IllegalArgumentException()
                }
                nodeVars.add(Variable(nameAndType[1], type))
            }
        }


        return Task(name, transition_in, transition_out, nodeVars.toList()).also {
            transition_in.end = it
            transition_out.start = it
            model.addTransition(transition_in)
            model.addTransition(transition_out)
        }
    }
}