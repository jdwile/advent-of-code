from day_04.solution import Solution
import unittest


class Test_Day4Solution(unittest.TestCase):
    name = "day_04"

    def test_part_one_example(self):
        sut = Solution(
            f"C:\\Users\\Jenner\\bench\\advent-of-code\\2021\\tests\\test_input\\{self.name}.txt"
        )
        self.assertEqual(sut.solve_part_one(), 4512)

    def test_part_two_example(self):
        sut = Solution(
            f"C:\\Users\\Jenner\\bench\\advent-of-code\\2021\\tests\\test_input\\{self.name}.txt"
        )
        self.assertEqual(sut.solve_part_two(), 1924)

    def test_part_one_solution(self):
        sut = Solution()
        self.assertEqual(sut.solve_part_one(), 49860)

    def test_part_two_solution(self):
        sut = Solution()
        self.assertEqual(sut.solve_part_two(), 24628)


if __name__ == "__main__":
    unittest.main()
