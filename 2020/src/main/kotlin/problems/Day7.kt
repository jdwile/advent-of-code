package problems

import common.*
import kotlin.collections.ArrayList

class Day7 : ISolution {
    override fun part1(): String {
        val lines = readFileAsStrings("${getPath()}/problems/input/7.in")

        val bags = getBags(lines)
        val bagsThatContainGold = bags.filter { it.value.contains("shiny gold", bags) }.count()

        return "Day 7, Part 1 - $bagsThatContainGold"
    }

    override fun part2(): String {
        val lines = readFileAsStrings("${getPath()}/problems/input/7.in")

        val bags = getBags(lines)
        val totalBagsContainedInGold = bags["shiny gold"]?.countInnerBags(bags)

        return "Day 7, Part 2 - $totalBagsContainedInGold"
    }

    private data class Bag(val name: String, val innerBags: Map<String, Int>) {
        fun contains(name: String, bags: Map<String, Bag>): Boolean {
            if (this.innerBags.keys.contains(name)) return true

            return this.innerBags.keys.any { bags[it]?.contains(name, bags) == true }
        }

        fun countInnerBags(bags: Map<String, Bag>): Int {
            return this.innerBags.entries.sumBy { (bags[it.key]?.countInnerBags(bags) ?: 0) * it.value + it.value }
        }
    }

    private fun getBags(lines: List<String>): Map<String, Bag> {
        val processedLines = ArrayList<String>()

        lines.filter { !it.contains("no other") }
                .forEach {
                    processedLines.add(
                            it.replace(" contain ", "-")
                                    .replace("[.]|\\sbags|\\sbag".toRegex(), "")
                                    .replace(", ", ":")
                    )
                }

        return processedLines.map { it.split("-", ":") }
                .map { it[0] to Bag(it[0], toInnerBags(it.subList(1, it.size))) }.toMap()
    }

    private fun toInnerBags(bags: List<String>) =
            bags.map { it.substring(2) to Integer.parseInt(it.substring(0, 1)) }.toMap()

}