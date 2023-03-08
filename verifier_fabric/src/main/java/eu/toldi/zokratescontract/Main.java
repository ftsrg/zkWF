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

package eu.toldi.zokratescontract;

import eu.toldi.zokratescontract.exceptions.ZokratesInstallFailed;
import eu.toldi.zokratescontract.exceptions.ZokratesNotInititialised;
import org.hyperledger.fabric.Logger;

import java.io.IOException;

public class Main {
    public static void main(String[] args) throws IOException, InterruptedException, ZokratesInstallFailed, ZokratesNotInititialised {
        VerificationKey vk = VerificationKey.fromJSONString("{\"alpha\":[\"0x0415c7f4fa9a3627a9a6741575a88c66ac28c8a89aba9a354b0fc855fa926bee\",\"0x03d4c63a196e9471f7e9f0fd6ba6919e42a409973a14c0a8fa85c2f284002857\"],\"beta\":[[\"0x0f1317895e7b7ec5ec22b6b5b08964c713056cb4ee5418d8558520023e0e50f5\",\"0x192fc78100545459301a73735c0972bd6e254aadaa6f5a89b30cc0f33ffd7016\"],[\"0x1761daf3be79dfd093eaec72fd28f22c1f67567c95c6b41c9c9d15084898f6c2\",\"0x123bfe0108fd5a41d53440c7bf8ec02a01b8172057dea4d08ec85db731795e19\"]],\"gamma\":[[\"0x0002184ed6627f349adf557804b01b37d68f4b169724038a13eedd42337bbb7d\",\"0x2709d7c95e243b7ef8f0a122208aa9d07035455934bdc4dd00f52f4696790187\"],[\"0x05eb29c4f76a3d27ef1e2b6d4cafd55627bf2b05a3f1bad02e70b11ae5018103\",\"0x09d85e2c7f161f9eb605e39cf37f2cafe466cd9b0035cc622c9b8207b17785e6\"]],\"delta\":[[\"0x1b08ba6b2bbfd5f7ead50a4cb534314ba972956f175eab551f4f1960f67c060e\",\"0x06deff975ae90b1c956d502a1c60e2d9be7a5575eda1128db02906ad7f4b8bc0\"],[\"0x07826cb269cf16302716a17f50aaa0d122bcc6c7db1b37ef238da67ece215c06\",\"0x0aa7a6dcd510e5f5ae710cfe0b4d6cd814be3e540f428f3bac29820a29295232\"]],\"gamma_abc\":[[\"0x10fad9db068896156ba8776fb290d773d7c8cd1ab84c5faf2adfea05a4f2f72e\",\"0x1ca7128ba8b4be0070b9837f03f9321034c7edb07dd4f8e6639aa3d8b614f394\"],[\"0x0aa6a354aaa90788ea23a2f9504f30c4fc0b95a7f4fc741f3a06cb4ab0c4e1bb\",\"0x2873bbd367c43752ba31394ea386f7a4faa5d078cb23816ad51c820e944e7e65\"]]}");
        zkProof p = zkProof.fromJSONString("{\"proof\":{\"a\":[\"0x08a7650fbf23dc49a951a66b86d7a644f7c30c32146ebf67a9b3f3f82ac2fdd4\",\"0x1cd97a7f659c07e306486903249116cad17d206be60d569876e2163645da44c9\"],\"b\":[[\"0x2fefd1e04e0a314667d655661fea6b871e003c305606a025f49c5fbdaa4e8f95\",\"0x2c28f5820fbe152a54b46c1414e97888ffae05d499a3a7f19de42e817521ca00\"],[\"0x1e89653794007282f8d751c8034492d07e7f5674e34bf1dbf14edec8986eef84\",\"0x0a36a03575d69edf7e8d2b5c6deb088addec27de3e268f73dd3545301e57c7ef\"]],\"c\":[\"0x2985658325dc4da07722f549532a2650d5995b835a9a12cb996560bc7a1d563a\",\"0x29cdeeff2ede14d250df61da4e56b391bf4ce505fa3f8c703a14843e47f69d42\"]},\"inputs\":[\"0x0000000000000000000000000000000000000000000000000000000000000002\"]}");
        ZoKrates zok = ZoKrates.INSTANCE;
        zok.initZokrates();
        Logger logger = Logger.getLogger("MainTest");
        boolean b = zok.verify(p,vk);
        if(b){
            logger.info("Success");
        }else {
            logger.info("Failed");
        }
    }
}
