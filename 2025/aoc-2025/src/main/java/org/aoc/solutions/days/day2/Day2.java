package org.aoc.solutions.days.day2;

import org.aoc.helpers.Pair;
import org.aoc.solutions.AbstractSolution;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Objects;
import java.util.regex.Pattern;

public class Day2 extends AbstractSolution {
  private final String input;
  private final Pattern INVALID_ID_PATTERN = Pattern.compile("^(\\d+)\\1$");
  private final Pattern INVALID_ID_PATTERN_2 = Pattern.compile("^(\\d+)\\1+$");

  public Day2() {
    this.input = this.readFromFile("day2/input.txt").getFirst();
  }

  public Day2(String input) {
    this.input = input;
  }

  @Override
  public String partOne() {
    List<Pair<Long, Long>> ranges =
      new ArrayList<>(Objects.requireNonNull(convertInputToRanges(input)));

    return "" + getInvalidIdSum(ranges, INVALID_ID_PATTERN);
  }

  @Override
  public String partTwo() {
    List<Pair<Long, Long>> ranges =
      new ArrayList<>(Objects.requireNonNull(convertInputToRanges(input)));

    return "" + getInvalidIdSum(ranges, INVALID_ID_PATTERN_2);
  }

  private long getInvalidIdSum(List<Pair<Long, Long>> ranges, Pattern INVALID_ID_PATTERN_2) {
    long invalidIdSum = 0;

    for (var range : ranges) {
      for (long i = range.getFirst(); i <= range.getSecond(); i++) {
        String stg = "" + i;
        if (INVALID_ID_PATTERN_2.matcher(stg).matches()) {
          invalidIdSum += i;
        }
      }
    }

    return invalidIdSum;
  }

  private List<Pair<Long, Long>> convertInputToRanges(String input) {
    if (Objects.isNull(input) || input.isEmpty()) return null;

    return Arrays.stream(input.split(","))
      .map(rawRange -> {
        var range = Arrays.stream(rawRange.split("-")).toList();
        return Pair.from(Long.parseLong(range.get(0)), Long.parseLong(range.get(1)));
      })
      .toList();
  }
}
