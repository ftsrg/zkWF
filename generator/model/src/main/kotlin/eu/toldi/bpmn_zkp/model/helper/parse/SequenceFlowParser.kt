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
import eu.toldi.bpmn_zkp.model.bpmn.Transition
import org.w3c.dom.Node

class SequenceFlowParser : ElementParser {
    override fun parse(model: Model, element: Node): Transition {
        val id = element.attributes.getNamedItem("id").textContent
        var transition = model.getTransition(id)
        val name = element.attributes.getNamedItem("name")
        if (name != null) {
            transition = Transition(id, name.textContent)
        }

        return transition.also {
            model.addTransition(it)
        }
    }
}