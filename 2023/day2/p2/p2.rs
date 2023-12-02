use std::fs;

fn main() {

    // Use fs to get all input to a Result<String>, which is then unwrapped
    let contents = fs::read_to_string("../input").unwrap();


    let mut total = 0;

    for line in contents.trim().split('\n').enumerate() {
        let first_split = line.1.split(':');
        let collection: Vec<&str> = first_split.collect();

        let mut red_max = 0;
        let mut blue_max = 0;
        let mut green_max = 0;
        let rounds: Vec<&str> = collection[1].split(";").collect();
        for round in rounds {
            let dice: Vec<&str> = round.trim().split(",").collect();
            for die in dice {
                let d: Vec<&str> = die.trim().split(" ").collect();
                let num = d[0].parse::<i32>().unwrap();
                match d[1] {
                    "red" => if num > red_max {red_max = num},
                    "blue" => if num > blue_max {blue_max = num},
                    "green" => if num > green_max {green_max = num},
                    _ => println!("Something went wrong")
                }
            }
        }
                        
        total += (red_max*blue_max*green_max)
    }

    println!("{}", total);
}