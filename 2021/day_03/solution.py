from utils.aoc import input_as_lines
from collections import Counter


class Solution:
    def __init__(
        self,
        filename="C:\\Users\\Jenner\\bench\\advent-of-code\\2021\\day_03\\input.txt",
    ) -> None:
        self.lines = [list(map(int, line)) for line in input_as_lines(filename)]

    def solve_part_one(self) -> str:
        lines_by_index = list(map(list, zip(*self.lines)))
        index_counts = map(Counter, lines_by_index)

        gamma, epsilon = "", ""
        for line in index_counts:
            gamma += str(self.get_filter_key(line))
            epsilon += str(self.get_filter_key(line, use_min=True))

        return int(gamma, 2) * int(epsilon, 2)

    def solve_part_two(self) -> str:
        num_len = len(self.lines[0])

        oxygen_rating = self.get_air_quality_ratings(self.lines, num_len)
        co2_rating = self.get_air_quality_ratings(self.lines, num_len, is_co2=True)

        return int(oxygen_rating, 2) * int(co2_rating, 2)

    def get_air_quality_ratings(self, ratings, num_len, is_co2=False) -> str:
        for i in range(num_len):
            index_count = Counter(line[i] for line in ratings)
            filter_key = self.get_filter_key(index_count, is_co2)
            ratings = list(filter(lambda n: n[i] == filter_key, ratings))

        return "".join(map(str, ratings[0]))

    def get_filter_key(self, index_count, use_min=False) -> str:
        if index_count.get(0) == index_count.get(1):
            if use_min:
                return 0
            return 1

        if use_min:
            return min(index_count, key=index_count.get)

        return max(index_count, key=index_count.get)
