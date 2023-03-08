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
import eu.toldi.bpmn_zkp.model.bpmn.MessageThrowEvent
import eu.toldi.bpmn_zkp.model.helper.TagNames
import org.w3c.dom.Node

class MessageThrowIntermediateEventParser : ElementParser {

    override fun parse(model: Model, element: Node): MessageThrowEvent {
        val id = element.attributes.getNamedItem("id").textContent
        val properties = element.childNodes
        var toId = ""
        var tiId = ""
        var messageRef = ""
        for (i in 0 until properties.length) {
            val node = properties.item(i)
            when (node.nodeName) {
                TagNames.INCOMING_TRANSITION, TagNames.INCOMING_TRANSITION2 -> tiId = node.textContent
                TagNames.OUTGOING_TRANSITION, TagNames.OUTGOING_TRANSITION2 -> toId = node.textContent
                TagNames.MESSAGE_EVENT_DEF -> messageRef = node.attributes.getNamedItem("messageRef").textContent
            }
        }
        val transition_in = model.getTransition(tiId)
        val transition_out = model.getTransition(toId)
        val message = model.getMessage(messageRef)
        return MessageThrowEvent(id, transition_in, transition_out, message).also {
            transition_in.end = it
            transition_out.start = it
            model.addTransition(transition_in)
            model.addTransition(transition_out)
            model.addMessage(message)
        }
    }
}
