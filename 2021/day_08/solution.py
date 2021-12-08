from utils.aoc import input_as_lines, input_as_string
import re


class Solution:
    def __init__(
        self,
        filename="C:\\Users\\Jenner\\bench\\advent-of-code\\2021\\day_08\\input.txt",
    ) -> None:
        self.signals, self.outputs = self.parse_signals(input_as_lines(filename))
        self.known_signals = {}

    def solve_part_one(self) -> str:
        relevant_counts = [2, 3, 4, 7]
        return sum(
            [
                len([x for x in output if len(x) in relevant_counts])
                for output in self.outputs
            ]
        )

    def solve_part_two(self) -> str:
        result = 0
        for i in range(len(self.signals)):
            signal_list = self.signals[i]
            output_list = self.outputs[i]

            for signal in signal_list:
                match len(signal):
                    case 2: self.known_signals[1] = signal
                    case 4: self.known_signals[4] = signal
                    case 3: self.known_signals[7] = signal
                    case 7: self.known_signals[8] = signal

            for signal in signal_list:
                match len(signal), self.common_with(1, signal), self.common_with(4, signal):
                    case 6, 2, 3: self.known_signals[0] = signal
                    case 5, 1, 2: self.known_signals[2] = signal
                    case 5, 2, _: self.known_signals[3] = signal
                    case 5, 1, 3: self.known_signals[5] = signal
                    case 6, 1, 3: self.known_signals[6] = signal
                    case 6, 2, 4: self.known_signals[9] = signal

            decode = {"".join(sorted(v)): str(k) for k, v in self.known_signals.items()}

            output = ""
            for o in output_list:
                output += decode["".join(sorted(o))]

            result += int(output)

        return result

    def common_with(self, num, pattern) -> int:
        return len([c for c in pattern if c in self.known_signals[num]])

    def parse_signals(self, lines) -> tuple:
        signals, outputs = [], []
        for line in lines:
            s, o = line.split("|")
            signals.append(s.split())
            outputs.append(o.split())
        return (signals, outputs)
