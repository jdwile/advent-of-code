from day_01.solution import Solution
import unittest

class Test_Day1Solution(unittest.TestCase):
    name = 'day_01'

    def test_part_one_example(self):
        sut = Solution(f'C:\\Users\\Jenner\\bench\\advent-of-code\\2021\\tests\\{self.name}\\test_input.txt')
        self.assertEqual(sut.solve_part_one(), 514579)

    def test_part_two_example(self):
        sut = Solution(f'C:\\Users\\Jenner\\bench\\advent-of-code\\2021\\tests\\{self.name}\\test_input.txt')
        self.assertEqual(sut.solve_part_two(), 241861950)

    def test_part_one_solution(self):
        sut = Solution()
        self.assertEqual(sut.solve_part_one(), 468051)

    def test_part_two_solution(self):
        sut = Solution()
        self.assertEqual(sut.solve_part_two(), 272611658)

if __name__ == '__main__':
    unittest.main()