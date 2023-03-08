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

package eu.toldi.bpmn_zkp.model.helper

object TagNames {
    const val START_EVENT2 = "bpmn2:startEvent"
    const val FINAL_EVENT2 = "bpmn2:endEvent"
    const val TASK2 = "bpmn2:task"
    const val PARALLEL_GATEWAY2 = "bpmn2:parallelGateway"
    const val EXCLUSIVE_GATEWAY2 = "bpmn2:exclusiveGateway"
    const val INCOMING_TRANSITION2 = "bpmn2:incoming"
    const val OUTGOING_TRANSITION2 = "bpmn2:outgoing"

    const val START_EVENT = "bpmn:startEvent"
    const val FINAL_EVENT = "bpmn:endEvent"
    const val TASK = "bpmn:task"
    const val PARALLEL_GATEWAY = "bpmn:parallelGateway"
    const val EXCLUSIVE_GATEWAY = "bpmn:exclusiveGateway"
    const val INCOMING_TRANSITION = "bpmn:incoming"
    const val OUTGOING_TRANSITION = "bpmn:outgoing"

    const val SEQUENCE_FLOW = "bpmn:sequenceFlow"
    const val SEQUENCE_FLOW2 = "bpmn2:sequenceFlow"
    const val MESSAGE_FLOW = "bpmn2:messageFlow"

    const val COLLABORATION = "bpmn2:collaboration"
    const val PARTICIPANT = "bpmn2:participant"
    const val PROCESS = "bpmn2:process"
    const val LANE = "bpmn2:lane"
    const val LANE_SET = "bpmn2:laneSet"
    const val THROW_EVENT = "bpmn2:intermediateThrowEvent"
    const val CATCH_EVENT = "bpmn2:intermediateCatchEvent"
    const val MESSAGE = "bpmn2:message"
    const val MESSAGE_EVENT_DEF = "bpmn2:messageEventDefinition"

    const val INTERMEDIATE_CATCH_EVENT = "bpmn2:intermediateCatchEvent"
    const val INTERMEDIATE_THROW_EVENT = "bpmn2:intermediateThrowEvent"

    const val PUBLIC_KEY = "zkp:publicKey"
}