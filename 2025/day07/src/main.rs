use std::{fs, io};

fn main() {
    let input_path;

    println!("input mode: ");
    let mut mode = String::new();
    io::stdin().read_line(&mut mode).expect("failed to read line");
    if mode.trim() == "i" {
        input_path = "day07/src/input.txt";
        
    }else if mode.trim() == "e" {
        input_path = "day07/src/example.txt";
        
    } else {
        println!("invalid input: either i or e is valid");
        return
    }
    
    let input = match fs::read_to_string(input_path) {
        Ok(input) => input,
        Err(_) => {
            println!("couldn't read input file");
            return;
        },
    };

    let mut matrix: Vec<Vec<char>> = Vec::new();
    for line in input.lines() {
        let mut temp: Vec<char> = Vec::new();
        for ch in line.chars() {
            temp.push(ch);
        }
        matrix.push(temp);
    }
    let mut start_pos: (usize, usize) = (0, 0);
    for (c, column) in matrix.iter().enumerate(){
        for (r, row)  in column.iter().enumerate() {
            if *row == 'S' {
                start_pos = (c, r);
            }
        }
    }
    let part_two_matrix = matrix.clone(); 
    println!("beginning from: {}, {}", start_pos.0, start_pos.1);

    println!("------------------------------");
    let part_one_answer = start_beam(&mut matrix, start_pos);
    for col in matrix {
        for row in col {
            print!("{row} ");
        }
        println!();
    }
    let part_two_answer = part_two(&part_two_matrix, start_pos) + 1;
    println!("------------------------------");
    println!("Part One: {part_one_answer}");
    println!("Part Two: {part_two_answer}");
}

/// start_pos: (col, row)
fn start_beam(matrix: &mut Vec<Vec<char>>, start_pos: (usize, usize)) -> i64 {
    let mut count = 0;
    for col in start_pos.0..matrix.len() {
        if matrix[col][start_pos.1] == '.' {
            matrix[col][start_pos.1] = '|';
        }
        if matrix[col][start_pos.1] == '^' {
            let mut split = 0;
            if matrix[col][start_pos.1-1] == '.' {
                println!("found a new beam: {}, {}", col, start_pos.1-1);
                split += 1;
                count += start_beam(matrix, (col, start_pos.1-1));
            }
            if matrix[col][start_pos.1+1] == '.' {
                println!("found a new beam: {}, {}", col, start_pos.1+1);
                split += 1;
                count += start_beam(matrix, (col, start_pos.1+1));
            }
            if split > 0 {
                count += 1;
            }
            return count;
        }
    }
    return count;
}

fn part_two(matrix: &Vec<Vec<char>>, start_pos: (usize, usize)) -> i64 {
    let mut count = 0;
    for col in start_pos.0..matrix.len() {
        if matrix[col][start_pos.1] == '^' {
            count += 1;
            count += part_two(matrix, (col, start_pos.1-1));
            count += part_two(matrix, (col, start_pos.1+1));
            return count;
        }
    }
    return count;
}
