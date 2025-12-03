package org.aoc.solutions.days.day3;

import org.aoc.solutions.AbstractSolution;

import java.util.*;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

public class Day3 extends AbstractSolution {
  private final List<List<Integer>> input;

  public Day3() {
    this.input = this.readFromFile("day3/input.txt").stream()
      .map(rawBattery ->
        Arrays.stream(rawBattery.split(""))
          .mapToInt(Integer::parseInt).boxed().toList()
      ).toList();
  }

  public Day3(List<List<Integer>> input) {
    this.input = input;
  }

  @Override
  public String partOne() {
    long totalJoltage = input.stream()
      .map(battery -> getHighJoltage(battery, 2))
      .mapToLong(Long::longValue)
      .sum();

    return "" + totalJoltage;
  }

  @Override
  public String partTwo() {
    long totalJoltage = input.stream()
      .map(battery -> getHighJoltage(battery, 12))
      .mapToLong(Long::longValue)
      .sum();

    return "" + totalJoltage;
  }

  private long getHighJoltage(List<Integer> battery, int length) {
    if (length > battery.size()) return -1;

    ArrayList<Integer> joltageIndices = new ArrayList<>();
    Set<Integer> ignored = new HashSet<>();

    for (int i = 0; i < length; i++) {
      int idx = findMaxIndexWithinRangeIgnoring(battery, i, battery.size() - length + i, ignored);
      joltageIndices.add(idx);
      ignored.addAll(IntStream.range(0, idx + 1).boxed().toList());
    }

    return calculateJoltage(battery, joltageIndices);
  }

  private long calculateJoltage(List<Integer> battery, List<Integer> highJoltageIndices) {
    return Long.parseLong(
      highJoltageIndices.stream()
        .map(battery::get)
        .map(String::valueOf)
        .collect(Collectors.joining()));
  }

  private int findMaxIndexWithinRangeIgnoring(List<Integer> battery, int start, int end, Set<Integer> ignored) {
    int maxVal = -1;
    int maxIndex = -1;

    for (int i = start; i <= end; i++) {
      if (ignored.contains(i)) continue;

      int cur = battery.get(i);
      if (maxVal < cur) {
        maxVal = cur;
        maxIndex = i;
      }
    }

    return maxIndex;
  }
}
