package problems

import common.*
import kotlin.math.abs

class Day12 : ISolution {
    override fun part1(): String {
        val instructions = readFileAsStrings("${getPath()}/problems/input/12.in")

        var x = 0
        var y = 0

        val dirs = arrayOf(Pair(1, 0), Pair(0, -1), Pair(-1, 0), Pair(0, 1))
        var dirIndex = 0

        instructions.map {
            val command = it[0]
            val len = it.substring(1).toInt()

            when (command) {
                'N' -> y += len
                'S' -> y -= len
                'E' -> x += len
                'W' -> x -= len
                'F' -> {
                    x += dirs[dirIndex].first * len
                    y += dirs[dirIndex].second * len
                }
                'R' -> dirIndex += len / 90
                'L' -> dirIndex -= len / 90
            }
            if (dirIndex < 0) dirIndex += 4
            dirIndex %= 4
        }

        val res = abs(x) + abs(y)
        return "Day 12, Part 1 - $res"
    }

    override fun part2(): String {
        val instructions = readFileAsStrings("${getPath()}/problems/input/12.in")

        var x = 0
        var y = 0

        var waypointX = 10
        var waypointY = 1

        instructions.map {
            val command = it[0]
            val len = it.substring(1).toInt()

            when (command) {
                'N' -> waypointY += len
                'S' -> waypointY -= len
                'E' -> waypointX += len
                'W' -> waypointX -= len
                'F' -> {
                    x += waypointX * len
                    y += waypointY * len
                }
                'R' -> {
                    for (i in 0 until (len / 90)) {
                        val tmp = -waypointX
                        waypointX = waypointY
                        waypointY = tmp
                    }
                }
                'L' -> {
                    for (i in 0 until (len / 90)) {
                        val tmp = -waypointY
                        waypointY = waypointX
                        waypointX = tmp
                    }
                }
            }
        }

        val res = abs(x) + abs(y)
        return "Day 12, Part 2 - $res"
    }
}