package org.aoc.solutions.days.day3;

import org.aoc.helpers.Pair;
import org.aoc.solutions.AbstractGridSolution;
import org.aoc.solutions.Solution;
import org.aoc.solutions.days.day2.Day2;
import org.junit.jupiter.api.Test;

import java.util.List;

import static org.junit.jupiter.api.Assertions.assertEquals;


public class Day3Test {
  @Test
  public void TestPartOne() {
    List<String> testInput = List.of(
      "467..114..",
      "...*......",
      "..35...633",
      "......#...",
      "617*......",
      ".....+.58.",
      "..592.....",
      "......755.",
      "...$.*....",
      ".664.598.."
    );

    Solution solution = new Day3(testInput);

    String result = solution.partOne();
    assertEquals(result, "4361");
  }
  @Test
  public void TestPartTwo() {
    List<String> testInput = List.of(
      "467..114..",
      "...*......",
      "..35...633",
      "......#...",
      "617*......",
      ".....+.58.",
      "..592.....",
      "......755.",
      "...$.*....",
      ".664.598.."
    );

    Solution solution = new Day3(testInput);

    String result = solution.partTwo();
    assertEquals(result, "467835");
  }
}
