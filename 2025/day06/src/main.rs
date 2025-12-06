use std::fs;

fn main() {
    let input_path = "day06/src/example.txt";
    let input_op_path = "day06/src/example_op.txt";
    /*
    let input_path = "day06/src/input.txt";
    let input_op_path = "day06/src/input_op.txt";
    */
    
    let input = match fs::read_to_string(input_path) {
        Ok(input) => input,
        Err(_) => {
            println!("couldn't read ranges file");
            return;
        },
    };
    let input_op = match fs::read_to_string(input_op_path) {
        Ok(input) => input,
        Err(_) => {
            println!("couldn't read ranges file");
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

