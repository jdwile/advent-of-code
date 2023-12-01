package org.aoc;

import org.aoc.solutions.Solution;
import org.aoc.solutions.days.day1.Day1;

public class Main {
  public static void main(String[] args) {
    Solution sol = new Day1();
    System.out.println("Part 1: " + sol.partOne());
    System.out.println("Part 2: " + sol.partTwo());
  }
}
