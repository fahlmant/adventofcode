use std::fs;

fn main() {

    // Use fs to get all input to a Result<String>, which is then unwrapped
    let contents = fs::read_to_string("../input").unwrap();
    let mut total = 0;

    for line in contents.trim().split('\n').enumerate() {
        let card: Vec<&str> = line.1.split(":").collect();
        let values: Vec<&str> = card[1].split("|").collect();
        let winners: Vec<i32> = values[0].split_whitespace().map(|x|->i32{x.parse().unwrap()}).collect();
        let numbers: Vec<i32> = values[1].split_whitespace().map(|x|->i32{x.parse().unwrap()}).collect();   

        let mut points = 0;
        for w in winners {
            if numbers.contains(&w) {
                if points == 0 {
                    points = 1
                } else {
                    points *= 2;
                }
            }
        }
        total += points;
    }

    println!("{}", total);
}
