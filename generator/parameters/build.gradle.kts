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
    implementation(project(":zokrates-wrapper"))
    implementation(kotlin("script-runtime"))
}