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
    println!("answer is {answer}");
}
