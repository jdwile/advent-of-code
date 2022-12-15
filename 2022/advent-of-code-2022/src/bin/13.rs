use std::{
    cmp::Ordering,
    fmt::{Display, Error, Formatter},
    str::FromStr,
};

#[aoc::main(13)]
pub fn main(input: &str) -> (usize, usize) {
    solve(input)
}

#[aoc::test(13)]
pub fn test(input: &str) -> (String, String) {
    let res = solve(input);
    (res.0.to_string(), res.1.to_string())
}

fn solve(input: &str) -> (usize, usize) {
    let p1 = part1(input);
    let p2 = part2(input);
    (p1, p2)
}

#[derive(Clone, Debug, Eq, PartialEq)]
enum Packet {
    Value(usize),
    List(Vec<Packet>),
}

impl PartialOrd for Packet {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        match (self, other) {
            (Self::Value(a), Self::Value(b)) => a.partial_cmp(b),
            (Self::List(a), Self::List(b)) => {
                let mut a = a.iter();
                let mut b = b.iter();

                loop {
                    match (a.next(), b.next()) {
                        (Some(a), Some(b)) => {
                            if let Some(ordering) = a.partial_cmp(b) {
                                if ordering != Ordering::Equal {
                                    return Some(ordering);
                                }
                            }
                        }
                        (Some(_), None) => return Some(Ordering::Greater),
                        (None, Some(_)) => return Some(Ordering::Less),
                        (None, None) => return Some(Ordering::Equal),
                    }
                }
            }
            (Self::Value(a), Self::List(_)) => Self::List(vec![Self::Value(*a)]).partial_cmp(other),
            (Self::List(_), Self::Value(b)) => self.partial_cmp(&Self::List(vec![Self::Value(*b)])),
        }
    }
}

impl Ord for Packet {
    fn cmp(&self, other: &Self) -> Ordering {
        self.partial_cmp(other).unwrap()
    }
}

impl Display for Packet {
    fn fmt(&self, f: &mut Formatter) -> Result<(), Error> {
        match self {
            Self::Value(n) => write!(f, "{n}"),
            Self::List(list) => {
                write!(f, "[")?;

                for (i, p) in list.iter().enumerate() {
                    write!(f, "{p}")?;

                    if i != list.len() - 1 {
                        write!(f, ",")?;
                    }
                }

                write!(f, "]")
            }
        }
    }
}

impl FromStr for Packet {
    type Err = ();

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let s = s.replace("10", "a");
        let mut packet = Self::List(Vec::new());
        let mut chars = s.chars();

        while let Some(c) = chars.next() {
            match c {
                '[' => {
                    let mut s = String::new();
                    let mut depth = 1;

                    while depth > 0 {
                        let c = chars.next().unwrap();

                        match c {
                            '[' => depth += 1,
                            ']' => depth -= 1,
                            _ => {}
                        }

                        s.push(c);
                    }

                    if let Ok(p) = s[..s.len() - 1].parse() {
                        if let Self::List(list) = &mut packet {
                            list.push(p);
                        }
                    }
                }
                ',' => {}
                'a' => {
                    if let Self::List(list) = &mut packet {
                        list.push(Self::Value(10));
                    }
                }
                _ => {
                    if let Self::List(list) = &mut packet {
                        list.push(Self::Value(c.to_digit(10).unwrap() as usize));
                    }
                }
            }
        }

        Ok(packet)
    }
}

fn part1(_input: &str) -> usize {
    let pairs = _input.split("\r\n\r\n");
    let mut index_sum = 0;

    for (i, pair) in pairs.enumerate() {
        let mut p = pair.lines().map(|packet| packet.parse::<Packet>().unwrap());
        let packet1 = p.next().unwrap();
        let packet2 = p.next().unwrap();

        if packet1 <= packet2 {
            index_sum += i + 1;
        }
    }

    index_sum
}

fn part2(_input: &str) -> usize {
    let pairs = _input.split("\r\n\r\n");
    let divider_one = Packet::List(vec![Packet::List(vec![Packet::Value(2)])]);
    let divider_two = Packet::List(vec![Packet::List(vec![Packet::Value(6)])]);

    let mut packets = vec![divider_one.clone(), divider_two.clone()];

    for pair in pairs {
        let mut p = pair.lines().map(|s| s.parse::<Packet>().unwrap());
        packets.push(p.next().unwrap());
        packets.push(p.next().unwrap());
    }

    packets.sort();

    let one_index = packets
        .iter()
        .position(|packet| packet == &divider_one)
        .unwrap();

    let two_index = packets
        .iter()
        .position(|packet| packet == &divider_two)
        .unwrap();

    (one_index + 1) * (two_index + 1)
}
