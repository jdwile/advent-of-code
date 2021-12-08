from utils.aoc import input_as_lines, input_as_string
import re


class Solution:
    def __init__(
        self,
        filename="C:\\Users\\Jenner\\bench\\advent-of-code\\2021\\day_08\\input.txt",
    ) -> None:
        self.signals, self.outputs = self.parse_signals(input_as_lines(filename))

    def solve_part_one(self) -> str:
        relevant_counts = [2, 3, 4, 7]
        return sum(
            [
                len([x for x in output if len(x) in relevant_counts])
                for output in self.outputs
            ]
        )

    def solve_part_two(self) -> str:
        relevant_counts = {2: 1, 3: 7, 4: 4, 7: 8}
        known_signals = {}
        result = 0
        for i in range(len(self.signals)):
            signal_list = self.signals[i]
            output_list = self.outputs[i]
            uniques = [
                signal
                for signal in signal_list
                if len(signal) in relevant_counts.keys()
            ]

            for unique in uniques:
                known_signals[relevant_counts[len(unique)]] = unique

            remaining_signals = [
                signal for signal in signal_list if signal not in uniques
            ]

            known_signals[9] = [
                signal
                for signal in remaining_signals
                if all([c in signal for c in known_signals[4]])
            ][0]
            remaining_signals.remove(known_signals[9])

            five_counts = [signal for signal in remaining_signals if len(signal) == 5]
            known_signals[3] = [
                signal
                for signal in five_counts
                if all([c in signal for c in known_signals[1]])
            ][0]
            five_counts.remove(known_signals[3])
            remaining_signals.remove(known_signals[3])

            known_signals[5] = [
                signal
                for signal in five_counts
                if len([c for c in signal if c in known_signals[4]]) == 3
            ][0]
            five_counts.remove(known_signals[5])
            remaining_signals.remove(known_signals[5])

            known_signals[2] = five_counts[0]
            remaining_signals.remove(known_signals[2])

            known_signals[0] = [
                signal
                for signal in remaining_signals
                if all([c in signal for c in known_signals[1]])
            ][0]

            remaining_signals.remove(known_signals[0])
            known_signals[6] = remaining_signals[0]

            for k in known_signals.keys():
                known_signals[k] = "".join(sorted(known_signals[k]))
            decode = {v: k for k, v in known_signals.items()}

            output = ""
            for o in output_list:
                output += str(decode["".join(sorted(o))])

            result += int(output)

        return result

    def parse_signals(self, lines) -> tuple:
        signals, outputs = [], []
        for line in lines:
            s, o = line.split("|")
            signals.append(s.split())
            outputs.append(o.split())
        return (signals, outputs)
