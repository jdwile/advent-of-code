use scan_fmt::scan_fmt;
use std::collections::HashMap;

#[aoc::main(18)]
pub fn main(input: &str) -> (usize, usize) {
    solve(input)
}

#[aoc::test(18)]
pub fn test(input: &str) -> (String, String) {
    let res = solve(input);
    (res.0.to_string(), res.1.to_string())
}

fn solve(input: &str) -> (usize, usize) {
    let mut grid = Grid {
        items: HashMap::<Point, Value>::new(),
    };

    input.lines().for_each(|line| {
        let coords = scan_fmt!(line, "{d},{d},{d}", i64, i64, i64).unwrap();
        grid.items.insert(coords, Value::Lava);
    });

    let p1 = part1(&mut grid);
    let p2 = part2(&mut grid);
    (p1, p2)
}

type Point = (i64, i64, i64);

#[derive(Eq, PartialEq, Debug, Clone, Default)]
enum Value {
    Lava,
    #[default]
    Air,
    Vaccuum,
}

#[derive(Eq, PartialEq, Debug, Clone)]
struct Grid {
    items: HashMap<Point, Value>,
}

impl Grid {
    fn count_sides(&mut self) -> usize {
        let x_min = self.items.keys().map(|p| p.0).min().unwrap();
        let y_min = self.items.keys().map(|p| p.1).min().unwrap();
        let z_min = self.items.keys().map(|p| p.2).min().unwrap();

        let x_max = self.items.keys().map(|p| p.0).max().unwrap();
        let y_max = self.items.keys().map(|p| p.1).max().unwrap();
        let z_max = self.items.keys().map(|p| p.2).max().unwrap();

        let mut sides: usize = 0;
        for x in x_min..x_max + 1 {
            for y in y_min..y_max + 1 {
                for z in z_min..z_max + 1 {
                    if self.items.entry((x, y, z)).or_default() == &mut Value::Lava {
                        if self.items.entry((x - 1, y, z)).or_default() == &mut Value::Air {
                            sides += 1;
                        }
                        if self.items.entry((x + 1, y, z)).or_default() == &mut Value::Air {
                            sides += 1;
                        }
                        if self.items.entry((x, y - 1, z)).or_default() == &mut Value::Air {
                            sides += 1;
                        }
                        if self.items.entry((x, y + 1, z)).or_default() == &mut Value::Air {
                            sides += 1;
                        }
                        if self.items.entry((x, y, z - 1)).or_default() == &mut Value::Air {
                            sides += 1;
                        }
                        if self.items.entry((x, y, z + 1)).or_default() == &mut Value::Air {
                            sides += 1;
                        }
                    }
                }
            }
        }
        sides
    }

    fn initialize_vaccuum(&mut self) {
        let x_min = self.items.keys().map(|p| p.0).min().unwrap();
        let y_min = self.items.keys().map(|p| p.1).min().unwrap();
        let z_min = self.items.keys().map(|p| p.2).min().unwrap();

        let x_max = self.items.keys().map(|p| p.0).max().unwrap();
        let y_max = self.items.keys().map(|p| p.1).max().unwrap();
        let z_max = self.items.keys().map(|p| p.2).max().unwrap();

        for x in x_min - 1..x_max + 2 {
            for y in y_min - 1..y_max + 2 {
                for z in z_min - 1..z_max + 2 {
                    let at_x_edge = x == x_min - 1 || x == x_max + 1;
                    let at_y_edge = y == y_min - 1 || y == y_max + 1;
                    let at_z_edge = z == z_min - 1 || z == z_max + 1;

                    if at_x_edge || at_y_edge || at_z_edge {
                        self.items.insert((x, y, z), Value::Air);
                    } else if self.items.entry((x, y, z)).or_default() != &mut Value::Lava {
                        self.items.insert((x, y, z), Value::Vaccuum);
                    }
                }
            }
        }
    }

    fn propogate_air(&mut self) {
        let x_min = self.items.keys().map(|p| p.0).min().unwrap();
        let y_min = self.items.keys().map(|p| p.1).min().unwrap();
        let z_min = self.items.keys().map(|p| p.2).min().unwrap();

        let x_max = self.items.keys().map(|p| p.0).max().unwrap();
        let y_max = self.items.keys().map(|p| p.1).max().unwrap();
        let z_max = self.items.keys().map(|p| p.2).max().unwrap();

        let mut changed = false;

        for x in x_min..x_max + 1 {
            for y in y_min..y_max + 1 {
                for z in z_min..z_max + 1 {
                    if self.items[&(x, y, z)] == Value::Vaccuum
                        && (self.items[&(x - 1, y, z)] == Value::Air
                            || self.items[&(x + 1, y, z)] == Value::Air
                            || self.items[&(x, y - 1, z)] == Value::Air
                            || self.items[&(x, y + 1, z)] == Value::Air
                            || self.items[&(x, y, z - 1)] == Value::Air
                            || self.items[&(x, y, z + 1)] == Value::Air)
                    {
                        self.items.insert((x, y, z), Value::Air);
                        changed = true;
                    }
                }
            }
        }

        if changed {
            self.propogate_air();
        }
    }
}

fn part1(grid: &mut Grid) -> usize {
    grid.count_sides()
}

fn part2(grid: &mut Grid) -> usize {
    grid.initialize_vaccuum();
    grid.propogate_air();
    grid.count_sides()
}
