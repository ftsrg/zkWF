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

package eu.toldi.bpmn_zkp.model.bpmn

import eu.toldi.bpmn_zkp.model.state.StateVectorElement

sealed class IntermediateMessageEvent(
    id: String,
    override val incomingTransition: Transition,
    override val outGoingTransition: Transition,
    val message: Message
) : Event(id), SingleInput, SingleOutput, StateVectorElement

class MessageCatchEvent(
    id: String,
    incomingTransition: Transition,
    outGoingTransition: Transition,
    message: Message
) : IntermediateMessageEvent(
    id = id,
    incomingTransition = incomingTransition,
    outGoingTransition = outGoingTransition,
    message = message
)

class MessageThrowEvent(
    id: String,
    incomingTransition: Transition,
    outGoingTransition: Transition,
    message: Message
) : IntermediateMessageEvent(
    id = id,
    incomingTransition = incomingTransition,
    outGoingTransition = outGoingTransition,
    message = message
)