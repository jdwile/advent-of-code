package org.aoc.solutions.days.day7;

import org.aoc.helpers.Pair;
import org.aoc.solutions.AbstractGridSolution;

import java.util.HashMap;
import java.util.Map;

public class Day7 extends AbstractGridSolution {
  private final Map<Pair<Integer, Integer>, Character> input;

  public Day7() {
    this.input = this.readFromFile("day7/input.txt");
  }

  public Day7(Map<Pair<Integer, Integer>, Character> input) {
    this.input = input;
  }

  @Override
  public String partOne() {
    int splits = 0;

    for (int y = 1; y <= MAX_Y; y++) {
      for (int x = 0; x <= MAX_X; x++) {
        char cur = input.getOrDefault(Pair.from(x, y), '.');
        char above = input.getOrDefault(Pair.from(x, y - 1), '.');

        if (above == '|' || above == 'S') {
          if (cur == '.') {
            input.put(Pair.from(x, y), '|');
          } else if (cur == '^') {
            input.put(Pair.from(x - 1, y), '|');
            input.put(Pair.from(x + 1, y), '|');
            splits++;
          }
        }
      }
    }
    return "" + splits;
  }

  @Override
  public String partTwo() {
    Map<Pair<Integer, Integer>, Long> dp = new HashMap<>();

    for (int x = 0; x < +MAX_X; x++) {
      if (input.getOrDefault(Pair.from(x, MAX_Y), '.') == '|') {
        dp.put(Pair.from(x, MAX_Y), 1L);
      }
    }

    for (int y = MAX_Y - 1; y >= 0; y--) {
      for (int x = 0; x <= MAX_X; x++) {
        char cur = input.getOrDefault(Pair.from(x, y), '.');
        char belowGrid = input.getOrDefault(Pair.from(x, y + 1), '.');
        long belowDP = dp.getOrDefault(Pair.from(x, y + 1), 0L);

        if (cur == '|' || cur == 'S') {
          if (belowGrid == '^') {
            dp.put(Pair.from(x, y), dp.getOrDefault(Pair.from(x - 1, y + 1), 0L) + dp.getOrDefault(Pair.from(x + 1, y + 1), 0L));
          } else {
            dp.put(Pair.from(x, y), belowDP);
          }
        }
      }
    }

    long max = 0;
    for (int x = 0; x <= MAX_X; x++) {
      max = Math.max(max, dp.getOrDefault(Pair.from(x, 0), 0L));
    }

    return "" + max;
  }
}
