plugins {
    kotlin("jvm") version "1.6.10"
}

repositories {
    mavenCentral()
}

subprojects {

    apply {
        plugin("kotlin")
    }


    repositories {
        mavenCentral()
    }

    dependencies {
/*        implementation(kotlin("stdlib"))
        testImplementation(kotlin("test"))
        implementation(kotlin("script-runtime"))*/
    }
}