use scan_fmt::scan_fmt;

#[aoc::main(04)]
pub fn main(input: &str) -> (usize, usize) {
    solve(input)
}

#[aoc::test(04)]
pub fn test(input: &str) -> (String, String) {
    let res = solve(input);
    (res.0.to_string(), res.1.to_string())
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

fn part1(input: &str) -> usize {
    input
        .lines()
        .map(|line| scan_fmt!(line, "{d}-{d},{d}-{d}", usize, usize, usize, usize).unwrap())
        .filter(|(a, b, c, d)| is_contained((*a, *b), (*c, *d)))
        .count()
}

fn part2(input: &str) -> usize {
    input
        .lines()
        .map(|line| scan_fmt!(line, "{d}-{d},{d}-{d}", usize, usize, usize, usize).unwrap())
        .filter(|(a, b, c, d)| has_overlap((*a, *b), (*c, *d)) || is_contained((*a, *b), (*c, *d)))
        .count()
}
