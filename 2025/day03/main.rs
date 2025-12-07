use std::fs;

fn main() {
    let input = match fs::read_to_string("day03/input.txt") {
        Ok(input) => input,
        Err(_) => {
            println!("couldn't read file");
            return;
        }
    };
    let batteries: Vec<&str> = input.split("\n").collect();
    let part_one_answer = part_one(batteries.clone());
    println!("------------------------------");
    let part_two_answer = part_two(batteries.clone());

    println!("Part One: {part_one_answer}");
    println!("Part Two: {part_two_answer}");
}

fn make_joltage(battery: &str, n: usize) -> i64 {
    let bytes = battery.as_bytes();
    let len = bytes.len();

    let mut stack: Vec<u32> = Vec::with_capacity(n);

    for (i, &byte) in bytes.iter().enumerate() {
        let digit = (byte as char).to_digit(10).unwrap();
        let remaining = len - 1 - i;

        while let Some(&top) = stack.last() {
            if digit > top && (stack.len() + remaining >= n) {
                stack.pop();
            } else {
                break;
            }
        }
        if stack.len() < n {
            stack.push(digit);
        }
    }

    return stack
        .into_iter()
        .fold(0i64, |acc, digit| acc * 10 + digit as i64);
}

fn part_one(batteries: Vec<&str>) -> i64 {
    let mut sum: i64 = 0;
    for battery in batteries.iter() {
        if battery.is_empty() {
            continue;
        };
        let joltage = make_joltage(&battery, 2);
        println!("battery: {battery} | joltage: {joltage}");
        sum += joltage;
    }
    return sum;
}

fn part_two(batteries: Vec<&str>) -> i64 {
    let mut sum: i64 = 0;
    for battery in batteries.iter() {
        if battery.is_empty() {
            continue;
        };
        println!("battery: {battery}");
        let joltage = make_joltage(&battery, 12);
        println!("joltage: {joltage}");
        sum += joltage;
    }
    return sum;
}
