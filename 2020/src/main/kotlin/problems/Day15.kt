package main.kotlin.problems

import main.kotlin.common.*

class Day15 : ISolution {
    override fun part1(): String {
        val startingNums = readFileAsStrings("15.in")[0].split(",").map(String::toInt)

        val record = HashMap<Int, Int>()

        for (i in startingNums.indices) {
            record[startingNums[i]] = i
        }

        var last = 6
        var next = 0
        for (turn in startingNums.size until 2020) {
            last = next
            next = when {
                record[last] == null -> {
                    0
                }
                else -> {
                    turn - record[last]!!
                }
            }
            record[last] = turn
        }

        return "Day 15, Part 1 - $last"
    }

    override fun part2(): String {
        val startingNums = readFileAsStrings("15.in")[0].split(",").map(String::toInt)

        val record = HashMap<Int, Int>()

        for (i in startingNums.indices) {
            record[startingNums[i]] = i
        }

        var last = 6
        var next = 0
        for (turn in startingNums.size until 30000000) {
            last = next
            next = when {
                record[last] == null -> {
                    0
                }
                else -> {
                    turn - record[last]!!
                }
            }
            record[last] = turn
        }

        return "Day 15, Part 2 - $last"
    }
}