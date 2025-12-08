use std::{cmp::Ordering, fs, io};

fn main() {
    let input_path;
    let pop_count;
    println!("input mode: ");
    let mut mode = String::new();
    io::stdin()
        .read_line(&mut mode)
        .expect("failed to read line");
    if mode.trim() == "i" {
        input_path = "day08/src/input.txt";
        pop_count = 1000;
    } else if mode.trim() == "e" {
        input_path = "day08/src/example.txt";
        pop_count = 10;
    } else {
        println!("invalid input: either i or e is valid");
        return;
    }

    let input = match fs::read_to_string(input_path) {
        Ok(input) => input,
        Err(_) => {
            println!("couldn't read input file");
            return;
        }
    };
    let mut boxes: Vec<&str> = Vec::new();

    for line in input.lines() {
        boxes.push(line);
    }

    println!("------------------------------");
    let list = make_distance_list(&boxes);

    for item in list {
        let (a, b, c) = item;
        println!("{a}: {b}, {c}");
    }
}

fn make_distance_list<'a>(boxes: &Vec<&'a str>) -> Vec<(f32, &'a str, &'a str)> {
    let len = boxes.len();

    let mut distances: Vec<(f32, &str, &str)> = Vec::new();
    for col in 0..len {
        for row in 0..len {
            distances.push((
                calculate_distance(boxes[col], boxes[row]),
                boxes[col],
                boxes[row],
            ));
        }
    }
    distances.sort_by(|a, b| match (a.0).total_cmp(&b.0) {
        Ordering::Less => Ordering::Less,
        Ordering::Greater => Ordering::Greater,
        Ordering::Equal => Ordering::Equal,
    });
    return distances;
}

fn calculate_distance(a: &str, b: &str) -> f32 {
    let a: Vec<i32> = a.trim().split(",").map(|c| c.parse().unwrap()).collect();
    let b: Vec<i32> = b.trim().split(",").map(|c| c.parse().unwrap()).collect();
    let (ax, ay, az) = (a[0], a[1], a[2]);
    let (bx, by, bz) = (b[0], b[1], b[2]);

    f32::sqrt(((ax - bx) * (ax - bx) + (ay - by) * (ay - by) + (az - bz) * (az - bz)) as f32)
}
