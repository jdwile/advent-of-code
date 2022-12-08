use std::collections::HashSet;

use itertools::Itertools;

#[aoc::main(08)]
pub fn main(input: &str) -> (usize, usize) {
    solve(input)
}

#[aoc::test(08)]
pub fn test(input: &str) -> (String, String) {
    let res = solve(input);
    (res.0.to_string(), res.1.to_string())
}

fn solve(input: &str) -> (usize, usize) {
    let p1 = part1(input.clone());
    let p2 = part2(input);
    (p1, p2)
}

fn part1(input: &str) -> usize {
    let trees = input
        .lines()
        .map(|line| {
            line.trim()
                .split("")
                .filter(|c| c != &"")
                .map(|c| c.parse::<i32>().unwrap())
                .collect_vec()
        })
        .collect_vec();

    let mut highest_tree: i32;
    let mut visible_trees = HashSet::<(usize, usize)>::new();

    for i in 0..trees.len() {
        highest_tree = -1;
        for j in 0..trees[i].len() {
            if trees[i][j] > highest_tree {
                // println!("VISIBLE FROM LEFT ({}, {})", i, j);
                highest_tree = trees[i][j];
                visible_trees.insert((i, j));
            }
        }
    }

    for i in 0..trees.len() {
        highest_tree = -1;
        for j in 0..trees[i].len() {
            // println!("VISIBLE FROM RIGHT ({}, {})", i, trees[i].len() - j - 1);
            if trees[i][trees[i].len() - j - 1] > highest_tree {
                highest_tree = trees[i][trees[i].len() - j - 1];
                visible_trees.insert((i, trees[i].len() - j - 1));
            }
        }
    }

    for i in 0..trees[0].len() {
        highest_tree = -1;
        for j in 0..trees.len() {
            if trees[j][i] > highest_tree {
                highest_tree = trees[j][i];
                visible_trees.insert((j, i));
            }
        }
    }

    for i in 0..trees[0].len() {
        highest_tree = -1;
        for j in 0..trees.len() {
            if trees[trees.len() - j - 1][i] > highest_tree {
                highest_tree = trees[trees.len() - j - 1][i];
                visible_trees.insert((trees.len() - j - 1, i));
            }
        }
    }

    visible_trees.len()
}

fn part2(input: &str) -> usize {
    let trees = input
        .lines()
        .map(|line| {
            line.trim()
                .split("")
                .filter(|c| c != &"")
                .map(|c| c.parse::<i32>().unwrap())
                .collect_vec()
        })
        .collect_vec();

    let mut highest_view = 0;

    for i in 1..trees.len() - 1 {
        for j in 1..trees[0].len() - 1 {
            let mut total_view_score = 1;
            let mut di = 1;
            let highest_tree = trees[i][j];
            let mut view_score = 0;

            while i + di < trees.len() {
                view_score += 1;
                if trees[i + di][j] >= highest_tree {
                    break;
                }
                di += 1;
            }

            total_view_score *= view_score;

            di = 1;
            view_score = 0;
            while i as i32 - di as i32 >= 0 {
                view_score += 1;
                if trees[i - di][j] >= highest_tree {
                    break;
                }
                di += 1;
            }

            total_view_score *= view_score;

            let mut dj = 1;
            view_score = 0;
            while j + dj < trees[i].len() {
                view_score += 1;
                if trees[i][j + dj] >= highest_tree {
                    break;
                }
                dj += 1;
            }

            total_view_score *= view_score;

            dj = 1;
            view_score = 0;
            while j as i32 - dj as i32 >= 0 {
                view_score += 1;
                if trees[i][j - dj] >= highest_tree {
                    break;
                }
                dj += 1;
            }

            total_view_score *= view_score;
            if highest_view < total_view_score {
                // println!("New highest view! ({}, {}) => {}", i, j, total_view_score);
                highest_view = total_view_score;
            }
        }
    }

    highest_view
}
