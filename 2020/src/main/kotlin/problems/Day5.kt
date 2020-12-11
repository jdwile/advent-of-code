package problems

import common.*

class Day5 : ISolution {
    override fun part1(): String {
        val barcodes = readFileAsStrings("${getPath()}/problems/input/5.in")

        val max = barcodes.map(::getSeatId).maxOrNull()!!

        return "Day 5, Part 1: $max"
    }

    override fun part2(): String {
        val barcodes = readFileAsStrings("${getPath()}/problems/input/5.in")

        val myId = barcodes
                .map(::getSeatId)
                .sorted()
                .zipWithNext()
                .first { it.second - it.first == 2 }
                .first + 1

        return "Day 5, Part 2: $myId"
    }

    private fun getSeatId(barcode: String): Int {
        return barcode
                .replace(Regex("[BR]"), "1")
                .replace(Regex("[FL]"), "0")
                .toInt(2)
    }
}