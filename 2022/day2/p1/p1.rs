use std::fs;

fn main() {


    
    // Use fs to get all input to a Result<String>, which is then unwrapped
    let contents = fs::read_to_string("../input").unwrap();

    let mut score = 0;

    for line in contents.trim().split('\n').enumerate() {
        let split = line.1.split(' ');
        let round: Vec<&str> = split.collect();
        
        // Rock is X and A
        // Paper is Y and B
        // Scissors is Z and C
        
        if round[1] == "X" {
            // Rock
            score += 1;
            match round[0] {
                "A" => score += 3,
                "C" => score += 6,
                _ => score += 0
            }
            
        } else if round[1] == "Y" {
            // Paper
            score += 2;
            match round[0] {
                "A" => score += 6,
                "B" => score += 3,
                _ => score += 0
            }
        
        } else if round[1] == "Z" {
            // Scissors
            score += 3;
            match round[0] {
                "B" => score += 6,
                "C" => score += 3,
                _ => score += 0
            }

        }
    }

    println!("{}", score)
    
}
