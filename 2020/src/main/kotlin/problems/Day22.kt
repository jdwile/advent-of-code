package main.kotlin.problems

import main.kotlin.common.*

class Day22 : ISolution {
    override fun part1(): String {
        val (deck1, deck2) = parseStartingDecks(readFileAsStrings("22.in"))

        val res = when (playGame(deck1, deck2)) {
            1 -> scoreDeck(deck1)
            2 -> scoreDeck(deck2)
            else -> throw Exception(":(")
        }

        return "Day 22, Part 1 - $res"
    }

    override fun part2(): String {
        val (player1, player2) = parseStartingDecks(readFileAsStrings("22.in"))

        val res = when (playRecursiveGame(player1, player2)) {
            1 -> scoreDeck(player1)
            2 -> scoreDeck(player2)
            else -> throw Exception(":(")
        }

        return "Day 22, Part 2 - $res"
    }

    private fun playGame(deck1: ArrayList<Int>, deck2: ArrayList<Int>): Int {
        while (deck1.isNotEmpty() && deck2.isNotEmpty()) {
            val play1 = deck1.removeFirst()
            val play2 = deck2.removeFirst()

            val whichPlayerWon = when (play1 > play2) {
                true -> 1
                false -> 2
            }

            when (whichPlayerWon) {
                1 -> deck1.addAll(listOf(play1, play2))
                2 -> deck2.addAll(listOf(play2, play1))
            }
        }

        return when (deck1.isNotEmpty()) {
            true -> 1
            false -> 2
        }
    }

    private fun playRecursiveGame(deck1: ArrayList<Int>, deck2: ArrayList<Int>): Int {
        val gameHistory = HashSet<String>()

        while (deck1.isNotEmpty() && deck2.isNotEmpty()) {
            val deckHash1 = deck1.joinToString(",")
            val deckHash2 = deck2.joinToString(",")

            if (gameHistory.contains(deckHash1) && gameHistory.contains(deckHash2)) {
                return 1
            }

            gameHistory.addAll(listOf(deckHash1, deckHash2))

            val play1 = deck1.removeFirst()
            val play2 = deck2.removeFirst()

            val whichPlayerWon = if (deck1.size >= play1 && deck2.size >= play2) {
                val deck1Subgame = ArrayList(deck1.slice(0 until play1))
                val deck2Subgame = ArrayList(deck2.slice(0 until play2))

                playRecursiveGame(deck1Subgame, deck2Subgame)
            } else {
                when (play1 > play2) {
                    true -> 1
                    false -> 2
                }
            }
            when (whichPlayerWon) {
                1 -> deck1.addAll(listOf(play1, play2))
                2 -> deck2.addAll(listOf(play2, play1))
            }
        }

        return when (deck1.isEmpty()) {
            true -> 2
            false -> 1
        }
    }

    private fun scoreDeck(deck: ArrayList<Int>): Int {
        return deck.zip(deck.indices.reversed()).fold(0) { acc, p -> acc + p.first * (p.second + 1) }
    }

    private fun parseStartingDecks(lines: ArrayList<String>): Pair<ArrayList<Int>, ArrayList<Int>> {
        var i = 1

        val deck1 = ArrayList<Int>()
        val deck2 = ArrayList<Int>()

        while (lines[i].isNotEmpty()) {
            deck1.add(lines[i].toInt())
            i++
        }
        i += 2
        while (i < lines.size) {
            deck2.add(lines[i].toInt())
            i++
        }

        return Pair(deck1, deck2)
    }
}