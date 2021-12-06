from utils.aoc import input_as_string


class Solution:
    def __init__(
        self,
        filename="C:\\Users\\Jenner\\bench\\advent-of-code\\2021\\day_06\\input.txt",
    ) -> None:
        self.input = input_as_string(filename).split(",")

    def solve_part_one(self) -> str:
        fish = self.simulate_fish(80)
        return sum(fish)

    def solve_part_two(self) -> str:
        fish = self.simulate_fish(256)
        return sum(fish)

    def simulate_fish(self, days) -> dict:
        fish = [*map(self.input.count, "012345678")]  # this is some dank shit
        print(fish)

        for _ in range(days):
            new_fish = [0] * 9
            for i in range(1, 9):
                new_fish[i - 1] += fish[i]
            new_fish[6] += fish[0]
            new_fish[8] += fish[0]

            fish = new_fish
        return fish
