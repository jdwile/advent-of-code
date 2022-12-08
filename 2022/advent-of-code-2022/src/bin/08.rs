use std::{collections::HashSet, ops::Range};

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
    let trees = parse_trees(input);
    let p1 = part1(trees.clone());
    let p2 = part2(trees);
    (p1, p2)
}

fn parse_trees(input: &str) -> Vec<Vec<i32>> {
    input
        .lines()
        .map(|line| {
            line.trim()
                .split("")
                .filter(|c| c != &"")
                .map(|c| c.parse::<i32>().unwrap())
                .collect_vec()
        })
        .collect_vec()
}

fn part1(trees: Vec<Vec<i32>>) -> usize {
    let mut visible_trees = HashSet::<(usize, usize)>::new();

    trees.iter().enumerate().for_each(|(i, row)| {
        let mut highest_tree = -1;
        row.iter().enumerate().for_each(|(j, &tree)| {
            if tree > highest_tree {
                highest_tree = tree;
                visible_trees.insert((i, j));
            }
        });
    });

    trees.iter().enumerate().for_each(|(i, row)| {
        let mut highest_tree = -1;
        row.iter().enumerate().for_each(|(j, _)| {
            if trees[i][trees[i].len() - j - 1] > highest_tree {
                highest_tree = trees[i][trees[i].len() - j - 1];
                visible_trees.insert((i, trees[i].len() - j - 1));
            }
        });
    });

    trees[0].iter().enumerate().for_each(|(i, _)| {
        let mut highest_tree = -1;
        trees.iter().enumerate().for_each(|(j, _)| {
            if trees[j][i] > highest_tree {
                highest_tree = trees[j][i];
                visible_trees.insert((j, i));
            }
        });
    });

    trees[0].iter().enumerate().for_each(|(i, _)| {
        let mut highest_tree = -1;
        trees.iter().enumerate().for_each(|(j, _)| {
            if trees[trees.len() - j - 1][i] > highest_tree {
                highest_tree = trees[trees.len() - j - 1][i];
                visible_trees.insert((trees.len() - j - 1, i));
            }
        });
    });

    visible_trees.len()
}

fn part2(trees: Vec<Vec<i32>>) -> usize {
    trees
        .iter()
        .enumerate()
        .map(|(i, row)| {
            row.iter()
                .enumerate()
                .map(|(j, &tree)| {
                    let mut scenic_score = 1;
                    let di_dj: Vec<(i32, i32)> = vec![(1, 0), (-1, 0), (0, 1), (0, -1)];

                    di_dj.iter().for_each(|(di, dj)| {
                        let mut step: i32 = 1;
                        let mut visible_trees = 0;
                        let i_range: Range<i32> = 0..trees.len() as i32;
                        let j_range: Range<i32> = 0..trees[i].len() as i32;

                        while i_range.contains(&(i as i32 + di * step))
                            && j_range.contains(&(j as i32 + dj * step))
                        {
                            visible_trees += 1;
                            if trees[(i as i32 + di * step) as usize]
                                [(j as i32 + dj * step) as usize]
                                >= tree
                            {
                                break;
                            }
                            step += 1;
                        }
                        scenic_score *= visible_trees;
                    });
                    scenic_score
                })
                .max()
                .unwrap()
        })
        .max()
        .unwrap()
}
