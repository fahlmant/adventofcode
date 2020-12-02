use std::fs;


struct Password {
    occurrences: (usize, usize),
    letter: char,
    password: String,
}

fn main() {


    // Use fs to get all input to a Result<String>, which is then unwrapped
    let contents = fs::read_to_string("../input").unwrap();
    /* Builds a Vec<Password> by
    1. Trimming the string of any leading or trailing whitespace (trim())
    2. Splitting out the lines based on \n (lines())
    3. Mapping each line to a var (p) and running parse_line on p
    4. Taking the resulting Password strucs from parse_line and collect()-int them into the Vec<Password>
    */
    let passwords: Vec<Password> = contents.trim().lines().map(|p| parse_line(p)).collect();

    let mut valid_count: i32 = 0;
    for word in passwords {
        let first_letter: char = word.password.chars().nth(word.occurrences.0 - 1).unwrap();
        let second_letter: char = word.password.chars().nth(word.occurrences.1 - 1).unwrap();
        if first_letter == word.letter && second_letter != word.letter {
            valid_count += 1;
        } else if second_letter == word.letter && first_letter != word.letter {
            valid_count += 1; 
        }
    }

    println!("Valid Passwords: {}", valid_count)

}

fn parse_line(content: &str) -> Password {

    // Splits the line by spaces (since input is something like "1-9 a: asdfasdf")
    let split: Vec<&str> = content.split(" ").collect();

    // Splits the first secion of line (number-number) by -, maps each part of split to i, which is then parsed to a usize, and both are collected into the Vec<usize>
    let occurrences: Vec<usize> = split[0].split("-").map(|i| i.parse::<usize>().unwrap()).collect();
    // Takes the second split ("a:"), and gets the first char from the string
    let letter: char = split[1].chars().nth(0).unwrap();
    // Takes the third split and converts to a string
    let password: String = String::from(split[2]);

    // Builds password struct and implicitly returns
    Password {
        occurrences: (occurrences[0], occurrences[1]),
        letter: letter,
        password: password
    }
}