use std::fs;

fn main() {

    // Use fs to get all input to a Result<String>, which is then unwrapped
    let contents = fs::read_to_string("../input").unwrap();

    // Initalize var to hold the highest calorie count
    let mut total = 0;

    for line in contents.trim().split('\n').enumerate() {
        let mut number_list: Vec<char> = Vec::new();
        let line_str = line.1.to_string()
                .replace("zero", "zero0zero")
                .replace("one", "one1one")
                .replace("two", "two2two")
                .replace("three", "three3three")
                .replace("four", "four4four")
                .replace("five", "five5five")
                .replace("six", "six6six")
                .replace("seven", "seven7seven")
                .replace("eight", "eight8eight")
                .replace("nine", "nine9nine");
        
        for character in line_str.chars() {
            if character.is_digit(10){
                number_list.push(character);
            }
        }
        let first = number_list[0];
        let last = number_list[number_list.len() - 1];
        total += format!("{first}{last}").parse::<u32>().expect("digits can be u32");
    }

    println!("{}", total);
}