package org.aoc.solutions.days.day2;

import org.aoc.solutions.AbstractSolution;
import org.aoc.solutions.Solution;

import java.util.*;
import java.util.concurrent.atomic.AtomicInteger;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class Day2 extends AbstractSolution implements Solution {
  private final List<String> raw_input;
  private Map<Integer, List<Map<String, Integer>>> game_info;
  private final Map<String, Integer> COLOUR_LIMITS = Map.of("red", 12, "green", 13, "blue", 14);

  public Day2() {
    this.raw_input = this.readFromFile("day2/input.txt");
    this.game_info = parseGames();
  }

  public Day2(List<String> input) {
    this.raw_input = input;
    this.game_info = parseGames();
  }

  Map<Integer, List<Map<String, Integer>>> parseGames() {
    Map<Integer, List<Map<String, Integer>>> rounds = new HashMap<>();
    Pattern p = Pattern.compile("^Game (\\d+): (.+)$");

    raw_input.forEach(game -> {
      Matcher m = p.matcher(game);
      if (!m.matches()) return;
      Integer roundNumber = Integer.parseInt(m.group(1));

      List<Map<String, Integer>> round = Arrays.stream(m.group(2).split("; ")).map(set -> {
        Map<String, Integer> hand = new HashMap<>();
        Arrays.stream(set.split(", ")).forEach(pull -> {
          String[] pullInfo = pull.split(" ");
          int count = Integer.parseInt(pullInfo[0]);
          String colour = pullInfo[1];
          hand.put(colour, count);
        });
        return hand;
      }).toList();

      rounds.put(roundNumber, round);
    });

    return rounds;
  }

  @Override
  public String partOne() {
    AtomicInteger roundSum = new AtomicInteger();
    this.game_info.keySet().forEach(round -> {
      Map<String, Integer> highestSeen = tallyColours(round);
      if (highestSeen.keySet().stream().anyMatch(colour -> highestSeen.get(colour) > COLOUR_LIMITS.get(colour))) {
        return;
      }
      roundSum.addAndGet(round);
    });

    return roundSum.toString();
  }

  @Override
  public String partTwo() {
    AtomicInteger roundSum = new AtomicInteger();
    this.game_info.keySet().forEach(round -> {
      Map<String, Integer> highestSeen = tallyColours(round);
      roundSum.addAndGet(highestSeen.values().stream().reduce(1, (acc, val) -> acc *= val));
    });

    return roundSum.toString();
  }

  private Map<String, Integer> tallyColours(Integer round) {
    Map<String, Integer> highestSeen = new HashMap<>(Map.of("red", 0, "green", 0, "blue", 0));
    game_info.get(round).forEach(hand -> {
      hand.keySet().forEach(colour -> {
        highestSeen.put(colour, Math.max(hand.get(colour), highestSeen.get(colour)));
      });
    });
    return highestSeen;
  }
}
