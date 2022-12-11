use itertools::Itertools;
use scan_fmt::scan_fmt;

#[aoc::main(11)]
pub fn main(input: &str) -> (u64, u64) {
    solve(input)
}

#[aoc::test(11)]
pub fn test(input: &str) -> (String, String) {
    let res = solve(input);
    (res.0.to_string(), res.1.to_string())
}

#[derive(Clone, Debug)]
struct Operation {
    op: char,
    n: String,
}

fn solve(input: &str) -> (u64, u64) {
    let monkies = get_monkies(input);
    let p1 = part1(monkies.clone());
    let p2 = part2(monkies);
    (p1, p2)
}

fn get_monkies(input: &str) -> Vec<(Vec<u64>, Operation, u64, usize, usize)> {
    input
        .trim()
        .split("\r\n\r\n")
        .map(|monkey_def| {
            let monkey = monkey_def.trim().lines().skip(1).collect_vec();
            let mut items_str = monkey[0].split(", ").collect_vec();
            items_str[0] = items_str[0].split_once(": ").unwrap().1;
            let items = items_str
                .iter()
                .map(|&item| item.trim().parse::<u64>().unwrap())
                .collect_vec();

            let op_str = monkey[1].split_once(" old ").unwrap().1;
            let (op, num_str) = op_str.split_once(' ').unwrap();
            let operation = Operation {
                op: *op.trim().chars().collect_vec().first().unwrap(),
                n: num_str.to_string(),
            };

            let test = scan_fmt!(monkey[2].trim(), "Test: divisible by {d}", u64).unwrap();
            let success =
                scan_fmt!(monkey[3].trim(), "If true: throw to monkey {d}", usize).unwrap();
            let failure =
                scan_fmt!(monkey[4].trim(), "If false: throw to monkey {d}", usize).unwrap();

            (items, operation, test, success, failure)
        })
        .collect_vec()
}

fn get_inspection_counts(
    mut monkies: Vec<(Vec<u64>, Operation, u64, usize, usize)>,
    rounds: u64,
    modification: impl Fn(u64) -> u64,
) -> Vec<u64> {
    let mut monkey_inspections: Vec<u64> = vec![0; monkies.len()];

    for _ in 0..rounds {
        for i in 0..monkies.len() {
            let (items, operation, test, success, fail) = monkies[i].clone();
            items.iter().for_each(|&(mut worry_level)| {
                monkey_inspections[i] += 1;

                match operation.op {
                    '+' => worry_level += operation.n.parse::<u64>().unwrap(),
                    '*' => match operation.n.as_str() {
                        "old" => worry_level *= worry_level,
                        _ => worry_level *= operation.n.parse::<u64>().unwrap(),
                    },
                    _ => panic!("Shouldn't happen"),
                }

                worry_level = modification(worry_level);

                if worry_level.rem_euclid(test) == 0 {
                    monkies[success].0.push(worry_level);
                } else {
                    monkies[fail].0.push(worry_level);
                }
            });
            monkies[i].0.clear();
        }
    }

    monkey_inspections
}

fn part1(monkies: Vec<(Vec<u64>, Operation, u64, usize, usize)>) -> u64 {
    let mut inspections = get_inspection_counts(monkies, 20, |x| x / 3);

    inspections.sort();
    inspections.reverse();
    inspections[0] * inspections[1]
}

fn part2(monkies: Vec<(Vec<u64>, Operation, u64, usize, usize)>) -> u64 {
    let modulus = monkies.iter().map(|m| m.2).product::<u64>();
    let mut inspections = get_inspection_counts(monkies, 10000, |x| x % modulus);

    inspections.sort();
    inspections.reverse();
    inspections[0] * inspections[1]
}
