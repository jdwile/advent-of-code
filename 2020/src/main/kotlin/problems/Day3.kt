package problems

import common.*
import java.math.BigInteger

class Day3: ISolution {
    override fun part1(): String {
        val grid = readFileAsCharArray("${getPath()}/problems/input/3.in")

        val treeCount = getTreeCount(grid)

        return "Day 3, Part 1: $treeCount"
    }

    override fun part2(): String {
        val grid = readFileAsCharArray("${getPath()}/problems/input/3.in")

        var treeCountProduct: BigInteger = BigInteger.ONE

        treeCountProduct *= getTreeCount(grid, 1, 1).toBigInteger()
        treeCountProduct *= getTreeCount(grid, 3, 1).toBigInteger()
        treeCountProduct *= getTreeCount(grid, 5, 1).toBigInteger()
        treeCountProduct *= getTreeCount(grid, 7, 1).toBigInteger()
        treeCountProduct *= getTreeCount(grid, 1, 2).toBigInteger()

        return "Day 3, Part 2: $treeCountProduct"
    }

    private fun getTreeCount(grid: Array<CharArray>, dx: Int = 3, dy: Int = 1): Int {
        var x = 0
        var y = 0
        var treeCount = 0

        while (y < grid.size) {
            if (grid[y][x] == '#') {
                treeCount++
            }
            x = (x + dx) % grid[y].size
            y += dy
        }

        return treeCount
    }
}