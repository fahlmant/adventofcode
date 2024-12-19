use std::fs;

fn main() {

    // Use fs to get all input to a Result<String>, which is then unwrapped
    let contents = fs::read_to_string("../input").unwrap();


    let mut total = 0;
    let red = 12;
    let green = 13;
    let blue = 14;

    // Game 1: 8 green; 5 green, 6 blue, 1 red; 2 green, 1 blue, 4 red; 10 green, 1 red, 2 blue; 2 blue, 3 red
    'label: for line in contents.trim().split('\n').enumerate() {
        let first_split = line.1.split(':');
        let collection: Vec<&str> = first_split.collect();
        let rounds: Vec<&str> = collection[1].split(";").collect();
        for round in rounds {
            let dice: Vec<&str> = round.trim().split(",").collect();
            for die in dice {
                let d: Vec<&str> = die.trim().split(" ").collect();
                let num = d[0].parse::<i32>().unwrap();
                match d[1] {
                    "red" => if num > red {continue 'label},
                    "blue" => if num > blue {continue 'label},
                    "green" => if num > green {continue 'label},
                    _ => println!("Something went wrong")
                }
            }
        }
                        
        total += line.0 +1
    }

    println!("{}", total);
}