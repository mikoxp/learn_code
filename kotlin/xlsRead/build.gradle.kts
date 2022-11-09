import org.jetbrains.kotlin.gradle.tasks.KotlinCompile

plugins {
    kotlin("jvm") version "1.5.10"
}

group = "me.pcoles"
version = "1.0-SNAPSHOT"

repositories {
    mavenCentral()
}
dependencies {
    // https://mvnrepository.com/artifact/org.apache.poi/poi
    implementation("org.apache.poi:poi:5.2.3")
// https://mvnrepository.com/artifact/org.apache.poi/poi-ooxml
    implementation("org.apache.poi:poi-ooxml:5.2.3")

}


tasks.withType<KotlinCompile>() {
    kotlinOptions.jvmTarget = "1.8"
}