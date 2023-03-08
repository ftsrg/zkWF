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

import com.owlike.genson.Genson;
import org.hyperledger.fabric.contract.annotation.DataType;
import org.hyperledger.fabric.contract.annotation.Property;

@DataType()
public class zkProof {

    private final static Genson genson = new Genson();

    @Property
    private Proof proof;
    @Property
    private String[] inputs;

   public Proof getProof() { return proof; }
   public void setProof(Proof value) { this.proof = value; }

    public String[] getInputs() { return inputs; }
    public void setInputs(String[] value) { this.inputs = value; }

    public String toJSONString() {
        return genson.serialize(this).toString();
    }

    public static zkProof fromJSONString(String json) {
        zkProof asset = genson.deserialize(json, zkProof.class);
        return asset;
    }

}
