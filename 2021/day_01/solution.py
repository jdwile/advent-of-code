from utils.aoc import input_as_ints


class Solution:
    def __init__(
        self,
        filename="C:\\Users\\Jenner\\bench\\advent-of-code\\2021\\day_01\\input.txt",
    ) -> None:
        self.lines = input_as_ints(filename)

    def solve_part_one(self) -> str:
        count = 0
        for i in range(1, len(self.lines)):
            count += self.lines[i] > self.lines[i - 1]

        return count

    def solve_part_two(self) -> str:
        count = 0
        for i in range(2, len(self.lines)):
            count += sum(self.lines[i - 3 : i]) > sum(self.lines[i - 3 - 1 : i - 1])

        return count
