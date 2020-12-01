package problems.day1

import common.ISolution
import common.getPath
import common.readFileAsInts

class Day1: ISolution {
    override fun part1() {
        val lines = readFileAsInts("${getPath()}/problems/day1/1.in")

        println("Part 1: " + getSumTo(lines, 2020))
    }

    fun getSumTo(lines: ArrayList<Int>, num: Int): Int
    {
        for (i in 0 .. lines.size - 2)
        {
            for (j in i until lines.size)
            {
                val iv = lines[i]
                val jv = lines[j]

                if (iv + jv == num)
                {
                    return iv * jv
                }
            }
        }
        return -1
    }

    override fun part2() {
        val lines = readFileAsInts("${getPath()}/problems/day1/1.in")
        var flag = false

        for (i in 0 .. lines.size - 3)
        {
            val res = getSumTo(lines, 2020 - lines[i])

            if (res > -1) {
                println("Part 2: " + (res * lines[i]))
                break
            }
        }
    }
}