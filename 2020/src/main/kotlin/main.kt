import common.ISolution
import problems.*
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

    val dec = DecimalFormat("###,###.##")
    dec.roundingMode = RoundingMode.HALF_UP

    solutions.forEach {
        val part1Time: Double = measureNanoTime {
            repeat(50) { _ -> it.part1() }
        }.toDouble()
        print(it.part1())
        println(" - " + dec.format(part1Time / 50 / 1000000) + "ms")

        val part2Time: Double = measureNanoTime {
            repeat(50) { _ -> it.part2() }
        }.toDouble()
        print(it.part2())
        println(" - " + dec.format(part2Time / 50 / 1000000) + "ms")
    }

//      println(solutions[10].part1())
//      println(solutions[10].part2())
}