from utils.aoc import input_as_lines


class Solution:
    def __init__(
        self,
        filename="C:\\Users\\Jenner\\bench\\advent-of-code\\2021\\day_04\\input.txt",
    ) -> None:
        self.lines = input_as_lines(filename)

    def solve_part_one(self) -> str:
        draws = [int(x) for x in self.lines[0].split(",")]
        boards = self.generate_boards()

        for draw in draws:
            for board in boards:
                set = False
                for row in range(len(board)):
                    for square in range(len(board[row])):
                        if board[row][square][0] == draw:
                            board[row][square] = (draw, True)
                            set = True
                            break
                    if set:
                        break

                if set:
                    if self.has_bingo(board):
                        return self.score_board(board) * draw

    def solve_part_two(self) -> str:
        draws = [int(x) for x in self.lines[0].split(",")]
        boards = self.generate_boards()

        last_board_won = None
        last_draw_won = None
        boards_won = []

        for draw in draws:
            for b in range(len(boards)):
                set = False
                if b in boards_won:
                    continue
                for row in range(len(boards[b])):
                    for square in range(len(boards[b][row])):
                        if boards[b][row][square][0] == draw:
                            boards[b][row][square] = (draw, True)
                            set = True
                            break
                    if set:
                        break

                if set:
                    if self.has_bingo(boards[b]):
                        boards_won.append(b)
                        last_board_won = b
                        last_draw_won = draw

        return self.score_board(boards[last_board_won]) * last_draw_won

    def generate_boards(self) -> list:
        boards = []
        i = 1
        while i < len(self.lines):
            i += 1
            board = []
            for j in range(5):
                board.append([(int(x), False) for x in self.lines[i].split()])
                i += 1
            boards.append(board)
        return boards

    def has_bingo(self, board) -> bool:
        if any([all(x[1] for x in row) for row in board]):
            return True

        if any(
            [all(x[1] for x in [row[i] for row in board]) for i in range(len(board))]
        ):
            return True

    def score_board(self, board) -> int:
        return sum(
            [
                sum([x[0] for x in list(filter(lambda space: not space[1], board[i]))])
                for i in range(len(board))
            ]
        )
