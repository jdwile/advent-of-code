package main.kotlin.problems

import main.kotlin.common.*

class Day16 : ISolution {
    override fun part1(): String {
        val (rules, _, otherTickets) = parseTicketInfo(readFileAsStrings("16.in"))
        var errors = 0

        otherTickets.map { ticket ->
            val invalid = ticket.map { num ->
                if (rules.any { it.validate(num) }) {
                    0
                } else {
                    num
                }
            }.fold(0) { acc, n -> acc + n }

            errors += invalid
        }

        return "Day 16, Part 1 - $errors"
    }

    override fun part2(): String {
        val (rules, myTicket, otherTickets) = parseTicketInfo(readFileAsStrings("16.in"))

        val validTickets = otherTickets.filter { ticket ->
            ticket.map { num ->
                rules.any { it.validate(num) }
            }.fold(true) { acc, b -> acc && b }
        } as ArrayList<List<Int>>
        validTickets.add(myTicket)

        val columnToRuleMappings = IntArray(myTicket.size) { -1 }

        var loop = true
        while (loop) {
            loop = false

            for (c in columnToRuleMappings.indices) {
                val validRules = ArrayList<Int>()
                for (r in rules.indices) {
                    if (r in columnToRuleMappings) continue

                    var ruleIsValid = true
                    for (t in validTickets.indices) {
                        if (!rules[r].validate(validTickets[t][c])) {
                            ruleIsValid = false
                            break
                        }
                    }

                    if (ruleIsValid) validRules.add(r)
                }

                if (validRules.size == 1) {
                    columnToRuleMappings[c] = validRules[0]
                    loop = true
                }
            }
        }

        var res = 1.toLong()
        for (r in columnToRuleMappings.indices) {
            if (rules[columnToRuleMappings[r]].name.contains("departure")) {
                res *= myTicket[r]
            }
        }

        return "Day 16, Part 2 - $res"
    }

    private data class Rule(val name: String, val ranges: List<IntRange>) {
        fun validate(num: Int): Boolean {
            return ranges.any { it.contains(num) }
        }

        override fun toString(): String {
            return "$name: $ranges"
        }
    }

    private fun parseTicketInfo(lines: ArrayList<String>): Triple<List<Rule>, List<Int>, List<List<Int>>> {
        val rules = ArrayList<Rule>()
        var i = 0
        while (lines[i].isNotEmpty()) {
            rules.add(parseRule(lines[i]))
            i++
        }
        i += 2

        val myTicket = lines[i].split(",").map(String::toInt)
        val otherTickets = ArrayList<List<Int>>()

        i += 3
        while (i < lines.size) {
            otherTickets.add(lines[i].split(",").map(String::toInt))
            i++
        }

        return Triple(rules, myTicket, otherTickets)
    }

    private fun parseRule(line: String): Rule {
        val parsedLine = Regex("^([\\w\\s]+): (\\d+)-(\\d+) or (\\d+)-(\\d+)\$").matchEntire(line)!!.groupValues
        return Rule(
                parsedLine[1],
                listOf(IntRange(parsedLine[2].toInt(), parsedLine[3].toInt()), IntRange(parsedLine[4].toInt(), parsedLine[5].toInt()))
        )
    }
}