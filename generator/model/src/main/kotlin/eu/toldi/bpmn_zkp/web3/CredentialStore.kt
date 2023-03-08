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

import org.web3j.crypto.Credentials

object CredentialStore {
    val credentialsList = listOf(
        //Credentials.create("ae816ff258665a9689d57baa16005790537c772a788dcc33f9aeba0a2e5c5fb0"),
        Credentials.create("0x4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d"),
        Credentials.create("0x6cbed15c793ce57650b9877cf6fa156fbef513c4e6134f022a85b1ffdd59b2a1"),
        Credentials.create("0x6370fd033278c143179d81c5526140625662b8daa446c22ee2d73db3707e620c"),
        Credentials.create("0x646f1ce2fdad0e6deeeb5c7e8e5543bdde65e86029e2fd9fc169899c440a7913"),
        Credentials.create("0xadd53f9a7e588d003326d1cbf9e4a43c061aadd9bc938c843a79e7b4fd2ad743"),
        Credentials.create("0x395df67f0c2d2d9fe1ad08d1bc8b6627011959b79c53d7dd6a3536a33ab8a4fd"),
        Credentials.create("0xe485d098507f54e7733a205420dfddbe58db035fa577fc294ebd14db90767a52"),
        Credentials.create("0xa453611d9419d0e56f499079478fd72c37b251a94bfde4d19872c44cf65386e3"),
        Credentials.create("0x829e924fdf021ba3dbbc4225edfece9aca04b929d6e75613329ca6f1d31c0bb4"),
        Credentials.create("0xb0057716d5917badaf911b193b12b910811c1497b5bada8d7711f758981c3773"),
    )
}