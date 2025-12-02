use std::fs;

fn main() {
    let input = match fs::read_to_string("input.txt") {
        Ok(input) => input,
        Err(_) => {
            println!("couldn't read file");
            return;
        },
    };
    let ranges_pair: Vec<&str>= input.split(",").collect();
    let part_one_answer = part_one(ranges_pair.clone());
    let part_two_answer = part_two(ranges_pair.clone());

    println!("Part One: {part_one_answer}");
    println!("Part Two: {part_two_answer}");
}

fn is_invalid_part_one(num: String) -> bool {
    let length = num.len();
    if length % 2 != 0 {
        return false
    }
    let mid = length / 2;
    let mut first = num; 
    let second = first.split_off(mid); // mid would be the start index for second substr (returned)
    return first == second
}

fn part_one(ranges_pair: Vec<&str>) -> i64 {
    let mut sum: i64 = 0;
    for pair in ranges_pair.iter() {
        let cleaned_pair = pair.trim();
        let range : Vec<&str>= cleaned_pair.split("-").collect();
        if range.len() != 2 {
            println!("something bad has happened: {pair}: {cleaned_pair}");
            return 0; 
        }
        let first: i64 = range[0].parse().unwrap();
        let second: i64 = range[1].parse().unwrap();
        for i in first..second + 1 {
            if is_invalid_part_one(i.to_string()) {
                sum += i;
            }
        }
        println!("processed: {cleaned_pair}");
    }
    return sum
}

fn is_invalid_part_two(num: String) -> bool {
    let length = num.len();
    if length % 2 != 0 {
        return false
    }
    let mid = length / 2;
    let mut first = num; 
    let second = first.split_off(mid); // mid would be the start index for second substr (returned)
    return first == second
}

fn part_two(ranges_pair: Vec<&str>) -> i64 {
    let mut sum: i64 = 0;
    for pair in ranges_pair.iter() {
        let cleaned_pair = pair.trim();
        let range : Vec<&str>= cleaned_pair.split("-").collect();
        if range.len() != 2 {
            println!("something bad has happened: {pair}: {cleaned_pair}");
            return 0; 
        }
        let first: i64 = range[0].parse().unwrap();
        let second: i64 = range[1].parse().unwrap();
        for i in first..second + 1 {
            if is_invalid_part_two(i.to_string()) {
                sum += i;
            }
        }
        println!("processed: {cleaned_pair}");
    }
    return sum
}
