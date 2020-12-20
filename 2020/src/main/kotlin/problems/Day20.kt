package main.kotlin.problems

import main.kotlin.common.*
import kotlin.math.sqrt

class Day20 : ISolution {
    override fun part1(): String {
        val tiles = parseTiles(readFileAsStrings("20.in"))

        val puzzleLength = sqrt(tiles.size.toDouble()).toInt()
        val tileMatrix = Array<Array<Tile?>>(puzzleLength) { Array(puzzleLength) { null } }
        val tilesInUse = HashSet<Long>()

        checkTile(tiles, tilesInUse, tileMatrix, 0, 0, puzzleLength)

        var res: Long = 1
        res *= tileMatrix[0][0]!!.id
        res *= tileMatrix[0][puzzleLength - 1]!!.id
        res *= tileMatrix[puzzleLength - 1][0]!!.id
        res *= tileMatrix[puzzleLength - 1][puzzleLength - 1]!!.id

        return "Day 20, Part 1 - $res"
    }

    override fun part2(): String {
        val tiles = parseTiles(readFileAsStrings("20.in"))

        val puzzleLength = sqrt(tiles.size.toDouble()).toInt()
        val tileMatrix = Array<Array<Tile?>>(puzzleLength) { Array(puzzleLength) { null } }
        val tilesInUse = HashSet<Long>()

        checkTile(tiles, tilesInUse, tileMatrix, 0, 0, puzzleLength)

        val imageWidth = puzzleLength * (tileMatrix[0][0]!!.grid.size - 2)
        var image = ArrayList<ArrayList<Char>>()
        repeat(imageWidth) {
            image.add(ArrayList())
        }

        tileMatrix.forEachIndexed { ri, row ->
            row.forEachIndexed { _, tile ->
                val grid = tile!!.versions[tile.matchingVersionIndex]
                val len = grid.size - 2

                for (i in 1..len) {
                    for (j in 1..len) {
                        val pixel = when (grid[i][j]) {
                            true -> '#'
                            false -> ' '
                        }
                        image[ri * len + i - 1].add(pixel)
                    }
                }
            }
        }
        val monster = arrayOf(
                Coord(0, 18),
                Coord(1, 0), Coord(1, 5), Coord(1, 6), Coord(1, 11), Coord(1, 12), Coord(1, 17), Coord(1, 18), Coord(1, 19),
                Coord(2, 1), Coord(2, 4), Coord(2, 7), Coord(2, 10), Coord(2, 13), Coord(2, 16)
        )

        var monsterCount = 0
        var res = 0
        var rotCount = 0

        while (monsterCount == 0) {
            monsterCount = countAndMarkMonsters(image, monster)

            if (monsterCount == 0) {
                if (rotCount < 4) {
                    image = Tile.rotate90(image) as ArrayList<ArrayList<Char>>
                    rotCount++
                } else {
                    rotCount = 0
                    image = Tile.flip(image) as ArrayList<ArrayList<Char>>
                }
            } else {
                res = image.map { it.count { a -> a == '#' } }.fold(0) { acc, n -> acc + n }

            }
        }

        return "Day 20, Part 2 - $res"
    }

    data class Coord(val x: Int, val y: Int)

    class Tile constructor(var id: Long, var grid: List<List<Boolean>>) {
        var versions: List<List<List<Boolean>>> = generateVersions()
        var matchingVersionIndex = 0
        var versionIndex: Int? = 0

        private fun generateVersions(): List<List<List<Boolean>>> {
            val res = ArrayList<List<List<Boolean>>>()

            var cur = grid
            repeat(4) {
                res.add(cur)
                cur = rotate90(cur)
            }

            cur = flip(cur)
            repeat(4) {
                res.add(cur)
                cur = rotate90(cur)
            }

            return res
        }

        companion object {
            @JvmName("rotate90bool")
            fun rotate90(g: List<List<Boolean>>): List<List<Boolean>> {
                return g.flatMap { it.withIndex() }
                        .groupBy({ (i, _) -> i }, { (_, v) -> v })
                        .map { (_, v) -> v.reversed() }
            }

            @JvmName("flipbool")
            fun flip(g: List<List<Boolean>>): List<List<Boolean>> {
                return g.reversed()
            }

            fun rotate90(g: List<List<Char>>): List<List<Char>> {
                return g.flatMap { it.withIndex() }
                        .groupBy({ (i, _) -> i }, { (_, v) -> v })
                        .map { (_, v) -> v.reversed() }
            }

            fun flip(g: List<List<Char>>): List<List<Char>> {
                return g.reversed()
            }
        }
    }

    private fun parseTiles(lines: ArrayList<String>): List<Tile> {
        val res = ArrayList<Tile>()

        var i = 0
        var curGrid = ArrayList<List<Boolean>>()
        var curId: Long = 0
        while (i in lines.indices) {
            when {
                lines[i].isEmpty() -> {
                    res.add(Tile(curId, curGrid))
                }
                lines[i][0] == 'T' -> {
                    curId = Regex("^Tile (\\d+):$").matchEntire(lines[i])!!.groupValues[1].toLong()
                    curGrid = ArrayList()
                }
                else -> curGrid.add(lines[i].map { it == '#' })
            }

            i++
        }

        res.add(Tile(curId, curGrid))

        return res
    }

    private fun checkLeftFit(leftTile: Tile, rightTile: Tile): Boolean {
        val leftVersion = leftTile.versions[leftTile.versionIndex!!]
        val rightVersion = rightTile.versions[rightTile.versionIndex!!]
        val len = leftVersion[0].size

        for (i in leftVersion.indices) {
            if (leftVersion[i][len - 1] != rightVersion[i][0]) return false
        }
        return true
    }

    private fun checkUpperFit(topTile: Tile, bottomTile: Tile): Boolean {
        val topVersion = topTile.versions[topTile.versionIndex!!]
        val bottomVersion = bottomTile.versions[bottomTile.versionIndex!!]
        val len = topVersion[0].size

        for (i in topVersion.indices) {
            if (topVersion[len - 1][i] != bottomVersion[0][i]) return false
        }
        return true
    }

    private fun checkTile(tiles: List<Tile>, tilesInUse: HashSet<Long>, tileMatrix: Array<Array<Tile?>>, row: Int, col: Int, len: Int): Boolean {
        if (row >= len || col >= len) return true

        for (tile in tiles) {
            if (tile.id in tilesInUse) continue

            tilesInUse.add(tile.id)
            tileMatrix[row][col] = tile

            tile.versions.forEachIndexed { i, _ ->
                tile.versionIndex = i

                var fitsLeft = true
                var fitsUp = true
                if (col > 0) {
                    fitsLeft = checkLeftFit(tileMatrix[row][col - 1]!!, tile)
                }
                if (row > 0) {
                    fitsUp = checkUpperFit(tileMatrix[row - 1][col]!!, tile)
                }

                if (fitsLeft && fitsUp) {
                    val nextCol = (col + 1) % len
                    val nextRow = (len * row + col + 1) / len

//                    println("now checking for ($nextRow, $nextCol)")
                    if (checkTile(tiles, tilesInUse, tileMatrix, nextRow, nextCol, len)) {
                        tile.matchingVersionIndex = i
                        return true
                    }
                }

                tile.versionIndex = null
            }
            tileMatrix[row][col] = null
            tilesInUse.remove(tile.id)
        }

        return false
    }

    private fun countAndMarkMonsters(image: ArrayList<ArrayList<Char>>, monster: Array<Coord>): Int {
        var count = 0
        for (x in image.indices) {
            for (y in image[x].indices) {
                var found = true
                for (coords in monster) {
                    val dx = x + coords.x
                    val dy = y + coords.y

                    if (dx >= image.size || dy >= image[x].size) {
                        found = false
                        break
                    }
                    if (image[dx][dy] != '#') {
                        found = false
                        break
                    }
                }

                if (found) {
                    count++
                    for (coords in monster) {
                        val dx = x + coords.x
                        val dy = y + coords.y
                        image[dx][dy] = 'O'
                    }
                }
            }
        }
        return count
    }
}