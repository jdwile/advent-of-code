from utils.aoc import input_as_lines
from itertools import cycle


class Solution:
    def __init__(
        self,
        filename="C:\\Users\\Jenner\\bench\\advent-of-code\\2021\\day_21\\input.txt",
    ) -> None:
        self.starting_positions = self.parse_starting_positions(
            input_as_lines(filename)
        )
        self.win_states = {}

    def solve_part_one(self) -> str:
        scores = [0, 0]
        positions = self.starting_positions.copy()
        die = cycle(list(range(1, 101)))
        turn = 0
        rolls = 0
        while all([score < 1000 for score in scores]):
            rolled_dice = [next(die), next(die), next(die)]
            positions[turn] = (positions[turn] + sum(rolled_dice)) % 10
            if positions[turn] == 0:
                positions[turn] += 10

            scores[turn] += positions[turn]

            turn = (turn + 1) % 2
            rolls += 3

        return min(scores) * rolls

    def solve_part_two(self) -> str:
        return max(self.count_wins(*self.starting_positions, 0, 0))

    def count_wins(self, pos_1, pos_2, score_1=0, score_2=0):
        if score_1 >= 21:
            return (1, 0)
        if score_2 >= 21:
            return (0, 1)

        if (pos_1, pos_2, score_1, score_2) in self.win_states:
            return self.win_states[(pos_1, pos_2, score_1, score_2)]

        wins = (0, 0)

        for die_1 in [1, 2, 3]:
            for die_2 in [1, 2, 3]:
                for die_3 in [1, 2, 3]:
                    new_pos_1 = (pos_1 + die_1 + die_2 + die_3) % 10
                    if new_pos_1 == 0:
                        new_pos_1 = 10
                    new_score_1 = score_1 + new_pos_1

                    u = self.count_wins(pos_2, new_pos_1, score_2, new_score_1)
                    wins = (wins[0] + u[1], wins[1] + u[0])
        self.win_states[(pos_1, pos_2, score_1, score_2)] = wins
        return wins

    def parse_starting_positions(self, lines):
        return [int(lines[0].split(" ")[-1]), int(lines[1].split(" ")[-1])]
