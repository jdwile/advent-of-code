package org.aoc.solutions.days.day4;

import org.aoc.helpers.Pair;
import org.aoc.solutions.AbstractSolution;

import java.util.*;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class Day4 extends AbstractSolution {
  private final List<String> raw_input;
  private final List<Scorecard> cards;

  public Day4() {
    this.raw_input = this.readFromFile("day4/input.txt");
    this.cards = parseCards();
  }

  private List<Scorecard> parseCards() {
    List<Scorecard> cards = new ArrayList<>();
    Pattern p = Pattern.compile("^Card (\\d+):(.*)\\|(.*)$");

    this.raw_input.forEach(line -> {
      Matcher m = p.matcher(line);
      if (m.matches()) {
        cards.add(new Scorecard(
          Integer.parseInt(m.group(1)),
          Arrays.stream(m.group(2).trim().split("\\s+")).map(Integer::parseInt).toList(),
          Arrays.stream(m.group(3).trim().split("\\s+")).map(Integer::parseInt).toList()
        ));
      }
    });
    return cards;
  }

  public Day4(List<String> input) {
    this.raw_input = input;
    this.cards = parseCards();
  }

  @Override
  public String partOne() {
    int points = 0;

    for (Scorecard card : this.cards) {
      int winningCount = card.getPickedNumbers()
        .stream().filter(card.getWinningNumbers()::contains)
        .toList()
        .size();

      if (winningCount > 0) {
        points += (int) Math.pow(2, winningCount - 1);
        System.out.printf("Card %d worth %d points\n", card.getId(), (int) Math.pow(2, winningCount - 1));
      }
    }

    return "" + points;
  }

  @Override
  public String partTwo() {
    return null;
  }
}
