import java.io.File

fun main() {
    var string: String = "Hello, World!"  // defining a variable
    println("$string")
    val data = readFileLineByLineUsingForEachLine("country.csv")
    println(data)
    writeLine("data"+System.currentTimeMillis()+".sql",data)
}

fun readFileLineByLineUsingForEachLine(fileName: String) :MutableList<String>{
    var data:MutableList<String> = ArrayList()
    File(fileName).forEachLine { data. add(it)}
    return data
}

fun writeLine(path: String,data:MutableList<String>) {
    val f = File(path!!)
    data.forEach({
        val split = it.split(";")
        if(split.size>2){
            val s = "(%s,%s,%s,%s)\n".format(
                split[1].replace(" ","").replace(",","."),
                split[2].replace(" ","").replace(",","."),
                split[3].replace(" ","").replace(",","."),
                        split[4].replace(" ","").replace(",",".")
            )
            println(s)
            f.appendText(s)
        }
    })
}
