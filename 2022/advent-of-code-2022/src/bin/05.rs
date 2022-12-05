use scan_fmt::scan_fmt;

use itertools::Itertools;

#[aoc::main(05)]
pub fn main(input: &str) -> (String, String) {
    solve(input)
}

#[aoc::test(05)]
pub fn test(input: &str) -> (String, String) {
    solve(input)
}

fn solve(input: &str) -> (String, String) {
    let (crates_str, procedure_str) = input.split_once("\r\n\r\n").unwrap();
    let crates = get_crates(crates_str);
    let procedure = get_procedure(procedure_str);

    let p1 = part1(crates.clone(), procedure.clone());
    let p2 = part2(crates, procedure);
    (p1, p2)
}

fn get_procedure(procedure: &str) -> Vec<(usize, usize, usize)> {
    procedure
        .lines()
        .map(|line| scan_fmt!(line, "move {d} from {d} to {d}", usize, usize, usize).unwrap())
        .collect_vec()
}

fn get_crates(crates: &str) -> Vec<Vec<char>> {
    let chunks = crates
        .lines()
        .map(|line| line.chars().collect_vec())
        .collect_vec();

    let mut crates = Vec::<Vec<char>>::new();
    for j in 0..chunks[0].len() {
        let mut new_crate = Vec::<char>::new();
        for crates in &chunks {
            if crates[j].is_alphanumeric() {
                new_crate.push(crates[j]);
            }
        }
        new_crate.reverse();
        crates.push(new_crate);
    }
    crates.retain(|_crate| !_crate.is_empty() && _crate[0].is_numeric());
    crates.iter_mut().for_each(|_crate| {
        _crate.remove(0);
    });
    crates
}

fn part1(crates: Vec<Vec<char>>, procedure: Vec<(usize, usize, usize)>) -> String {
    let mut crates = crates;

    procedure.iter().for_each(|&(num, from, to)| {
        for _ in 0..num {
            let item = crates[from - 1].pop().unwrap();
            crates[to - 1].push(item);
        }
    });

    crates.iter().map(|c| c.last().unwrap()).join("")
}

fn part2(crates: Vec<Vec<char>>, procedure: Vec<(usize, usize, usize)>) -> String {
    let mut crates = crates;

    procedure.iter().for_each(|&(num, from, to)| {
        let len = crates[from - 1].len();
        let mut moved_crates: Vec<char> = crates[from - 1].drain(len - num..len).collect();
        crates[to - 1].append(&mut moved_crates);
    });

    crates.iter().map(|c| c.last().unwrap()).join("")
}
