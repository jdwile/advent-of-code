package main.kotlin.problems

import main.kotlin.common.*

class Day17 : ISolution {
    override fun part1(): String {
        var (space: HashMap<Point3D, Boolean>, dim: Int) = createStartingSpace(readFileAsStrings("17.in"))

        for (cycles in 1..6) {
            var coords = -(dim + cycles + 1)..(dim + cycles + 1)
            val next = HashMap<Point3D, Boolean>()
            for (z in -cycles..cycles) {
                for (y in coords) {
                    for (x in coords) {
                        val n = countNeighbors(Point3D(x, y, z), space)

                        if (space[Point3D(x, y, z)] == true) {
                            if (n in arrayOf(2, 3)) {
                                next[Point3D(x, y, z)] = true
                            }
                        } else {
                            if (n == 3) {
                                next[Point3D(x, y, z)] = true
                            }
                        }
                    }
                }
            }
            space = next
        }

        val res = space.values.size

        return "Day 17, Part 1 - $res"
    }

    override fun part2(): String {
        var (space: HashMap<Point4D, Boolean>, dim: Int) = createStartingSpace4D(readFileAsStrings("17.in"))

        for (cycles in 1..6) {
            var coords = -(dim + cycles + 1)..(dim + cycles + 1)
            val next = HashMap<Point4D, Boolean>()
            for (z in -cycles..cycles) {
                for (y in coords) {
                    for (x in coords) {
                        for (w in -cycles..cycles) {
                            val n = countNeighbors4D(Point4D(w, x, y, z), space)

                            if (space[Point4D(w, x, y, z)] == true) {
                                if (n in arrayOf(2, 3)) {
                                    next[Point4D(w, x, y, z)] = true
                                }
                            } else {
                                if (n == 3) {
                                    next[Point4D(w, x, y, z)] = true
                                }
                            }
                        }
                    }
                }
            }
            space = next
        }

        val res = space.values.size

        return "Day 17, Part 2 - $res"
    }

    private fun countNeighbors(point: Point3D, space: HashMap<Point3D, Boolean>): Int {
        val diffs = arrayOf(-1, 0, 1)
        var res = 0

        for (dz in diffs) {
            for (dy in diffs) {
                for (dx in diffs) {
                    if (dx == 0 && dy == 0 && dz == 0) continue
                    if (space[Point3D(point.x + dx, point.y + dy, point.z + dz)] == true) res++
                }
            }
        }

        return res
    }

    private fun countNeighbors4D(point: Point4D, space: HashMap<Point4D, Boolean>): Int {
        val diffs = arrayOf(-1, 0, 1)
        var res = 0

        for (dz in diffs) {
            for (dy in diffs) {
                for (dx in diffs) {
                    for (dw in diffs) {
                        if (dx == 0 && dy == 0 && dz == 0 && dw == 0) continue
                        if (space[Point4D(point.w + dw, point.x + dx, point.y + dy, point.z + dz)] == true) res++
                    }
                }
            }
        }

        return res
    }

//    private fun printSpace(space: HashMap<Point3D, Boolean>, dim: Int, zMax: Int) {
//        for (z in -zMax..zMax) {
//            println("z=$z")
//            for (y in -dim..dim) {
//                for (x in -dim..dim) {
//                    if (space[Point3D(x, y, z)] == true) {
//                        print('#')
//                    } else {
//                        print('.')
//                    }
//                }
//                println("")
//            }
//            println("")
//        }
//    }

    private fun createStartingSpace(lines: ArrayList<String>): Pair<HashMap<Point3D, Boolean>, Int> {
        val dim = (lines[0].length - 1) / 2
        val grid = HashMap<Point3D, Boolean>()

        for (y in lines.indices) {
            for (x in lines[0].indices) {
                if (lines[y][x] == '#') {
                    grid[Point3D(x - dim, y - dim, 0)] = true
                }
            }
        }

        return Pair(grid, dim)
    }

    private fun createStartingSpace4D(lines: ArrayList<String>): Pair<HashMap<Point4D, Boolean>, Int> {
        val dim = (lines[0].length - 1) / 2
        val grid = HashMap<Point4D, Boolean>()

        for (y in lines.indices) {
            for (x in lines[0].indices) {
                if (lines[y][x] == '#') {
                    grid[Point4D(0, x - dim, y - dim, 0)] = true
                }
            }
        }

        return Pair(grid, dim)
    }

    private data class Point3D(val x: Int, val y: Int, val z: Int)
    private data class Point4D(val w: Int, val x: Int, val y: Int, val z: Int)
}