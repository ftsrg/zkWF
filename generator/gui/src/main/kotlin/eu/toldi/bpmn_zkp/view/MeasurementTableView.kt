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

package eu.toldi.bpmn_zkp.view

import eu.toldi.bpmn_zkp.model.helper.MeasuredTimes

import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.GlobalScope
import kotlinx.coroutines.launch
import kotlinx.coroutines.withContext
import tornadofx.*
import java.io.File


class MeasurementTableView(val measurements: List<MeasuredTimes>) : View("Measurement Table") {
    override val root = vbox {
        minHeight = 800.0
        minWidth = 800.0
        button("Export to csv") {
            action {
                GlobalScope.launch {
                    val file = File("results.csv")
                    val sb = StringBuilder()
                    sb.appendLine("Compile Time;Setup Time;Witness Time;Proof time")
                    measurements.forEach {
                        sb.appendLine("${it.compileTime};${it.setUpTime};${it.witnessTime};${it.proofTime}")
                    }
                    file.writeText(sb.toString())
                    withContext(Dispatchers.Main) {
                        information("Success!", "Results written to results.csv file!")
                    }
                }

            }
        }
        tableview<MeasuredTimes> {
            minHeight = 800.0
            minWidth = 800.0
            isEditable = false
            val measuredTimes = measurements.toObservable()
            items = measuredTimes
            readonlyColumn("Test ID", MeasuredTimes::ID)
            readonlyColumn("Compile Time", MeasuredTimes::compileTime)
            readonlyColumn("Setup Time", MeasuredTimes::setUpTime)
            readonlyColumn("Witness Time", MeasuredTimes::witnessTime)
            readonlyColumn("Proof time", MeasuredTimes::proofTime)
        }
    }
}