package org.aoc.solutions.days.day4;

import java.util.List;
import java.util.Objects;
import java.util.stream.Collectors;

public class Scorecard {

  private List<Integer> winningNumbers;

  private List<Integer> pickedNumbers;
  private int id;

  public Scorecard(int id, List<Integer> winningNumbers, List<Integer> pickedNumbers) {
    this.id = id;
    this.winningNumbers = winningNumbers;
    this.pickedNumbers = pickedNumbers;
  }

  public int getId() {
    return id;
  }

  public void setId(int id) {
    this.id = id;
  }

  public List<Integer> getPickedNumbers() {
    return pickedNumbers;
  }

  public void setPickedNumbers(List<Integer> pickedNumbers) {
    this.pickedNumbers = pickedNumbers;
  }
  public List<Integer> getWinningNumbers() {
    return winningNumbers;
  }

  public void setWinningNumbers(List<Integer> winningNumbers) {
    this.winningNumbers = winningNumbers;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) return true;
    if (o == null || getClass() != o.getClass()) return false;
    Scorecard scorecard = (Scorecard) o;
    return Objects.equals(getWinningNumbers(), scorecard.getWinningNumbers())
      && Objects.equals(getPickedNumbers(), scorecard.getPickedNumbers())
      && getId() == scorecard.getId();
  }

  @Override
  public int hashCode() {
    return Objects.hash(getWinningNumbers(), getPickedNumbers());
  }

  @Override
  public String toString() {
    return "Scorecard " + getId() + " { " +
       getWinningNumbers().stream().map(Object::toString).collect(Collectors.joining(" ")) +
      " | "
      + getPickedNumbers().stream().map(Object::toString).collect(Collectors.joining(" ")) +
      " }";
  }
}
