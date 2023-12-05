package org.aoc.solutions.days.day3;

import org.aoc.helpers.Pair;
import org.aoc.solutions.AbstractGridSolution;
import org.jetbrains.annotations.NotNull;

import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;

public class Day3 extends AbstractGridSolution {
  private final Map<Pair<Integer, Integer>, Character> grid;
  private final String DIGITS = "0123456789";
  private final String DIGITS_AND_DOTS = "." + DIGITS;
  private final List<Pair<Integer, List<Pair<Integer, Integer>>>> partNumberLocationPairs;

  public Day3() {
    this.grid = this.readFromFile("day3/input.txt");
    this.partNumberLocationPairs = getPartNumberLocationPairs();
  }

  public Day3(List<String> input) {
    this.grid = this.generateGridFromLines(input);
    this.partNumberLocationPairs = getPartNumberLocationPairs();
  }

  @Override
  public String partOne() {
    int sum = getValidPartNumbersAndLocations().stream()
      .map(Pair::getFirst)
      .reduce(Integer::sum)
      .get();

    return "" + sum;
  }

  @Override
  public String partTwo() {
    int sum = 0;

    List<Pair<Integer, List<Pair<Integer, Integer>>>>  possibleTrueParts = getValidPartNumbersAndLocations();

    for (Pair<Integer, Integer> gearLocation : getGearLocations()) {
      List<Pair<Integer, List<Pair<Integer, Integer>>>> neighboringParts = possibleTrueParts.stream()
        .filter((partNumberLocationPair) ->
          partNumberLocationPair.getSecond()
            .stream()
            .map(this::getNeighboringCoordinates)
            .anyMatch(l -> l.contains(gearLocation))
        )
        .toList();

      if (neighboringParts.size() == 2) {
        sum += neighboringParts.get(0).getFirst() * neighboringParts.get(1).getFirst();
      }
    }

    return "" + sum;
  }

  @NotNull
  private List<Pair<Integer, List<Pair<Integer, Integer>>>> getPartNumberLocationPairs() {
    List<Pair<Integer, List<Pair<Integer, Integer>>>> partNumbers = new ArrayList<>();

    for (int y = 0; y < this.MAX_Y; y++) {
      boolean buildingPartNumber = false;
      StringBuilder partNumber = new StringBuilder();
      List<Pair<Integer, Integer>> partLocations = new ArrayList<>();

      for (int x = 0; x < this.MAX_X; x++) {
        Pair<Integer, Integer> currentLocation = Pair.from(x, y);
        if (DIGITS.indexOf(this.grid.get(currentLocation)) >= 0) {
          if (!buildingPartNumber) {
            buildingPartNumber = true;
          }

          partNumber.append(grid.get(currentLocation));
          partLocations.add(currentLocation);
        } else if (buildingPartNumber) {
          buildingPartNumber = false;
          partNumbers.add(Pair.from(Integer.parseInt(partNumber.toString()), partLocations));
          partLocations = new ArrayList<>();
          partNumber = new StringBuilder();
        }
      }

      if (buildingPartNumber) {
        partNumbers.add(Pair.from(Integer.parseInt(partNumber.toString()), partLocations));
      }
    }
    return partNumbers;
  }

  @NotNull
  private List<Pair<Integer, List<Pair<Integer, Integer>>>> getValidPartNumbersAndLocations() {
    List<Pair<Integer, List<Pair<Integer, Integer>>>> validPartNumbersAndLocations = new ArrayList<>();
    for (Pair<Integer, List<Pair<Integer, Integer>>> partNumberLocationPair : partNumberLocationPairs) {
      Integer partNumber = partNumberLocationPair.getFirst();
      List<Pair<Integer, Integer>> partLocations = partNumberLocationPair.getSecond();
      List<Pair<Integer, Integer>> neighboringLocations = partLocations
        .stream()
        .map(this::getNeighboringCoordinates)
        .map(coords -> coords.stream().filter(this.grid::containsKey).toList())
        .flatMap(List::stream)
        .toList();

      if (neighboringLocations.stream().anyMatch(location -> DIGITS_AND_DOTS.indexOf(grid.get(location)) == -1)) {
        validPartNumbersAndLocations.add(partNumberLocationPair);
      }
    }
    return validPartNumbersAndLocations;
  }

  @NotNull
  private List<Pair<Integer, Integer>> getGearLocations() {
    List<Pair<Integer, Integer>> gearLocations = new ArrayList<>();

    for (int y = 0; y < this.MAX_Y; y++) {
      for (int x = 0; x < this.MAX_X; x++) {
        Pair<Integer, Integer> currentLocation = Pair.from(x, y);
        if (this.grid.get(currentLocation) == '*') {
          gearLocations.add(currentLocation);
        }
      }
    }
    return gearLocations;
  }
}
