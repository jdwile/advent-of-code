import re
from collections import Counter
from utils.aoc import input_as_lines


class Solution:
    def __init__(
        self,
        filename="C:\\Users\\Jenner\\bench\\advent-of-code\\2021\\day_22\\input.txt",
    ) -> None:
        self.input = input_as_lines(filename)

    def solve_part_one(self) -> str:
        return self.count_cubes(init=True)

    def solve_part_two(self) -> str:
        return self.count_cubes()

    def count_cubes(self, init=False):
        cubes = Counter()

        for line in self.input:
            sign = 1 if line.split(" ")[0] == "on" else 0
            m = re.search(
                r"x=(-?\d+)\.\.(-?\d+),y=(-?\d+)\.\.(-?\d+),z=(-?\d+)\.\.(-?\d+)", line
            )
            new_x0, new_x1, new_y0, new_y1, new_z0, new_z1 = map(int, m.groups())

            if init and (
                any([v < -50 for v in [new_x0, new_y0, new_z0]])
                or any([v > 50 for v in [new_x1, new_y1, new_z1]])
            ):
                continue

            update = Counter()
            for (
                cur_x0,
                cur_x1,
                cur_y0,
                cur_y1,
                cur_z0,
                cur_z1,
            ), cur_sign in cubes.items():
                inter_x0 = max(new_x0, cur_x0)
                inter_x1 = min(new_x1, cur_x1)
                inter_y0 = max(new_y0, cur_y0)
                inter_y1 = min(new_y1, cur_y1)
                inter_z0 = max(new_z0, cur_z0)
                inter_z1 = min(new_z1, cur_z1)
                if (
                    inter_x0 <= inter_x1
                    and inter_y0 <= inter_y1
                    and inter_z0 <= inter_z1
                ):
                    update[
                        (inter_x0, inter_x1, inter_y0, inter_y1, inter_z0, inter_z1)
                    ] -= cur_sign

            if sign > 0:
                update[(new_x0, new_x1, new_y0, new_y1, new_z0, new_z1)] += sign

            cubes.update(update)

        return sum(
            (x1 - x0 + 1) * (y1 - y0 + 1) * (z1 - z0 + 1) * sign
            for (x0, x1, y0, y1, z0, z1), sign in cubes.items()
        )
