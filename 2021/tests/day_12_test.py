from day_12.solution import Solution
import unittest


class Test_Day12Solution(unittest.TestCase):
    name = "day_12"

    def test_part_one_example(self):
        sut = Solution(
            f"C:\\Users\\Jenner\\bench\\advent-of-code\\2021\\tests\\test_input\\{self.name}.txt"
        )
        self.assertEqual(sut.solve_part_one(), 10)

    # def test_part_two_example(self):
    #     sut = Solution(
    #         f"C:\\Users\\Jenner\\bench\\advent-of-code\\2021\\tests\\test_input\\{self.name}.txt"
    #     )
    #     self.assertEqual(sut.solve_part_two(), 195)

    def test_part_one_solution(self):
        sut = Solution()
        self.assertEqual(sut.solve_part_one(), 3497)

    # def test_part_two_solution(self):
    #     sut = Solution()
    #     self.assertEqual(sut.solve_part_two(), 232)


if __name__ == "__main__":
    unittest.main()
