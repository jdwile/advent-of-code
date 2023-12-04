package org.aoc.solutions.days.day2;

import org.aoc.solutions.Solution;
import org.aoc.solutions.days.day1.Day1;
import org.junit.jupiter.api.Test;

import java.util.List;

import static org.junit.jupiter.api.Assertions.assertEquals;


public class Day2Test {
  @Test
  public void TestPartOne() {
    List<String> testInput = List.of("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
      "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
      "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
      "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
      "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"
    );

    Solution solution = new Day2(testInput);

    String result = solution.partOne();
    assertEquals(result, "8");
  }

  @Test
  public void TestPartTwo() {
    List<String> testInput = List.of("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
      "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
      "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
      "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
      "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"
    );

    Solution solution = new Day2(testInput);

    String result = solution.partTwo();
    assertEquals(result, "2286");
  }
}
