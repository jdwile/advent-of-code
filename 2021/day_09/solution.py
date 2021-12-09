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

        while len(low_points) > 0:
            to_visit = [low_points.pop()]
            basin = []

            while len(to_visit) > 0:
                current = to_visit.pop()
                if current in visited:
                    continue
                current_height = self.heightmap[current]
                visited.append(current)
                basin.append(current)

                neighbors = self.get_neighbors(*current)
                unvisited_neighbors = [n for n in neighbors if n not in visited]
                valid_neighbors = [
                    n
                    for n in unvisited_neighbors
                    if self.heightmap[n] > current_height and self.heightmap[n] < 9
                ]

                for v in valid_neighbors:
                    to_visit.append(v)
            basins.append(basin)

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
        neighbors = []
        neighbors.append((i - 1, j))
        neighbors.append((i + 1, j))
        neighbors.append((i, j - 1))
        neighbors.append((i, j + 1))
        return [n for n in neighbors if n in self.heightmap]

    def parse_heightmap(self, lines) -> dict:
        heightmap = {}
        for i, line in enumerate(lines):
            for j, height in enumerate(line):
                heightmap[(i, j)] = int(height)
        return heightmap
