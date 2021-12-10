from utils.aoc import input_as_lines
import math


class Solution:
    def __init__(
        self,
        filename="C:\\Users\\Jenner\\bench\\advent-of-code\\2021\\day_10\\input.txt",
    ) -> None:
        self.lines = input_as_lines(filename)
        self.error_points_lookup = {")": 3, "]": 57, "}": 1197, ">": 25137}
        self.edit_points_lookup = {")": 1, "]": 2, "}": 3, ">": 4}
        self.open_close_mapping = {"(": ")", "[": "]", "{": "}", "<": ">"}

    def solve_part_one(self) -> str:
        return sum([self.calculate_syntax_error_score(line) for line in self.lines])

    def solve_part_two(self) -> str:
        incomplete_lines = [
            line for line in self.lines if self.calculate_syntax_error_score(line) == 0
        ]
        closing_lists = [self.get_closing_chars(line) for line in incomplete_lines]
        closing_scores = sorted(
            [self.calculate_edit_score(lst) for lst in closing_lists]
        )
        return closing_scores[len(closing_scores) // 2]

    def calculate_syntax_error_score(self, line) -> int:
        closers = []
        for c in line:
            if c in self.open_close_mapping.keys():
                closers.append(self.open_close_mapping[c])
            elif c != closers.pop():
                return self.error_points_lookup[c]
        return 0

    def get_closing_chars(self, line) -> list:
        closers = []
        for c in line:
            if c in self.open_close_mapping.keys():
                closers.append(self.open_close_mapping[c])
                continue
            closers.pop()
        return list(reversed(closers))

    def calculate_edit_score(self, chars) -> int:
        total = 0
        for char in chars:
            total *= 5
            total += self.edit_points_lookup[char]
        return total
