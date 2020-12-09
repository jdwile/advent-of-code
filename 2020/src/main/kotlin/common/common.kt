package common

import java.io.File
import java.math.BigInteger

fun getPath(): String { return System.getProperty("user.dir") + "/src/main/kotlin" }

fun readFileAsStrings(fileName: String): ArrayList<String> {
    val res = ArrayList<String>()

    File(fileName).forEachLine { res.add(it) }

    return res
}

fun readFileAsInts(fileName: String): ArrayList<Int> {
    val res = ArrayList<Int>()

    readFileAsStrings(fileName).forEach { res.add(it.toInt()) }

    return res
}

fun readFileAsLongs(fileName: String): ArrayList<Long> {
    val res = ArrayList<Long>()

    readFileAsStrings(fileName).forEach { res.add(it.toLong()) }

    return res
}

fun readFileAsCharArray(fileName: String): Array<CharArray> {
    val res = ArrayList<CharArray>()

    File(fileName).forEachLine { res.add(it.toCharArray()) }

    return res.toTypedArray()
}
