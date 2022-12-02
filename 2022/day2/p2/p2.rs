use std::fs;

fn main() {


    
    // Use fs to get all input to a Result<String>, which is then unwrapped
    let contents = fs::read_to_string("../input").unwrap();

    let mut score = 0;

    for line in contents.trim().split('\n').enumerate() {
        let split = line.1.split(' ');
        let round: Vec<&str> = split.collect();
    
        // Rock is worth 1 point, and represented by an A 
        // Paper is worth 2 points, and represented by a B
        // Scissors is worth 3 points, and represented by a C
        
        if round[1] == "X" {
            // lose
            // No points added for a loss
            match round[0] {
                // Scissors loses against rock
                "A" => score += 3,
                // Rock loses against paper
                "B" => score += 1,
                // Paper loses against scissors 
                "C" => score += 2,
                _ => println!("oops")
            }
            
        } else if round[1] == "Y" {
            // draw
            // add 3 to score for a draw
            score += 3;
            match round[0] {
                // Rock draws with rock
                "A" => score += 1,
                // Paper draws with paper
                "B" => score += 2,
                // Scissors draws with scissors
                "C" => score += 3,
                _ => println!("oops")
            }
        
        } else if round[1] == "Z" {
            // win
            // Add 6 to score for a win
            score += 6;
            match round[0] {
                // Paper wins against rock
                "A" => score += 2,
                // Scissors wins against paper
                "B" => score += 3,
                // Rock wins against scissors
                "C" => score += 1,
                _ => println!("oops")
            }

        }
    }

    println!("{}", score)
    
}
