from utils.aoc import input_as_string


class Solution:
    def __init__(
        self,
        filename="C:\\Users\\Jenner\\bench\\advent-of-code\\2021\\day_06\\input.txt",
    ) -> None:
        self.input = [int(x) for x in input_as_string(filename).split(",")]

    def solve_part_one(self) -> str:
        fish = self.simulate_fish(80)
        return sum(fish.values())

    def solve_part_two(self) -> str:
        fish = self.simulate_fish(256)
        return sum(fish.values())

    def simulate_fish(self, days) -> dict:
        fish = dict.fromkeys(range(9), 0)
        for f in self.input:
            fish[f] += 1

        for _ in range(days):
            new_fish = dict.fromkeys(range(9), 0)
            for i in range(1, 9):
                new_fish[i - 1] += fish[i]
            new_fish[6] += fish[0]
            new_fish[8] += fish[0]

            fish = new_fish
        return fish
