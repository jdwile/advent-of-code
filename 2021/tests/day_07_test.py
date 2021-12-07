from day_07.solution import Solution
import unittest


class Test_Day7Solution(unittest.TestCase):
    name = "day_07"

    def test_part_one_example(self):
        sut = Solution(
            f"C:\\Users\\Jenner\\bench\\advent-of-code\\2021\\tests\\test_input\\{self.name}.txt"
        )
        self.assertEqual(sut.solve_part_one(), 37)

    def test_part_two_example(self):
        sut = Solution(
            f"C:\\Users\\Jenner\\bench\\advent-of-code\\2021\\tests\\test_input\\{self.name}.txt"
        )
        self.assertEqual(sut.solve_part_two(), 168)

    def test_part_one_solution(self):
        sut = Solution()
        self.assertEqual(sut.solve_part_one(), 325528)

    def test_part_two_solution(self):
        sut = Solution()
        self.assertEqual(sut.solve_part_two(), 85015836)


if __name__ == "__main__":
    unittest.main()
