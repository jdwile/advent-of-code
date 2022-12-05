use itertools::Itertools;

#[aoc::main(02)]
pub fn main(input: &str) -> (usize, usize) {
    solve(input)
}

#[aoc::test(02)]
pub fn test(input: &str) -> (String, String) {
    let res = solve(input);
    (res.0.to_string(), res.1.to_string())
}

fn solve(input: &str) -> (usize, usize) {
    let moves = get_moves(input);

    let p1 = part1(moves.clone());
    let p2 = part2(moves);

    (p1, p2)
}

static WIN: u32 = 6;
static DRAW: u32 = 3;
static LOSE: u32 = 0;

fn get_winning_move(mv: char) -> char {
    char::from_u32((mv as u32 - ('A' as u32) + 1).rem_euclid(3) + ('X' as u32)).unwrap()
}

fn get_losing_move(mv: char) -> char {
    char::from_u32((mv as u32 - ('A' as u32) + 2).rem_euclid(3) + ('X' as u32)).unwrap()
}

fn get_moves(input: &str) -> Vec<(char, char)> {
    input
        .lines()
        .map(|line| {
            let moves = line.split_once(' ').unwrap();
            (
                moves.0.chars().next().unwrap(),
                moves.1.chars().next().unwrap(),
            )
        })
        .collect_vec()
}

fn part1(moves: Vec<(char, char)>) -> usize {
    moves
        .iter()
        .map(|&(opp_move, my_move)| {
            (my_move as u32 - 'X' as u32 + 1)
                + match my_move {
                    _win if get_winning_move(opp_move) == my_move => WIN,
                    _lose if get_losing_move(opp_move) == my_move => LOSE,
                    _ => DRAW,
                }
        })
        .sum::<u32>() as usize
}

fn part2(moves: Vec<(char, char)>) -> usize {
    moves
        .iter()
        .map(|&(opp_move, strategy)| match strategy {
            'X' => get_losing_move(opp_move) as u32 - 'X' as u32 + 1 + LOSE,
            'Y' => opp_move as u32 - 'A' as u32 + 1 + DRAW,
            'Z' => get_winning_move(opp_move) as u32 - 'X' as u32 + 1 + WIN,
            _ => panic!("Invalid move"),
        })
        .sum::<u32>() as usize
}
