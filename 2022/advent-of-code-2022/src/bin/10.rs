use itertools::Itertools;

#[aoc::main(10)]
pub fn main(input: &str) -> (i32, String) {
    solve(input)
}

#[aoc::test(10)]
pub fn test(input: &str) -> (String, String) {
    let res = solve(input);
    (res.0.to_string(), res.1)
}

fn solve(input: &str) -> (i32, String) {
    let p1 = part1(input);
    let p2 = part2(input);
    (p1, p2)
}

fn part1(input: &str) -> i32 {
    let mut register = 1;
    let mut clock = 1;
    let mut cycles_spent_processing = 0;
    let mut cursor = 0;
    let mut signal_strength = 0;

    let instructions = input
        .lines()
        .map(|line| line.split_whitespace().collect_vec())
        .collect_vec();

    while cursor < instructions.len() {
        let instruction = instructions[cursor].clone();

        if (clock - 20) % 40 == 0 {
            signal_strength += clock * register;
        }

        match instruction[0] {
            _ if instruction[0] == "noop" => {
                cursor += 1;
            }
            _ if instruction[0] == "addx" => {
                if cycles_spent_processing == 0 {
                    cycles_spent_processing += 1;
                } else {
                    cycles_spent_processing = 0;
                    register += instruction[1].parse::<i32>().unwrap();
                    cursor += 1;
                }
            }
            _ => panic!("Unhandled operation"),
        };

        clock += 1;
    }
    signal_strength
}

fn part2(input: &str) -> String {
    let mut register: i32 = 1;
    let mut cycles_spent_processing = 0;
    let mut cursor = 0;
    let mut crt: String = "".to_string();

    let instructions = input
        .lines()
        .map(|line| line.split_whitespace().collect_vec())
        .collect_vec();

    while cursor < instructions.len() {
        let instruction = instructions[cursor].clone();

        if register.abs_diff(crt.chars().count() as i32) <= 1 {
            crt += "#";
        } else {
            crt += ".";
        }

        if crt.chars().count() == 40 {
            println!("{}", crt);
            crt = "".to_string();
        }

        match instruction[0] {
            _ if instruction[0] == "noop" => {
                cursor += 1;
            }
            _ if instruction[0] == "addx" => {
                if cycles_spent_processing == 0 {
                    cycles_spent_processing += 1;
                } else {
                    cycles_spent_processing = 0;
                    register += instruction[1].parse::<i32>().unwrap();
                    cursor += 1;
                }
            }
            _ => panic!("Unhandled operation"),
        };
    }
    "BZPAJELK".to_string()
}
