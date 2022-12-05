#[aoc::main(01)]
pub fn main(input: &str) -> (usize, usize) {
    solve(input)
}

#[aoc::test(01)]
pub fn test(input: &str) -> (String, String) {
    let res = solve(input);
    (res.0.to_string(), res.1.to_string())
}

fn solve(input: &str) -> (usize, usize) {
    let calories = get_calories(input);

    let p1 = part1(calories.clone());
    let p2 = part2(calories.clone());

    (p1, p2)
}

fn get_calories(input: &str) -> Vec<usize> {
    input
        .split("\r\n\r\n")
        .map(|elf| {
            elf.lines()
                .map(|cal_str| cal_str.parse::<usize>().unwrap())
                .sum::<usize>()
        })
        .collect()
}

fn part1(calories: Vec<usize>) -> usize {
    *calories.iter().max().unwrap()
}

fn part2(calories: Vec<usize>) -> usize {
    let mut calories = calories;
    calories.sort();
    calories.iter().rev().take(3).sum()
}
