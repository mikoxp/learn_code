import java.io.File
import java.nio.file.Files
import java.nio.file.Paths

fun main(args: Array<String>) {
    println("Copy files:")
    var pathIn: String  = "test_in"
    var pathOut: String  = "test_out"
    createDirectory(pathOut)
    File(pathIn).walkBottomUp().forEach {
        if(it.isFile) {
            val name=it.name.split(".")[0]
            createDirectory("$pathOut/$name")
            val outPath="$pathOut/$name/${it.name}"
            if(!File(outPath).exists()) {
                it.copyTo(File(outPath))
                println(outPath)
            }
        }
    }
}

fun createDirectory(path:String){
    Files.createDirectories(Paths.get(path));
}