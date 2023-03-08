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
import org.hyperledger.fabric.contract.Context;
import org.hyperledger.fabric.contract.ContractInterface;
import org.hyperledger.fabric.contract.annotation.Contract;
import org.hyperledger.fabric.contract.annotation.Default;
import org.hyperledger.fabric.contract.annotation.Transaction;
import org.hyperledger.fabric.contract.annotation.Contact;
import org.hyperledger.fabric.contract.annotation.Info;
import org.hyperledger.fabric.contract.annotation.License;


import java.io.IOException;

import static java.nio.charset.StandardCharsets.UTF_8;

@Contract(name = "VeriferContract",
        info = @Info(title = "Example Verifier contract",
                description = "Hyperledger Fabric Smart-contract, that can verify proofs generated with ZoKrates",
                version = "0.0.1",
                license =
                @License(name = "GPL-3.0-or-later"
                ),
                contact = @Contact(email = "balazs.toldi@edu.bme.hu",
                        name = "verifier",
                        url = "https://git.toldi.eu/Bazsalanszky/fabric-ZoKrates-verifier")))
@Default
public class VerifierContract implements ContractInterface {



    public VerifierContract() throws IOException, InterruptedException, ZokratesInstallFailed {
            ZoKrates.INSTANCE.initZokrates();
    }

    @Transaction()
    public boolean proofExists(Context ctx, String proofId) {
        byte[] buffer = ctx.getStub().getState(proofId);
        return (buffer != null && buffer.length > 0);
    }

    @Transaction
    public void addVerifierKey(Context ctx,String verifierKeyID,VerificationKey verificationKey){
        boolean exists = proofExists(ctx, verifierKeyID);
        if (exists) {
            throw new RuntimeException("The verification key " + verifierKeyID + " already exists");
        }
        ctx.getStub().putState(verifierKeyID, verificationKey.toJSONString().getBytes(UTF_8));
    }

    @Transaction()
    public boolean addProof(Context ctx,String verifierKeyID, String proofId, zkProof proof) throws ZokratesNotInititialised, IOException, InterruptedException {
        boolean exists = proofExists(ctx, proofId);
        if (exists) {
            throw new RuntimeException("The asset " + proofId + " already exists");
        }

        exists = proofExists(ctx, verifierKeyID);
        if (!exists) {
            throw new RuntimeException("The verification key " + proofId + " does not exists!");
        }
        VerificationKey vk = readVerificationKey(ctx, verifierKeyID);
        boolean result = ZoKrates.INSTANCE.verify(proof, vk);
        Logger.getLogger("VerifierContract").info("Proof verified");
        if (result)
            ctx.getStub().putState(proofId, proof.toJSONString().getBytes(UTF_8));
        return result;
    }

    @Transaction()
    public zkProof readProof(Context ctx, String proofId) {
        boolean exists = proofExists(ctx, proofId);
        if (!exists) {
            throw new RuntimeException("The asset " + proofId + " does not exist");
        }

        return zkProof.fromJSONString(new String(ctx.getStub().getState(proofId), UTF_8));
    }

    @Transaction()
    public VerificationKey readVerificationKey(Context ctx, String verifierKeyID) {
        boolean exists = proofExists(ctx, verifierKeyID);
        if (!exists) {
            throw new RuntimeException("The asset " + verifierKeyID + " does not exist");
        }

        return VerificationKey.fromJSONString(new String(ctx.getStub().getState(verifierKeyID), UTF_8));
    }


}
