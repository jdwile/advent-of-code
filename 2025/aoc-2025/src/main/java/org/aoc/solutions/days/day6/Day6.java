package org.aoc.solutions.days.day6;

import com.google.common.collect.Range;
import org.aoc.solutions.AbstractSolution;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class Day6 extends AbstractSolution {
  private final List<String> input;
  private final List<String> OPERATORS = List.of("*", "+");

  public Day6() {
    this.input = this.readFromFile("day6/input.txt");
  }

  public Day6(List<String> input) {
    this.input = input;
  }

  @Override
  public String partOne() {
    List<List<Long>> operands = new ArrayList<>();
    List<String> operators = new ArrayList<>(
      Arrays.stream(
          input.getLast().trim().split("\\s+")
        )
        .toList()
    );

    long result = 0L;

    for (int i = 0; i < input.size() - 1; i++) {
      operands.add(
        Arrays.stream(
            input.get(i).trim().split("\\s+")
          )
          .map(Long::parseLong)
          .toList());
    }

    for (int i = 0; i < operators.size(); i++) {
      int index = i;
      String operator = operators.get(index);

      if (operator.equals("*")) {
        result += operands
          .stream()
          .map(line -> line.get(index))
          .reduce(1L, (acc, val) -> acc * val);
      } else if (operator.equals("+")) {
        result += operands
          .stream()
          .map(line -> line.get(index))
          .reduce(0L, Long::sum);
      }
    }
    return "" + result;
  }

  @Override
  public String partTwo() {
    List<List<Long>> operands = new ArrayList<>();
    List<String> operators = new ArrayList<>(
      Arrays.stream(
          input.getLast().trim().split("\\s+")
        )
        .toList()
        .reversed()
    );

    long result = 0L;


    List<Long> curOperandList = new ArrayList<>();
    for (int x = input.getFirst().length() - 1; x >= 0; x--) {
      StringBuilder longString = new StringBuilder();
      for (int y = 0; y < input.size() - 1; y++) {
        longString.append(input.get(y).charAt(x));
      }

      String finalLongString = longString.toString().trim();
      if (finalLongString.isEmpty()) {
        operands.add(curOperandList);
        curOperandList = new ArrayList<>();
      } else {
        curOperandList.add(Long.parseLong(finalLongString));
      }
    }
    operands.add(curOperandList);

    for (int i = 0; i < operators.size(); i++) {
      String operator = operators.get(i);

      if (operator.equals("*")) {
        result += operands
          .get(i)
          .stream()
          .reduce(1L, (acc, val) -> acc * val);
      } else if (operator.equals("+")) {
        result += operands
          .get(i)
          .stream()
          .reduce(0L, Long::sum);
      }
    }
    return "" + result;
  }
}
