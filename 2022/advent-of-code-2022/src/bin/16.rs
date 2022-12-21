use itertools::Itertools;
use std::collections::HashMap;

#[aoc::main(16)]
pub fn main(input: &str) -> (i64, i64) {
    solve(input)
}

#[aoc::test(16)]
pub fn test(input: &str) -> (String, String) {
    let res = solve(input);
    (res.0.to_string(), res.1.to_string())
}

fn solve(input: &str) -> (i64, i64) {
    let p1 = part1(input);
    let p2 = part2(input);
    (p1, p2)
}

#[derive(Eq, PartialEq, Hash, Debug, Clone)]
struct Valve {
    flow_rate: i64,
    tunnels: Vec<String>,
}

fn get_valves(input: &str) -> HashMap<&str, Valve> {
    let mut valves = HashMap::<&str, Valve>::new();

    input.lines().for_each(|line| {
        let (valve_def, tunnel_def) = line.split_once(';').unwrap();
        let (_, rate_str) = valve_def.split_once('=').unwrap();

        valves.insert(
            valve_def
                .split_ascii_whitespace()
                .skip(1)
                .into_iter()
                .next()
                .unwrap()
                .trim(),
            Valve {
                flow_rate: rate_str.parse().unwrap(),
                tunnels: match tunnel_def.find(',') {
                    Some(_) => {
                        let mut tunnels = tunnel_def.split(',').collect_vec();
                        tunnels[0] = tunnels[0]
                            .split_ascii_whitespace()
                            .into_iter()
                            .rev()
                            .next()
                            .unwrap()
                            .trim();

                        tunnels
                            .into_iter()
                            .map(|s| s.trim().to_string())
                            .collect_vec()
                    }
                    None => vec![tunnel_def
                        .split_whitespace()
                        .into_iter()
                        .rev()
                        .next()
                        .unwrap()
                        .to_string()],
                },
            },
        );
    });

    valves
}

fn get_flow_rate(
    current_valve: &str,
    open_valves: Vec<&str>,
    minute_count: i64,
    valves: HashMap<&str, Valve>,
) -> i64 {
    if minute_count == 30 {
        return 0;
    }

    // println!("{}", minute_count);

    let mut flow_rates = Vec::<i64>::new();

    if valves[&current_valve].flow_rate > 0 && !open_valves.contains(&current_valve) {
        let mut new_open_valves = open_valves.clone();
        new_open_valves.push(current_valve);
        // println!("Opening valve {}", current_valve);
        flow_rates.push(
            (minute_count - 1) * valves[&current_valve].flow_rate
                + get_flow_rate(
                    current_valve.clone(),
                    new_open_valves,
                    minute_count + 1,
                    valves.clone(),
                ),
        );
    }

    valves[&current_valve].tunnels.iter().for_each(|new_valve| {
        // println!("Moving to valve {}", new_valve);
        flow_rates.push(get_flow_rate(
            new_valve,
            open_valves.clone(),
            minute_count + 1,
            valves.clone(),
        ));
    });

    if flow_rates.len() == 0 {
        return 0;
    }

    flow_rates.into_iter().max().unwrap()
}

fn part1(input: &str) -> i64 {
    let valves = get_valves(input);
    for key in valves.keys() {
        println!("{}: {:?}", key, valves[key]);
    }
    get_flow_rate("AA", Vec::<&str>::new(), 1, valves)
}

fn part2(_input: &str) -> i64 {
    0
}
