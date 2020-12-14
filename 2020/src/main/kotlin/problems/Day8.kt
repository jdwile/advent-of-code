package main.kotlin.problems

import main.kotlin.common.ISolution
import main.kotlin.common.readFileAsStrings

class Day8 : ISolution {
    override fun part1(): String {
        val program = readFileAsStrings("8.in")
                .map { it.split(" ") }
                .map { Instruction(it[0], it[1].toInt()) }

        val (acc, _) = runProgram(program)

        return "Day 8, Part 1 - $acc"
    }

    override fun part2(): String {
        val program = readFileAsStrings("8.in")
                .map { it.split(" ") }
                .map { Instruction(it[0], it[1].toInt()) }

        val attempts = HashSet<Int>()

        while (true) {
            var flippedIndex = 0
            while (flippedIndex < program.size && (program[flippedIndex].operation == "acc" || flippedIndex in attempts)) flippedIndex++
            attempts.add(flippedIndex)

            val testProgram = program.map(Instruction::clone)
            when (testProgram[flippedIndex].operation) {
                "nop" -> testProgram[flippedIndex].operation = "jmp"
                "jmp" -> testProgram[flippedIndex].operation = "nop"
            }

            val (acc, hasLooped) = runProgram(testProgram)

            if (!hasLooped) {
                return "Day 8, Part 2 - $acc"
            }
        }
    }

    class Instruction constructor(var operation: String, val arg: Int, var visited: Boolean = false) {
        override fun toString() = "$operation: $arg (visited=$visited)"
        fun clone(): Instruction {
            return Instruction(operation, arg, visited)
        }
    }

    private fun runProgram(program: List<Instruction>): Pair<Int, Boolean> {
        var i = 0
        var acc = 0

        var hasLooped = false

        while (true) {
            if (i == program.size) break

            if (program[i].visited) {
                hasLooped = true
                break
            }

            program[i].visited = true

            when (program[i].operation) {
                "acc" -> acc += program[i].arg
                "jmp" -> i += program[i].arg - 1
                "nop" -> {
                }
            }
            i++
        }

        return Pair(acc, hasLooped)
    }
}