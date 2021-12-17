from utils.aoc import input_as_string
import re


class Solution:
    def __init__(
        self,
        filename="C:\\Users\\Jenner\\bench\\advent-of-code\\2021\\day_17\\input.txt",
    ) -> None:
        self.x_min, self.x_max, self.y_min, self.y_max = self.parse_target(
            input_as_string(filename)
        )
        self.x_range = range(self.x_min, self.x_max + 1)
        self.y_range = range(self.y_min, self.y_max + 1)

    def solve_part_one(self) -> str:
        highest = 0
        for x in range(0, self.x_max + 1):
            for y in range(0, -self.y_min):
                max_height, in_target = self.simulate_probe(x, y)
                if in_target:
                    if max_height > highest:
                        highest = max_height
        return highest

    def solve_part_two(self) -> str:
        count, n = 0, 200
        for x in range(0, self.x_max + 1):
            for y in range(self.y_min, -self.y_min):
                _, in_target = self.simulate_probe(x, y)
                if in_target:
                    count += 1
        return count

    def simulate_probe(self, x_v, y_v):
        max_height = 0
        dx = -1 if x_v > 0 else 1
        x, y = 0, 0

        def missed_target():
            return x_v == 0 and (x < self.x_min or x > self.x_max)

        def impossible_to_hit():
            return y < self.y_min and y_v < 0 or x > self.x_max and x_v > 0

        while not (missed_target() or impossible_to_hit()):
            x += x_v
            y += y_v

            if not x_v == 0:
                x_v += dx
            y_v -= 1

            max_height = max(max_height, y)

            if x in self.x_range and y in self.y_range:
                return max_height, True

        return max_height, False

    def parse_target(self, input):
        m = re.search(r"x=(-?\d+)\.\.(-?\d+), y=(-?\d+)\.\.(-?\d+)", input)
        return [int(x) for x in m.groups()]
