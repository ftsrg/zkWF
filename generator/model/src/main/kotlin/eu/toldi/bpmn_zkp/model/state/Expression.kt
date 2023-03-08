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

package eu.toldi.bpmn_zkp.model.state

import kotlin.text.Regex.Companion.escape

// Boolean expression transitions
class Expression(name: String) {

    companion object {
        val acceptedExpressions = arrayOf("==", ">=", "<=", ">", "<", "!=", "true", "false", "!")
        val numbers = charArrayOf('0', '1', '2', '3', '4', '5', '6', '7', '8', '9')
    }

    private val expression = mutableListOf<String>()

    init {
        val parts = name.split(Regex("${escape("||")}|&&"))
        var len = 0
        parts.forEach {

            val part = it.trim()
            val regexStringBuilder = StringBuilder()
            var prefix = ""
            acceptedExpressions.forEach { exp ->
                regexStringBuilder.append(prefix)
                prefix = "|"
                regexStringBuilder.append(escape(exp))
            }
            val variables = part.split(Regex(regexStringBuilder.toString()))

            var sublen = 0
            variables.filter { it.isNotBlank() && !acceptedExpressions.contains(it) }.forEach { vr ->
                val vrt = vr.trim()
                if (!numbers.contains(vrt[0]))
                    if (part.startsWith("!"))
                        expression.add("!v_next.$vrt")
                    else
                        expression.add("v_next.$vrt")
                else
                    expression.add(vrt)
                sublen += vrt.length
                if (it.length > sublen + 3) {
                    var e = it.substring(sublen + 1, sublen + 3).trim()
                    if (acceptedExpressions.contains(e)) {
                        expression.add(e)
                        sublen += 2
                    } else {
                        e = it.substring(sublen, sublen + 1)
                        if (acceptedExpressions.contains(e)) {
                            expression.add(e)
                            sublen += 1
                        }
                    }
                } else if (it.length > sublen + 1 && it.startsWith("!").not()) {
                    val e = it.substring(sublen, sublen + 1)
                    if (acceptedExpressions.contains(e) && e != "!") {
                        expression.add(e)
                        sublen += 1
                    }
                }
            }
            len += it.length
            if (name.length > len + 3) {
                expression.add(name.substring(len, len + 2))
                len += 2
            }
        }
    }

    override fun toString(): String {
        val sb = StringBuilder()
        var prefix = ""
        expression.forEach {
            sb.append(prefix)
            prefix = " "
            sb.append(it)
        }
        return sb.toString()
    }
}