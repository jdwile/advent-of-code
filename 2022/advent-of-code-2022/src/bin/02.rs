#[aoc::main(02)]
fn main(input: &str) -> (usize, usize) {
    let p1 = part1(input);
    let p2 = part2(input);

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

fn part1(input: &str) -> usize {
    input
        .lines()
        .map(|line| {
            let moves = line.split_once(" ").unwrap();
            (
                moves.0.chars().nth(0).unwrap(),
                moves.1.chars().nth(0).unwrap(),
            )
        })
        .map(|(opp_move, my_move)| {
            (my_move as u32 - 'X' as u32 + 1)
                + match my_move {
                    _win if get_winning_move(opp_move) == my_move => WIN,
                    _lose if get_losing_move(opp_move) == my_move => LOSE,
                    _ => DRAW,
                }
        })
        .sum::<u32>() as usize
}

fn part2(input: &str) -> usize {
    input
        .lines()
        .map(|line| {
            let moves = line.split_once(" ").unwrap();
            (
                moves.0.chars().nth(0).unwrap(),
                moves.1.chars().nth(0).unwrap(),
            )
        })
        .map(|(opp_move, strategy)| match strategy {
            'X' => get_losing_move(opp_move) as u32 - 'X' as u32 + 1 + LOSE,
            'Y' => opp_move as u32 - 'A' as u32 + 1 + DRAW,
            'Z' => get_winning_move(opp_move) as u32 - 'X' as u32 + 1 + WIN,
            _ => panic!("Invalid move"),
        })
        .sum::<u32>() as usize
}
