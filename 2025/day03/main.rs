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
    let part_two_answer = part_two(batteries.clone());

    println!("Part One: {part_one_answer}");
    println!("Part Two: {part_two_answer}");
}

fn make_joltage(battery: &str, n: i32, index: usize, number: i64) -> i64 {
    if index >= battery.len() {
        if n == 0 {
            return number;
        } 
        return 0;
    } 
    if n == 0 {
        return number;
    }
    let take_num = (battery.as_bytes()[index] as char).to_digit(10).unwrap();
    let take = make_joltage(battery, n-1, index+1, number*10 + (take_num as i64));
    let no_take = make_joltage(battery, n, index+1, number);

    if take > no_take {
        return take
    }
    return no_take
}

fn part_one(batteries: Vec<&str>) -> i64 {
    let mut sum: i64 = 0;
    for battery in batteries.iter() {
        if battery.is_empty() {
            continue
        };
        let joltage = make_joltage(&battery, 2, 0, 0);
        println!("battery: {battery} | joltage: {joltage}");
        sum += joltage;
    }
    return sum
}

fn part_two(batteries: Vec<&str>) -> i64 {
    let mut sum: i64 = 0;
    for battery in batteries.iter() {
        if battery.is_empty() {
            continue
        };
        println!("battery: {battery}");
        let joltage = make_joltage(&battery, 12, 0, 0);
        println!("joltage: {joltage} \n");
        sum += joltage;
    }
    return sum
}
