package org.aoc.helpers;

import java.util.Objects;

public class Pair<T1, T2> {
  private T1 first;
  private T2 second;

  public Pair(T1 first, T2 second) {
    this.first = first;
    this.second = second;
  }

  public static <T1, T2> Pair<T1, T2> from(T1 first, T2 second) {
    return new Pair<>(first, second);
  }

  public T1 getFirst() { return this.first; }
  public T2 getSecond() { return this.second; }

  @Override
  public boolean equals(Object obj) {
    if (obj == null) return false;
    if (obj.getClass() != this.getClass()) return false;

    final Pair otherPair = (Pair) obj;
    if (otherPair.getFirst().getClass() != this.getFirst().getClass()) return false;
    if (otherPair.getSecond().getClass() != this.getSecond().getClass()) return false;

    final Pair<T1, T2> other = (Pair<T1, T2>) otherPair;
    return other.getFirst().equals(this.first) && other.getSecond().equals(this.second);
  }

  @Override
  public int hashCode() {
    return Objects.hash(first, second);
  }

  @Override
  public String toString() {
    return String.format("(%s, %s)", this.first.toString(), this.second.toString());
  }
}
