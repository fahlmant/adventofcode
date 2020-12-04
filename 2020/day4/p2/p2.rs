extern crate regex;
use std::fs;
use regex::Regex;
use std::collections::HashMap;

const REQUIRED_KEYS: [&str; 7] = ["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"];
const EYE_COLORS: [&str; 7] = ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"];

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

    for f in &REQUIRED_KEYS {
        // If any of the fields are missing (not including cid), passport is not valid
        if !passport.contains_key(*f) {
            valid = false;
            continue;
        }

        // Ensure that each key's value complies to restrictions
        match *f {
            "byr" => {let byr = passport.get("byr").unwrap().to_string().parse::<i32>().unwrap(); if byr > 2002 || byr < 1920 {valid = false}}
            "iyr" => {let iyr = passport.get("iyr").unwrap().to_string().parse::<i32>().unwrap(); if iyr > 2020 || iyr < 2010 {valid = false}}
            "eyr" => {let eyr = passport.get("eyr").unwrap().to_string().parse::<i32>().unwrap(); if eyr > 2030 || eyr < 2020 {valid = false}}
            "hgt" => {let hgt = passport.get("hgt").unwrap(); if !valid_hgt(hgt.clone()){valid = false}}
            "hcl" => {let hcl = passport.get("hcl").unwrap(); if !valid_hcl(hcl.clone()){valid = false}}
            "ecl" => {let ecl = passport.get("ecl").unwrap(); if !valid_ecl(ecl.clone()){valid = false}}
            "pid" => {let pid = passport.get("pid").unwrap(); if !valid_pid(pid.clone()){valid = false}}
            _ => continue
        }
    }

    return valid;
}

fn valid_hgt(hgt: String) -> bool {

    let mut valid = false;

    // hgt can contain cm or in
    if hgt.contains("cm") {
        // cm will be between 150cm and 193cm, so it should be 5 characters
        if hgt.len() == 5 {
            let num: i32 = hgt[..3].to_string().parse::<i32>().unwrap();
            if num <= 193 && num >= 150 {
                valid = true;
            }
        }
    } else if hgt.contains("in") {
        // in will be between 59 and 76, so it should be 4 characters
        if hgt.len() == 4 {
            let num: i32 = hgt[..2].to_string().parse::<i32>().unwrap();
            if num <=76 && num >= 59 {
                valid = true;
            }
        }

    }

    return valid;
}

fn valid_hcl(hcl: String) -> bool {
    let re: Regex = Regex::new(r"#[\d|a-f]{6}").unwrap();

    let valid = re.is_match(&hcl.to_string());

    return valid;
}

fn valid_ecl(ecl: String) -> bool {
    let mut valid = false; 

    for color in &EYE_COLORS {
        if &ecl.to_string() == color {
            valid = true;
        }
    }

    return valid;
}

fn valid_pid(pid: String)  -> bool {
    let re: Regex = Regex::new(r"[\d]{9}").unwrap();

    // Make sure pid is exactly 9 character
    let valid = re.is_match(&pid.to_string()) && pid.len() == 9;

    return valid;
}