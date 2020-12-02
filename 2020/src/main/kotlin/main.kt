import common.ISolution
import problems.*

fun main() {
   val solutions = ArrayList<ISolution>()
   solutions.add(Example())
   solutions.add(Day1())
   solutions.add(Day2())

//   solutions.forEach {
//      println(it.part1())
//      println(it.part2())
//   }
   println(solutions[2].part1())
   println(solutions[2].part2())
}