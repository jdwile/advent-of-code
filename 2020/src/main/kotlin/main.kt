import main.kotlin.common.ISolution
import main.kotlin.problems.*
import java.math.RoundingMode
import java.text.DecimalFormat
import kotlin.system.measureNanoTime

fun main() {
    val solutions = ArrayList<ISolution>()
    solutions.add(Day1())
    solutions.add(Day2())
    solutions.add(Day3())
    solutions.add(Day4())
    solutions.add(Day5())
    solutions.add(Day6())
    solutions.add(Day7())
    solutions.add(Day8())
    solutions.add(Day9())
    solutions.add(Day10())
    solutions.add(Day11())
    solutions.add(Day12())
    solutions.add(Day13())
    solutions.add(Day14())
    solutions.add(Day15())
    solutions.add(Day16())
    solutions.add(Day17())
    solutions.add(Day18())
    solutions.add(Day19())
    solutions.add(Day20())
    solutions.add(Day21())
    solutions.add(Day22())
    solutions.add(Day23())
    solutions.add(Day24())
    solutions.add(Day25())

    val dec = DecimalFormat("###,###.##")
    dec.roundingMode = RoundingMode.HALF_UP

    val times = 10

    val totalTime: Double = measureNanoTime {

        solutions.forEach {
            val part1Time: Double = measureNanoTime {
                repeat(times) { _ -> it.part1() }
            }.toDouble()
            print(it.part1())
            println(" - " + dec.format(part1Time / times / 1000000) + "ms")

            val part2Time: Double = measureNanoTime {
                repeat(times) { _ -> it.part2() }
            }.toDouble()
            print(it.part2())
            println(" - " + dec.format(part2Time / times / 1000000) + "ms")
        }
    }.toDouble()

    println("\nTotal Time: ${dec.format(totalTime / times / 1000000 / 1000)} seconds (${dec.format(totalTime / times / 1000000)} ms)")
    println("Average Time per problem: ${dec.format(totalTime / times / 1000000 / solutions.size)} ms")

//      println(solutions[20].part1())
//      println(solutions[20].part2())
}