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
    char::from_u32((mv as u32 - ('A' as u32) + 1) % 3 + ('X' as u32)).unwrap()
}

fn get_losing_move(mv: char) -> char {
    char::from_u32((mv as u32 - ('A' as u32) + 2) % 3 + ('X' as u32)).unwrap()
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
        .map(|round| {
            (round.1 as u32 - 'X' as u32 + 1)
                + if get_losing_move(round.0) == round.1 {
                    LOSE
                } else if get_winning_move(round.0) == round.1 {
                    WIN
                } else {
                    DRAW
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
        .map(|round| match round.1 {
            'X' => get_losing_move(round.0) as u32 - 'X' as u32 + 1 + LOSE,
            'Y' => round.0 as u32 - 'A' as u32 + 1 + DRAW,
            'Z' => get_winning_move(round.0) as u32 - 'X' as u32 + 1 + WIN,
            _ => panic!("Invalid move"),
        })
        .sum::<u32>() as usize
}
