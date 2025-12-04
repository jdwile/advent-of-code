package org.aoc.solutions.days.day4;

import org.aoc.helpers.Pair;
import org.aoc.solutions.AbstractGridSolution;

import java.util.List;
import java.util.Map;

public class Day4 extends AbstractGridSolution {
  private final Map<Pair<Integer, Integer>, Character> input;

  public Day4() {
    this.input = this.readFromFile("day4/input.txt");
  }

  public Day4(Map<Pair<Integer, Integer>, Character> input) {
    this.input = input;
  }

  @Override
  public String partOne() {
    return "" + getAccessiblePaperRolls().size();
  }

  @Override
  public String partTwo() {
    List<Pair<Integer, Integer>> accessibleRolls = getAccessiblePaperRolls();
    int removedRolls = 0;

    while (!accessibleRolls.isEmpty()) {
      accessibleRolls.forEach(paperRoll -> input.put(paperRoll, '.'));
      removedRolls += accessibleRolls.size();
      accessibleRolls = getAccessiblePaperRolls();
    }

    return "" + removedRolls;
  }

  private List<Pair<Integer, Integer>> getAccessiblePaperRolls() {
    return input.keySet().stream()
      .filter(coord -> input.getOrDefault(coord, '.') == '@')
      .filter(coord ->
        getNeighboringCoordinates(coord).stream()
          .filter(neighbor -> input.getOrDefault(neighbor, '.') == '@')
          .toList().size() < 4)
      .toList();
  }
}
