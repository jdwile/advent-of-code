package problems

import common.*

class Day6: ISolution {
    override fun part1(): String {
        val lines = readFileAsStrings("${getPath()}/problems/input/6.in")

        val res = lines
                .let(::parseGroups)
                .map { it.reduce(Set<Char>::union).size }
                .sum()

        return "Day 6, Part 1 - $res"
    }

    override fun part2(): String {
        val lines = readFileAsStrings("${getPath()}/problems/input/6.in")

        val res = lines
                .let(::parseGroups)
                .map { it.reduce(Set<Char>::intersect).size }
                .sum()

        return "Day 6, Part 2 - $res"
    }

    private fun parseGroups(entries: List<String>): List<List<Set<Char>>> {
        val groups = ArrayList<ArrayList<Set<Char>>>()
        var group = ArrayList<Set<Char>>()

        entries.forEachIndexed { i, entry ->
            if (entry.isNotEmpty()) {
                group.add(entry.toSet())
            }
            if (entry.isEmpty() || i == entries.lastIndex) {
                groups.add(group)
                group = ArrayList()
            }
        }

        return groups
    }
}