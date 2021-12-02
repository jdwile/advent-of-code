from day_02.solution import Solution
import unittest


class Test_Day2Solution(unittest.TestCase):
    name = "day_02"

    def test_part_one_example(self):
        sut = Solution(
            f"C:\\Users\\Jenner\\bench\\advent-of-code\\2021\\tests\\test_input\\{self.name}.txt"
        )
        self.assertEqual(sut.solve_part_one(), 150)

    def test_part_two_example(self):
        sut = Solution(
            f"C:\\Users\\Jenner\\bench\\advent-of-code\\2021\\tests\\test_input\\{self.name}.txt"
        )
        self.assertEqual(sut.solve_part_two(), 900)

    def test_part_one_solution(self):
        sut = Solution()
        self.assertEqual(sut.solve_part_one(), 1924923)

    def test_part_two_solution(self):
        sut = Solution()
        self.assertEqual(sut.solve_part_two(), 1982495697)


if __name__ == "__main__":
    unittest.main()
