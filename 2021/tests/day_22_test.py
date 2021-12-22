from day_22.solution import Solution
import unittest


class Test_Day22Solution(unittest.TestCase):
    name = "day_22"

    def test_part_one_example(self):
        sut = Solution(
            f"C:\\Users\\Jenner\\bench\\advent-of-code\\2021\\tests\\test_input\\{self.name}.txt"
        )
        self.assertEqual(sut.solve_part_one(), 474140)

    def test_part_two_example(self):
        sut = Solution(
            f"C:\\Users\\Jenner\\bench\\advent-of-code\\2021\\tests\\test_input\\{self.name}.txt"
        )
        self.assertEqual(sut.solve_part_two(), 2758514936282235)

    def test_part_one_solution(self):
        sut = Solution()
        self.assertEqual(sut.solve_part_one(), 570915)

    def test_part_two_solution(self):
        sut = Solution()
        self.assertEqual(sut.solve_part_two(), 1268313839428137)


if __name__ == "__main__":
    unittest.main()
