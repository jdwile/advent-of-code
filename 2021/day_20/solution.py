from collections import defaultdict
from utils.aoc import input_as_lines


class Solution:
    def __init__(
        self,
        filename="C:\\Users\\Jenner\\bench\\advent-of-code\\2021\\day_20\\input.txt",
    ) -> None:
        self.enhance, self.image = self.parse(input_as_lines(filename))
        self.neighbor_map = {}
        self.default_value = "."

    def solve_part_one(self) -> str:
        image = self.run_enhancement_simulation(2)
        return len([x for x in image.values() if x == "#"])

    def solve_part_two(self) -> str:
        image = self.run_enhancement_simulation(50)
        return len([x for x in image.values() if x == "#"])

    def run_enhancement_simulation(self, n):
        image = self.image.copy()

        coords = image.keys()
        x_range, y_range = [c[0] for c in coords], [c[1] for c in coords]
        x_min, x_max = min(x_range), max(x_range)
        y_min, y_max = min(y_range), max(y_range)

        for _ in range(n):
            x_max += 1
            x_min -= 1
            y_max += 1
            y_min -= 1

            default_value = self.enhance_position(x_max + 100, y_max + 100, image)
            new_image = {}

            for i in range(x_min, x_max + 1):
                for j in range(y_min, y_max + 1):
                    new_image[(i, j)] = self.enhance_position(i, j, image)

            image = new_image
            self.default_value = default_value

        return image

    def enhance_position(self, x, y, image):
        pixel_map = {".": "0", "#": "1"}
        neighbor_values = [
            pixel_map[self.get_pixel(*n, image)] for n in self.get_neighbors(x, y)
        ]
        enhancement_index = int("".join(neighbor_values), 2)
        return self.enhance[enhancement_index]

    def get_pixel(self, i, j, image):
        if (i, j) in image.keys():
            return image[(i, j)]
        return self.default_value

    def get_neighbors(self, i, j) -> list:
        if (i, j) in self.neighbor_map:
            return self.neighbor_map[(i, j)]
        dx = dy = {-1, 0, 1}
        neighbors = []
        for x in dx:
            for y in dy:
                neighbors.append((i + x, j + y))
        self.neighbor_map[(i, j)] = sorted(neighbors)
        return self.neighbor_map[(i, j)]

    def parse(self, lines):
        enhance = lines[0]

        image = {}
        for i, line in enumerate(lines[2:]):
            for j, c in enumerate(line):
                image[(i, j)] = c

        return enhance, image
