use std::fs;

fn main() {

    // Use fs to get all input to a Result<String>, which is then unwrapped
    let contents = fs::read_to_string("../input").unwrap();

    // Initalize var to hold the highest calorie count
    let mut highest_calories = 0;
    let mut current_calories = 0;

    for line in contents.trim().split('\n').enumerate() {
        if line.1 == "" {
            if current_calories > highest_calories {
                highest_calories = current_calories;
            }
            current_calories = 0;
            continue;
        }

        current_calories += line.1.parse::<i32>().unwrap();
    }

    println!("{}", highest_calories)
    
}