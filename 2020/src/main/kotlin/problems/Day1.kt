package problems

import common.*

class Day1 : ISolution {
    override fun part1(): String {
        val lines = readFileAsInts("${getPath()}/problems/input/1.in")

        return "Day 1, Part 1: " + findPairSum(lines, 2020)
    }

    override fun part2(): String {
        val lines = readFileAsInts("${getPath()}/problems/input/1.in")

        for (i in 0 until lines.size - 2) {
            val res = findPairSum(lines, 2020 - lines[i])

            if (res > -1) {
                return "Day 1, Part 2: " + (res * lines[i])
            }
        }

        return "uh oh"
    }

    fun findPairSum(nums: ArrayList<Int>, goal: Int): Int {
        for (i in 0 until nums.size - 1) {
            for (j in i until nums.size) {
                val iv = nums[i]
                val jv = nums[j]

                if (iv + jv == goal) {
                    return iv * jv
                }
            }
        }
        return -1
    }
}