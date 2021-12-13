from typing import Tuple
from utils.aoc import input_as_lines


class Solution:
    def __init__(
        self,
        filename="C:\\Users\\Jenner\\bench\\advent-of-code\\2021\\day_13\\input.txt",
    ) -> None:
        self.paper, self.folds = self.parse_input(input_as_lines(filename))

    def solve_part_one(self) -> str:
        paper = self.fold_paper(self.paper.copy(), self.folds, once=True)
        return len(paper)

    def solve_part_two(self) -> str:
        paper = self.fold_paper(self.paper.copy(), self.folds)

        max_x = max([d[0] for d in paper]) + 1
        max_y = max([d[1] for d in paper]) + 1

        code = ""
        for i in range(max_x):
            line = "\n"
            for j in range(max_y):
                line += "#" if (i, j) in paper else " "
            code += line
        return code

    def fold_paper(self, paper, folds, once=False) -> set:
        for axis, v in folds:
            if axis == "y":
                folded_dots = [dot for dot in paper if dot[0] > v]
                for dot in folded_dots:
                    paper.remove(dot)
                    paper.add((v - (dot[0] - v), dot[1]))
            else:
                folded_dots = [dot for dot in paper if dot[1] > v]
                for dot in folded_dots:
                    paper.remove(dot)
                    paper.add((dot[0], v - (dot[1] - v)))
            if once:
                break
        return paper

    def parse_input(self, lines) -> Tuple:
        paper, folds = set(), []
        for line in lines:
            if "," in line:
                y, x = [int(n) for n in line.split(",")]
                paper.add((x, y))
            elif "=" in line:
                axis, n = line.split("=")
                axis = axis.split(" ")[2]
                folds.append((axis, int(n)))
        return paper, folds
