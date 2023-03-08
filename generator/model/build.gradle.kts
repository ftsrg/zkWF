plugins {
    kotlin("jvm")
}

group = "eu.toldi"
version = "1.0-SNAPSHOT"

repositories {
    mavenCentral()
}

dependencies {
    implementation(kotlin("stdlib"))
    implementation("com.beust:klaxon:5.6")
    implementation("org.web3j:core:4.9.1")
}