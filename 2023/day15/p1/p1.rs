use std::fs;

fn main() {

    // Use fs to get all input to a Result<String>, which is then unwrapped
    let contents = fs::read_to_string("../input").unwrap();
    let mut total = 0;

    let sequence: Vec<&str> = contents.trim().split(',').collect(); 

    for s in sequence {
        let mut value = 0;
        for c in s.chars() {
            value += c as u32;
            value *= 17;
            value %= 256;
        }
        total += value;
    }

    println!("{}", total);
}

