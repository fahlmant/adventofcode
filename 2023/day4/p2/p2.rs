use std::fs;

fn main() {

    // Use fs to get all input to a Result<String>, which is then unwrapped
    let contents = fs::read_to_string("../input").unwrap();
    let mut total = 0;

    let lines = contents.trim().split('\n');
    
    let mut card_count: Vec<i32> = vec![1 as i32; lines.clone().count()];

    let mut current_index = 0;
    for line in lines.enumerate() {
        let card: Vec<&str> = line.1.split(":").collect();
        let values: Vec<&str> = card[1].split("|").collect();
        let winners: Vec<i32> = values[0].split_whitespace().map(|x|->i32{x.parse().unwrap()}).collect();
        let numbers: Vec<i32> = values[1].split_whitespace().map(|x|->i32{x.parse().unwrap()}).collect();   

        let mut winner_numbers = 0;
        for w in winners {
            if numbers.contains(&w) {
                winner_numbers += 1;
            }
        }

        for i in current_index+1..current_index+winner_numbers+1 {
            card_count[i] += card_count[current_index];
        }

        current_index += 1;
    }

    total = card_count.iter().sum();
    println!("{}", total);
}
