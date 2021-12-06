from day_06.solution import Solution as Solution
from utils.aoc import timer

day6 = Solution()

solutions = {6: day6}


@timer
def run_part_one(day: int):
    print(f"Day {day+1}, Part 1: {solutions[day].solve_part_one()}")


@timer
def run_part_two(day: int):
    print(f"Day {day+1}, Part 2: {solutions[day].solve_part_two()}")


for day in solutions:
    run_part_one(day)
    run_part_two(day)
