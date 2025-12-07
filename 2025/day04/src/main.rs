use std::{fs, usize};

fn main() {
    let input = match fs::read_to_string("day04/src/input.txt") {
        Ok(input) => input,
        Err(_) => {
            println!("couldn't read file");
            return;
        }
    };
    let mut matrix: Vec<Vec<char>> = Vec::new();
    for row in input.lines() {
        let processed_row: Vec<char> = row.chars().collect();
        matrix.push(processed_row);
    }

    let part_one_answer = part_one(&matrix);
    println!("------------------------------");
    let part_two_answer = part_two(&matrix);

    println!("Part One: {part_one_answer}");
    println!("Part Two: {part_two_answer}");
}

fn is_accessible(matrix: &Vec<Vec<char>>, x: i32, y: i32) -> bool {
    let mut threshold = 0;
    let directions: Vec<(i32, i32)> = vec![
        (-1, -1),
        (0, -1),
        (1, -1),
        (1, 0),
        (1, 1),
        (0, 1),
        (-1, 1),
        (-1, 0),
    ];
    for direction in directions {
        let new_x = x + direction.0;
        let new_y = y + direction.1;
        if new_x < 0 || matrix.len() as i32 <= new_x {
            continue;
        }
        if new_y < 0 || matrix[0].len() as i32 <= new_y {
            continue;
        }
        if matrix[new_x as usize][new_y as usize] == '@' {
            threshold += 1;
        }
    }
    if threshold < 4 {
        return true;
    }
    return false;
}

fn part_one(matrix: &Vec<Vec<char>>) -> i64 {
    let mut sum: i64 = 0;
    for x in 0..matrix.len() {
        for y in 0..matrix[x].len() {
            if matrix[x][y] == '@' {
                println!("testing: {x}, {y}");
                if is_accessible(matrix, x as i32, y as i32) {
                    println!("accessible!");
                    sum += 1
                }
            }
        }
    }
    return sum;
}

fn part_two(m: &Vec<Vec<char>>) -> i64 {
    let mut sum: i64 = 0;
    let mut matrix = m.clone();
    let mut prev_ans = 0;
    loop {
        let mut coords: Vec<(usize, usize)> = Vec::new();
        for x in 0..matrix.len() {
            for y in 0..matrix[x].len() {
                if matrix[x][y] == '@' {
                    println!("testing: {x}, {y}");
                    if is_accessible(&matrix, x as i32, y as i32) {
                        println!("accessible!");
                        coords.push((x, y));
                        sum += 1
                    }
                }
            }
        }
        if sum == prev_ans {
            break;
        }
        prev_ans = sum;
        for coord in coords {
            matrix[coord.0][coord.1] = '.';
        }
    }
    return sum;
}
