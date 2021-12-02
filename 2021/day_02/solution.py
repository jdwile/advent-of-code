from utils.aoc import input_as_lines


class Solution:
    def __init__(
        self,
        filename="C:\\Users\\Jenner\\bench\\advent-of-code\\2021\\day_02\\input.txt",
    ) -> None:
        self.instructions = [(item.split(' ')[0], int(item.split(' ')[1])) for item in input_as_lines(filename)]

    def solve_part_one(self) -> str:
        x, y = 0, 0
        for dir, mag in self.instructions:
            match dir:
                case 'forward':
                    x += mag
                case 'up':
                    y -= mag
                case 'down':
                    y += mag

        return x * y

    def solve_part_two(self) -> str:
        x, y, aim = 0, 0, 0
        for dir, mag in self.instructions:
            match dir:
                case 'forward':
                    x += mag
                    y += aim * mag
                case 'up':
                    aim -= mag
                case 'down':
                    aim += mag

        return x * y
