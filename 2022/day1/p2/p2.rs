use std::fs;

fn main() {

    // Use fs to get all input to a Result<String>, which is then unwrapped
    let contents = fs::read_to_string("../input").unwrap();

    // Initalize current calorie counter
    let mut current_calories = 0;
    // Initalize vector to hold each elf's calorie count
    let mut elf_calorie_list: Vec<i32> = Vec::new();

    for line in contents.trim().split('\n').enumerate() {
        if line.1 == "" {
            elf_calorie_list.push(current_calories);
            current_calories = 0;
            continue;
        }

        current_calories += line.1.parse::<i32>().unwrap();
    }

    elf_calorie_list.sort();
    elf_calorie_list.reverse();

    println!("{}", elf_calorie_list[0] + elf_calorie_list[1] + elf_calorie_list[2])
}