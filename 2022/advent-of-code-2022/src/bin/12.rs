use parse_display::Display;
use std::collections::HashMap;

use itertools::Itertools;

#[aoc::main(12)]
pub fn main(input: &str) -> (usize, usize) {
    solve(input)
}

#[aoc::test(12)]
pub fn test(input: &str) -> (String, String) {
    let res = solve(input);
    (res.0.to_string(), res.1.to_string())
}

fn solve(input: &str) -> (usize, usize) {
    let grid = input
        .lines()
        .map(|line| line.chars().collect_vec())
        .collect_vec();

    let p1 = part1(grid.clone());
    let p2 = part2(grid);
    (p1, p2)
}

#[derive(Clone, Copy, Debug)]
enum Direction {
    Uphill,
    Downhill,
}

#[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, Display)]
#[display("{elevation}: ({x}, {y})")]
struct Point {
    x: usize,
    y: usize,
    elevation: char,
}

impl Point {
    fn can_move_to(&self, destination: Point, direction: Direction) -> bool {
        match direction {
            Direction::Uphill => {
                destination.get_elevation_value() - self.get_elevation_value() <= 1
            }
            Direction::Downhill => {
                self.get_elevation_value() - destination.get_elevation_value() <= 1
            }
        }
    }

    fn get_elevation_value(&self) -> i32 {
        match self.elevation {
            'S' => 0,
            'E' => 26,
            _ => self.elevation as i32 - 'a' as i32,
        }
    }

    fn get_neighbors(&self, grid: Vec<Vec<char>>) -> Vec<Point> {
        let mut neighbors = Vec::<Point>::new();
        let dx_dy: Vec<(i32, i32)> = vec![(0, 1), (0, -1), (1, 0), (-1, 0)];
        for (dx, dy) in dx_dy {
            let new_x = self.x as i32 + dx;
            let new_y = self.y as i32 + dy;
            if new_x >= 0 && new_x < grid.len() as i32 && new_y >= 0 && new_y < grid[0].len() as i32
            {
                neighbors.push(Point {
                    x: new_x as usize,
                    y: new_y as usize,
                    elevation: grid[new_x as usize][new_y as usize],
                })
            }
        }
        neighbors
    }
}

fn a_star(start: Point, goal: char, direction: Direction, grid: Vec<Vec<char>>) -> usize {
    let mut came_from = HashMap::<Point, Point>::new();
    let mut possible_points = vec![start];

    let mut g_score = HashMap::<Point, usize>::new();
    g_score.insert(start, 0);

    let mut f_score = HashMap::<Point, usize>::new();
    f_score.insert(start, 0);

    while !possible_points.is_empty() {
        let (cur_index, &cur) = possible_points
            .iter()
            .enumerate()
            .reduce(|(i, a), (j, b)| {
                f_score.entry(*a).or_insert(usize::MAX);
                f_score.entry(*b).or_insert(usize::MAX);
                if f_score[a] <= f_score[b] {
                    (i, a)
                } else {
                    (j, b)
                }
            })
            .unwrap();

        if cur.elevation == goal {
            let mut path = Vec::<Point>::new();
            let mut node = cur;
            path.push(node);

            while node != start {
                node = came_from[&node];
                path.push(node);
            }

            return path.len();
        }

        possible_points.remove(cur_index);

        cur.get_neighbors(grid.clone())
            .iter()
            .filter(|&point| cur.can_move_to(*point, direction))
            .for_each(|neighbor| {
                let tentative_g_score = g_score[&cur] + 1;
                g_score.entry(*neighbor).or_insert(usize::MAX);

                if tentative_g_score < g_score[neighbor] {
                    came_from.insert(*neighbor, cur);
                    g_score.insert(*neighbor, tentative_g_score);
                    g_score.insert(*neighbor, tentative_g_score + 1);

                    if !possible_points.contains(neighbor) {
                        possible_points.push(*neighbor);
                    }
                }
            });
    }

    0
}

fn part1(grid: Vec<Vec<char>>) -> usize {
    let mut start: Point = Point {
        x: 0,
        y: 0,
        elevation: 'S',
    };

    grid.iter().enumerate().for_each(|(x, row)| {
        row.iter().enumerate().for_each(|(y, &point)| {
            if point == 'S' {
                start = Point {
                    x,
                    y,
                    elevation: point,
                };
            }
        })
    });

    a_star(start, 'E', Direction::Uphill, grid) - 1
}

fn part2(grid: Vec<Vec<char>>) -> usize {
    let mut start: Point = Point {
        x: 0,
        y: 0,
        elevation: 'E',
    };

    grid.iter().enumerate().for_each(|(x, row)| {
        row.iter().enumerate().for_each(|(y, &point)| {
            if point == 'E' {
                start = Point {
                    x,
                    y,
                    elevation: point,
                };
            }
        })
    });

    a_star(start, 'a', Direction::Downhill, grid) - 1
}
