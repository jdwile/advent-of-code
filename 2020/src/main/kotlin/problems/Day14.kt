package main.kotlin.problems

import main.kotlin.common.ISolution
import main.kotlin.common.readFileAsStrings
import java.lang.Long.toBinaryString

class Day14 : ISolution {
    override fun part1(): String {
        val lines = readFileAsStrings("14.in")
        val memoryRegex = Regex("^mem\\[([0-9]*)\\] = ([0-9]*)\$")
        val maskRegex = Regex("^mask = (.*)\$")
        var bitmask = ""
        val memory = HashMap<Long, Long>()

        lines.forEach { line ->
            if (maskRegex.matches(line)) {
                bitmask = maskRegex.matchEntire(line)!!.groupValues[1]
            } else {
                val (memLoc, memVal) = memoryRegex.matchEntire(line)!!.groupValues.subList(1, 3).map { it.toLong() }

                memory[memLoc] = memVal.overwriteWithBinaryMask(bitmask)
            }
        }

        val res = memory.values.sum()

        return "Day 14, Part 1 - $res"
    }

    override fun part2(): String {
        val lines = readFileAsStrings("14.in")
        val memoryRegex = Regex("^mem\\[([0-9]*)\\] = ([0-9]*)\$")
        val maskRegex = Regex("^mask = (.*)\$")
        var bitmask = ""
        val memory = HashMap<Long, Long>()

        lines.forEach { line ->
            if (maskRegex.matches(line)) {
                bitmask = maskRegex.matchEntire(line)!!.groupValues[1]
            } else {
                val (memLoc, memVal) = memoryRegex.matchEntire(line)!!.groupValues.subList(1, 3).map { it.toLong() }

                val locations = arrayListOf(memLoc.floatingBinaryMask(bitmask))

                while (locations[0].contains('X')) {
                    val cur = locations[0]
                    locations.removeAt(0)
                    locations.add(cur.replaceFirst('X', '0'))
                    locations.add(cur.replaceFirst('X', '1'))
                }

                locations.map { it.toLong(2) }
                        .map {
                            memory[it] = memVal
                        }
            }
        }

        val res = memory.values.sum()

        return "Day 14, Part 2 - $res"
    }

    private fun Long.to36bitString(): String =
            toBinaryString(this).padStart(36, '0')

    private fun Long.overwriteWithBinaryMask(mask: String): Long =
            this.to36bitString().mapIndexed { i, c ->
                when (mask[i]) {
                    '0' -> '0'
                    '1' -> '1'
                    else -> c
                }
            }.joinToString("")
                    .toLong(2)

    private fun Long.floatingBinaryMask(mask: String): String =
            this.to36bitString().mapIndexed { i, c ->
                when (mask[i]) {
                    '0' -> c
                    '1' -> '1'
                    else -> 'X'
                }
            }.joinToString("")
}