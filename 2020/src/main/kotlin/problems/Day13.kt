package problems

import common.*

class Day13 : ISolution {
    override fun part1(): String {
        val input = readFileAsStrings("${getPath()}/problems/input/13.in")
        val earlyTime = input[0].toInt()
        val busLines = input[1].split(",")
                .filter { it != "x" }
                .map { it.toInt() }

        var time = earlyTime
        var res = 0
        var loop = true

        while (loop) {
            for (i in busLines.indices) {
                if (time % busLines[i] == 0) {
                    res = busLines[i] * (time - earlyTime)
                    loop = false
                    break
                }
            }
            time++
        }

        return "Day 13, Part 1 - $res"
    }

    override fun part2(): String {
        val input = readFileAsStrings("${getPath()}/problems/input/13.in")
        val busLines = input[1].split(",")
        val buses = busLines.zip(busLines.indices)
                .filter { it.first != "x" }
                .map { Pair(it.first.toLong(), it.second) }

        var time = buses[0].first
        var step: Long = 1

        buses.forEach { (bus, i) ->
            time = generateSequence(time) { it + step }
                    .takeWhileInclusive { (it + i) % bus != 0.toLong() }
                    .last()
            step *= bus
        }

        return "Day 13, Part 2 - $time"
    }
}