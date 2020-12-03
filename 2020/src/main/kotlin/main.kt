import common.ISolution
import problems.*
import kotlin.system.measureTimeMillis

fun main() {
   val solutions = ArrayList<ISolution>()
   solutions.add(Example())
   solutions.add(Day1())
   solutions.add(Day2())
   solutions.add(Day3())
   solutions.add(Day4())

   solutions.forEach {
      val part1Time = measureTimeMillis {
         print(it.part1())
      }
      println(" - $part1Time ms")

      val part2Time = measureTimeMillis {
         print(it.part2())
      }
      println(" - $part2Time ms")
   }

//   println(solutions[4].part1())
//   println(solutions[4].part2())
}