from typing import List
import time


def input_as_string(filename: str) -> str:
    """returns the content of the input file as a string"""
    with open(filename) as f:
        return f.read().rstrip("\n")


def input_as_lines(filename: str) -> List[str]:
    """Return a list where each line in the input file is an element of the list"""
    return input_as_string(filename).split("\n")


def input_as_ints(filename: str) -> List[int]:
    """Return a list where each line in the input file is an element of the list, converted into an integer"""
    lines = input_as_lines(filename)
    line_as_int = lambda l: int(l.rstrip("\n"))
    return list(map(line_as_int, lines))


def timer(func):
    def wrapper(*args, **kwargs):
        start_time = time.time()
        result = func(*args, **kwargs)
        print(f"\nTime required: {(time.time() - start_time)*1000:.2f} ms\n")
        return result

    return wrapper
