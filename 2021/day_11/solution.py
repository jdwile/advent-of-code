from utils.aoc import input_as_lines


class Solution:
    def __init__(
        self,
        filename="C:\\Users\\Jenner\\bench\\advent-of-code\\2021\\day_11\\input.txt",
    ) -> None:
        self.energy_levels = self.parse_energy_levels(input_as_lines(filename))
        self.neighbor_map = {}

    def solve_part_one(self) -> str:
        flashes = 0
        energy_levels = self.energy_levels.copy()

        for _ in range(100):
            octopi_flashed = set()
            octopi_to_increment = list(energy_levels.keys())

            for octopus in octopi_to_increment:
                energy_levels[octopus] += 1
                if energy_levels[octopus] > 9:
                    if octopus in octopi_flashed:
                        continue
                    octopi_flashed.add(octopus)
                    for neighbor in self.get_neighbors(*octopus):
                        octopi_to_increment.append(neighbor)

            flashes += len(octopi_flashed)
            for octopus in octopi_flashed:
                energy_levels[octopus] = 0

        return flashes

    def solve_part_two(self) -> str:
        flashes = 0
        energy_levels = self.energy_levels.copy()
        steps = 0

        while True:
            steps += 1
            octopi_flashed = set()
            octopi_to_increment = list(energy_levels.keys())

            for octopus in octopi_to_increment:
                energy_levels[octopus] += 1
                if energy_levels[octopus] > 9:
                    if octopus in octopi_flashed:
                        continue
                    octopi_flashed.add(octopus)
                    for neighbor in self.get_neighbors(*octopus):
                        octopi_to_increment.append(neighbor)

            flashes += len(octopi_flashed)
            for octopus in octopi_flashed:
                energy_levels[octopus] = 0

            if len(octopi_flashed) == len(energy_levels.keys()):
                return steps

    def get_neighbors(self, i, j) -> list:
        if (i, j) in self.neighbor_map:
            return self.neighbor_map[(i, j)]
        dx = dy = {-1, 0, 1}
        neighbors = []
        for x in dx:
            for y in dy:
                neighbors.append((i + x, j + y))
        self.neighbor_map[(i, j)] = [
            n for n in neighbors if n in self.energy_levels and n != (i, j)
        ]
        return self.neighbor_map[(i, j)]

    def parse_energy_levels(self, lines) -> dict:
        energy_levels = {}
        for i, line in enumerate(lines):
            for j, energy in enumerate(line):
                energy_levels[(i, j)] = int(energy)
        return energy_levels
