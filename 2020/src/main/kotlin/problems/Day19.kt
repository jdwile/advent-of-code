package main.kotlin.problems

import main.kotlin.common.*
import java.util.*
import kotlin.collections.ArrayList
import kotlin.collections.HashMap

class Day19 : ISolution {
    private val digitRegex = Regex("\\d+")

    override fun part1(): String {
        val (rawRules, messages) = parseInput(readFileAsStrings("19.in"))

        val rules = HashMap<Int, String>()

        rawRules.forEach {
            val groups = it.split(": ")
            rules[groups[0].toInt()] = groups[1]
        }

        var s = rules[0]!!
        while (digitRegex.containsMatchIn(s)) {
            s = digitRegex.replace(s) { m -> "(${rules[m.value.toInt()]})" }
        }

        val regex = Regex("^" + s.replace(" ", "").replace("\"", "") + "\$")

        val res = messages.filter { regex.matches(it) }.size

        return "Day 19, Part 1 - $res"
    }

    override fun part2(): String {
        val (rawRules, messages) = parseInput(readFileAsStrings("19.in"))

        val rules = rawRules.associate { s ->
            s.split(": ").let { (a, b) ->
                a.toInt() to (b.removeSurrounding("\"").split(" | ").map { it.split(' ') })
            }
        }

        val res = messages.count { it.checkWith(rules).contains("") }

        return "Day 19, Part 2 - $res"
    }

    private fun String.checkWith(rules: Map<Int, List<List<String>>>, cur: Int = 0): Set<String> {
        val res = mutableSetOf<String>()

        when (cur) {
            8 -> {
                val next = 42
                var rest = setOf(this)

                while (rest.isNotEmpty()) {
                    rest = rest.flatMap { it.checkWith(rules, next) }.toSet()
                    res += rest
                }
            }

            11 -> {
                val first = 42
                val second = 31
                var rest = setOf(this)
                rest = rest.flatMap { it.checkWith(rules, first) }.toSet()

                val stack: Stack<Set<String>> = Stack()
                while (rest.isNotEmpty()) {
                    stack.push(rest)
                    rest = rest.flatMap { it.checkWith(rules, first) }.toSet()
                }

                while (stack.isNotEmpty()) {
                    rest = stack.pop()
                    repeat(stack.size + 1) { _ ->
                        rest = rest.flatMap { it.checkWith(rules, second) }.toSet()
                    }
                    res += rest
                }
            }

            else -> {
                val options = rules[cur]!!
                for (option in options) {
                    var rest = setOf(this)
                    for (rule in option) {
                        if (rule.toIntOrNull() != null) {
                            rest = rest.flatMap { it.checkWith(rules, rule.toInt()) }.toSet()
                        } else {
                            if (this.isNotEmpty() && this.first() == rule.first()) {
                                return setOf(this.drop(1))
                            } else {
                                return setOf()
                            }
                        }
                    }

                    res += rest
                }
            }
        }

        return res
    }

    private fun parseInput(lines: ArrayList<String>): Pair<ArrayList<String>, ArrayList<String>> {
        val rules = ArrayList<String>()
        val messages = ArrayList<String>()

        var i = 0
        while (lines[i].isNotEmpty()) {
            rules.add(lines[i])
            i++
        }
        i++
        while (i < lines.size) {
            messages.add(lines[i])
            i++
        }

        return Pair(rules, messages)
    }
}