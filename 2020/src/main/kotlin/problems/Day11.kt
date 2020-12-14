package main.kotlin.problems

import main.kotlin.common.ISolution
import main.kotlin.common.readFileAsCharArray

class Day11 : ISolution {
    override fun part1(): String {
        var seats = readFileAsCharArray("11.in")
        seats = runGameOfAirplaneSeats(seats)

        val res = seats.hash().count { it == '#' }

        return "Day 11, Part 1 - $res"
    }

    override fun part2(): String {
        var seats = readFileAsCharArray("11.in")
        seats = runGameOfAirplaneSeats(seats, true)

        val res = seats.hash().count { it == '#' }

        return "Day 11, Part 2 - $res"
    }

    private fun Array<CharArray>.hash(): String {
        return this.joinToString("\n") { it.joinToString("") }
    }

    private fun runGameOfAirplaneSeats(plane: Array<CharArray>, seesExtraFar: Boolean = false): Array<CharArray> {
        var last: String
        val diffs = arrayOf(-1, 0, 1)
        var seats = plane

        do {
            last = seats.hash()
            val next = Array(seats.size) { CharArray(seats[0].size) }

            for (x in seats.indices) {
                for (y in seats[x].indices) {
                    next[x][y] = seats[x][y]

                    if (seats[x][y] == '.') continue
                    var numOccupiedNeighbors = 0

                    for (dx in diffs) {
                        if (x + dx !in seats.indices) continue
                        for (dy in diffs) {
                            if ((dx == 0 && dy == 0) || y + dy !in seats[x].indices) continue

                            var multi = 1
                            while (seesExtraFar &&
                                    x + dx * multi in seats.indices
                                    && y + dy * multi in seats[x].indices
                                    && seats[x + dx * multi][y + dy * multi] == '.') {
                                multi++
                            }
                            if (x + dx * multi in seats.indices
                                    && y + dy * multi in seats[x].indices
                                    && seats[x + dx * multi][y + dy * multi] == '#') {
                                numOccupiedNeighbors++
                            }
                        }
                    }
                    next[x][y] = calculateSeatBehavior(seats[x][y], numOccupiedNeighbors, seesExtraFar)
                }
            }
            seats = next.toList().toTypedArray()
        } while (seats.hash() != last)
        return seats
    }

    private fun calculateSeatBehavior(seat: Char, numOccupiedNeighbors: Int, tolerant: Boolean = false): Char {
        if (seat == 'L' && numOccupiedNeighbors == 0) {
            return '#'
        }

        if (seat == '#' && numOccupiedNeighbors >= 4) {
            if (tolerant && numOccupiedNeighbors == 4) {
                return '#'
            }
            return 'L'
        }

        return seat
    }
}