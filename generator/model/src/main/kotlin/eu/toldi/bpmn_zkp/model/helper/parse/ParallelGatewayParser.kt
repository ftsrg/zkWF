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
import eu.toldi.bpmn_zkp.model.bpmn.ParallelGateway
import eu.toldi.bpmn_zkp.model.bpmn.ParallelGatewayEnd
import eu.toldi.bpmn_zkp.model.bpmn.ParallelGatewayStart
import eu.toldi.bpmn_zkp.model.bpmn.Transition
import eu.toldi.bpmn_zkp.model.helper.TagNames
import org.w3c.dom.Node

class ParallelGatewayParser : ElementParser {
    override fun parse(model: Model, element: Node): ParallelGateway {
        val name = element.attributes.getNamedItem("id").textContent
        val properties = element.childNodes
        val toId: MutableList<String> = mutableListOf()
        val tiId: MutableList<String> = mutableListOf()
        for (j in 0 until properties.length) {
            val node = properties.item(j)
            when (node.nodeName) {
                TagNames.INCOMING_TRANSITION, TagNames.INCOMING_TRANSITION2 -> tiId.add(node.textContent)
                TagNames.OUTGOING_TRANSITION, TagNames.OUTGOING_TRANSITION2 -> toId.add(node.textContent)
            }
        }
        assert(toId.size > 1 && tiId.size == 1 || toId.size == 1 && tiId.size > 1)
        val transitions_in = Array<Transition>(tiId.size) {
            model.getTransition(tiId[it])
        }.toList()
        val transitions_out = Array<Transition>(toId.size) {
            model.getTransition(toId[it])
        }.toList()

        if (toId.size > 1 && tiId.size == 1) {
            return ParallelGatewayStart(name, transitions_in[0], transitions_out).apply {
                incomingTransition.end = this
                outGoingTransitions.forEach {
                    it.start = this
                    model.addTransition(it)
                }
                model.addTransition(incomingTransition)
            }
        } else {
            return ParallelGatewayEnd(name, transitions_in, transitions_out[0]).apply {
                outGoingTransition.start = this
                incomingTransitions.forEach {
                    it.end = this
                    model.addTransition(it)
                }
                model.addTransition(outGoingTransition)
            }
        }
    }
}