package common

import java.io.File

public fun getPath(): String { return System.getProperty("user.dir") + "/src/main/kotlin" }

public fun readFile(fileName: String): ArrayList<String> {
    var res = ArrayList<String>()

    File(fileName).forEachLine { res.add(it) }

    return res
}

public fun readFileAsInts(fileName: String): ArrayList<Int> {
    var res = ArrayList<Int>()

    readFile(fileName).forEach { res.add(it.toInt()) }

    return res
}
