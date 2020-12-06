use std::fs;
use std::collections::HashSet;

fn main() {

    let contents = fs::read_to_string("../input").unwrap();
    let values: Vec<String> = contents.split("\n\n").map(|l| l.replace("\n", " ")).map(|l| l.replace(" ", "")).collect();

    let mut yes_count: usize = 0;
    for party in values {
        let mut questions: HashSet<char> = HashSet::new();
        for character in party.chars() {
            questions.insert(character);
        }
        yes_count += questions.len();
    }

    println!("{}", yes_count)
}