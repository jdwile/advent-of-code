package problems.example

import common.*

class Example : ISolution {
    override fun part1(): String {
        val lines = readFileAsInts("${getPath()}/problems/example/ex.in")

        var res = 0
        lines.forEach { res += (it / 3) - 2 }

        return "Part 1: $res"
    }

    override fun part2(): String {
        val lines = readFileAsInts("${getPath()}/problems/example/ex.in")

        var res = 0
        lines.forEach {
            var cur = (it / 3) - 2
            do {
                res += cur
                cur = (cur / 3) - 2
            } while (cur > 0)
        }

        return "Part 2: $res"
    }
}