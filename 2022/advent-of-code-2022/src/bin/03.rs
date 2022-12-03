use itertools::Itertools;
use sets::Set;

#[aoc::main(03)]
pub fn main(input: &str) -> (usize, usize) {
    solve(input)
}

#[aoc::test(03)]
pub fn test(input: &str) -> (usize, usize) {
    solve(input)
}

fn solve(input: &str) -> (usize, usize) {
    let p1 = part1(input);
    let p2 = part2(input);

    (p1, p2)
}

fn get_item_priority(item: char) -> usize {
    match item {
        _ if item.is_uppercase() => item as usize - 'A' as usize + 27,
        _ => item as usize - 'a' as usize + 1,
    }
}

fn get_unique_chars(pack: &str) -> Vec<char> {
    Set::new_unordered(&pack.chars().collect::<Vec<char>>())
        .nonrepeat()
        .data
}

fn part1(input: &str) -> usize {
    input
        .lines()
        .map(|line| -> usize {
            let (one, two) = line.split_at(line.len() / 2);

            get_unique_chars(one)
                .iter()
                .map(|item| match two.find(*item) {
                    Some(_) => get_item_priority(*item),
                    None => 0,
                })
                .sum()
        })
        .sum()
}

fn part2(input: &str) -> usize {
    input
        .lines()
        .tuples()
        .map(|(one, two, three)| -> usize {
            get_unique_chars(one)
                .iter()
                .map(|item| match two.find(*item) {
                    Some(_) => match three.find(*item) {
                        Some(_) => get_item_priority(*item),
                        None => 0,
                    },
                    None => 0,
                })
                .sum()
        })
        .sum()
}
