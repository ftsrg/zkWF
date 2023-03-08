package web3;

import eu.toldi.bpmn_zkp.web3.StaticArray2;
import org.web3j.abi.FunctionEncoder;
import org.web3j.abi.TypeReference;
import org.web3j.abi.datatypes.*;
import org.web3j.abi.datatypes.generated.Uint256;
import org.web3j.crypto.Credentials;
import org.web3j.protocol.Web3j;
import org.web3j.protocol.core.RemoteCall;
import org.web3j.protocol.core.RemoteFunctionCall;
import org.web3j.protocol.core.methods.response.TransactionReceipt;
import org.web3j.tx.Contract;
import org.web3j.tx.TransactionManager;
import org.web3j.tx.gas.ContractGasProvider;

import java.math.BigInteger;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;

/**
 * <p>Auto generated code.
 * <p><strong>Do not modify!</strong>
 * <p>Please use the <a href="https://docs.web3j.io/command_line.html">web3j command line tools</a>,
 * or the org.web3j.codegen.SolidityFunctionWrapperGenerator in the
 * <a href="https://github.com/web3j/web3j/tree/master/codegen">codegen module</a> to update.
 *
 * <p>Generated with web3j version 4.8.8.
 */
@SuppressWarnings("rawtypes")
public class Model extends Contract {
    public static String BINARY;

    public static final String FUNC_GETCIPHERTEXT = "getCiphertext";

    public static final String FUNC_GETCURRENTHASH = "getCurrentHash";

    public static final String FUNC_GETLASTSIGNATURE = "getLastSignature";

    public static final String FUNC_STEPMODEL = "stepModel";

    public static final String FUNC_VERIFYTX = "verifyTx";

    @Deprecated
    protected Model(String BINARY,String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    protected Model(String BINARY,String contractAddress, Web3j web3j, Credentials credentials, ContractGasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    @Deprecated
    protected Model(String BINARY,String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }

    protected Model(String BINARY,String contractAddress, Web3j web3j, TransactionManager transactionManager, ContractGasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public RemoteFunctionCall<String> getCiphertext() {
        final Function function = new Function(FUNC_GETCIPHERTEXT, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Utf8String>() {}));
        return executeRemoteCallSingleValueReturn(function, String.class);
    }

    public RemoteFunctionCall<Hash> getCurrentHash() {
        final Function function = new Function(FUNC_GETCURRENTHASH, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Hash>() {}));
        return executeRemoteCallSingleValueReturn(function, Hash.class);
    }

    public RemoteFunctionCall<Signature> getLastSignature() {
        final Function function = new Function(FUNC_GETLASTSIGNATURE, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Signature>() {}));
        return executeRemoteCallSingleValueReturn(function, Signature.class);
    }

    public RemoteFunctionCall<TransactionReceipt> stepModel(Hash hash, String ciphertext, Signature sig_new, Proof p) {
        final Function function = new Function(
                FUNC_STEPMODEL, 
                Arrays.<Type>asList(hash, 
                new org.web3j.abi.datatypes.Utf8String(ciphertext), 
                sig_new, 
                p), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<Boolean> verifyTx(Proof proof, List<BigInteger> input) {
        final Function function = new Function(FUNC_VERIFYTX, 
                Arrays.<Type>asList(proof, 
                new org.web3j.abi.datatypes.generated.StaticArray19<org.web3j.abi.datatypes.generated.Uint256>(
                        org.web3j.abi.datatypes.generated.Uint256.class,
                        org.web3j.abi.Utils.typeMap(input, org.web3j.abi.datatypes.generated.Uint256.class))), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Bool>() {}));
        return executeRemoteCallSingleValueReturn(function, Boolean.class);
    }

    @Deprecated
    public static Model load(String BINARY,String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        return new Model(BINARY,contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    @Deprecated
    public static Model load(String BINARY,String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        return new Model(BINARY,contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }

    public static Model load(String BINARY,String contractAddress, Web3j web3j, Credentials credentials, ContractGasProvider contractGasProvider) {
        return new Model(BINARY,contractAddress, web3j, credentials, contractGasProvider);
    }

    public static Model load(String BINARY,String contractAddress, Web3j web3j, TransactionManager transactionManager, ContractGasProvider contractGasProvider) {
        return new Model(BINARY,contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public static RemoteCall<Model> deploy(Web3j web3j, Credentials credentials, ContractGasProvider contractGasProvider, Hash start_hash, String start_ciphertext) {
        String encodedConstructor = FunctionEncoder.encodeConstructor(Arrays.<Type>asList(start_hash, 
                new org.web3j.abi.datatypes.Utf8String(start_ciphertext)));
        return deployRemoteCall(Model.class, web3j, credentials, contractGasProvider, BINARY, encodedConstructor);
    }

    public static RemoteCall<Model> deploy(Web3j web3j, TransactionManager transactionManager, ContractGasProvider contractGasProvider, Hash start_hash, String start_ciphertext) {
        String encodedConstructor = FunctionEncoder.encodeConstructor(Arrays.<Type>asList(start_hash, 
                new org.web3j.abi.datatypes.Utf8String(start_ciphertext)));
        return deployRemoteCall(Model.class, web3j, transactionManager, contractGasProvider, BINARY, encodedConstructor);
    }

    @Deprecated
    public static RemoteCall<Model> deploy(Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit, Hash start_hash, String start_ciphertext) {
        String encodedConstructor = FunctionEncoder.encodeConstructor(Arrays.<Type>asList(start_hash, 
                new org.web3j.abi.datatypes.Utf8String(start_ciphertext)));
        return deployRemoteCall(Model.class, web3j, credentials, gasPrice, gasLimit, BINARY, encodedConstructor);
    }

    @Deprecated
    public static RemoteCall<Model> deploy(Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit, Hash start_hash, String start_ciphertext) {
        String encodedConstructor = FunctionEncoder.encodeConstructor(Arrays.<Type>asList(start_hash, 
                new org.web3j.abi.datatypes.Utf8String(start_ciphertext)));
        return deployRemoteCall(Model.class, web3j, transactionManager, gasPrice, gasLimit, BINARY, encodedConstructor);
    }

    public static class Hash extends StaticStruct {
        public Uint256 a;

        public Uint256 b;

        public Uint256 c;

        public Uint256 d;

        public Uint256 e;

        public Uint256 f;

        public Uint256 g;

        public Uint256 h;


        public Hash(Uint256 a, Uint256 b, Uint256 c, Uint256 d, Uint256 e, Uint256 f, Uint256 g, Uint256 h) {
            super(a,b,c,d,e,f,g,h);
            this.a = a;
            this.b = b;
            this.c = c;
            this.d = d;
            this.e = e;
            this.f = f;
            this.g = g;
            this.h = h;
        }
    }

    public static class Signature extends MyStaticStruct {
        public StaticArray2<Uint256> R;

        public Uint256 S;

        public Signature(List<Uint256> R, Uint256 S) {
            super(new StaticArray2<Uint256>(R),S);
            this.R = new StaticArray2<Uint256>(R);
            this.S = S;
        }

    }

    public static class G1Point extends StaticStruct {
        public Uint256 X;

        public Uint256 Y;

        public G1Point(Uint256 X, Uint256 Y) {
            super(X,Y);
            this.X = X;
            this.Y = Y;
        }

    }

    public static class G2Point extends MyStaticStruct {
        public StaticArray2<Uint256> X;

        public StaticArray2<Uint256> Y;

        public G2Point(List<Uint256> X, List<Uint256> Y) {
            super(new org.web3j.abi.datatypes.generated.StaticArray2<org.web3j.abi.datatypes.generated.Uint256>(X),new org.web3j.abi.datatypes.generated.StaticArray2<org.web3j.abi.datatypes.generated.Uint256>(Y));
            this.X = new StaticArray2<Uint256>(X);
            this.Y = new StaticArray2<Uint256>(Y);
        }

        public G2Point(StaticArray2<Uint256> X, StaticArray2<Uint256> Y) {
            super(X,Y);
            this.X = X;
            this.Y = Y;
        }
    }

    public static class Proof extends MyStaticStruct {
        public G1Point a;

        public G2Point b;

        public G1Point c;

        public Proof(G1Point a, G2Point b, G1Point c) {
            super(a,b,c);
            this.a = a;
            this.b = b;
            this.c = c;
        }
    }
}
