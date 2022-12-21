use std::{collections::HashMap, hash::Hash};

#[aoc::main(21)]
pub fn main(input: &str) -> (i64, i64) {
    solve(input)
}

#[aoc::test(21)]
pub fn test(input: &str) -> (String, String) {
    let res = solve(input);
    (res.0.to_string(), res.1.to_string())
}

fn solve(input: &str) -> (i64, i64) {
    let p1 = part1(input);
    let p2 = part2(input);
    (p1, p2)
}

#[derive(Hash, Eq, PartialEq, Debug, Clone, Copy)]
enum Operation {
    Addition,
    Subtraction,
    Multiplication,
    Division,
    Number,
}

#[derive(Hash, Eq, PartialEq, Debug, Clone)]
struct Expression {
    operation: Operation,
    operator: String,
    operand: String,
    result: i64,
}

fn solve_monkey(
    monkey: String,
    monkies: HashMap<String, Expression>,
    monkies_memo: &mut HashMap<String, i64>,
) -> i64 {
    if monkies_memo.contains_key(&monkey) {
        return monkies_memo[&monkey];
    }

    let expression = monkies[&monkey].clone();

    let result = match expression.operation {
        Operation::Number => expression.result,
        Operation::Addition => {
            let operator_result = solve_monkey(expression.operator, monkies.clone(), monkies_memo);
            let operand_result = solve_monkey(expression.operand, monkies, monkies_memo);

            operator_result + operand_result
        }
        Operation::Subtraction => {
            let operator_result = solve_monkey(expression.operator, monkies.clone(), monkies_memo);
            let operand_result = solve_monkey(expression.operand, monkies, monkies_memo);

            operator_result - operand_result
        }
        Operation::Multiplication => {
            let operator_result = solve_monkey(expression.operator, monkies.clone(), monkies_memo);
            let operand_result = solve_monkey(expression.operand, monkies, monkies_memo);

            operator_result * operand_result
        }
        Operation::Division => {
            let operator_result = solve_monkey(expression.operator, monkies.clone(), monkies_memo);
            let operand_result = solve_monkey(expression.operand, monkies, monkies_memo);

            operator_result / operand_result
        }
    };

    monkies_memo.insert(monkey, result);
    result
}

fn part1(input: &str) -> i64 {
    let mut monkies = HashMap::<String, Expression>::new();

    input.lines().for_each(|line| {
        let (monkey, operation) = line.split_once(": ").unwrap();
        let expression: Expression;

        if operation.contains('+') {
            let (operator, operand) = operation.split_once(" + ").unwrap();

            expression = Expression {
                operation: Operation::Addition,
                operator: operator.to_string(),
                operand: operand.to_string(),
                result: 0,
            };
        } else if operation.contains('-') {
            let (operator, operand) = operation.split_once(" - ").unwrap();

            expression = Expression {
                operation: Operation::Subtraction,
                operator: operator.to_string(),
                operand: operand.to_string(),
                result: 0,
            };
        } else if operation.contains('*') {
            let (operator, operand) = operation.split_once(" * ").unwrap();

            expression = Expression {
                operation: Operation::Multiplication,
                operator: operator.to_string(),
                operand: operand.to_string(),
                result: 0,
            };
        } else if operation.contains('/') {
            let (operator, operand) = operation.split_once(" / ").unwrap();

            expression = Expression {
                operation: Operation::Division,
                operator: operator.to_string(),
                operand: operand.to_string(),
                result: 0,
            };
        } else {
            expression = Expression {
                operation: Operation::Number,
                operator: "".to_string(),
                operand: "".to_string(),
                result: operation.parse().unwrap(),
            };
        }

        monkies.insert(monkey.to_string(), expression);
    });

    // monkies
    //     .keys()
    //     .for_each(|monkey| println!("{}: {:?}", monkey, monkies[monkey]));
    solve_monkey(
        "root".to_string(),
        monkies,
        &mut HashMap::<String, i64>::new(),
    )
}

fn part2(_input: &str) -> i64 {
    0
}
