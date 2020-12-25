package main.kotlin.problems

import main.kotlin.common.*

class Day25 : ISolution {
    private val divider = 20201227

    override fun part1(): String {
        val (cardPublicKey, doorPublicKey) = readFileAsLongs("25.in").zipWithNext()[0]

        val subjectNum = 7
        val cardLoopSize: Int

        var cardLoop = 1
        var cardSubjectNumber: Long = 1
        while (true) {
            cardSubjectNumber = (cardSubjectNumber * subjectNum) % divider

            if (cardSubjectNumber == cardPublicKey) {
                cardLoopSize = cardLoop
                break
            }

            cardLoop++
        }

        var encryptionKey: Long = 1
        for (i in 1..cardLoopSize) {
            encryptionKey = (encryptionKey * doorPublicKey) % divider
        }

        return "Day 25, Part 1 - $encryptionKey"
    }

    override fun part2(): String {
        return "Merry Christmas :)"
    }

}