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
import eu.toldi.bpmn_zkp.model.bpmn.StartEvent
import eu.toldi.bpmn_zkp.model.helper.TagNames
import org.w3c.dom.Node

class StartEventParser : ElementParser {
    override fun parse(model: Model, element: Node): StartEvent {
        val name = element.attributes.getNamedItem("id").textContent
        val properties = element.childNodes
        var tId = ""
        for (i in 0 until properties.length) {
            val node = properties.item(i)
            if (node.nodeName == TagNames.OUTGOING_TRANSITION || node.nodeName == TagNames.OUTGOING_TRANSITION2) {
                tId = node.textContent
                break
            }
        }
        val transition = model.getTransition(tId)
        return StartEvent(name, transition).also {
            transition.start = it
            model.addTransition(transition)
        }
    }
}