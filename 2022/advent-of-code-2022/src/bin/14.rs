use std::collections::HashMap;

use itertools::Itertools;

#[aoc::main(14)]
pub fn main(input: &str) -> (usize, usize) {
    solve(input)
}

#[aoc::test(14)]
pub fn test(input: &str) -> (String, String) {
    let res = solve(input);
    (res.0.to_string(), res.1.to_string())
}

fn solve(input: &str) -> (usize, usize) {
    let p1 = part1(input);
    let p2 = part2(input);
    (p1, p2)
}

#[derive(Eq, PartialEq, Hash, Debug, Clone, Copy)]
struct Point {
    x: i32,
    y: i32,
}

impl Point {
    fn has_object_beneath(&self, grid: &HashMap<Point, char>) -> bool {
        grid.keys().any(|item| item.x == self.x && item.y > self.y)
    }

    fn can_fall_down(&self, grid: &HashMap<Point, char>) -> bool {
        !grid
            .keys()
            .any(|item| item.x == self.x && item.y == self.y + 1)
    }

    fn can_fall_left(&self, grid: &HashMap<Point, char>) -> bool {
        !grid
            .keys()
            .any(|item| item.x == self.x - 1 && item.y == self.y + 1)
    }

    fn can_fall_right(&self, grid: &HashMap<Point, char>) -> bool {
        !grid
            .keys()
            .any(|item| item.x == self.x + 1 && item.y == self.y + 1)
    }
}

fn parse_grid(input: &str) -> HashMap<Point, char> {
    let mut grid = HashMap::<Point, char>::new();
    grid.insert(Point { x: 500, y: 0 }, '+');

    input.lines().for_each(|line| {
        let coords = line
            .split(" -> ")
            .map(|pair| {
                let coord_str = pair.split_once(',').unwrap();
                Point {
                    x: coord_str.0.parse::<i32>().unwrap(),
                    y: coord_str.1.parse::<i32>().unwrap(),
                }
            })
            .collect_vec();

        let mut cur = coords[0];
        grid.insert(cur, '#');

        coords.iter().skip(1).for_each(|&dest| {
            let mut dx = 0;
            let mut dy = 0;

            if cur.x != dest.x {
                dx = (dest.x - cur.x).signum();
            }
            if cur.y != dest.y {
                dy = (dest.y - cur.y).signum();
            }

            while cur != dest {
                cur.x += dx;
                cur.y += dy;

                grid.insert(cur, '#');
            }
        });
    });

    grid
}

#[allow(dead_code)]
fn print_grid(grid: &HashMap<Point, char>) {
    let x_min = grid.keys().map(|p| p.x).min().unwrap();
    let x_max = grid.keys().map(|p| p.x).max().unwrap();
    let y_min = grid.keys().map(|p| p.y).min().unwrap();
    let y_max = grid.keys().map(|p| p.y).max().unwrap();

    for y in y_min..y_max + 1 {
        let mut line = "".to_string();
        for x in x_min..x_max + 1 {
            let point = Point { x, y };
            if grid.contains_key(&point) {
                line += format!("{}", *grid.get(&point).unwrap()).as_str();
            } else {
                line += ".";
            }
        }
        println!("{}", line);
    }
}

fn part1(input: &str) -> usize {
    let mut grid = parse_grid(input);
    let mut falling_into_void = false;

    while !falling_into_void {
        let mut sand = Point { x: 500, y: 0 };
        let mut is_at_rest = false;

        while !is_at_rest {
            if !sand.has_object_beneath(&grid) {
                falling_into_void = true;
                break;
            }

            if sand.can_fall_down(&grid) {
                sand.y += 1;
            } else if sand.can_fall_left(&grid) {
                sand.x -= 1;
                sand.y += 1;
            } else if sand.can_fall_right(&grid) {
                sand.x += 1;
                sand.y += 1;
            } else {
                grid.insert(sand, 'o');
                is_at_rest = true;
            }
        }
    }
    grid.values().filter(|&c| *c == 'o').count()
}

fn part2(input: &str) -> usize {
    let mut grid = parse_grid(input);
    let mut is_piled = false;

    let floor = grid.keys().map(|p| p.y).max().unwrap() + 2;

    while !is_piled {
        let mut sand = Point { x: 500, y: 0 };
        let mut is_at_rest = false;

        while !is_at_rest {
            if sand.y + 1 == floor {
                grid.insert(sand, 'o');
                is_at_rest = true;
                continue;
            }

            if sand.can_fall_down(&grid) {
                sand.y += 1;
            } else if sand.can_fall_left(&grid) {
                sand.x -= 1;
                sand.y += 1;
            } else if sand.can_fall_right(&grid) {
                sand.x += 1;
                sand.y += 1;
            } else {
                grid.insert(sand, 'o');
                is_at_rest = true;
                if sand.x == 500 && sand.y == 0 {
                    is_piled = true;
                }
            }
        }
    }
    grid.values().filter(|&c| *c == 'o').count()
}
