/*
 * SPDX-License-Identifier: Apache License 2.0
 */

package eu.toldi;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.junit.jupiter.api.Assertions.assertTrue;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.verify;


public final class VerifierContractTest {
/*
    @Nested
    class AssetExists {
        @Test
        public void noProperProof() throws IOException, InterruptedException, ZokratesInstallFailed {

            VerifierContract contract = new VerifierContract();
            Context ctx = mock(Context.class);
            ChaincodeStub stub = mock(ChaincodeStub.class);
            when(ctx.getStub()).thenReturn(stub);

            when(stub.getState("10001")).thenReturn(new byte[] {});
            boolean result = contract.proofExists(ctx,"10001");

            assertFalse(result);
        }

        @Test
        public void proofExists() throws IOException, InterruptedException, ZokratesInstallFailed {

            VerifierContract contract = new VerifierContract();
            Context ctx = mock(Context.class);
            ChaincodeStub stub = mock(ChaincodeStub.class);
            when(ctx.getStub()).thenReturn(stub);

            when(stub.getState("10001")).thenReturn(new byte[] {42});
            boolean result = contract.proofExists(ctx,"10001");

            assertTrue(result);

        }

        @Test
        public void noKey() throws IOException, InterruptedException, ZokratesInstallFailed {
            VerifierContract contract = new VerifierContract();
            Context ctx = mock(Context.class);
            ChaincodeStub stub = mock(ChaincodeStub.class);
            when(ctx.getStub()).thenReturn(stub);

            when(stub.getState("10002")).thenReturn(null);
            boolean result = contract.proofExists(ctx,"10002");

            assertFalse(result);

        }

    }

    @Nested
    class AssetCreates {

        @Test
        public void verifyValidProof() throws ZokratesNotInititialised, IOException, InterruptedException, ZokratesInstallFailed {
            VerifierContract contract = new VerifierContract();
            Context ctx = mock(Context.class);
            ChaincodeStub stub = mock(ChaincodeStub.class);
            when(ctx.getStub()).thenReturn(stub);

            zkProof p = zkProof.fromJSONString("{\"proof\":{\"a\":[\"0x08a7650fbf23dc49a951a66b86d7a644f7c30c32146ebf67a9b3f3f82ac2fdd4\",\"0x1cd97a7f659c07e306486903249116cad17d206be60d569876e2163645da44c9\"],\"b\":[[\"0x2fefd1e04e0a314667d655661fea6b871e003c305606a025f49c5fbdaa4e8f95\",\"0x2c28f5820fbe152a54b46c1414e97888ffae05d499a3a7f19de42e817521ca00\"],[\"0x1e89653794007282f8d751c8034492d07e7f5674e34bf1dbf14edec8986eef84\",\"0x0a36a03575d69edf7e8d2b5c6deb088addec27de3e268f73dd3545301e57c7ef\"]],\"c\":[\"0x2985658325dc4da07722f549532a2650d5995b835a9a12cb996560bc7a1d563a\",\"0x29cdeeff2ede14d250df61da4e56b391bf4ce505fa3f8c703a14843e47f69d42\"]},\"inputs\":[\"0x0000000000000000000000000000000000000000000000000000000000000002\"]}");

            boolean result = contract.addProof(ctx,"10001",p);
            assertTrue(result);
        }

        @Test
        public void alreadyExists() throws IOException, InterruptedException, ZokratesInstallFailed {
            VerifierContract contract = new VerifierContract();
            Context ctx = mock(Context.class);
            ChaincodeStub stub = mock(ChaincodeStub.class);
            when(ctx.getStub()).thenReturn(stub);

            when(stub.getState("10002")).thenReturn(new byte[] { 42 });
            zkProof p = zkProof.fromJSONString("{\"proof\":{\"a\":[\"0x08a7650fbf23dc49a951a66b86d7a644f7c30c32146ebf67a9b3f3f82ac2fdd4\",\"0x1cd97a7f659c07e306486903249116cad17d206be60d569876e2163645da44c9\"],\"b\":[[\"0x2fefd1e04e0a314667d655661fea6b871e003c305606a025f49c5fbdaa4e8f95\",\"0x2c28f5820fbe152a54b46c1414e97888ffae05d499a3a7f19de42e817521ca00\"],[\"0x1e89653794007282f8d751c8034492d07e7f5674e34bf1dbf14edec8986eef84\",\"0x0a36a03575d69edf7e8d2b5c6deb088addec27de3e268f73dd3545301e57c7ef\"]],\"c\":[\"0x2985658325dc4da07722f549532a2650d5995b835a9a12cb996560bc7a1d563a\",\"0x29cdeeff2ede14d250df61da4e56b391bf4ce505fa3f8c703a14843e47f69d42\"]},\"inputs\":[\"0x0000000000000000000000000000000000000000000000000000000000000002\"]}");
            Exception thrown = assertThrows(RuntimeException.class, () -> contract.addProof(ctx, "10002", p));

            assertEquals(thrown.getMessage(), "The asset 10002 already exists");

        }


        @Test
        public void verifyInvalidProof() throws ZokratesNotInititialised, IOException, InterruptedException, ZokratesInstallFailed {
            VerifierContract contract = new VerifierContract();
            Context ctx = mock(Context.class);
            ChaincodeStub stub = mock(ChaincodeStub.class);
            when(ctx.getStub()).thenReturn(stub);
            String json = "{\"proof\":{\"a\":[\"0x08a7650fbf23dc49a951a66b86d7a644f7c30c32146ebf67a9b3f3f82ac2fdd4\",\"0x1cd97a7f659c07e306486903249116cad17d206be60d569876e2163645da44c9\"],\"b\":[[\"0x2fefd1e04e0a314667d655661fea6b871e003c305606a025f49c5fbdaa4e8f95\",\"0x2c28f5820fbe152a54b46c1414e97888ffae05d499a3a7f19de42e817521ca00\"],[\"0x1e89653794007282f8d751c8034492d07e7f5674e34bf1dbf14edec8986eef84\",\"0x0a36a03575d69edf7e8d2b5c6deb088addec27de3e268f73dd3545301e57c7ef\"]],\"c\":[\"0x2985658325dc4da07722f549532a2650d5995b835a9a12cb996560bc7a1d563a\",\"0x29cdeeff2ede14d250df61da4e56b391bf4ce505fa3f8c703a14843e47f69d42\"]},\"inputs\":[\"0x0000000000000000000000000000000000000000000000000000000000000003\"]}";
            zkProof p = zkProof.fromJSONString(json);

            boolean result = contract.addProof(ctx,"10001",p);
            assertFalse(result);
        }
    }

    @Test
    public void verifyValidProof() throws ZokratesNotInititialised, IOException, InterruptedException, ZokratesInstallFailed {
        VerifierContract contract = new VerifierContract();
        Context ctx = mock(Context.class);
        ChaincodeStub stub = mock(ChaincodeStub.class);
        when(ctx.getStub()).thenReturn(stub);

        zkProof p = zkProof.fromJSONString("{\"proof\":{\"a\":[\"0x08a7650fbf23dc49a951a66b86d7a644f7c30c32146ebf67a9b3f3f82ac2fdd4\",\"0x1cd97a7f659c07e306486903249116cad17d206be60d569876e2163645da44c9\"],\"b\":[[\"0x2fefd1e04e0a314667d655661fea6b871e003c305606a025f49c5fbdaa4e8f95\",\"0x2c28f5820fbe152a54b46c1414e97888ffae05d499a3a7f19de42e817521ca00\"],[\"0x1e89653794007282f8d751c8034492d07e7f5674e34bf1dbf14edec8986eef84\",\"0x0a36a03575d69edf7e8d2b5c6deb088addec27de3e268f73dd3545301e57c7ef\"]],\"c\":[\"0x2985658325dc4da07722f549532a2650d5995b835a9a12cb996560bc7a1d563a\",\"0x29cdeeff2ede14d250df61da4e56b391bf4ce505fa3f8c703a14843e47f69d42\"]},\"inputs\":[\"0x0000000000000000000000000000000000000000000000000000000000000002\"]}");

        boolean result = contract.addProof(ctx,"10001",p);
        assertTrue(result);
    }
    */
}
