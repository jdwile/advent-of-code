from utils.aoc import input_as_lines
from collections import defaultdict


class Solution:
    def __init__(
        self,
        filename="C:\\Users\\Jenner\\bench\\advent-of-code\\2021\\day_12\\input.txt",
    ) -> None:
        self.cave_system = self.parse_caves(input_as_lines(filename))

    def solve_part_one(self) -> str:
        num_paths = self.traverse_cave("start", set())
        return num_paths

    def solve_part_two(self) -> str:
        num_paths = self.traverse_cave("start", set(), visitation_exception=True)
        return num_paths

    def traverse_cave(self, cave, visited, visitation_exception=False) -> int:
        if cave[0].islower():
            visited.add(cave)

        if cave == "end":
            return 1

        num_paths = 0
        for next in self.cave_system[cave]:
            if next == "start":
                continue

            if next.isupper() or next.islower() and next not in visited:
                num_paths += self.traverse_cave(
                    next, visited.copy(), visitation_exception
                )
            elif next.islower() and next in visited and visitation_exception:
                num_paths += self.traverse_cave(next, visited.copy(), False)
        return num_paths

    def parse_caves(self, lines) -> dict:
        caves = defaultdict(list)
        for line in lines:
            start, end = line.split("-")
            caves[start].append(end)
            caves[end].append(start)
        return caves
