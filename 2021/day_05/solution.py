from utils.aoc import input_as_lines
import re


class Solution:
    def __init__(
        self,
        filename="C:\\Users\\Jenner\\bench\\advent-of-code\\2021\\day_05\\input.txt",
    ) -> None:
        self.lines = input_as_lines(filename)

    def solve_part_one(self) -> str:
        vents = self.parse_vents()
        seafloor = {}

        for vent in vents:
            x1, y1, x2, y2 = vent
            if not (x1 == x2 or y1 == y2):
                continue  # check if straight line

            for x in range(min(x1, x2), max(x1, x2) + 1):
                for y in range(min(y1, y2), max(y1, y2) + 1):
                    if not (x, y) in seafloor.keys():
                        seafloor[(x, y)] = 0
                    seafloor[(x, y)] = seafloor[(x, y)] + 1

        return len(list(filter(lambda x: x > 1, seafloor.values())))

    def solve_part_two(self) -> str:
        vents = self.parse_vents()
        seafloor = {}

        for vent in vents:
            x1, y1, x2, y2 = vent
            if x1 == x2 or y1 == y2:  # straight line
                for x in range(min(x1, x2), max(x1, x2) + 1):
                    for y in range(min(y1, y2), max(y1, y2) + 1):
                        if not (x, y) in seafloor.keys():
                            seafloor[(x, y)] = 0
                        seafloor[(x, y)] = seafloor[(x, y)] + 1
            else:  # diagonal line
                x, y = x1, y1
                dx = 1 if x2 > x1 else -1
                dy = 1 if y2 > y1 else -1

                while x != x2 and y != y2:
                    if not (x, y) in seafloor.keys():
                        seafloor[(x, y)] = 0
                    seafloor[(x, y)] = seafloor[(x, y)] + 1
                    x += dx
                    y += dy

                if not (x, y) in seafloor.keys():
                    seafloor[(x, y)] = 0
                seafloor[(x, y)] = seafloor[(x, y)] + 1

        return len(list(filter(lambda x: x > 1, seafloor.values())))

    def parse_vents(self) -> list:
        vents = []
        for line in self.lines:
            m = re.match(r"(\d+),(\d+) -> (\d+),(\d+)", line)
            vents.append([int(x) for x in m.groups()])
        return vents
