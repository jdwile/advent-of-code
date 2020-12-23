package main.kotlin.problems

import main.kotlin.common.*

class Day23 : ISolution {
    override fun part1(): String {
        val cupsList = readFileAsStrings("23.in")[0].map { it.toString().toInt() } as ArrayList<Int>
        val cups = Array(10) { 0 }

        cupsList.zipWithNext().forEach { (a, b) -> cups[a] = b }
        cups[cupsList[cupsList.size - 1]] = cupsList[0]

        playCrabGame(cups, cupsList[0])

        var res = ""
        var cur = cups[1]
        for (i in 0..8) {
            res += cur.toString()
            cur = cups[cur]
        }

        return "Day 23, Part 1 - $res"
    }

    override fun part2(): String {
        val cupsList = readFileAsStrings("23.in")[0].map { it.toString().toInt() } as ArrayList<Int>
        val cups = Array(1000001) { i -> i + 1 }

        cupsList.zipWithNext().forEach { (a, b) -> cups[a] = b }
        cups[cupsList[cupsList.size - 1]] = cupsList.size + 1
        cups[1000000] = cupsList[0]
        cups[0] = 0

        playCrabGame(cups, cupsList[0], 10000000)

        val one = cups[1]
        val two = cups[one].toLong()

        return "Day 23, Part 2 - ${one * two}"
    }

    private fun playCrabGame(cups: Array<Int>, first: Int, maxMoves: Int = 100) {
        var curCup = first
        for (moves in 1..maxMoves) {
            var value = curCup
            val one = cups[curCup]
            val two = cups[one]
            val three = cups[two]

            do {
                value--
                if (value == 0) value = cups.size - 1
            } while (value in listOf(one, two, three))

            cups[curCup] = cups[three]
            cups[three] = cups[value]
            cups[value] = one

            curCup = cups[curCup]
        }
    }
}