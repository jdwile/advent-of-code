package org.aoc.solutions;

import java.io.*;
import java.util.List;

public abstract class AbstractSolution implements Solution {
  protected List<String> readFromFile(String filePath) {
      String file ="C:\\Users\\Jenner\\bench\\advent-of-code\\2023\\aoc-2023\\src\\main\\java\\org\\aoc\\solutions\\days\\" + filePath;
      List<String> lines = List.of();

      try {
        BufferedReader reader = new BufferedReader(new FileReader(file));
        lines = reader.lines().toList();
        reader.close();
      } catch (Exception ignored) {
      }

      return lines;
  }

  abstract public String partOne();

  abstract public String partTwo();
}
