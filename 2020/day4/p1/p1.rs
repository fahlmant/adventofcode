use std::fs;
use std::collections::HashMap;

fn main() {

    let contents = fs::read_to_string("../input").unwrap();

    // Split on line breaks, strip newlines for each entry and collect into Vector
    let values: Vec<String> = contents.split("\n\n").map(|l| l.replace("\n", " ")).collect();

    let mut valid_passport_count: usize = 0;

    for value in values {
        if validate_passport(value) {
            valid_passport_count += 1;
        }
    }

    println!("{}", valid_passport_count);

}

fn validate_passport(passport_data: String) -> bool {

    // I hate this but I'm not sure how to make this global/constant
    let fields: Vec<&str> = vec!["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"];

    // Create a new HashMap of String: String
    let mut passport: HashMap<String, String> = HashMap::new();

    // Split by spaces to get each field and value
    let data: Vec<&str> = passport_data.split(" ").collect();
    
    for d in data {
        // Split the key:value in input by : and insert it into the HasMap
        let split: Vec<String> = d.split(":").map(|s| String::from(s)).collect();
        // Cloning seems funky here, but not sure what to do
        passport.insert(split[0].clone(), split[1].clone());

    }

    let mut valid: bool = true;

    for f in fields.clone() {
        // If any of the fields are missing (not including cid), passport is not valid
        if !passport.contains_key(f) {
            valid = false;
        }
    }

    return valid;
}