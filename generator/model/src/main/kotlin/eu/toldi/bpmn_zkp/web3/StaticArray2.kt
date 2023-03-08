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

package eu.toldi.bpmn_zkp.web3

import org.web3j.abi.datatypes.StaticArray
import org.web3j.abi.datatypes.Type


class StaticArray2<T : Type<*>> : StaticArray<T> {
    @Deprecated("")
    constructor(values: List<T>) : super(2, values) {
    }

    @Deprecated("")
    @SafeVarargs
    constructor(vararg values: T) : super(2, *values) {
    }

    constructor(type: Class<T>, values: List<T>?) : super(type, 2, values) {}

    @SafeVarargs
    constructor(type: Class<T>, vararg values: T) : super(type, 2, *values) {
    }

    override fun getTypeAsString(): String {
        return value[0].typeAsString + "[" + value.size + "]"
    }


}
