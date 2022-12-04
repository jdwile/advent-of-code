#[aoc::main(04)]
pub fn main(input: &str) -> (usize, usize) {
    solve(input)
}

#[aoc::test(04)]
pub fn test(input: &str) -> (usize, usize) {
    solve(input)
}

fn solve(input: &str) -> (usize, usize) {
    let p1 = part1(input);
    let p2 = part2(input);

    (p1, p2)
}

fn is_contained(a: (usize, usize), b: (usize, usize)) -> bool {
    a.0 <= b.0 && a.1 >= b.1 || b.0 <= a.0 && b.1 >= a.1
}

fn has_overlap(a: (usize, usize), b: (usize, usize)) -> bool {
    b.0 <= a.0 && b.1 >= a.0 || b.0 <= a.1 && b.1 >= a.1
}

fn get_range(r: &str) -> (usize, usize) {
    let range_str = r.split_once('-').unwrap();
    (range_str.0.parse().unwrap(), range_str.1.parse().unwrap())
}

fn part1(input: &str) -> usize {
    input
        .lines()
        .map(|line| line.split_once(',').unwrap())
        .map(|(first, second)| (get_range(first), get_range(second)))
        .map(|(a, b)| is_contained(a, b))
        .filter(|o| *o)
        .count()
}

fn part2(input: &str) -> usize {
    input
        .lines()
        .map(|line| line.split_once(',').unwrap())
        .map(|(first, second)| (get_range(first), get_range(second)))
        .map(|(a, b)| has_overlap(a, b) || is_contained(a, b))
        .filter(|o| *o)
        .count()
}
