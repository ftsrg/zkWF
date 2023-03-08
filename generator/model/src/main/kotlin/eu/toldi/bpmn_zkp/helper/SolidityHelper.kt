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

package eu.toldi.bpmn_zkp.helper


import org.web3j.abi.datatypes.Function
import org.web3j.abi.datatypes.Type
import org.web3j.crypto.Credentials
import org.web3j.crypto.RawTransaction
import org.web3j.crypto.TransactionEncoder
import org.web3j.protocol.Web3j
import org.web3j.protocol.core.DefaultBlockParameterName
import org.web3j.protocol.core.methods.response.EthGetTransactionCount
import org.web3j.protocol.core.methods.response.EthGetTransactionReceipt
import org.web3j.protocol.core.methods.response.EthSendTransaction
import org.web3j.protocol.core.methods.response.TransactionReceipt
import org.web3j.tx.RawTransactionManager
import org.web3j.tx.TransactionManager
import org.web3j.tx.gas.DefaultGasProvider
import org.web3j.tx.response.PollingTransactionReceiptProcessor
import org.web3j.tx.response.TransactionReceiptProcessor
import org.web3j.utils.Numeric
import web3.MyFunctionEncoder
import java.io.File
import java.math.BigInteger


class SolidityHelper(val web3j: Web3j, val credentials: Credentials) {

    fun compileContract(file: File): Int {
        val pb = ProcessBuilder(
            "solc",
            file.absoluteFile.absolutePath,
            "--bin",
            "--abi",
            "--optimize",
            "--overwrite",
            "-o",
            "build/"
        )
        val p = pb.start()
        return p.waitFor()
    }

    fun deployContract(bin: String, encodedParameters: String = ""): String {
        val ethGetTransactionCount: EthGetTransactionCount =
            web3j.ethGetTransactionCount(credentials.address, DefaultBlockParameterName.LATEST)
                .sendAsync().get()


        val nonce: BigInteger = ethGetTransactionCount.getTransactionCount()
        val binary = String(File(bin).readBytes())
        val rawTransaction: RawTransaction = RawTransaction.createContractTransaction(
            nonce,
            DefaultGasProvider.GAS_PRICE,
            BigInteger.valueOf(5_000_000),
            BigInteger.valueOf(0),
            "0x$binary$encodedParameters"
        )
        val signedMessage: ByteArray = TransactionEncoder.signMessage(rawTransaction, credentials)
        val hexValue: String = Numeric.toHexString(signedMessage)

        val transactionResponse: EthSendTransaction = web3j.ethSendRawTransaction(hexValue).sendAsync().get()

        var pendingTransaction = true
        if (transactionResponse.hasError()) {
            System.out.println(transactionResponse.error.data)
            System.out.println(transactionResponse.error.message)
        }
        val transactionHash = transactionResponse.transactionHash

        var transactionReceipt: EthGetTransactionReceipt = web3j.ethGetTransactionReceipt(transactionHash).send()

        while (true) {
            transactionReceipt = if (!transactionReceipt.getTransactionReceipt().isPresent()) {
                Thread.sleep(100);
                web3j.ethGetTransactionReceipt(transactionHash).send()
            } else {
                break
            }
        }
        val transactionReceiptFinal: TransactionReceipt = transactionReceipt.getTransactionReceipt().get()

        return transactionReceiptFinal.contractAddress
    }

    fun callContract(contract: String, functionName: String, arguments: List<Type<Any>>): TransactionReceipt {

        val function = Function(
            functionName,
            arguments, emptyList()
        )


        val txData: String = MyFunctionEncoder().encodeFunction(function)
        println("$function\n$txData")
        val txManager: TransactionManager = RawTransactionManager(web3j, credentials)

        val txHash: String = txManager.sendTransaction(
            DefaultGasProvider.GAS_PRICE,
            BigInteger.valueOf(5_000_000),
            contract,
            txData,
            BigInteger.ZERO
        ).transactionHash

        val receiptProcessor: TransactionReceiptProcessor = PollingTransactionReceiptProcessor(
            web3j,
            TransactionManager.DEFAULT_POLLING_FREQUENCY,
            TransactionManager.DEFAULT_POLLING_ATTEMPTS_PER_TX_HASH
        )
        val txReceipt: TransactionReceipt = receiptProcessor.waitForTransactionReceipt(txHash)
        println(txReceipt.status)
        return txReceipt
    }
}