package org.aoc.solutions.days.day1;

import org.aoc.helpers.Pair;
import org.aoc.solutions.AbstractSolution;

import java.util.*;
import java.util.concurrent.atomic.AtomicInteger;

public class Day1 extends AbstractSolution {
  private final List<String> input;

  public Day1() {
    this.input = this.readFromFile("day1/input.txt");
  }

  public Day1(List<String> input) {
    this.input = input;
  }

  @Override
  public String partOne() {
    AtomicInteger restedAtZeroCount = new AtomicInteger();
    AtomicInteger dial = new AtomicInteger(50);

    input
      .stream()
      .map(this::convertInputToDirectionPair)
      .filter(Objects::nonNull)
      .forEach(val -> {
        dial.addAndGet(val.getFirst() * val.getSecond());
        while (dial.get() < 0) dial.addAndGet(100);
        while (dial.get() >= 100) dial.addAndGet(-100);
        if (dial.get() == 0) restedAtZeroCount.getAndIncrement();
      });

    return restedAtZeroCount.toString();
  }

  @Override
  public String partTwo() {
    AtomicInteger restedAtZeroCount = new AtomicInteger();
    AtomicInteger dial = new AtomicInteger(50);

    input
      .stream()
      .map(this::convertInputToDirectionPair)
      .filter(Objects::nonNull)
      .forEach(val -> {
        int curDial = dial.get();

        int dir = val.getFirst();
        int rotation = val.getSecond();

        while (rotation > 99) {
          restedAtZeroCount.addAndGet(1);
          rotation -= 100;
        }

        for (int i = 0; i < rotation; i++) {
          curDial += dir;

          if (curDial < 0) curDial += 100;
          if (curDial > 99) curDial -= 100;
          if (curDial == 0) restedAtZeroCount.addAndGet(1);
        }

        dial.set(curDial);
      });

    return restedAtZeroCount.toString();
  }

  private Pair<Integer, Integer> convertInputToDirectionPair(String input) {
    if (Objects.isNull(input) || input.isEmpty()) return null;

    Integer direction = input.substring(0, 1).equals("L") ? -1 : 1;
    Integer rest = Integer.parseInt(input.substring(1));

    return new Pair<>(direction, rest);
  }
}
