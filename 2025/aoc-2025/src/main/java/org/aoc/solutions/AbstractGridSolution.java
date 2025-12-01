package org.aoc.solutions;

import org.aoc.helpers.Pair;

import java.io.BufferedReader;
import java.io.FileReader;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

public abstract class AbstractGridSolution implements Solution {
  protected int MAX_X, MAX_Y;
  protected Map<Pair<Integer, Integer>, List<Pair<Integer, Integer>>> neighbor_cache = new ConcurrentHashMap<>();
  protected Map<Pair<Integer, Integer>, Character> readFromFile(String filePath) {
    String file ="C:\\Users\\Jenner\\bench\\advent-of-code\\2025\\aoc-2025\\src\\main\\java\\org\\aoc\\solutions\\days\\" + filePath;
    List<String> lines = List.of();

    try {
      BufferedReader reader = new BufferedReader(new FileReader(file));
      lines = reader.lines().toList();
      reader.close();
    } catch (Exception ignored) {
    }

    return generateGridFromLines(lines);
  }
  protected Map<Pair<Integer, Integer>, Character> generateGridFromLines(List<String> lines) {
    Map<Pair<Integer, Integer>, Character> grid = new HashMap<>();
    for (int y = 0; y < lines.size(); y++) {
      for (int x = 0; x < lines.get(y).length(); x++) {
        grid.put(Pair.from(x, y), lines.get(y).charAt(x));
      }
    }

    this.MAX_X = lines.get(0).length();
    this.MAX_Y = lines.size();

    return grid;
  }

  protected List<Pair<Integer, Integer>> getNeighboringCoordinates(Pair<Integer, Integer> coordinate) {
    return neighbor_cache.computeIfAbsent(coordinate, coord -> {
      List<Pair<Integer, Integer>> neighbors = new ArrayList<>();
      for (int dx = -1; dx <= 1; dx++) {
        for (int dy = -1; dy <= 1; dy++) {
          if (dx == 0 && dy == 0) continue;
          neighbors.add(Pair.from(coord.getFirst() + dx, coord.getSecond() + dy));
        }
      }
      return neighbors;
    });
  }

  abstract public String partOne();

  abstract public String partTwo();
}
