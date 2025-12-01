use std::fs;

fn main() {
    let input = match fs::read_to_string("input.txt") {
        Ok(input) => input,
        Err(_) => {
            println!("couldn't read input.txt");
            return;
        },
    };
    let parts: Vec<&str> = input.split("\n").collect();
    let part_one_answer = part_one(&parts);
    println!("part one: {part_one_answer}");
    let part_two_answer = part_two(&parts);
    println!("part two: {part_two_answer}");
}

fn part_one(parts: &Vec<&str>) -> i32 {
    let mut answer = 0; 
    let mut current_dial: i32 = 50; 

    for part in parts.iter() {
        if current_dial == 0 {
            answer = answer + 1;
        }

        let direction: char = match part.chars().nth(0) {
            Some(dir) => dir,
            None => {
                println!("Something is wrong");
                break;
            }
        };
        let step: i32 = part[1..].trim().parse().expect("should've been a number");

        if direction == 'R' {
            current_dial = (current_dial + step) % 100;
        } else if direction == 'L' {
            current_dial = (current_dial - step) % 100;
        }

    }
    return answer;
}

fn part_two(parts: &Vec<&str>) -> i32 {
    let mut answer = 0; 
    let mut current_dial: i32 = 50; 

    for part in parts.iter() {
        if part.is_empty() {
            break;
        }
        let direction: char = match part.chars().nth(0) {
            Some(dir) => dir,
            None => {
                println!("Something is wrong");
                break;
            }
        };
        let step: i32 = part[1..].trim().parse().expect("should've been a number");

        for _ in 1..step+1 {
            if direction == 'R' {
                current_dial = (current_dial + 1) % 100;
            }
            if direction == 'L' {
                current_dial = (current_dial - 1 + 100) % 100;
            }
            if current_dial == 0 {
                answer += 1;
            }
        }
    }
    return answer;

}
