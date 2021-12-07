from utils.aoc import input_as_string
from statistics import median


class Solution:
    def __init__(
        self,
        filename="C:\\Users\\Jenner\\bench\\advent-of-code\\2021\\day_07\\input.txt",
    ) -> None:
        self.crabs = [int(x) for x in input_as_string(filename).split(",")]

    def solve_part_one(self) -> str:
        alignment = median(self.crabs)
        cost = sum([*map(lambda crab: abs(alignment - crab), self.crabs)])
        return int(cost)

    def solve_part_two(self) -> str:
        alignment = round(sum(self.crabs) / len(self.crabs))
        cost = self.align_crabs(alignment)

        new_alignment = alignment
        while new_alignment in range(len(self.crabs)):
            new_alignment = new_alignment + 1

            new_cost = self.align_crabs(new_alignment)

            if new_cost > cost:
                break

            cost = new_cost

        new_alignment = alignment
        while new_alignment in range(len(self.crabs)):
            new_alignment = new_alignment - 1

            new_cost = self.align_crabs(new_alignment)

            if new_cost > cost:
                break

            cost = new_cost

        return int(cost)

    def align_crabs(self, alignment) -> int:
        return sum(
            [
                *map(
                    lambda crab: sum(range(1, abs(alignment - crab) + 1)),
                    self.crabs,
                )
            ]
        )
