use std::{cmp::max, cmp::min, fs};

fn main() {
    /*
    let input_ranges_path = "day05/src/example_ranges.txt";
    let input_ids_path = "day05/src/example_ids.txt";
    */
    let input_ranges_path = "day05/src/input_ranges.txt";
    let input_ids_path = "day05/src/input_ids.txt";
    let input_ranges = match fs::read_to_string(input_ranges_path) {
        Ok(input) => input,
        Err(_) => {
            println!("couldn't read ranges file");
            return;
        }
    };
    let mut ranges: Vec<(i64, i64)> = Vec::new();
    for range in input_ranges.lines() {
        let pair: Vec<&str> = range.split("-").collect();
        let first: i64 = pair[0].parse().unwrap();
        let second: i64 = pair[1].parse().unwrap();

        ranges.push((first, second));
    }

    let input_ids = match fs::read_to_string(input_ids_path) {
        Ok(input) => input,
        Err(_) => {
            println!("couldn't read ranges file");
            return;
        }
    };
    let mut ids: Vec<i64> = Vec::new();
    for id in input_ids.lines() {
        ids.push(id.parse().unwrap());
    }

    ranges.sort();
    let part_one_answer = part_one(&ids, &ranges);
    println!("------------------------------");
    let part_two_answer = part_two(&ranges);

    println!("Part One: {part_one_answer}");
    println!("Part One: {part_two_answer}");
}

fn is_fresh(id: i64, ranges: &Vec<(i64, i64)>) -> bool {
    for range in ranges {
        if range.0 <= id && id <= range.1 {
            println!("fresh id found: {id}");
            return true;
        }
    }
    return false;
}

fn part_one(ids: &Vec<i64>, ranges: &Vec<(i64, i64)>) -> i64 {
    let mut answer = 0;
    for id in ids {
        println!("checking id: {id}");
        if is_fresh(*id, &ranges) {
            answer += 1;
        }
    }
    return answer;
}

fn is_overlapping(a: (i64, i64), b: (i64, i64)) -> bool {
    if b.0 <= a.0 && a.0 <= b.1 {
        return true;
    }
    if b.0 <= a.1 && a.1 <= b.1 {
        return true;
    }
    if a.0 <= b.0 && b.1 <= a.1 {
        return true;
    }
    if b.0 <= a.0 && a.1 <= b.1 {
        return true;
    }
    return false;
}

fn merge_two_ranges(a: (i64, i64), b: (i64, i64)) -> (i64, i64) {
    return (min(a.0, b.0), max(a.1, b.1));
}

fn merge_ranges(ranges: &Vec<(i64, i64)>) -> Vec<(i64, i64)> {
    let mut merged: Vec<(i64, i64)> = Vec::new();

    merged.push(ranges[0]);
    for i in 1..ranges.len() {
        let last_index = merged.len() - 1;
        let last = merged[last_index];
        if is_overlapping(ranges[i], last) {
            let merged_ranges = merge_two_ranges(ranges[i], last);
            merged[last_index] = merged_ranges;
        } else {
            merged.push(ranges[i]);
        }
    }
    return merged;
}

fn part_two(ranges: &Vec<(i64, i64)>) -> i64 {
    let mut count = 0;
    let ranges_clone = ranges.clone();
    let new_ranges = merge_ranges(&ranges_clone);

    for range in new_ranges {
        let start = range.0;
        let end = range.1;
        println!("range: {start} - {end}");
        count += end - start + 1;
    }
    return count;
}
