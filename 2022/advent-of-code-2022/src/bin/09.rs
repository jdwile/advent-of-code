use std::collections::HashSet;

#[aoc::main(09)]
pub fn main(input: &str) -> (usize, usize) {
    solve(input)
}

#[aoc::test(09)]
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
struct Position {
    x: i32,
    y: i32,
}

impl Position {
    fn is_touching(&self, other: Position) -> bool {
        self.x.abs_diff(other.x) <= 1 && self.y.abs_diff(other.y) <= 1
    }
    fn is_covering(&self, other: Position) -> bool {
        self.x == other.x && self.y == other.y
    }

    fn move_toward(&mut self, other: Position) {
        if self.x.abs_diff(other.x) >= 1 && self.y.abs_diff(other.y) >= 1 {
            self.x += (other.x - self.x).signum();
            self.y += (other.y - self.y).signum();
        } else if self.x.abs_diff(other.x) >= 1 {
            self.x += (other.x - self.x).signum();
        } else {
            self.y += (other.y - self.y).signum();
        }
    }
}

fn part1(input: &str) -> usize {
    let mut head = Position { x: 0, y: 0 };
    let mut tail = Position { x: 0, y: 0 };

    let mut tail_history = HashSet::<Position>::new();
    tail_history.insert(tail.clone());

    input.lines().for_each(|line| {
        let (dir, mag_str) = line.split_once(' ').unwrap();
        let mag = mag_str.parse::<i32>().unwrap();
        let mut dest = head;
        match dir {
            _ if dir == "U" => dest.x += mag,
            _ if dir == "D" => dest.x -= mag,
            _ if dir == "L" => dest.y -= mag,
            _ if dir == "R" => dest.y += mag,
            _ => panic!("Shouldn't happen"),
        };

        while !head.is_covering(dest) {
            head.move_toward(dest);

            while !head.is_touching(tail) {
                tail.move_toward(head);
                tail_history.insert(tail);
            }
        }
    });
    tail_history.len()
}

fn part2(input: &str) -> usize {
    let mut rope = Vec::<Position>::new();
    for _ in 0..10 {
        rope.push(Position { x: 0, y: 0 });
    }

    let mut tail_history = HashSet::<Position>::new();
    tail_history.insert(rope[rope.len() - 1]);

    input.lines().for_each(|line| {
        let (dir, mag_str) = line.split_once(' ').unwrap();
        let mag = mag_str.parse::<i32>().unwrap();

        let mut dest = rope[0];
        match dir {
            _ if dir == "U" => dest.x += mag,
            _ if dir == "D" => dest.x -= mag,
            _ if dir == "L" => dest.y -= mag,
            _ if dir == "R" => dest.y += mag,
            _ => panic!("Shouldn't happen"),
        };

        while !rope[0].is_covering(dest) {
            rope[0].move_toward(dest);

            for i in 1..rope.len() {
                let dest = rope[i - 1];
                while !rope[i].is_touching(dest) {
                    rope[i].move_toward(dest);
                    tail_history.insert(rope[rope.len() - 1]);
                }
            }
        }
    });
    tail_history.len()
}
