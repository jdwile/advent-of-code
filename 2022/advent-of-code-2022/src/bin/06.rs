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

fn is_unique_marker(marker: &str) -> bool {
    let unique_chars = marker
        .chars()
        .fold(Vec::<char>::new(), |mut acc: Vec<char>, c| {
            if !acc.contains(&c) {
                acc.push(c);
            }
            acc
        });
    unique_chars.len() == marker.len()
}

fn find_marker(signal: &str, marker_size: usize) -> usize {
    let signal = signal;
    for i in 0..signal.len() - marker_size {
        let marker = &signal[i..i + marker_size];
        if is_unique_marker(marker) {
            return i + marker_size;
        }
    }
    0
}

fn part1(signal: &str) -> usize {
    find_marker(signal, 4)
}

fn part2(signal: &str) -> usize {
    find_marker(signal, 14)
}
