package main.kotlin.common

import java.io.File

fun getPath(): String {
    return System.getProperty("user.dir") + "/src/main/kotlin/problems/input/"
}

fun readFileAsStrings(fileName: String): ArrayList<String> {
    val res = ArrayList<String>()

    File("${getPath()}$fileName").forEachLine { res.add(it) }

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

    File("${getPath()}$fileName").forEachLine { res.add(it.toCharArray()) }

    return res.toTypedArray()
}


fun <T> Sequence<T>.takeWhileInclusive(pred: (T) -> Boolean): Sequence<T> {
    var shouldContinue = true
    return takeWhile {
        val result = shouldContinue
        shouldContinue = pred(it)
        result
    }
}