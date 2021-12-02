from day_01.solution import Solution as Solution1
from day_02.solution import Solution as Solution2
from utils.aoc import timer

day1 = Solution1()
day2 = Solution2()

days = [day1, day2]


@timer
def run_part_one(day: int):
    print(f"Day {day+1}, Part 1: {days[day].solve_part_one()}")


@timer
def run_part_two(day: int):
    print(f"Day {day+1}, Part 2: {days[day].solve_part_two()}")


for day in range(len(days)):
    run_part_one(day)
    run_part_two(day)
