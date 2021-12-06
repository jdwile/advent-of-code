from day_06.solution import Solution
import unittest


class Test_Day6Solution(unittest.TestCase):
    name = "day_06"

    def test_part_one_example(self):
        sut = Solution(
            f"C:\\Users\\Jenner\\bench\\advent-of-code\\2021\\tests\\test_input\\{self.name}.txt"
        )
        self.assertEqual(sut.solve_part_one(), 5934)

    def test_part_two_example(self):
        sut = Solution(
            f"C:\\Users\\Jenner\\bench\\advent-of-code\\2021\\tests\\test_input\\{self.name}.txt"
        )
        self.assertEqual(sut.solve_part_two(), 26984457539)

    def test_part_one_solution(self):
        sut = Solution()
        self.assertEqual(sut.solve_part_one(), 377263)

    def test_part_two_solution(self):
        sut = Solution()
        self.assertEqual(sut.solve_part_two(), 1695929023803)


if __name__ == "__main__":
    unittest.main()
