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

package eu.toldi.bpmn_zkp.utils

import org.w3c.dom.Node
import org.w3c.dom.NodeList


fun NodeList.foreach(func: (Node) -> Unit) {
    for (i in 0..this.length) {
        if (this.item(i) != null && this.item(i).nodeName != "#text")
            func(this.item(i))
    }
}

operator fun NodeList.get(index: Int): Node {
    return item(index)
}