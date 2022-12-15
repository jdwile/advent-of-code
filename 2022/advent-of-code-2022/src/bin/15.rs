use itertools::Itertools;
use scan_fmt::scan_fmt;
use std::cmp::max;

#[aoc::main(15)]
pub fn main(input: &str) -> (i64, i64) {
    solve(input)
}

#[aoc::test(15)]
pub fn test(input: &str) -> (String, String) {
    let res = solve(input);
    (res.0.to_string(), res.1.to_string())
}

fn solve(input: &str) -> (i64, i64) {
    let p1 = part1(input);
    let p2 = part2(input);
    (p1, p2)
}

#[derive(Eq, PartialEq, Hash, Debug, Clone, Copy)]
struct Sensor {
    location: Point,
    range: i64,
}

type Point = (i64, i64);

fn parse_grid(input: &str) -> (Vec<Sensor>, i64, i64) {
    let mut sensors = Vec::<Sensor>::new();

    let bounds_str = input.lines().next().unwrap().split_once(',').unwrap();
    let y = bounds_str.0.parse::<i64>().unwrap();
    let search_space = bounds_str.1.parse::<i64>().unwrap();

    input
        .lines()
        .skip(1)
        .map(|line| {
            scan_fmt!(
                line,
                "Sensor at x={d}, y={d}: closest beacon is at x={d}, y={d}",
                i64,
                i64,
                i64,
                i64
            )
            .unwrap()
        })
        .for_each(|(sx, sy, bx, by)| {
            sensors.push(Sensor {
                location: (sx, sy),
                range: (sx - bx).abs() + (sy - by).abs(),
            });
        });

    (sensors, y, search_space)
}

fn get_impossible_ranges(sensors: &[Sensor], y: i64) -> Vec<Point> {
    let ranges = sensors
        .iter()
        .filter_map(|sensor| {
            let delta = sensor.range - (y - sensor.location.1).abs();
            if delta >= 0 {
                Some((sensor.location.0 - delta, sensor.location.0 + delta))
            } else {
                None
            }
        })
        .collect_vec();
    merge_ranges(ranges)
}

fn merge_ranges(mut ranges: Vec<Point>) -> Vec<Point> {
    ranges.sort_by(|r1, r2| r1.0.cmp(&r2.0));
    let mut index = 0;
    for i in 1..ranges.len() {
        if ranges[index].1 >= ranges[i].0 {
            ranges[index].1 = max(ranges[index].1, ranges[i].1);
        } else {
            index += 1;
            ranges[index] = ranges[i];
        }
    }
    ranges.resize(index + 1, (0, 0));
    ranges
}

fn find_unblocked_space(sensors: Vec<Sensor>, y: i64, search_space: i64) -> Option<i64> {
    let mut ranges = Vec::<Point>::new();

    sensors.iter().for_each(|sensor| {
        let delta = sensor.range - (y - sensor.location.1).abs();
        if delta >= 0 {
            ranges.push((sensor.location.0 - delta, sensor.location.0 + delta));
        }
    });

    ranges.sort_by(|(range_one_x, _), (range_two_x, _)| range_one_x.cmp(range_two_x));

    let mut max_x: i64 = 0;
    for point in ranges {
        if point.0 > max_x {
            return Some(point.0 - 1);
        }
        if point.1 > search_space {
            return None;
        }

        max_x = max(max_x, point.1);
    }
    None
}

fn part1(input: &str) -> i64 {
    let (sensors, y, _) = parse_grid(input);
    let impossible_ranges = get_impossible_ranges(&sensors, y);

    impossible_ranges
        .iter()
        .map(|(x_min, x_max)| x_max - x_min)
        .sum()
}

fn part2(input: &str) -> i64 {
    let (sensors, _, search_space) = parse_grid(input);

    for y in 0..search_space {
        if let Some(x_val) = find_unblocked_space(sensors.clone(), y, search_space) {
            return x_val * 4_000_000 + y;
        }
    }
    0
}
