use std::{cmp::Ordering, collections::HashMap, fs, io};

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

    let distances = make_distance_list(&boxes);
    let part_one_answer = part_one(&boxes, &distances, pop_count);
    let part_two_answer = part_two(&boxes, &distances);

    println!("\n------------------------------");
    println!("Part One: {part_one_answer}");
    println!("Part Two: {part_two_answer}");
}
fn part_two(boxes: &Vec<&str>, distances: &Vec<(f32, usize, usize)>) -> i64 {
    let len = boxes.len();
    let mut union_find = UnionFind::new(len);
    let mut last_connection = (0, 0);

    for d in distances {
        if union_find.find(d.1) == union_find.find(d.2) {
            continue;
        }
        union_find.union(d.1, d.2);
        last_connection = (d.1, d.2);

        let root = union_find.find(0);
        let mut all_connected = true;
        for idx in 1..len {
            if union_find.find(idx) != root {
                all_connected = false;
                break;
            }
        }

        if all_connected {
            break;
        }
    }

    let box_a = boxes[last_connection.0];
    let box_b = boxes[last_connection.1];
    let x_a: i64 = box_a.split(",").next().unwrap().trim().parse().unwrap();
    let x_b: i64 = box_b.split(",").next().unwrap().trim().parse().unwrap();

    let result = x_a * x_b;
    println!(
        "Last connection: box {} ({}) and box {} ({})",
        last_connection.0, box_a, last_connection.1, box_b
    );
    println!("X coordinates: {} * {} = {}", x_a, x_b, result);

    result
}

fn part_one(boxes: &Vec<&str>, distances: &Vec<(f32, usize, usize)>, pop_count: usize) -> i64 {
    let len = boxes.len();
    let mut counter = 0;
    let mut union_find = UnionFind::new(len);
    for d in distances {
        if counter == pop_count {
            break;
        }
        counter += 1;
        if union_find.find(d.1) == union_find.find(d.2) {
            continue;
        }
        union_find.union(d.1, d.2);
    }
    let map = union_find.unions();

    for (key, val) in map.iter() {
        print!("\n{key}: ");
        for v in val {
            print!("{v} ");
        }
    }
    let mut sizes: Vec<usize> = map.values().map(|v| v.len()).collect();
    sizes.sort_by(|a, b| b.cmp(a));

    let result: i64 = sizes[0] as i64 * sizes[1] as i64 * sizes[2] as i64;
    println!(
        "Top 3 sizes: {} * {} * {} = {}",
        sizes[0], sizes[1], sizes[2], result
    );
    result
}

struct UnionFind {
    parent: Vec<usize>,
}

impl UnionFind {
    pub fn new(len: usize) -> Self {
        let mut parent: Vec<usize> = Vec::with_capacity(len);
        for idx in 0..len {
            parent.push(idx);
        }
        Self { parent }
    }
    pub fn union(&mut self, a: usize, b: usize) {
        let parent_a = self.find(a);
        let parent_b = self.find(b);
        if parent_a != parent_b {
            self.parent[parent_a] = parent_b;
        }
    }
    pub fn find(&mut self, element: usize) -> usize {
        if self.parent[element] != element {
            self.parent[element] = self.find(self.parent[element]);
        }
        self.parent[element]
    }
    pub fn unions(&mut self) -> HashMap<usize, Vec<usize>> {
        let mut map: HashMap<usize, Vec<usize>> = HashMap::new();
        for idx in 0..self.parent.len() {
            let root = self.find(idx);
            map.entry(root).or_insert(Vec::new()).push(idx)
        }
        map
    }
}

fn make_distance_list(boxes: &Vec<&str>) -> Vec<(f32, usize, usize)> {
    let len = boxes.len();

    let mut distances: Vec<(f32, usize, usize)> = Vec::new();
    for col in 0..len {
        for row in col + 1..len {
            if boxes[col] == boxes[row] {
                continue;
            }
            distances.push((calculate_distance(boxes[col], boxes[row]), col, row));
        }
    }
    distances.sort_by(|a, b| match (a.0).total_cmp(&b.0) {
        Ordering::Less => Ordering::Less,
        Ordering::Greater => Ordering::Greater,
        Ordering::Equal => Ordering::Equal,
    });
    for (idx, d) in distances.clone().iter().enumerate() {
        if idx < 30 {
            println!("a: {} | b: {} | distance: {}", d.1, d.2, d.0);
        } else {
            break;
        };
    }
    return distances;
}

fn calculate_distance(a: &str, b: &str) -> f32 {
    let a: Vec<i64> = a.trim().split(",").map(|c| c.parse().unwrap()).collect();
    let b: Vec<i64> = b.trim().split(",").map(|c| c.parse().unwrap()).collect();
    let (ax, ay, az) = (a[0], a[1], a[2]);
    let (bx, by, bz) = (b[0], b[1], b[2]);

    f32::sqrt(((ax - bx) * (ax - bx) + (ay - by) * (ay - by) + (az - bz) * (az - bz)) as f32)
}
