package org.aoc.solutions.days.day3;

import org.aoc.solutions.AbstractSolution;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
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
    int totalJoltage = input.stream()
      .map(this::getHighJoltage)
      .mapToInt(Integer::intValue)
      .sum();

    return "" + totalJoltage;
  }

  @Override
  public String partTwo() {
    long totalJoltage = input.stream()
      .map(battery -> calculateJoltage(battery, recursiveGetHighJoltage(battery, 12, new ArrayList<>())))
      .mapToLong(Long::longValue)
      .sum();

    return "" + totalJoltage;
  }

  private int getHighJoltage(List<Integer> battery) {
    int highestIndex = findMaxIndexIgnoring(battery, List.of());
    int highestValue = battery.get(highestIndex);
    int secondHighestIndex = findMaxIndexIgnoring(battery, List.of(highestIndex));
    int secondHighestValue = battery.get(secondHighestIndex);

    int highJoltage = highestValue * 10 + secondHighestValue;

    if (highestIndex > secondHighestIndex) {
      highJoltage = secondHighestValue * 10 + highestValue;

      List<Integer> ignoreList = IntStream.range(0, highestIndex + 1).boxed().toList();
      int highestAfterIndex = findMaxIndexIgnoring(battery, ignoreList);

      if (highestAfterIndex > -1) {
        int highestAfterValue = battery.get(highestAfterIndex);

        int contenderJoltage = highestValue * 10 + highestAfterValue;
        if (contenderJoltage > highJoltage) highJoltage = contenderJoltage;
      }
    }
    return highJoltage;
  }

  private ArrayList<Integer> recursiveGetHighJoltage(List<Integer> battery, int digits, ArrayList<Integer> ignored) {
    int highestIndex = findMaxIndexIgnoring(battery, ignored);
    if (highestIndex < 0) return ignored;

    ArrayList<Integer> highJoltageIndices = new ArrayList<>();
    highJoltageIndices.add(highestIndex);
    highJoltageIndices.addAll(ignored);
    if (digits == 1) {
      return highJoltageIndices;
    }

    if (digits > 1) {
      highJoltageIndices = recursiveGetHighJoltage(battery, digits - 1, highJoltageIndices);
      highJoltageIndices.sort(Integer::compareTo);
    }

    long highJoltage = calculateJoltage(battery, highJoltageIndices);

    if (highJoltageIndices.size() >= 2) {
      while (battery.get(highJoltageIndices.get(0)) < battery.get(highJoltageIndices.get(1))) {
        ArrayList<Integer> ignoreList = new ArrayList<>();
        ignoreList.addAll(IntStream.range(0, highestIndex + 1).boxed().toList());
        ignoreList.addAll(highJoltageIndices);
        int highestAfterIndex = findMaxIndexIgnoring(battery, ignoreList);
        if (highestAfterIndex < 0) {
          break;
        }
        ArrayList<Integer> newJoltageIndices = new ArrayList<>(highJoltageIndices);
        newJoltageIndices.removeFirst();
        newJoltageIndices.add(highestAfterIndex);
        newJoltageIndices.sort(Integer::compareTo);

        long highJoltageCandidate = calculateJoltage(battery, newJoltageIndices);
        if (highJoltageCandidate > highJoltage) {
          highJoltage = highJoltageCandidate;
          highJoltageIndices = newJoltageIndices;
        } else {
          highJoltageIndices = ignored;
          break;
        }
      }
    }

//    if (digits == 12)
//      System.out.printf("%s -> %d\n", battery.stream().map(String::valueOf).collect(Collectors.joining()), highJoltage);

    return highJoltageIndices;
  }

  private long calculateJoltage(List<Integer> battery, List<Integer> highJoltageIndices) {
    long highJoltage = 0;

    for (int i = highJoltageIndices.size() - 1; i >= 0; i--) {
      highJoltage += (long) (battery.get(highJoltageIndices.get(i)) * Math.pow(10, highJoltageIndices.size() - i - 1));
    }

    return highJoltage;
  }

  private int findMaxIndexIgnoring(List<Integer> battery, List<Integer> ignored) {
    int maxIndex = -1;
    int maxVal = -1;

    for (int i = 0; i < battery.size(); i++) {
      if (ignored.contains(i)) continue;

      int cur = battery.get(i);
      if (maxVal <= cur) {
        maxVal = cur;
        maxIndex = i;
      }
    }

    return maxIndex;
  }
}
