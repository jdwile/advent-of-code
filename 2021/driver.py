from day_22.solution import Solution as Solution
from utils.aoc import timer

day = Solution()

solutions = {22: day}


@timer
def run_part_one(day: int):
    print(f"Day {day}, Part 1: {solutions[day].solve_part_one()}")


@timer
def run_part_two(day: int):
    print(f"Day {day}, Part 2: {solutions[day].solve_part_two()}")


for day in solutions:
    run_part_one(day)
    run_part_two(day)
