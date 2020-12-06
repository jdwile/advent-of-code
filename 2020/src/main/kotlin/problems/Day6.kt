package problems

import common.*

class Day6: ISolution {
    override fun part1(): String {
        val lines = readFileAsStrings("${getPath()}/problems/input/6.in")

        val res = lines
                .let(::parseGroups)
                .map{ it.answers.entries.size }
                .fold(0) { acc, it -> acc + it }

        return "Day 6, Part 1 - $res"
    }

    override fun part2(): String {
        val lines = readFileAsStrings("${getPath()}/problems/input/6.in")

        val res = lines
                .let(::parseGroups)
                .map { group -> group.answers.count{ it.value == group.size }}
                .fold(0) { acc, it -> acc + it }

        return "Day 6, Part 2 - $res"
    }

    private fun parseGroups(entries: List<String>): List<Group> {
        val groups = ArrayList<Group>()
        var group = HashMap<String, Int>()
        var groupSize = 0

        entries.forEachIndexed { i, entry ->
            if (entry.isNotEmpty()) {
                groupSize++
                entry.toCharArray()
                        .map{ it.toString() }
                        .groupingBy { it }.eachCount()
                        .forEach { group.merge(it.key, it.value, Math::addExact) }
            }
            if (entry.isEmpty() || i == entries.lastIndex) {
                groups.add(Group(group, groupSize))
                groupSize = 0
                group = HashMap()
            }
        }

        return groups
    }

    private data class Group(val answers: Map<String, Int>, val size: Int)
}