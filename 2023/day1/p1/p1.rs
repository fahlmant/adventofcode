use std::fs;

fn main() {

    // Use fs to get all input to a Result<String>, which is then unwrapped
    let contents = fs::read_to_string("../input").unwrap();

    let mut total = 0;

    for line in contents.trim().split('\n').enumerate() {
        let mut number_list: Vec<char> = Vec::new();
        for character in line.1.chars() {
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