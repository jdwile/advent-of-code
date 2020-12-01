package problems.day1

import common.ISolution
import common.getPath
import common.readFileAsInts

class Day1: ISolution {
    override fun part1() {
        val lines = readFileAsInts("${getPath()}/problems/day1/1.in")
        var flag = false

        for (i in lines)
        {
            for (j in lines)
            {
                if (i + j == 2020)
                {
                    println("Part 1: " + (i * j))
                    flag = true
                    break
                }
            }
            if (flag) break
        }
    }

    override fun part2() {
        val lines = readFileAsInts("${getPath()}/problems/day1/1.in")
        var flag = false

        for (i in lines)
        {
            for (j in lines)
            {
                for (k in lines)
                {
                    if (i + j + k == 2020)
                    {
                        println("Part 2: " + (i * j * k))
                        flag = true
                        break
                    }
                }
                if (flag) break
            }
            if (flag) break
        }
    }
}