from utils.aoc import input_as_lines
import math


class Solution:
    def __init__(
        self,
        filename="C:\\Users\\Jenner\\bench\\advent-of-code\\2021\\day_09\\input.txt",
    ) -> None:
        self.heightmap = self.parse_heightmap(input_as_lines(filename))

    def solve_part_one(self) -> str:
        low_points = self.get_low_points()
        return sum([self.heightmap[point] + 1 for point in low_points])

    def solve_part_two(self) -> str:
        low_points = self.get_low_points()
        visited = []
        basins = []

        def generate_basin(position):
            if position in visited:
                return []

            visited.append(position)

            unvisited_neighbors = [
                n for n in self.get_neighbors(*position) if n not in visited
            ]
            valid_neighbors = [
                n
                for n in unvisited_neighbors
                if self.heightmap[n] > self.heightmap[position]
                and self.heightmap[n] < 9
            ]

            basin = [position]
            for v in valid_neighbors:
                for r in generate_basin(v):
                    basin.append(r)

            return basin

        for starting_point in low_points:
            basins.append(generate_basin(starting_point))

        return math.prod(sorted([len(basin) for basin in basins], reverse=True)[0:3])

    def get_low_points(self) -> list:
        low_points = []
        for pos in self.heightmap.keys():
            height = self.heightmap[pos]
            neighbors = self.get_neighbors(*pos)
            if all([height < self.heightmap[n] for n in neighbors]):
                low_points.append(pos)
        return low_points

    def get_neighbors(self, i, j) -> list:
        neighbors = [(i - 1, j), (i + 1, j), (i, j - 1), (i, j + 1)]
        return [n for n in neighbors if n in self.heightmap]

    def parse_heightmap(self, lines) -> dict:
        heightmap = {}
        for i, line in enumerate(lines):
            for j, height in enumerate(line):
                heightmap[(i, j)] = int(height)
        return heightmap
