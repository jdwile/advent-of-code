use itertools::Itertools;

#[aoc::main(06)]
pub fn main(input: &str) -> (usize, usize) {
    solve(input)
}

#[aoc::test(06)]
pub fn test(input: &str) -> (String, String) {
    let res = solve(input);
    (res.0.to_string(), res.1.to_string())
}

fn solve(input: &str) -> (usize, usize) {
    let p1 = part1(input);
    let p2 = part2(input);
    (p1, p2)
}

fn find_marker(signal: &str, marker_size: usize) -> usize {
    signal
        .chars()
        .collect_vec()
        .windows(marker_size)
        .enumerate()
        .filter(|(_, window)| window.into_iter().all_unique())
        .map(|(i, _)| i + marker_size)
        .next()
        .unwrap()
}

fn part1(signal: &str) -> usize {
    find_marker(signal, 4)
}

fn part2(signal: &str) -> usize {
    find_marker(signal, 14)
}
