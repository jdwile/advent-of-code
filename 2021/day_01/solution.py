from utils.aoc import input_as_ints
import itertools


class Solution:
    def __init__(
        self,
        filename="C:\\Users\\Jenner\\bench\\advent-of-code\\2021\\day_01\\input.txt",
    ) -> None:
        self.lines = input_as_ints(filename)

    def solve_part_one(self) -> str:
        for a, b in itertools.product(self.lines, self.lines):
            if a + b == 2020:
                return a * b

    def solve_part_two(self) -> str:
        for a, b, c in itertools.product(self.lines, self.lines, self.lines):
            if a + b + c == 2020:
                return a * b * c
