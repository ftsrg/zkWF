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

package eu.toldi.bpmn_zkp.model


import eu.toldi.bpmn_zkp.model.bpmn.*
import eu.toldi.bpmn_zkp.model.helper.TagNames
import eu.toldi.bpmn_zkp.model.helper.parse.*
import eu.toldi.bpmn_zkp.model.state.Expression
import eu.toldi.bpmn_zkp.model.state.StateVectorElement
import eu.toldi.bpmn_zkp.model.state.Variable
import eu.toldi.bpmn_zkp.model.state.VariableType
import eu.toldi.bpmn_zkp.utils.foreach
import eu.toldi.bpmn_zkp.utils.get
import org.w3c.dom.Node
import java.io.File
import javax.xml.parsers.DocumentBuilderFactory
import kotlin.math.absoluteValue


class Model(file: File) {
    val participants = mutableListOf<Participant>()
    val processes = mutableListOf<Process>()

    val messageFlows = mutableListOf<MessageFlow>()
    val messages = mutableListOf<Message>()
    val transitions = mutableSetOf<Transition>()
    val events = mutableListOf<Event>()
    val publicKeys = mutableListOf<String>()
    val variables = mutableMapOf<String, Variable>()
    val variableWritePermission = mutableMapOf<Int, Set<Variable>>()
    val laneKeys = mutableMapOf<String, String>()
    init {
        val factory = DocumentBuilderFactory.newDefaultInstance().apply {
            isIgnoringElementContentWhitespace = true
        }
        val builder = factory.newDocumentBuilder()
        val doc = builder.parse(file)



        if (doc.getElementsByTagName(TagNames.COLLABORATION).length == 0) {
            throw IllegalArgumentException("Invalid model! The model must be a collaboration!")
        } else {
            val collaboration = doc.getElementsByTagName(TagNames.COLLABORATION).item(0).childNodes
            for (i in 0 until collaboration.length) {
                parseCollaborationNode(collaboration.item(i))

            }

            val processes = doc.getElementsByTagName(TagNames.PROCESS)
            for (i in 0 until processes.length) {
                val process = processes.item(i)
                parseProcess(process)
            }
        }
    }

    fun parseCollaborationNode(node: Node) {
        when (node.nodeName) {
            TagNames.PARTICIPANT -> {
                participants.add(
                    Participant(
                        id = node.attributes.getNamedItem("id").textContent,
                        name = node.attributes.getNamedItem("name").textContent ?: "participant",
                        processRef = node.attributes.getNamedItem("processRef").textContent,
                        publicKey = node.attributes.getNamedItem(TagNames.PUBLIC_KEY)?.let { it.textContent }
                    )
                )
            }
            TagNames.MESSAGE_FLOW -> {
                messageFlows.add(
                    MessageFlow(
                        id = node.attributes.getNamedItem("id").textContent,
                        sourceRef = node.attributes.getNamedItem("sourceRef").textContent,
                        targetRef = node.attributes.getNamedItem("targetRef").textContent
                    )
                )
            }
        }
    }

    fun parseProcess(process: Node) {
        val id = process.attributes.getNamedItem("id").textContent
        val key = participants.filter { it.processRef == id }.first().publicKey
        var hasLanes = false
        val laneMap: MutableMap<String, MutableSet<String>> = mutableMapOf()

        for (i in 0 until process.childNodes.length) {
            val node_name = process.childNodes.item(i).nodeName
            if (node_name == TagNames.LANE_SET) {
                process.childNodes[i].childNodes.foreach { lane ->

                    hasLanes = true
                    val lane_id = lane.attributes.getNamedItem("id").textContent
                    laneKeys[lane_id] = lane.attributes.getNamedItem(TagNames.PUBLIC_KEY).textContent
                    laneMap[lane_id] = mutableSetOf()
                    lane.childNodes.foreach { node ->
                        laneMap[lane_id]!!.add(node.textContent)
                    }
                }
            }
        }

        process.childNodes.foreach { node ->
            val element = eu.toldi.bpmn_zkp.model.Model.Companion.mapping[node.nodeName]?.let {
                it.parse(
                    model = this,
                    element = node
                )
            }
            if (element != null) {
                addElement(element)
                val filtered = events.filter { it is StateVectorElement }
                if (hasLanes && element is StateVectorElement) {
                    val lane =
                        laneMap.toList().asSequence().filter { it.second.contains(element.id) }.map { it.first }.first()

                    publicKeys.add(filtered.indexOf(element), laneKeys[lane]!!)
                } else if (element is StateVectorElement) {
                    val index = filtered.indexOf(element)
                    publicKeys.add(index, key!!)
                }
            }

        }

    }


    companion object {
        val emptyTransition = listOf(0, -1)
        val mapping = mutableMapOf(
            TagNames.START_EVENT to StartEventParser(),
            TagNames.START_EVENT2 to StartEventParser(),
            TagNames.TASK to TaskParser(),
            TagNames.TASK2 to TaskParser(),
            TagNames.FINAL_EVENT to FinalEventParser(),
            TagNames.FINAL_EVENT2 to FinalEventParser(),
            TagNames.PARALLEL_GATEWAY to ParallelGatewayParser(),
            TagNames.PARALLEL_GATEWAY2 to ParallelGatewayParser(),
            TagNames.EXCLUSIVE_GATEWAY to ExclusiveGatewayParser(),
            TagNames.EXCLUSIVE_GATEWAY2 to ExclusiveGatewayParser(),
            TagNames.SEQUENCE_FLOW to SequenceFlowParser(),
            TagNames.SEQUENCE_FLOW2 to SequenceFlowParser(),
            TagNames.MESSAGE to MessageParser(),
            TagNames.INTERMEDIATE_CATCH_EVENT to MessageCatchIntermediateEventParser(),
            TagNames.INTERMEDIATE_THROW_EVENT to MessageThrowIntermediateEventParser(),
        )
    }

    fun getInitialStateVector(): String {
        val stateVector = IntArray(events.filter { it is StateVectorElement }.size) { 0 }
        val startEvents = events.filter { it is StartEvent }
        startEvents.map { it as StartEvent }.forEach { start ->
            stateVector[events.filter { it is StateVectorElement }.indexOf(start.outGoingTransition.end)] = 1
        }
        return buildString {
            stateVector.forEach {
                append("$it ")
            }
        }
    }

    fun getTransition(id: String): Transition {
        return if (transitions.filter { it.id == id }.size == 1) {
            transitions.first { it.id == id }
        } else {
            Transition(id)
        }
    }

    fun getMessage(id: String): Message {
        return if (messages.filter { it.id == id }.size == 1) {
            messages.first { it.id == id }
        } else {
            Message(id)
        }
    }

    fun addTransition(transition: Transition) {
        val tr = transitions.firstOrNull { it.id == transition.id }
        transitions.remove(tr)
        transitions.add(transition)
    }

    fun addMessage(message: Message) {
        val msg = messages.firstOrNull { it.id == message.id }
        messages.remove(msg)
        messages.add(message)
    }

    fun addElement(e: Element) {
        when (e) {
            is Task -> {
                events.add(e)
                e.variables.forEach {
                    if (!variables.containsKey(it.name)) {
                        variables[it.name] = it
                    }
                }
                variableWritePermission[events.filter { it is StateVectorElement }.indexOf(e)] = e.variables.toSet()
            }
            is Event -> events.add(e)
            is MessageFlow -> messageFlows.add(e)
            is Transition -> addTransition(e)
            is Message -> addMessage(e)
        }
    }


    fun generateArrayP(): List<List<Int>> {
        val result: MutableList<List<Int>> = mutableListOf()
        val tasks = events.filter { it is StateVectorElement }
        events.forEach { event ->
            when (event) {
                is Task -> {
                    println(event.id)
                    result.add(listOf(-1, tasks.indexOf(event.incomingTransition.start)))
                    result.add(listOf(1, tasks.indexOf(event)))
                    result.add(emptyTransition)
                }
                is StartEvent -> {
                    result.add(listOf(1, tasks.indexOf(event)))
                    result.add(emptyTransition)
                    result.add(emptyTransition)
                }

                is MessageThrowEvent -> {
                    result.add(listOf(-1, tasks.indexOf(event.incomingTransition.start)))
                    result.add(listOf(1, tasks.indexOf(event)))
                    result.add(emptyTransition)
                }
                is MessageCatchEvent -> {
                    result.add(listOf(-1, tasks.indexOf(event.incomingTransition.start)))
                    result.add(listOf(1, tasks.indexOf(event)))
                    result.add(emptyTransition)
                }

                is FinalEvent -> {
                    event.incomingTransitions.forEach { incomingTransition ->
                        result.add(listOf(-1, tasks.indexOf(incomingTransition.start)))
                        result.add(emptyTransition)
                        result.add(emptyTransition)
                    }
                }

                is ParallelGatewayStart -> {
                    result.add(listOf(-1, tasks.indexOf(event.incomingTransition.start)))
                    event.outGoingTransitions.forEach {
                        if(tasks.contains(it.end))
                            result.add(listOf(1, tasks.indexOf(it.end)))
                        else
                            result.add(emptyTransition)
                    }
                }

                is ParallelGatewayEnd -> {
                    event.incomingTransitions.forEach {
                        result.add(listOf(-1, tasks.indexOf(it.start)))
                    }
                    result.add(listOf(1, tasks.indexOf(event.outGoingTransition.end)))

                    result.add(listOf(-1, tasks.indexOf(event.incomingTransitions[0].start)))
                    result.add(listOf(1, tasks.indexOf(event.outGoingTransition.end)))
                    result.add(emptyTransition)

                    result.add(listOf(-1, tasks.indexOf(event.incomingTransitions[1].start)))
                    result.add(listOf(1, tasks.indexOf(event.outGoingTransition.end)))
                    result.add(emptyTransition)
                }

                is ExclusiveGatewayStart -> {
                    event.outGoingTransitions.forEach {
                        result.add(listOf(-1, tasks.indexOf(event.incomingTransition.start)))
                        result.add(listOf(1, tasks.indexOf(it.end)))
                        result.add(emptyTransition)
                    }
                }
                is ExclusiveGatewayEnd -> {
                    event.incomingTransitions.forEach {
                        result.add(listOf(-1, tasks.indexOf(it.start)))
                        result.add(listOf(1, tasks.indexOf(event.outGoingTransition.end)))
                        result.add(emptyTransition)
                    }
                }
            }
        }
        return result
    }

    fun printModel() {
        events.forEachIndexed { index, event ->
            print(index)
            print(' ')
            print(event.id + " -> ")
            when (event) {
                is SingleOutput -> {
                    print(event.outGoingTransition.end.id)
                }

                is MultiOutput -> {
                    print("[ ")
                    event.outGoingTransitions.forEach {
                        print(it.end.id)
                        print(' ')
                    }
                    print(']')
                }
            }
            println()
        }
        messages.forEachIndexed { index, message ->
            println("$index ${message.id} ${message.name}")
        }
    }


    fun generateZokratesCode() {
        val rootTMP = File("root.zok.template").readText()
        val stateChange = File("stateChange.zok.template").readText()
        val hashTMP = File("hash.zok.template").readText()

        val P = generateArrayP()
        printModel()
        println(P)
        var i = 0
        val builder = StringBuilder()
        @Suppress("UNCHECKED_CAST") val stateVectorElements: List<StateVectorElement> =
            events.filter { it is StateVectorElement } as List<StateVectorElement>

        builder.append("const u32 len_w = ${P.size / 3}")
        builder.appendLine()
        builder.append("const u32 len_V = ${stateVectorElements.size}")
        builder.appendLine()
        builder.append("const u32 len_E = ${transitions.size}")
        builder.appendLine()
        builder.append("const signed_field[len_w][3][2] p = [")
        P.forEach {
            if (i % 3 == 0) builder.append('[')
            builder.append('[')
            var prefix = ' '
            it.forEach {
                builder.append(prefix)
                builder.append("signed_field_create(")
                builder.append(it.absoluteValue)
                builder.append(',')
                builder.append(it >= 0)
                builder.append(')')
                prefix = ','
            }
            builder.append(']')
            if (i % 3 == 2) builder.append(']')
            builder.append(',')
            i++
        }
        builder.setLength(builder.length - 1)
        builder.append(']')
        builder.appendLine()

        val stateZok = File("stateChange.zok")
        val st = stateChange.replace("[[ !!!REPLACE THIS WITH CONSTANTS!!! ]]", builder.toString())
        stateZok.writeText(st)

        builder.clear()

        val len_V: Int = stateVectorElements.size
        //val len_V = 16
        builder.append("const u32 len_V = $len_V")
        builder.appendLine()
        builder.append("const field[${len_V}][2] keys = [")
        var prefix = ""
        publicKeys.forEach {
            builder.append(prefix)
            builder.append('[')
            prefix = ","
            builder.append(it)
            builder.append(']')
        }
        builder.append(']')

        val sb = StringBuilder()
        sb.append("hash = sha256h([")
        prefix = ""
        val hash_array_size = 16
        val s_n = IntArray(stateVectorElements.size + 1) { 0 }.toList().chunked(hash_array_size)

        val variableSB = StringBuilder()
        variables.forEach {
            val vari = it.value
            variableSB.appendLine("\t$vari")
        }

        val assertsSB = StringBuilder()
        if (variables.isNotEmpty()) {
            assertsSB.appendLine("\t// Assertions for variable write permissions")
            events.forEachIndexed { index, event ->
                if (event is Task) {
                    val canChange = variableWritePermission[index] ?: emptySet()
                    val cannotChage = variables.filter { !canChange.contains(it.value) }
                        .toList()
                        .map { it.second }
                    if(cannotChage.isNotEmpty()) {
                        assertsSB.append("\tassert( state != $index ")
                        var prefix1 = "||"
                        cannotChage.forEach {
                            assertsSB.append(prefix1)
                            assertsSB.append(" v_curr.${it.name} == v_next.${it.name} ")
                            prefix1 = "&&"
                        }
                        assertsSB.append(')')
                        assertsSB.appendLine()
                    }
                }
            }
        }
        events.filter { it is MessageThrowEvent }.forEach { event ->
            val throwEvent = event as MessageThrowEvent
            val throwIndex = stateVectorElements.indexOf(event)
            event as MessageThrowEvent

            assertsSB.appendLine("\tassert( state == $throwIndex && msg_curr.${event.message.id} != msg_next.${event.message.id} || msg_curr.${event.message.id} == msg_next.${event.message.id} )")
        }

        events.filter { it is MessageCatchEvent }.forEach { event ->
            val catchEvent = event as MessageCatchEvent
            val catchIndex = stateVectorElements.indexOf(event)
            val throwIntex = stateVectorElements.asSequence()
                .indexOfFirst { it is MessageThrowEvent && it.message.id == catchEvent.message.id }
            assertsSB.appendLine("\tassert( state != $catchIndex || s_next[state] != 2 || s_next[$throwIntex] == 2 )")
        }

        val tasks = events.filter { it is StateVectorElement }
        events.filter { it is ExclusiveGatewayStart }.forEach { event ->
            val exclusive = event as ExclusiveGatewayStart
            val startIndex = tasks.indexOf(exclusive.incomingTransition.start)
            exclusive.outGoingTransitions.filter { it.id != (exclusive.default?.id ?: true) }.forEach {
                val transition = transitions.first { t -> t.id == it.id }
                val endIndex = tasks.indexOf(it.end)
                assert(transition.name != null)
                val expression = Expression(transition.name!!)
                if (startIndex < endIndex) {
                    assertsSB.appendLine("\tassert( changes[0] != $startIndex || changes[1] != $endIndex || ${expression})")
                } else {
                    assertsSB.appendLine("\tassert( changes[1] != $startIndex || changes[0] != $endIndex || ${expression})")
                }
            }
            if (exclusive.default != null) {
                val defTransition = transitions.first { it.id == exclusive.default.id }
                val endIndex = events.indexOf(defTransition.end)
                if (startIndex < endIndex) {
                    assertsSB.append("\tassert( (changes[0] != $startIndex || changes[1] != $endIndex) ")
                } else {
                    assertsSB.append("\tassert( (changes[1] != $startIndex || changes[0] != $endIndex) ")
                }
                var prefix1 = "||"
                exclusive.outGoingTransitions.filter { it.id != defTransition.id }.forEach {
                    val transition = transitions.first { t -> t.id == it.id }
                    val expression = Expression(transition.name!!)
                    assertsSB.append(" $prefix1 !(${expression})")
                    prefix1 = "&&"
                }
                assertsSB.appendLine(')')
            }

        }
        val hashStringBuilder = StringBuilder()
        val bools = variables.toList().map { it.second }.filter { it.type == VariableType.BOOL }.chunked(256)
        val u32s = variables.toList().map { it.second }.filter { it.type == VariableType.U32 }.chunked(2)
        val fields = variables.toList().map { it.second }.filter { it.type == VariableType.FIELD }.chunked(16)

        hashStringBuilder.append("sha256h([")
        s_n.forEachIndexed { index, chunk ->
            if (index > 0) {
                hashStringBuilder.append(',')
            }
            hashStringBuilder.append('[')
            if (index == s_n.size - 1) {
                if (chunk.size < 16) {
                    hashStringBuilder.append("...s_n[${index * 16}..${index * 16 + (chunk.size) - 1}],random,...[0;${16 - chunk.size}]")
                } else {
                    hashStringBuilder.append("...s_n[${index * 16}..${(index + 1) * 16 - 1}],random")
                }
            } else {
                hashStringBuilder.append("...s_n[${index * 16}..${(index + 1) * 16}]")

            }
            hashStringBuilder.append(']')
        }
        if (bools.isNotEmpty()) {
            hashStringBuilder.append(',')
        }
        bools.forEachIndexed { index, chunk ->
            if (index % 2 == 0) {
                if (index > 0) {
                    hashStringBuilder.append(',')
                }
                hashStringBuilder.append('[')
            }
            hashStringBuilder.append("...bool_to_u32([")
            var prefix3 = ""
            chunk.forEach {
                hashStringBuilder.append(prefix3)
                prefix3 = ","
                hashStringBuilder.append("v.${it.name}")
            }
            if (chunk.size < 256) {
                hashStringBuilder.append(",...[false;${256 - chunk.size}]")
            }
            hashStringBuilder.append("])")
            if (index % 2 == 1) {
                hashStringBuilder.append(']')
            }
        }
        //Bool group padding
        if (bools.size % 2 == 1) {
            hashStringBuilder.append(",...[0;8]]")
        }

        if (u32s.isNotEmpty()) {
            hashStringBuilder.append(',')
        }

        u32s.forEachIndexed { index, chunk ->
            if (index > 0) {
                hashStringBuilder.append(',')
            }
            hashStringBuilder.append('[')
            var prefix3 = ""
            chunk.forEach {
                hashStringBuilder.append(prefix3)
                prefix3 = ","
                hashStringBuilder.append("v.${it.name}")
            }
            if (chunk.size < 16) {
                hashStringBuilder.append(",...[0;${16 - chunk.size}]")
            }
            hashStringBuilder.append(']')
        }

        if (fields.isNotEmpty()) {
            hashStringBuilder.append(',')
        }

        fields.forEachIndexed { index, chunk ->
            if (index % 2 == 0) {
                if (index > 0) {
                    hashStringBuilder.append(',')
                }
                hashStringBuilder.append('[')
            }

            var prefix3 = ""
            chunk.forEach {
                hashStringBuilder.append(prefix3)
                prefix3 = ","
                hashStringBuilder.append("...field_to_u32(v.${it.name})")
            }
            if (chunk.size < 2) {
                hashStringBuilder.append(",...[0;8]]")
            } else
                hashStringBuilder.append("]")
        }
        hashStringBuilder.append("])\n")

        val messageHashStringBuilder = buildString {
            messages.forEach {
                appendLine("\tu32[8] ${it.id}")
            }
        }


        println("$sb")
        val rootZok = File("root.zok")
        val root = rootTMP
            .replace("[[ !!!REPLACE THIS WITH CONSTANTS!!! ]]", builder.toString())
            .replace("[[ !!! REPLACE THIS WITH HASH FUNCTION !!!] ]", hashStringBuilder.toString())
            .replace("[[ !!!REPLACE THIS WITH VARIABLES!!! ]]", variableSB.toString())
            .replace("[[ !!! REPLACE THIS WITH HASH VARIABLE ASSERTION !!!]]", assertsSB.toString())
            .replace("[[ !!!REPLACE THIS WITH MESSAGES!!! ]]", messageHashStringBuilder)

        rootZok.writeText(root)

        val hashZok = File("hash.zok")
        val hash = hashTMP
            .replace("[[ !!!REPLACE THIS WITH CONSTANTS!!! ]]", builder.toString())
            .replace("[[ !!! REPLACE THIS WITH HASH FUNCTION !!!] ]", hashStringBuilder.toString())
            .replace("[[ !!!REPLACE THIS WITH VARIABLES!!! ]]", variableSB.toString())
            .replace("[[ !!!REPLACE THIS WITH MESSAGES!!! ]]", messageHashStringBuilder)


        hashZok.writeText(hash)
    }

}
