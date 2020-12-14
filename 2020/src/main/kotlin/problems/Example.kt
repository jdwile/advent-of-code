package problems

import main.kotlin.common.ISolution
import main.kotlin.common.getPath
import main.kotlin.common.readFileAsInts

class Example : ISolution {
    override fun part1(): String {
        val lines = readFileAsInts("ex.in")

        var res = 0
        lines.forEach { res += (it / 3) - 2 }

        return "Example, Part 1: $res"
    }

    override fun part2(): String {
        val lines = readFileAsInts("ex.in")

        var res = 0
        lines.forEach {
            var cur = (it / 3) - 2
            do {
                res += cur
                cur = (cur / 3) - 2
            } while (cur > 0)
        }

        return "Example, Part 2: $res"
    }
}