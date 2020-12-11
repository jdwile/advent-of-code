package problems

import common.*

class Day2 : ISolution {
    override fun part1(): String {
        val lines = readFileAsStrings("${getPath()}/problems/input/2.in")
        var total = 0

        lines.forEach {
            val entries = it.split("([ \\-:])+".toRegex())
            val (low, high, letter, password) = entries.toCollection(ArrayList())

            if (isValidPassword1(low.toInt(), high.toInt(), letter[0], password)) {
                total++
            }
        }

        return "Day 2, Part 1: $total"
    }

    override fun part2(): String {
        val lines = readFileAsStrings("${getPath()}/problems/input/2.in")
        var total = 0

        lines.forEach {
            val entries = it.split("([ \\-:])+".toRegex())
            val (low, high, letter, password) = entries.toCollection(ArrayList())

            if (isValidPassword2(low.toInt(), high.toInt(), letter[0], password)) {
                total++
            }
        }

        return "Day 2, Part 2: $total"
    }

    private fun isValidPassword1(low: Int, high: Int, letter: Char, password: String): Boolean {
        return password.count { it == letter } in low..high
    }

    private fun isValidPassword2(low: Int, high: Int, letter: Char, password: String): Boolean {
        return (password[low - 1] == letter).xor(password[high - 1] == letter)
    }
}