use std::fs;

fn main() {
    let input = match fs::read_to_string("day03/input.txt") {
        Ok(input) => input,
        Err(_) => {
            println!("couldn't read file");
            return;
        },
    };
    let batteries: Vec<&str>= input.split("\n").collect();
    let part_one_answer = part_one(batteries.clone());
    println!("------------------------------");

    println!("Part One: {part_one_answer}");
}

fn make_joltage(battery: &str) -> u32 {
    let mut highest_num: u32 = 0;
    let mut highest_num_index: u32 = 0;
    for (i, b) in battery.char_indices() {
        let number = b.to_digit(10).unwrap();
        if number > highest_num {
            highest_num_index = i as u32;
            highest_num = b.to_digit(10).unwrap();
        }
    }

    let mut second_highest_num: u32 = 0;
    let mut second_highest_num_index: u32 = 0;
    for (i, b) in battery.char_indices() {
        let number = b.to_digit(10).unwrap();
        if number > second_highest_num && i as u32 != highest_num_index {
            second_highest_num_index = i as u32;
            second_highest_num = b.to_digit(10).unwrap();
        }
    }
    if highest_num_index > second_highest_num_index {
        let mut max_joltage = second_highest_num * 10 + highest_num;
        for (i, b) in battery.char_indices() {
            if (i as u32) <= highest_num_index {
                continue;
            }
            let number = b.to_digit(10).unwrap();
            let current_joltate = highest_num * 10 + number;
            if current_joltate > max_joltage {
                max_joltage = current_joltate;
            }
        }
        return max_joltage
    } 
    return highest_num * 10 + second_highest_num;
}

fn part_one(batteries: Vec<&str>) -> i64 {
    let mut sum: i64 = 0;
    for battery in batteries.iter() {
        if battery.is_empty() {
            continue
        };
        let joltage = make_joltage(&battery) as i64;
        println!("battery: {battery} | joltage: {joltage}");
        sum += joltage;
    }
    return sum
}

