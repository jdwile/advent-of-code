from utils.aoc import input_as_lines
from math import ceil


class Solution:
    def __init__(
        self,
        filename="C:\\Users\\Jenner\\bench\\advent-of-code\\2021\\day_18\\input.txt",
    ) -> None:
        self.homework = [eval(line) for line in input_as_lines(filename)]

    def solve_part_one(self) -> str:
        number = self.homework[0]
        for i in range(1, len(self.homework)):
            number = self.add(number, self.homework[i])

        return self.magnitude(number)

    def solve_part_two(self) -> str:
        largest_magnitude = 0
        for i in range(len(self.homework)):
            for j in range(len(self.homework)):
                if i == j:
                    continue

                largest_magnitude = max(
                    self.magnitude(self.add(self.homework[i], self.homework[j])),
                    largest_magnitude,
                )

        return largest_magnitude

    def add_left(self, number, n):
        if n is None:
            return number

        if type(number) is int:
            return number + n

        return [self.add_left(number[0], n), number[1]]

    def add_right(self, number, n):
        if n is None:
            return number

        if type(number) is int:
            return number + n

        return [number[0], self.add_right(number[1], n)]

    def explode(self, number, depth=0):
        if type(number) is int:
            return False, None, number, None

        if depth == 4:
            return True, number[0], 0, number[1]

        left_num, right_num = number
        reduced, previous_list, left_num, next_list = self.explode(left_num, depth + 1)
        if reduced:
            return (
                True,
                previous_list,
                [left_num, self.add_left(right_num, next_list)],
                None,
            )

        reduced, previous_list, right_num, next_list = self.explode(
            right_num, depth + 1
        )
        if reduced:
            return (
                True,
                None,
                [self.add_right(left_num, previous_list), right_num],
                next_list,
            )

        return False, None, number, None

    def split(self, number):
        if type(number) is int:
            if number >= 10:
                return True, [number // 2, ceil(number / 2)]

            return False, number

        left, right = number
        reduced, left = self.split(left)
        if reduced:
            return True, [left, right]

        reduced, right = self.split(right)
        return reduced, [left, right]

    def add(self, left, right):
        numbers = [left, right]
        while True:
            reduced, _, numbers, _ = self.explode(numbers)
            if reduced:
                continue

            reduced, numbers = self.split(numbers)
            if not reduced:
                break

        return numbers

    def magnitude(self, number):
        if type(number) is int:
            return number
        return 3 * self.magnitude(number[0]) + 2 * self.magnitude(number[1])
