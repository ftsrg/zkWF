plugins {
    id("com.github.johnrengelman.shadow") version "6.1.0"
    application
    java
    kotlin("jvm") version "1.0.0"
}

group = "eu.toldi"
version = "1.0-SNAPSHOT"

application {
    mainClass.set("eu.toldi.bpmn_zkp.CLIMain.kt")
    project.setProperty("mainClassName", "eu.toldi.bpmn_zkp.CLIMainKt")
}

repositories {
    mavenCentral()
    jcenter()
}

dependencies {
    implementation("com.andreapivetta.kolor:kolor:1.0.0")
    implementation(kotlin("stdlib"))
    implementation(project(":model"))
    implementation(project(":zokrates-wrapper"))
    val coroutines_version = "1.6.4"
    implementation("org.jetbrains.kotlinx:kotlinx-coroutines-core:$coroutines_version")
    implementation("org.jetbrains.kotlinx:kotlinx-coroutines-javafx:$coroutines_version")
    implementation("com.beust:klaxon:5.6")
    implementation("org.web3j:core:4.9.1")
    testImplementation("org.junit.jupiter:junit-jupiter-api:5.9.0")
    testRuntimeOnly("org.junit.jupiter:junit-jupiter-engine:5.9.0")
}

tasks.getByName<Test>("test") {
    useJUnitPlatform()
}