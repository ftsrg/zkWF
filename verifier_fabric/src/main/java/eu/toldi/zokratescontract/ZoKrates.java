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
import eu.toldi.zokratescontract.exceptions.ZokratesVerificationFailed;
import org.hyperledger.fabric.Logger;

import java.io.*;
import java.net.URL;
import java.util.Objects;

public enum ZoKrates {
    INSTANCE;
    private boolean isInitDone = false;
    private final Logger logger = Logger.getLogger("Zokrates");
    private String zokratesPath;


    private void checkInit() throws ZokratesNotInititialised {
        if (!isInitDone) {
            throw new ZokratesNotInititialised();
        }
    }

    private void installZokrates() throws IOException, InterruptedException, ZokratesInstallFailed {
        File tmp = new File(System.getProperty("user.dir") + "/tmp");
        tmp.mkdir();
        File dlFile = new File(tmp,"zokrates.sh");
        logger.info(dlFile.getAbsoluteFile().getAbsolutePath().toString());
        try (BufferedInputStream inputStream = new BufferedInputStream(new URL("https://raw.githubusercontent.com/ZoKrates/ZoKrates/master/scripts/one_liner.sh").openStream());
             FileOutputStream fileOS = new FileOutputStream(dlFile)) {
            byte data[] = new byte[1024];
            int byteContent;
            while ((byteContent = inputStream.read(data, 0, 1024)) != -1) {
                fileOS.write(data, 0, byteContent);
            }
        } catch (IOException e) {
            // handles IO exceptions
        }
        dlFile.setExecutable(true);
        ProcessBuilder pb = new ProcessBuilder(dlFile.getAbsoluteFile().getAbsolutePath().toString(), "--to", ".zokrates");
        Process process = pb.start();
        int exitCode = process.waitFor();
        if (exitCode != 0) {
            throw new ZokratesInstallFailed();
        }
        dlFile.delete();
        tmp.delete();
        logger.info("ZoKrates installed");
    }

    public void initZokrates() throws IOException, InterruptedException, ZokratesInstallFailed {
        if ( !(new File(System.getProperty("user.dir") + "/.zokrates/bin/zokrates").exists())) {
            installZokrates();
        }
        zokratesPath = System.getProperty("user.dir") + "/.zokrates/bin/zokrates";
        isInitDone = true;
        logger.info("ZoKrates initialized");
    }

    public boolean verify(zkProof p, VerificationKey v) throws ZokratesNotInititialised, IOException, InterruptedException {
        checkInit();
        File verificationKey = new File(System.getProperty("user.dir") +"/verification.key");
        File proofJson = new File(System.getProperty("user.dir") +"/proof.json");
        writeToFile(verificationKey,v.toJSONString());
        writeToFile(proofJson,p.toJSONString());

        ProcessBuilder pb = new ProcessBuilder(zokratesPath, "verify");
        Process process = pb.start();
        int exitCode = process.waitFor();
        String result = null;
        if (exitCode == 0) {
            BufferedReader reader =
                    new BufferedReader(new InputStreamReader(process.getInputStream()));
            StringBuilder builder = new StringBuilder();
            String line = null;
            while ( (line = reader.readLine()) != null) {
                builder.append(line);
                builder.append(System.getProperty("line.separator"));
            }
            result = builder.toString();

        }else {
            throw new ZokratesVerificationFailed();
        }

        proofJson.delete();
        verificationKey.delete();
        return Objects.equals(result.split("\n")[1], "PASSED");
    }

    private void writeToFile(File file,String data){
        try(FileOutputStream fos = new FileOutputStream(file);
            BufferedOutputStream bos = new BufferedOutputStream(fos)) {
            //convert string to byte array
            byte[] bytes = data.getBytes();
            //write byte array to file
            bos.write(bytes);
            bos.close();
            fos.close();
            System.out.print("Data written to file successfully.");
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
