from utils.aoc import input_as_lines


class Solution:
    def __init__(
        self,
        filename="C:\\Users\\Jenner\\bench\\advent-of-code\\2021\\day_15\\input.txt",
    ) -> None:
        self.cave = [[int(x) for x in line] for line in input_as_lines(filename)]

    def solve_part_one(self) -> str:
        return self.find_least_risky_path(self.cave)

    def solve_part_two(self) -> str:
        return self.find_least_risky_path(self.magnify_cave(self.cave))

    def find_least_risky_path(self, cave) -> int:
        costs = {}
        q = [(0, 0, 0)]

        while len(q) > 0:
            cost, x, y = q.pop()

            if x == len(cave) - 1 and y == len(cave[x]) - 1:
                break

            for nx, ny in self.get_neighbors(x, y, cave):
                new_cost = cost + cave[nx][ny]
                if (nx, ny) in costs and costs[(nx, ny)] <= new_cost:
                    continue
                costs[(nx, ny)] = new_cost
                q.append((new_cost, nx, ny))

            q = sorted(q, reverse=True)

        return costs[(len(cave) - 1, len(cave[0]) - 1)]

    def magnify_cave(self, cave, magnitude=5) -> list:
        def increment(val, m):
            for _ in range(m):
                val = val % 9 + 1
            return val

        new_cave = []

        for mx in range(magnitude):
            for i in range(len(cave)):
                row = []
                for j in range(len(cave[i])):
                    row.append(increment(cave[i][j], mx))
                new_cave.append(row)
            for i in range(len(cave)):
                for my in range(magnitude - 1):
                    for j in range(len(cave[i % magnitude])):
                        new_cave[i + mx * len(cave[i])].append(
                            new_cave[i + mx * len(cave[i])][
                                j + len(cave[i % magnitude]) * my
                            ]
                            % 9
                            + 1
                        )
        return new_cave

    def get_neighbors(self, i, j, cave) -> list:
        dx = dy = {-1, 0, 1}
        neighbors = []
        for x in dx:
            for y in dy:
                if (
                    abs(x) + abs(y) == 1
                    and i + x >= 0
                    and i + x < len(cave)
                    and j + y >= 0
                    and j + y < len(cave[x])
                ):
                    neighbors.append((i + x, j + y))
        return neighbors
