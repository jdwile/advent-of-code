package main.kotlin.problems

import main.kotlin.common.ISolution
import main.kotlin.common.readFileAsInts
import java.math.BigInteger

class Day10 : ISolution {
    override fun part1(): String {
        val adapters = readFileAsInts("10.in")
        adapters.sort()
        adapters.add(adapters[adapters.lastIndex] + 3)
        adapters.add(0, 0)

        var ones = 0
        var threes = 0
        adapters
                .zipWithNext()
                .map {
                    val diff = it.second - it.first
                    if (diff == 1) ones++
                    if (diff == 3) threes++
                }

        val res = ones * threes
        return "Day 10, Part 1 - $res"
    }

    override fun part2(): String {
        val adapters = readFileAsInts("10.in")
        adapters.sort()
        adapters.add(adapters[adapters.lastIndex] + 3)
        adapters.add(0, 0)

        val DP = HashMap<Int, BigInteger>()

        fun dp(i: Int): BigInteger {
            if (i == 0) return BigInteger.ONE
            if (DP[i] != null) return DP[i]!!

            var res = BigInteger.ZERO

            for (j in 0 until i) {
                if (adapters[i] - adapters[j] <= 3) {
                    res += dp(j)
                }
            }

            DP[i] = res
            return res
        }

        return "Day 10, Part 1 - ${dp(adapters.lastIndex)}"
    }
}