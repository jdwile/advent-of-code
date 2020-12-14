package main.kotlin.problems

import main.kotlin.common.ISolution
import main.kotlin.common.readFileAsLongs

class Day9 : ISolution {
    private var target: Long = 0

    override fun part1(): String {
        val nums = readFileAsLongs("9.in")

        for (i in 25 until nums.size) {
            val prevList = nums.subList(i - 25, i)
            val cur = nums[i]

            if (prevList.find { (cur - it) in prevList && it != cur } == null) {
                target = cur
                break
            }
        }

        return "Day 9, Part 1 - $target"
    }

    override fun part2(): String {
        val nums = readFileAsLongs("9.in")

        var low = 0
        var high = 0
        var total: Long = 0

        while (total < target) {
            while (total <= target) {
                total += nums[high]
                high++
            }

            while (total > target) {
                total -= nums[low]
                low++
            }
        }

        val rng = nums.subList(low, high)
        val res = rng.maxOrNull()!! + rng.minOrNull()!!

        return "Day 9, Part 2 - $res"
    }
}