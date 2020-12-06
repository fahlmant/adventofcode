use std::fs;
use std::collections::HashSet;

fn main() {

    let contents = fs::read_to_string("../input").unwrap();
    let values: Vec<String> = contents.split("\n\n").map(|l| l.replace("\n", " ")).collect();

    let mut yes_count: usize = 0;
    for party in values {
        let mut questions: HashSet<char> = HashSet::new();
        for character in party.chars() {
            questions.insert(character);
        }
        for c in questions {
            if party.matches(c).count() == party.matches(" ").count() + 1 {
                yes_count += 1;
            }
        }
    }

    println!("{}", yes_count)
}