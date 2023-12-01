package org.aoc.solutions.days.day1;

import org.aoc.solutions.Solution;
import org.junit.jupiter.api.Test;

import java.util.List;

import static org.junit.jupiter.api.Assertions.assertEquals;


public class Day1Test {
  @Test
  public void TestPartOne() {
    List<String> testInput = List.of("1abc2",
      "pqr3stu8vwx",
      "a1b2c3d4e5f",
      "treb7uchet"
    );

    Solution solution = new Day1(testInput);

    String result = solution.partOne();
    assertEquals(result, "142");
  }

  @Test
  public void TestPartTwo() {
    List<String> testInput = List.of("two1nine",
      "eightwothree",
      "abcone2threexyz",
      "xtwone3four",
      "4nineeightseven2",
      "zoneight234",
      "7pqrstsixteen"
    );

    Solution solution = new Day1(testInput);

    String result = solution.partTwo();
    assertEquals(result, "281");
  }
}
