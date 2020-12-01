package problems.day1

import common.ISolution
import common.getPath
import common.readFileAsInts

class Day1: ISolution {
    override fun part1() {
        val lines = readFileAsInts("${getPath()}/problems/day1/1.in")
        var flag = false

        for (i in 0 .. lines.size - 2)
        {
            for (j in i until lines.size)
            {
                val iv = lines[i]
                val jv = lines[j]

                if (iv + jv == 2020)
                {
                    println("Part 1: " + (iv * jv))
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

        for (i in 0 .. lines.size - 3)
        {
            for (j in i .. lines.size - 2)
            {
                for (k in j until lines.size)
                {
                    val iv = lines[i]
                    val jv = lines[j]
                    val kv = lines[k]

                    if (iv + jv + kv == 2020)
                    {
                        println("Part 2: " + (iv * jv * kv))
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