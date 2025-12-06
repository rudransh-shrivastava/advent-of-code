use std::{fs, io};

fn main() {
    let input_path;
    let input_num_path;
    let input_op_path;

    println!("input mode: ");
    let mut mode = String::new();
    io::stdin().read_line(&mut mode).expect("failed to read line");
    if mode.trim() == "i" {
        input_path = "day06/src/input.txt";
        input_num_path = "day06/src/input_num.txt";
        input_op_path = "day06/src/input_op.txt";
        
    }else if mode.trim() == "e" {
        input_path = "day06/src/example.txt";
        input_num_path = "day06/src/example_num.txt";
        input_op_path = "day06/src/example_op.txt";
        
    } else {

        println!("invalid input: either i or e is valid");
        return
    }

    solve_part_one(&input_num_path, &input_op_path);
    solve_part_two(&input_path);
}

fn solve_part_two(input_path: &str) { 
    let input = match fs::read_to_string(input_path) {
        Ok(input) => input,
        Err(_) => {
            println!("couldn't read input file");
            return;
        },
    };
    let mut matrix: Vec<Vec<char>> = Vec::new();
    for line in input.lines() {
        let mut row: Vec<char> = Vec::new();
        for ch in line.chars() {
            row.push(ch);
        }
        matrix.push(row); 
    }
    
    let mut sum = 0;
    let mut numbers: Vec<i64> = Vec::new();
    for (row, _) in matrix[0].iter().enumerate().rev() {
        let mut num: i64 = 0;
        for col in 0..matrix.len() {
            if matrix[col][row] == '*' {
                println!("appending to list: {num}");
                numbers.push(num);
                let mut answer = 1;
                for num in numbers {
                    answer *= num;
                }
                sum += answer;
                println!("answer: mul: {answer}");
                numbers = Vec::new();
                num = 0;
            }
            if matrix[col][row] == '+' {
                println!("appending to list: {num}");
                numbers.push(num);
                let mut answer = 0;
                for num in numbers {
                    answer += num;
                }
                println!("answer: sum: {answer}");
                sum += answer;
                numbers = Vec::new();
                num = 0;
            }

            if matrix[col][row].is_numeric() {
                num = num * 10 + matrix[col][row].to_digit(10).unwrap() as i64;
            }
        }
        if num != 0 {
            println!("appending to list: {num}");
            numbers.push(num);
        }
    }

    println!("Part Two: {sum}");
}

fn part_one(matrix: &Vec<Vec<i32>>, operators: &Vec<char>) -> i64 {
    let mut sum: i64 = 0;
    for (row, op) in operators.iter().enumerate(){
        let mut result: i64 = matrix[0][row] as i64;
        for col in 1..matrix.len() {
            if *op == '*'{
                let num = matrix[col][row];
                println!("mul: {num}");
                result *= matrix[col][row] as i64;
            }
            if *op == '+' {
                let num = matrix[col][row];
                println!("add: {num}");
                result += matrix[col][row] as i64;
            }
        }
        sum += result as i64;
    } 
    return sum;
}

fn solve_part_one(input_path: &str, input_op_path: &str) {
    let input = match fs::read_to_string(input_path) {
        Ok(input) => input,
        Err(_) => {
            println!("couldn't read input file");
            return;
        },
    };
    let input_op = match fs::read_to_string(input_op_path) {
        Ok(input) => input,
        Err(_) => {
            println!("couldn't read op file");
            return;
        },
    };
    let mut matrix: Vec<Vec<i32>> = Vec::new();
    let mut operators: Vec<char> = Vec::new();

    for line in input.lines() {
        let split = line.split(" ");
        let mut row: Vec<i32> = Vec::new();
        for item in split {
            let num: i32 = match item.parse() {
                Ok(num) => num,
                Err(_) => continue,
            };
            row.push(num);
        }
        matrix.push(row);
    }
    for line in input_op.lines() {
        let split = line.split(" ");
        for item in split {
            let ch: char = match item.parse() {
                Ok(op) => op,
                Err(_) => continue,
            };
            if ch == ' ' {
                continue;
            }
            if ch == ' ' {
                continue;
            }
            operators.push(ch);
        }
    }

    if matrix[0].len() != operators.len() {
        println!("something bad has happened why is length not the same?");
        return
    }
    let part_one_answer = part_one(&matrix, &operators);
    println!("------------------------------");

    println!("Part One: {part_one_answer}");
}
