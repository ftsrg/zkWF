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
import com.owlike.genson.annotation.JsonProperty;
import org.hyperledger.fabric.contract.annotation.DataType;
import org.hyperledger.fabric.contract.annotation.Property;

@DataType()
public class VerificationKey {

    private final static Genson genson = new Genson();

    @Property
    private String[] alpha;
    @Property
    private String[][] beta;
    @Property
    private String[][] gamma;
    @Property
    private String[][] delta;
    @Property
    @JsonProperty("gamma_abc")
    private String[][] gammaAbc;

    public String[] getAlpha() { return alpha; }
    public void setAlpha(String[] value) { this.alpha = value; }

    public String[][] getBeta() { return beta; }
    public void setBeta(String[][] value) { this.beta = value; }

    public String[][] getGamma() { return gamma; }
    public void setGamma(String[][] value) { this.gamma = value; }

    public String[][] getDelta() { return delta; }
    public void setDelta(String[][] value) { this.delta = value; }

    public String[][] getGammaAbc() { return gammaAbc; }
    public void setGammaAbc(String[][] value) { this.gammaAbc = value; }

    public String toJSONString() {
        return genson.serialize(this).toString();
    }

    public static VerificationKey fromJSONString(String json) {
        VerificationKey asset = genson.deserialize(json, VerificationKey.class);
        return asset;
    }
}
