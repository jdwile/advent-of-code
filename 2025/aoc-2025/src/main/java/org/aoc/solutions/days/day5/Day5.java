package org.aoc.solutions.days.day5;

import com.google.common.collect.Range;
import org.aoc.solutions.AbstractSolution;

import java.util.ArrayList;
import java.util.List;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class Day5 extends AbstractSolution {
  private final List<String> input;
  private final List<Range<Long>> ranges;
  private final List<Long> ingredients;

  public Day5() {
    this.input = this.readFromFile("day5/input.txt");
    this.ranges = parseRanges(input);
    this.ingredients = parseIngredients(input);
  }

  public Day5(List<String> input) {
    this.input = input;
    this.ranges = parseRanges(input);
    this.ingredients = parseIngredients(input);
  }

  @Override
  public String partOne() {
    return "" + ingredients.stream()
      .filter(this::isFresh)
      .toList()
      .size();
  }

  @Override
  public String partTwo() {
    return "" + ranges
      .stream()
      .map(range -> range.upperEndpoint() - range.lowerEndpoint() + 1)
      .mapToLong(Long::longValue)
      .sum();
  }

  private List<Long> parseIngredients(List<String> input) {
    List<Long> ingredients = new ArrayList<>();

    for (String line : input) {
      if (line.contains("-") || line.isEmpty()) continue;

      ingredients.add(Long.parseLong(line));
    }

    return ingredients;
  }

  private List<Range<Long>> parseRanges(List<String> input) {
    Pattern rangePattern = Pattern.compile("^(\\d+)-(\\d+)$");
    List<Range<Long>> ranges = new ArrayList<>();

    for (String line : input) {
      if (line.isEmpty()) break;

      Matcher m = rangePattern.matcher(line);
      if (m.matches()) {
        Long low = Long.parseLong(m.group(1));
        Long high = Long.parseLong(m.group(2));
        ranges.add(Range.closed(low, high));
      }
    }

    return compressRanges(ranges);
  }

  private List<Range<Long>> compressRanges(List<Range<Long>> ranges) {
    List<Range<Long>> minimalRanges = new ArrayList<>(ranges);
    boolean modified = true;

    while (modified) {
      modified = false;

      for (int i = 0; i < minimalRanges.size() - 1; i++) {
        Range<Long> curRange = minimalRanges.get(i);

        for (int j = i + 1; j < minimalRanges.size(); j++) {
          Range<Long> existing = minimalRanges.get(j);

          if (existing.encloses(curRange)) {
            minimalRanges.remove(curRange);
            modified = true;
            break;
          }

          if (existing.isConnected(curRange)) {
            Range<Long> span = existing.span(curRange);
            if (span.encloses(curRange) || span.encloses(existing)) {
              minimalRanges.remove(existing);
              minimalRanges.remove(curRange);
              minimalRanges.add(span);
              modified = true;
              break;
            }
          }
        }
        if (modified) break;
      }
    }
    return minimalRanges;
  }

  private boolean isFresh(Long ingredient) {
    return ranges
      .stream()
      .anyMatch(range -> range.contains(ingredient));
  }
}
