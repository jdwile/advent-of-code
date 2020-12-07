package problems

import common.*
import kotlin.collections.ArrayList
import kotlin.collections.HashMap

class Day7: ISolution {
    override fun part1(): String {
        val lines = readFileAsStrings("${getPath()}/problems/input/7.in")

        val adjacencyList = getAdjacencyList(lines)
        val bagsThatContainGold =  adjacencyList.keys
                                    .filter { it != "shiny gold" }
                                    .map { if (bagContainsGold(adjacencyList, it, ArrayList(), ArrayList())) 1 else 0 }
                                    .sum()

        return "Day 7, Part 1 - $bagsThatContainGold"
    }

    override fun part2(): String {
        val lines = readFileAsStrings("${getPath()}/problems/input/7.in")

        val adjacencyList = getAdjacencyList(lines)
        val totalBagsContainedInGold = getTotalBagsContainedIn(adjacencyList, "shiny gold")
        return "Day 7, Part 2 - $totalBagsContainedInGold"
    }

    private fun getTotalBagsContainedIn(adjacencyList: Map<String, List<Bag>>, currentBag: String): Int {
        var total = 0

        adjacencyList[currentBag]?.let { subBags ->
            subBags.forEach {
                total += it.count * (1 + getTotalBagsContainedIn(adjacencyList, it.name))
            }
        }

        return total
    }

    private tailrec fun bagContainsGold(adjacencyList: Map<String, List<Bag>>, bag: String, next: ArrayList<String>, visited: ArrayList<String>): Boolean {
        if (bag == "shiny gold") {
            return true
        }

        next.remove(bag)
        visited.add(bag)

        adjacencyList[bag]?.let { cur ->
            next.addAll(cur.filter { !visited.contains(it.name) && !next.contains(it.name) }.map { it.name })
        }

        if (next.size == 0) return false

        return bagContainsGold(adjacencyList, next[0], next, visited)
    }

    private data class Bag(val count: Int, val name: String)

    private fun getAdjacencyList(lines: List<String>): Map<String, List<Bag>> {
        val res = HashMap<String, ArrayList<Bag>>()
        val processedLines = ArrayList<String>()

        lines.forEach {
            processedLines.add(
                    it.replace("contain", "-")
                            .replace("[\\.]|\\sbags|\\sbag".toRegex(), "")
                            .replace(",", " :")
            )
        }

        processedLines.forEach { line ->
            val (cur, rest) = line.split(" - ")

            if (rest == "no other") {
                res[cur] = ArrayList()
            } else {
                rest.split(" : ").forEach {
                    val bagArr = it.split(" ")
                    val count = bagArr[0].toInt()
                    val bagName = bagArr.subList(1, bagArr.lastIndex+1).joinToString(" ")

                    if (res[cur].isNullOrEmpty()) {
                        res[cur] = ArrayList()
                    }
                    res[cur]?.add(Bag(count, bagName))
                }
            }
        }

        return res
    }
}