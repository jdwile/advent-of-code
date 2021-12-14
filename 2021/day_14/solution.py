from utils.aoc import input_as_lines
from collections import Counter


class Solution:
    def __init__(
        self,
        filename="C:\\Users\\12692\\bench\\advent-of-code\\2021\\day_14\\input.txt",
    ) -> None:
        self.template, self.rules = self.parse_input(input_as_lines(filename))

    def solve_part_one(self) -> str:
        element_counts = self.run_reaction(10)
        
        return max(element_counts) - min(element_counts)

    def solve_part_two(self) -> str:
        element_counts = self.run_reaction(40)
        
        return max(element_counts) - min(element_counts)

    def run_reaction(self, iterations):
        polymer = ["".join(x) for x in zip(self.template, self.template[1:])]
        element_occurances = Counter(self.template)
        pairs = Counter(polymer)

        for step in range(iterations):
            for pair, count in pairs.copy().items():
                element = self.rules[pair]
                pairs[pair] -= count
                pairs[pair[0] + element] += count
                pairs[element + pair[1]] += count
                element_occurances[element] += count
        
        return element_occurances.values()

    def parse_input(self, lines):
        template = lines[0]
        rules = {}

        for line in lines[2:]:
            pair, element = line.split(" -> ")
            rules[pair] = element
        
        return (template, rules)
