from utils.aoc import input_as_lines
import math


class Solution:
    def __init__(
        self,
        filename="C:\\Users\\Jenner\\bench\\advent-of-code\\2021\\day_09\\input.txt",
    ) -> None:
        self.heightmap = [[int(x) for x in line] for line in input_as_lines(filename)]

    def solve_part_one(self) -> str:
        low_points = self.get_low_points()
        return sum([self.heightmap[i][j] + 1 for i, j in low_points])

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
                current_height = self.heightmap[current[0]][current[1]]
                visited.append(current)
                basin.append(current)

                neighbors = self.get_neighbors(*current)
                unvisited_neighbors = [n for n in neighbors if n not in visited]
                valid_neighbors = [
                    n
                    for n in unvisited_neighbors
                    if self.heightmap[n[0]][n[1]] > current_height
                    and self.heightmap[n[0]][n[1]] < 9
                ]

                for v in valid_neighbors:
                    to_visit.append(v)
            basins.append(basin)

        return math.prod(sorted([len(basin) for basin in basins], reverse=True)[0:3])

    def get_low_points(self) -> list:
        low_points = []
        for i, row in enumerate(self.heightmap):
            for j, height in enumerate(row):
                neighbors = self.get_neighbors(i, j)
                if all([height < self.heightmap[ni][nj] for ni, nj in neighbors]):
                    low_points.append((i, j))
        return low_points

    def get_neighbors(self, i, j) -> list:
        neighbors = []
        if i - 1 >= 0:
            neighbors.append((i - 1, j))
        if i + 1 < len(self.heightmap):
            neighbors.append((i + 1, j))
        if j - 1 >= 0:
            neighbors.append((i, j - 1))
        if j + 1 < len(self.heightmap[i]):
            neighbors.append((i, j + 1))
        return neighbors
