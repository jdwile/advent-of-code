from day_13.solution import Solution
import unittest


class Test_Day13Solution(unittest.TestCase):
    name = "day_13"

    def test_part_one_example(self):
        sut = Solution(
            f"C:\\Users\\Jenner\\bench\\advent-of-code\\2021\\tests\\test_input\\{self.name}.txt"
        )
        self.assertEqual(sut.solve_part_one(), 17)

    # def test_part_two_example(self):
    #     sut = Solution(
    #         f"C:\\Users\\Jenner\\bench\\advent-of-code\\2021\\tests\\test_input\\{self.name}.txt"
    #     )
    #     self.assertEqual(sut.solve_part_two(), 195)

    # def test_part_one_solution(self):
    #     sut = Solution()
    #     self.assertEqual(sut.solve_part_one(), 3497)

    # def test_part_two_solution(self):
    #     sut = Solution()
    #     self.assertEqual(sut.solve_part_two(), 232)


if __name__ == "__main__":
    unittest.main()
