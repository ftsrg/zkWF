import org.jetbrains.kotlin.gradle.tasks.KotlinCompile

plugins {
    //kotlin("jvm") version "1.6.10"
    id("com.github.johnrengelman.shadow") version "6.1.0"
    id("org.openjfx.javafxplugin") version "0.0.12"
    /*id ("org.web3j") version "4.8.8"
    id ("com.github.node-gradle.node") version "3.1.1"
    id ("org.web3j.solidity") version "0.3.2"*/
    application
    java
    kotlin("jvm")
}


group = "eu.toldi"
version = "1.0-SNAPSHOT"

repositories {
    mavenCentral()
}

application {
    mainClass.set("eu.toldi.bpmn_zkp.GUI.kt")
    project.setProperty("mainClassName", "eu.toldi.bpmn_zkp.GUIKt")
}

dependencies {
    implementation(kotlin("stdlib"))
    implementation(project(":model"))
    implementation(project(":zokrates-wrapper"))
    val coroutines_version = "1.6.4"
    implementation("org.jetbrains.kotlinx:kotlinx-coroutines-core:$coroutines_version")
    implementation("org.jetbrains.kotlinx:kotlinx-coroutines-javafx:$coroutines_version")
    implementation("com.beust:klaxon:5.6")
    implementation("no.tornado:tornadofx:1.7.20")
    implementation("org.web3j:core:4.9.1")
}

javafx {
    modules("javafx.controls", "javafx.fxml", "javafx.web")
    version = "11"
}

tasks {
    shadowJar {

        archiveBaseName.set("bpmn_zkp")
        archiveClassifier.set("")
        archiveVersion.set("")
    }
}

tasks.withType<KotlinCompile>() {
    kotlinOptions.jvmTarget = "11"
}


