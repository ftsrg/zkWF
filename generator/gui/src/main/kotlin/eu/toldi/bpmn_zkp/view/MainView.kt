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


import eu.toldi.bpmn_zkp.event.Event
import eu.toldi.bpmn_zkp.event.EventBus
import eu.toldi.bpmn_zkp.event.EventListener
import eu.toldi.bpmn_zkp.event.EventType
import eu.toldi.bpmn_zkp.helper.SolidityHelper
import eu.toldi.bpmn_zkp.model.Model
import eu.toldi.bpmn_zkp.model.StateProof
import eu.toldi.bpmn_zkp.model.testing.State
import eu.toldi.bpmn_zkp.web3.CredentialStore
import eu.toldi.`zokrates-wrapper`.Zokrates
import javafx.concurrent.Worker
import javafx.event.EventHandler
import javafx.geometry.Pos
import javafx.scene.Parent
import javafx.scene.Scene
import javafx.scene.control.Label
import javafx.scene.control.ProgressIndicator
import javafx.scene.control.TextInputDialog
import javafx.scene.input.ScrollEvent
import javafx.scene.layout.HBox
import javafx.scene.layout.Priority
import javafx.scene.layout.VBox
import javafx.scene.web.WebView
import javafx.stage.FileChooser
import javafx.stage.Modality
import javafx.stage.Stage
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.GlobalScope
import kotlinx.coroutines.launch
import kotlinx.coroutines.withContext
import org.web3j.abi.FunctionEncoder
import org.web3j.abi.datatypes.Type
import org.web3j.abi.datatypes.Utf8String
import org.web3j.abi.datatypes.generated.Uint256
import org.web3j.protocol.Web3j
import org.web3j.protocol.http.HttpService
import org.web3j.tx.gas.DefaultGasProvider
import tornadofx.*
import java.io.File
import java.math.BigInteger
import java.nio.charset.StandardCharsets
import kotlin.system.exitProcess
import kotlin.system.measureTimeMillis


class MainView : View("BPMN Zokrates Generator"), EventListener {
    var selectedFile: File? = null
    var codeGenerated = false
    var compileing = false
    var compiled = false
    var setupDone = false

    private val publicModels = File("../editor/public/models/").apply {
        mkdir()
    }
    val web3j = Web3j.build(HttpService("http://127.0.0.1:8545"))


    val solidityHelper = SolidityHelper(web3j, CredentialStore.credentialsList[0])

    private lateinit var loading: HBox
    private lateinit var parameters: VBox
    private lateinit var parameters2: VBox
    private lateinit var measurementTable: HBox
    private lateinit var webView: WebView
    private var baseURL = ""
    private var selectedModel: Model? = null
    private var deployedContract: String? = null
    private lateinit var currentState: State
    override val root: Parent = drawer {

        item("Modeller") {
            webview {
                engine.load("http://127.0.0.1:8080/")
                engine.loadWorker.stateProperty().onChange {
                    val wh = engine.executeScript("document.documentElement.scrollHeight").toString().toDouble()
                    prefHeight = wh
                }
                engine.loadWorker.stateProperty().addListener { ov, old, new ->
                    if (new == Worker.State.SCHEDULED) {
                        println("Success")
                        println(engine.location)

                        val files = chooseFile(
                            title = "Select a BPMN file",

                            filters = listOf<FileChooser.ExtensionFilter>(
                                FileChooser.ExtensionFilter(
                                    "BPMN models (*.bpmn)",
                                    "*.bpmn"
                                )
                            ).toTypedArray(),
                            mode = FileChooserMode.Save
                        )
                        if (files.size == 1) {
                            files[0].writeText(
                                java.net.URLDecoder.decode(
                                    engine.location.split(",")[1],
                                    StandardCharsets.UTF_8.name()
                                )
                            )
                        }
                    }
                }
            }
        }
        item("Test") {
            hbox {
                vbox {
                    paddingAll = 16.0
                    hbox {
                        paddingAll = 6.0
                        hbox {
                            paddingAll = 5.0
                            label("Select File:")
                        }
                        button("select...").apply {
                            action {
                                val files = chooseFile(
                                    "Select a BPMN file",
                                    listOf<FileChooser.ExtensionFilter>(
                                        FileChooser.ExtensionFilter(
                                            "BPMN models (*.bpmn)",
                                            "*.bpmn"
                                        )
                                    ).toTypedArray(),
                                    File("../models").absoluteFile
                                )
                                if (files.isNotEmpty())
                                    selectedFile = files[0]
                            }
                        }
                    }
                    hbox {
                        paddingAll = 6.0
                        hbox {
                            paddingAll = 6.0
                            button("Generate ZoKrates Code").apply {
                                action {
                                    if (selectedFile != null) {
                                        selectedFile?.let {
                                            try {
                                                loading += ProgressIndicator()
                                                val model = Model(it)
                                                model.generateZokratesCode()
                                                println("${model.variables}, ${model.variableWritePermission}")
                                                loading.clear()
                                                information("Success!", "ZoKrates Code generated to root.zok")
                                                codeGenerated = true
                                                parameters.clear()
                                                parameters += ComputeView(model, measurementTable)
                                            } catch (e: Exception) {
                                                loading.clear()
                                                error("Error occurred", "Error! Please select a valid BPMN file!")
                                                e.printStackTrace()
                                            }

                                        }
                                    } else {
                                        error("No file selected", "Error! Please select a valid BPMN file!")
                                    }
                                }
                            }
                        }
                        hbox {
                            paddingAll = 6.0
                            button("Compile ZoKrates Code").apply {
                                action {
                                    compileCode()
                                }
                            }

                        }

                        hbox {
                            paddingAll = 6.0
                            button("ZoKrates Setup").apply {
                                action {
                                    runSetup()
                                }
                            }

                        }
                    }
                    loading = hbox {

                    }
                }
                parameters = vbox {

                }

                measurementTable = hbox {}
            }
        }
        item("Deploy", expanded = true) {
            menubar {
                menu("File") {
                    item("Open...").apply {
                        action {
                            val files = chooseFile(
                                "Select a BPMN file",
                                listOf<FileChooser.ExtensionFilter>(
                                    FileChooser.ExtensionFilter(
                                        "BPMN models (*.bpmn)",
                                        "*.bpmn"
                                    )
                                ).toTypedArray(),
                                File("../models").absoluteFile
                            )
                            if (files.isNotEmpty()) {
                                selectedFile = files[0]
                                files[0].copyTo(File(publicModels, files[0].name), true)
                                try {
                                    loading += ProgressIndicator()
                                    selectedModel = Model(files[0])
                                    selectedModel!!.generateZokratesCode()
                                    println("${selectedModel!!.variables}, ${selectedModel!!.variableWritePermission}")
                                    loading.clear()
                                    information("Success!", "ZoKrates Code generated to root.zok")
                                    codeGenerated = true
                                    parameters2.clear()
                                    parameters2 += ComputeView(selectedModel!!, measurementTable)
                                    baseURL =
                                        "http://127.0.0.1:8080/embededViewer.html?url=http://127.0.0.1:8080/models/${files[0].name}"
                                    webView.engine.load("$baseURL&state=${selectedModel!!.getInitialStateVector()}")
                                    EventBus.emitEvent(
                                        Event(
                                            EventType.MODEL_SELECTED,
                                            files[0].name
                                        )
                                    )
                                } catch (e: Exception) {
                                    loading.clear()
                                    error("Error occurred", "Error! Please select a valid BPMN file!")
                                    e.printStackTrace()
                                }


                            }
                        }
                    }
                    item("Save")
                    item("Quit").apply {
                        action {
                            exitProcess(0)
                        }
                    }
                }
                menu("Zokrates") {
                    item("Setup").action {
                        runSetup()
                    }
                    item("Deploy Verifier").action {
                        runVerifierSetup()
                    }
                }
                menu("Smart-Contract") {
                    item("Use deployed").apply {
                        action {
                            val dialog = TextInputDialog("")
                            dialog.title = "Contract address"
                            // dialog.headerText = "Please login, then enter your authentication token"
                            dialog.contentText = "Contract address"
                            val result = dialog.showAndWait()
                            if (result.isPresent) {
                                val contract = web3.Model.load(
                                    String(File("build/Model.bin").readBytes()),
                                    result.get(),
                                    web3j,
                                    CredentialStore.credentialsList[0],
                                    DefaultGasProvider.GAS_PRICE,
                                    DefaultGasProvider.GAS_LIMIT
                                )
                                val text = contract.ciphertext.send()
                                println(text)
                                val state = State.fromJson(text)
                                EventBus.emitEvent(
                                    Event(EventType.CONTRACT_DEPLOYED, result.get())
                                )
                                EventBus.emitEvent(
                                    Event(EventType.CURRENT_STATE_CHANGED, state.toJson())
                                )
                            }
                        }
                    }
                }
            }
            toolbar {
                label("No model selected").apply {
                    EventBus.registerEventListener(EventListener { event ->
                        if (event.type == EventType.MODEL_SELECTED)
                            this.text = event.data
                    })
                }
                label("") {
                    EventBus.registerEventListener(EventListener { event ->
                        if (event.type == EventType.CONTRACT_DEPLOYED) {
                            this.text = "Deployed contract: ${event.data}"
                            deployedContract = event.data
                        }
                    })
                }
                hbox { hgrow = Priority.ALWAYS }
                val proof = combobox<String> {
                    items = mutableListOf<String>().also { list ->
                        list.addAll(
                            File(".").listFiles()
                                .filter { it.name.startsWith("stateProof") && it.name.endsWith(".json") }
                                .map { it.name })
                    }.toObservable()
                    EventBus.registerEventListener { event ->
                        if (event.type == EventType.PROOF_GENERATED) {
                            items.add(event.data)
                        }
                    }
                }
                val accounts = combobox<String> {
                    items = CredentialStore.credentialsList.map { it.address }.toObservable()
                }
                button("Step").apply {
                    action {
                        val dialog = loadingDialog("Calling the Smart-Contract...")
                        GlobalScope.launch(Dispatchers.IO) {
                            val solidityHelper = SolidityHelper(
                                web3j,
                                CredentialStore.credentialsList[CredentialStore.credentialsList.indexOf(CredentialStore.credentialsList.first { it.address == accounts.selectedItem })]
                            )
                            EventBus.emitEvent(Event(EventType.STEP_INPROGRESS, ""))
                            val stateProof = StateProof.fromJson(
                                File(proof.selectedItem!!).readText()
                            )
                            val proofObj = stateProof.proof



                            println(proofObj.toJson())
                            val newHash = web3.Model.Hash(
                                Uint256(BigInteger(proofObj.inputs[11].drop(2), 16)),
                                Uint256(BigInteger(proofObj.inputs[12].drop(2), 16)),
                                Uint256(BigInteger(proofObj.inputs[13].drop(2), 16)),
                                Uint256(BigInteger(proofObj.inputs[14].drop(2), 16)),
                                Uint256(BigInteger(proofObj.inputs[15].drop(2), 16)),
                                Uint256(BigInteger(proofObj.inputs[16].drop(2), 16)),
                                Uint256(BigInteger(proofObj.inputs[17].drop(2), 16)),
                                Uint256(BigInteger(proofObj.inputs[18].drop(2), 16)),
                            )
                            val signature = web3.Model.Signature(
                                listOf(
                                    Uint256(BigInteger(proofObj.inputs[8].drop(2), 16)),
                                    Uint256(BigInteger(proofObj.inputs[9].drop(2), 16))
                                ),
                                Uint256(BigInteger(proofObj.inputs[10].drop(2), 16))
                            )
                            val ciphertext = Utf8String(stateProof.state.toJson())
                            val proofParam = web3.Model.Proof(
                                web3.Model.G1Point(
                                    Uint256(BigInteger(proofObj.proof.a[0].drop(2), 16)),
                                    Uint256(BigInteger(proofObj.proof.a[1].drop(2), 16))
                                ),
                                web3.Model.G2Point(
                                    proofObj.proof.b[0].map { Uint256(BigInteger(it.drop(2), 16)) },
                                    proofObj.proof.b[1].map { Uint256(BigInteger(it.drop(2), 16)) }),
                                web3.Model.G1Point(
                                    Uint256(BigInteger(proofObj.proof.c[0].drop(2), 16)),
                                    Uint256(BigInteger(proofObj.proof.c[1].drop(2), 16))
                                )
                            )

                            val reciept = solidityHelper.callContract(
                                deployedContract!!,
                                "stepModel", listOf(newHash, ciphertext, signature, proofParam) as List<Type<Any>>
                            )
                            withContext(Dispatchers.Main) {
                                dialog.close()
                                if (reciept.isStatusOK) {
                                    currentState = stateProof.state
                                    EventBus.emitEvent(
                                        Event(
                                            EventType.CURRENT_STATE_CHANGED,
                                            stateProof.state.toJson()
                                        )
                                    )
                                    information("Stepping Model!")
                                }
                                EventBus.emitEvent(
                                    Event(
                                        EventType.STEP_COMPLETE,
                                        ""
                                    )
                                )
                            }
                        }
                    }
                }

            }
            borderpane {
                right {
                    parameters2 = vbox {

                    }
                }
                center {
                    webView = webview {
                        zoom = 10.0
                        engine.load("http://127.0.0.1:8080/embededViewer.html")
                        addEventFilter(ScrollEvent.SCROLL) { e: ScrollEvent ->
                            val deltaY = e.deltaY
                            if (e.isControlDown && deltaY > 0) {
                                zoom *= 1.1
                                e.consume()
                            } else if (e.isControlDown && deltaY < 0) {
                                zoom /= 1.1
                                prefWidth = width / 1.1
                                e.consume()
                            }
                        }

                        engine.loadWorker.stateProperty().onChange {
                            /*val wh =
                                engine.executeScript("document.documentElement.scrollHeight").toString().toDouble()
                            prefHeight = wh*/
                        }
                    }
                }
            }

        }

    }

    private fun runSetup() {
        if (selectedFile != null && codeGenerated) {
            selectedFile?.let {
                try {
                    val dialog = loadingDialog("Zokrates setup in progress...")
                    GlobalScope.launch(Dispatchers.IO) {
                        if (!compiled) {
                            Zokrates.compile("root.zok")
                            Zokrates.compile("hash.zok", "hash")
                        }
                        val elapsed = measureTimeMillis {
                            Zokrates.setup()
                        }
                        compileing = false
                        withContext(Dispatchers.Main) {
                            dialog.close()
                            information(
                                "Success!",
                                "ZoKrates setup done in ${elapsed.toDouble() / 1000}s !"
                            )
                        }
                        setupDone = true
                    }
                } catch (e: Exception) {
                    loading.clear()
                    error("Error occurred", "Error! Please select a valid BPMN file!")
                    e.printStackTrace()
                }

            }
        } else {
            error(
                "Unavailable",
                "Error! Code compilation is unavailable before code generation!"
            )
        }
    }

    private fun runVerifierSetup() {
        //  if (setupDone) {
        val loadingDialog = loadingDialog("Deploying verifier")
        GlobalScope.launch(Dispatchers.IO) {
            // val Random = random;
            Zokrates.exportVerifier()

            solidityHelper.compileContract(File("./model.sol"))
            val random = 1675454832 /* TODO: Replace this constant with a secure random number (for demo purposes it's better to have a constant)
                // kotlin.math.abs(Random().nextInt() * 10000)
                */
            println(random)
            val initialState = State(
                selectedModel!!.getInitialStateVector().split(" ").filter { it != "" },
                (random).toString(),
                Array(selectedModel!!.variables.size) { "0" }.asList(),
                Array(selectedModel!!.messages.size) { Array(8) { "0" }.asList() }.asList()
            )
            val hashValue = Zokrates.computeWithness(initialState.toArgs(), "hash", "hash_deploy.result")
            val hash = web3.Model.Hash(
                Uint256(BigInteger(hashValue[0])),
                Uint256(BigInteger(hashValue[1])),
                Uint256(BigInteger(hashValue[2])),
                Uint256(BigInteger(hashValue[3])),
                Uint256(BigInteger(hashValue[4])),
                Uint256(BigInteger(hashValue[5])),
                Uint256(BigInteger(hashValue[6])),
                Uint256(BigInteger(hashValue[7]))
            )

            val ciphertext = Utf8String(initialState.toJson())
            val dataParams = listOf(hash, ciphertext)

            val parameter = FunctionEncoder.encodeConstructor(
                dataParams
            )

            val contract = solidityHelper.deployContract("build/Model.bin", parameter)

            println("Contract deployed: $contract")
            withContext(Dispatchers.Main) {
                EventBus.emitEvent(Event(EventType.CONTRACT_DEPLOYED, contract))
                EventBus.emitEvent(Event(EventType.CURRENT_STATE_CHANGED, initialState.toJson()))
                loadingDialog.close()
            }
        }
        /*  } else {
              error(
                  "Setup needed",
                  "Error! Setup needed in order to export verifier contract!"
              )
          }*/
    }

    init {
        measurementTable += MeasurementTableView(emptyList())
        EventBus.registerEventListener(
            this
        )
    }

    fun compileCode() {
        if (selectedFile != null && codeGenerated) {
            selectedFile?.let {
                try {
                    compileing = true
                    loading += ProgressIndicator()
                    GlobalScope.launch(Dispatchers.IO) {
                        val elapsed = measureTimeMillis {
                            Zokrates.compile("root.zok")
                        }
                        compileing = false
                        withContext(Dispatchers.Main) {
                            loading.clear()
                            information(
                                "Success!",
                                "ZoKrates Code compiled in ${elapsed.toDouble() / 1000}s !"
                            )
                        }
                    }
                } catch (e: Exception) {
                    loading.clear()
                    error("Error occurred", "Error! Please select a valid BPMN file!")
                    e.printStackTrace()
                }

            }
        } else {
            error(
                "Unavailable",
                "Error! Code compilation is unavailable before code generation!"
            )
        }
    }

    fun buildParamaets(m: Model) {

        parameters.clear()
        //parameters +=
    }

    override fun onEvent(event: Event) {
        when {
            event.type == EventType.CURRENT_STATE_CHANGED && baseURL.startsWith("http://") -> {
                val state = buildString {
                    State.fromJson(event.data).stateVector.forEach {
                        append("$it ")
                    }
                }.trim()
                currentState = State.fromJson(event.data)
                webView.engine.load("$baseURL&state=${state}")
            }
            event.type == EventType.SETUP_DONE -> {
                setupDone = true
            }
            event.type == EventType.CONTRACT_DEPLOYED -> {
                GlobalScope.launch(Dispatchers.IO) {
                    Thread.sleep(10000)
                    while (true) {
                        try {
                            val contract = web3.Model.load(
                                String(File("build/Model.bin").readBytes()),
                                deployedContract!!,
                                web3j,
                                CredentialStore.credentialsList[0],
                                DefaultGasProvider.GAS_PRICE,
                                DefaultGasProvider.GAS_LIMIT
                            )
                            val text = contract.ciphertext.send()
                            println(text)
                            val state = State.fromJson(text)
                            if (state != currentState) {
                                withContext(Dispatchers.Main) {
                                    EventBus.emitEvent(
                                        Event(EventType.CURRENT_STATE_CHANGED, state.toJson())
                                    )
                                }
                            }
                            Thread.sleep(3000)
                        } catch (e: Exception) {
                            e.printStackTrace()
                            break
                        }
                    }
                }
            }
        }
    }


    private fun loadingDialog(text: String): Stage {
        val stage = Stage()
        stage.initModality(Modality.APPLICATION_MODAL)
        val label = Label(text)
        val layout = VBox(label, ProgressIndicator())
        layout.alignment = Pos.CENTER
        val scene = Scene(layout, 200.0, 100.0)
        stage.onCloseRequest = EventHandler { event ->
            event.consume()
        }
        stage.width = 500.0
        stage.height = 200.0
        stage.isResizable = true
        stage.setTitle(text)
        stage.setScene(scene)
        stage.show()
        return stage
    }

}
