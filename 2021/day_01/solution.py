from utils.aoc import input_as_ints
import itertools
from collections import deque


class Solution:
    def __init__(
        self,
        filename="C:\\Users\\Jenner\\bench\\advent-of-code\\2021\\day_01\\input.txt",
    ) -> None:
        self.lines = input_as_ints(filename)

    def solve_part_one(self) -> str:
        rot = deque(self.lines)
        rot.rotate(1)
        count = 0
        for a, b in zip(self.lines[1:], list(rot)[1:]):
            count += a > b

        return count

    def solve_part_two(self) -> str:
        rot_1, rot_2 = deque(self.lines), deque(self.lines)
        rot_1.rotate(1)
        rot_2.rotate(2)
        count = 0
        for i in range(3, len(self.lines)):
            count += (
                self.lines[i] + rot_1[i] + rot_2[i]
                > self.lines[i - 1] + rot_1[i - 1] + rot_2[i - 1]
            )

        return count
