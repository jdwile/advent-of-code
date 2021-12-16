from utils.aoc import input_as_string
from math import prod


class Solution:
    def __init__(
        self,
        filename="C:\\Users\\Jenner\\bench\\advent-of-code\\2021\\day_16\\input.txt",
    ) -> None:
        self.packet = self.hex_to_bin(input_as_string(filename))
        self.version_sum = 0

    def solve_part_one(self) -> str:
        version_sum = 0
        packets = [self.evaluate_packet([self.packet])]
        while len(packets) > 0:
            packet = packets.pop()
            version_sum += int(packet.version, 2)
            for subpacket in packet.subpackets:
                packets.append(subpacket)
        return version_sum

    def solve_part_two(self) -> str:
        return self.evaluate_packet([self.packet]).literal

    def evaluate_packet(self, data):
        version = self.consume(data, 3)
        type_id = int(self.consume(data, 3), 2)

        packet = Packet(version, type_id)

        if packet.type_id == 4: # This is a literal, lets parse it out
            literal = ""
            while True:
                marker, *chunk = self.consume(data, 5)
                literal += "".join(chunk)
                if marker == "0":
                    break
            packet.set_literal(int(literal, 2))
            return packet

        length_type_id = self.consume(data, 1)[0]
        if length_type_id == "0": # This specifies the length of the subpacket data
            subpacket_len = int(self.consume(data, 15), 2)
            subpacket_data = [self.consume(data, subpacket_len)]
            while subpacket_data[0]:
                packet.add_subpacket(self.evaluate_packet(subpacket_data))
        else: # This specifies the number of subpackets
            num_subpackets = int(self.consume(data, 11), 2)
            for _ in range(num_subpackets):
                packet.add_subpacket(self.evaluate_packet(data))

        match packet.type_id:
            case 0:
                packet.set_literal(sum([p.literal for p in packet.subpackets]))
            case 1:
                packet.set_literal(prod([p.literal for p in packet.subpackets]))
            case 2:
                packet.set_literal(min([p.literal for p in packet.subpackets]))
            case 3:
                packet.set_literal(max([p.literal for p in packet.subpackets]))
            case 5:
                packet.set_literal(1 if packet.subpackets[0].literal > packet.subpackets[1].literal else 0)
            case 6:
                packet.set_literal(1 if packet.subpackets[0].literal < packet.subpackets[1].literal else 0)
            case 7:
                packet.set_literal(1 if packet.subpackets[0].literal == packet.subpackets[1].literal else 0)
        return packet

    def hex_to_bin(self, hex) -> str:
        return bin(int("1" + hex, 16))[3:]

    def consume(self, data, n):
        res = data[0][:n]
        data[0] = data[0][n:]
        return res


class Packet:
    def __init__(self, version: str, type_id: int):
        self.version = version
        self.type_id = type_id
        self.subpackets = []
        self.literal = 0

    def __str__(self):
        res = f"{self.version}: {self.type_id}"
        if self.type_id == 4:
            res += f" with literal {self.literal}"
        else:
            res += f" with {len(self.subpackets)} subpackets and result {self.literal}"
        return res

    def add_subpacket(self, subpacket):
        self.subpackets.append(subpacket)

    def set_literal(self, literal):
        self.literal = literal
