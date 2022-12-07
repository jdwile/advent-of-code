use parse_display::{Display, FromStr};
use std::collections::HashMap;

#[aoc::main(07)]
pub fn main(input: &str) -> (usize, usize) {
    solve(input)
}

#[aoc::test(07)]
pub fn test(input: &str) -> (String, String) {
    let res = solve(input);
    (res.0.to_string(), res.1.to_string())
}

fn solve(input: &str) -> (usize, usize) {
    let file_system = generate_file_system(input);
    let p1 = part1(file_system.clone());
    let p2 = part2(file_system.clone());
    (p1, p2)
}

#[derive(Debug, Clone)]
struct FileSystem {
    tree: HashMap<String, Directory>,
}

impl FileSystem {
    fn calculate_directory_size(&self, dir: &String) -> usize {
        if let Some(dir) = self.tree.get(dir) {
            let file_total = dir.files.iter().map(|f| f.size).sum::<usize>();
            let dir_total = dir
                .subdirectories
                .iter()
                .map(|d| self.calculate_directory_size(d))
                .sum::<usize>();

            file_total + dir_total
        } else {
            0
        }
    }
}

#[derive(Debug, Clone)]
struct Directory {
    files: Vec<File>,
    subdirectories: Vec<String>,
}

#[derive(Display, FromStr, PartialEq, Debug, Clone)]
#[display("{size} {name}")]
struct File {
    name: String,
    size: usize,
}

fn generate_file_system(input: &str) -> FileSystem {
    let mut directory_structure = vec![];
    let mut file_tree = HashMap::new();

    for group in input.split("$").skip(1) {
        let command = group.lines().next().unwrap().trim();

        if command.starts_with("cd") {
            let path = command.split_once(" ").unwrap().1;
            match path {
                _ if path == "/" => {
                    directory_structure.clear();
                    directory_structure.push("");
                }
                _ if path == ".." => {
                    directory_structure.pop();
                }
                _ => directory_structure.push(path.trim()),
            };
        } else if command.starts_with("ls") {
            let mut subdirectories = vec![];
            let mut files = vec![];

            group.lines().skip(1).for_each(|line| {
                let (size, name) = line.split_once(" ").unwrap();

                match size {
                    _ if size == "dir" => {
                        subdirectories.push(format!("{}/{}", directory_structure.join("/"), name))
                    }
                    _ => files.push(File {
                        name: name.to_string(),
                        size: size.parse().unwrap(),
                    }),
                }
            });

            file_tree.insert(
                directory_structure.join("/"),
                Directory {
                    files,
                    subdirectories,
                },
            );
        } else {
            panic!("Shouldn't get here");
        }
    }
    FileSystem { tree: file_tree }
}

fn part1(file_system: FileSystem) -> usize {
    file_system
        .tree
        .keys()
        .map(|path| file_system.calculate_directory_size(path))
        .filter(|&size| size <= 100_000)
        .sum()
}

fn part2(file_system: FileSystem) -> usize {
    let free_space = 70_000_000 - file_system.calculate_directory_size(&String::from(""));
    file_system
        .tree
        .keys()
        .map(|path| file_system.calculate_directory_size(path))
        .filter(|&size| free_space + size >= 30_000_000)
        .min()
        .unwrap()
}
