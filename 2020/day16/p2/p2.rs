use std::fs;

fn main() {

    let contents = fs::read_to_string("../input").unwrap();

    let lines: Vec<String> = contents.trim().lines().map(|l| l.to_string()).collect();
}