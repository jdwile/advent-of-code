package org.aoc.solutions.days.day1;

import org.aoc.solutions.AbstractSolution;
import org.aoc.solutions.Solution;

import java.util.*;

public class Day1 extends AbstractSolution {
  private final List<String> input;

  public Day1() {
    this.input = this.readFromFile("day1/input.txt");
  }

  public Day1(List<String> input) {
    this.input = input;
  }

  @Override
  public String partOne() {
    return input.stream()
      .map(this::convertCalibrationToDigitPair)
      .reduce(0, Integer::sum)
      .toString();
  }

  @Override
  public String partTwo() {
    return input.stream()
      .map(this::replaceSpellingsWithDigits)
      .map(this::convertCalibrationToDigitPair).reduce(0, Integer::sum)
      .toString();
  }

  private Integer convertCalibrationToDigitPair(String calibrationValue) {
    List<Character> digits = Arrays.stream(calibrationValue.split(""))
      .map(c -> c.charAt(0))
      .filter(Character::isDigit)
      .toList();

    return Integer.parseInt(digits.get(0).toString() + digits.get(digits.size() - 1));
  }

  private String replaceSpellingsWithDigits(String calibrationValue) {
    Map<String, Integer> spellings = Map.of(
      "one", 1,
      "two", 2,
      "three", 3,
      "four", 4,
      "five", 5,
      "six", 6,
      "seven", 7,
      "eight", 8,
      "nine", 9
    );

    for (String spelling : spellings.keySet()) {
      calibrationValue = calibrationValue.replaceAll(spelling, spelling + spellings.get(spelling) + spelling);
    }

    return calibrationValue;
  }
}
