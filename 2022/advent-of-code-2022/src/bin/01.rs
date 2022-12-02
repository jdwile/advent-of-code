#[aoc::main(01)]
pub fn main(input: &str) -> (usize, usize) {
    solve(input)
}

#[aoc::test(01)]
pub fn test(input: &str) -> (usize, usize) {
    solve(input)
}

fn solve(input: &str) -> (usize, usize) {
    let p1 = part1(input);
    let p2 = part2(input);

    (p1, p2)
}

fn part1(input: &str) -> usize {
    input
        .split("\r\n\r\n")
        .map(|elf| {
            elf.lines()
                .map(|cal_str| cal_str.parse::<usize>().unwrap())
                .sum::<usize>()
        })
        .max()
        .unwrap()
}

fn part2(input: &str) -> usize {
    let mut elf_calories: Vec<usize> = input
        .split("\r\n\r\n")
        .map(|elf| {
            elf.lines()
                .map(|cal_str| cal_str.parse::<usize>().unwrap())
                .sum::<usize>()
        })
        .collect();
    elf_calories.sort();
    elf_calories.iter().rev().take(3).sum()
}
