package main.kotlin.problems

import main.kotlin.common.*

class Day24 : ISolution {
    override fun part1(): String {
        val instructions = parseInstructions(readFileAsStrings("24.in"))

        val tiles = HashMap<Coord, Boolean>()
        tiles[Coord(0, 0, 0)] = true

        for (instruction in instructions) {
            var x = 0
            var y = 0
            var z = 0

            for (direction in instruction) {
                val d = direction.getDir()
                x += d.x
                y += d.y
                z += d.z

                if (tiles[Coord(x, y, z)] == null) tiles[Coord(x, y, z)] = true
            }

            val cur = Coord(x, y, z)
            tiles[cur] = when (tiles[cur]) {
                null -> true
                else -> !tiles[cur]!!
            }
        }

        val res = tiles.values.count { !it }

        return "Day 24, Part 1 - $res"
    }

    override fun part2(): String {
        val instructions = parseInstructions(readFileAsStrings("24.in"))

        var tiles = HashMap<Coord, Boolean>()
        tiles[Coord(0, 0, 0)] = true

        for (instruction in instructions) {
            var x = 0
            var y = 0
            var z = 0

            for (direction in instruction) {
                val d = direction.getDir()
                x += d.x
                y += d.y
                z += d.z

                if (tiles[Coord(x, y, z)] == null) tiles[Coord(x, y, z)] = true
            }

            val cur = Coord(x, y, z)
            tiles[cur] = when (tiles[cur]) {
                null -> true
                else -> !tiles[cur]!!
            }
        }

        for (i in 1..100) {
            val checkForNext = HashSet<Coord>()
            for (coord in tiles.keys) {
                checkForNext.add(coord)
                for (dir in Direction.values().map { it.getDir() }) {
                    val cur = Coord(coord.x + dir.x, coord.y + dir.y, coord.z + dir.z)
                    checkForNext.add(cur)
                }
            }

            val nextDayTiles = HashMap<Coord, Boolean>()
            for (coord in checkForNext) {
                val neighbors = Direction.values().map { it.getDir() }.map { Coord(it.x + coord.x, it.y + coord.y, it.z + coord.z) }
                val total = neighbors.map {
                    when (tiles[it] == null || tiles[it] == true) {
                        true -> 0
                        false -> 1
                    }
                }.reduce { acc, n -> acc + n }
                if (tiles[coord] == null || tiles[coord] == true) {
                    nextDayTiles[coord] = total != 2
                } else {
                    nextDayTiles[coord] = total == 0 || total > 2
                }
            }

            tiles = nextDayTiles
        }

        val res = tiles.values.count { !it }

        return "Day 24, Part 2 - $res"
    }

    data class Coord(val x: Int, val y: Int, val z: Int)

    enum class Direction {
        East {
            override fun getDir(): Coord = Coord(1, -1, 0)
        },
        West {
            override fun getDir(): Coord = Coord(-1, 1, 0)
        },
        NorthEast {
            override fun getDir(): Coord = Coord(1, 0, -1)
        },
        NorthWest {
            override fun getDir(): Coord = Coord(0, 1, -1)
        },
        SouthEast {
            override fun getDir(): Coord = Coord(0, -1, 1)
        },
        SouthWest {
            override fun getDir(): Coord = Coord(-1, 0, 1)
        };

        abstract fun getDir(): Coord

        companion object {
            fun parse(name: String): Direction {
                return when (name) {
                    "e" -> East
                    "w" -> West
                    "ne" -> NorthEast
                    "nw" -> NorthWest
                    "se" -> SouthEast
                    "sw" -> SouthWest
                    else -> throw Exception(":(")
                }
            }
        }
    }

    private fun parseInstructions(lines: ArrayList<String>): ArrayList<ArrayList<Direction>> {
        val instructions = ArrayList<ArrayList<Direction>>()
        for (line in lines) {
            val directions = ArrayList<Direction>()

            var i = 0
            while (i in line.indices) {
                var cur = line[i].toString()

                if (cur in listOf("n", "s")) {
                    i += 1
                    cur += line[i].toString()
                }

                directions.add(Direction.parse(cur))

                i++
            }

            instructions.add(directions)
        }

        return instructions
    }
}