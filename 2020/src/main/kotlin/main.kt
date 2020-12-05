import common.ISolution
import problems.*
import java.math.RoundingMode
import java.text.DecimalFormat
import java.text.NumberFormat
import javax.swing.text.NumberFormatter
import kotlin.system.measureNanoTime
import kotlin.system.measureTimeMillis

fun main() {
   val solutions = ArrayList<ISolution>()
   solutions.add(Day1())
   solutions.add(Day2())
   solutions.add(Day3())
   solutions.add(Day4())
   solutions.add(Day5())
   solutions.add(Day6())

   val dec = DecimalFormat("###,###.##")
   dec.roundingMode = RoundingMode.HALF_UP

   val time = measureTimeMillis {
      solutions.forEach {
         val part1Time: Double = measureNanoTime {
            repeat(5) { _ -> it.part1() }
         }.toDouble()
         print(it.part1())
         println(" - " + dec.format(part1Time / 5 / 1000000) + "ms")

         val part2Time: Double = measureNanoTime {
            repeat(5) { _ -> it.part2() }
         }.toDouble()
         print(it.part2())
         println(" - " + dec.format(part2Time / 5 / 1000000) + "ms")
      }

//      println(solutions[6].part1())
//      println(solutions[6].part2())
   }
   println(" --- $time ms total")
}