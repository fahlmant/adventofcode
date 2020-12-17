use std::fs;

fn main() {

    let contents = fs::read_to_string("../input").unwrap();

    let numbers: Vec<usize> = contents.trim().split(",").map(|l| l.to_string()).map(|l| l.parse::<usize>().unwrap()).collect();

    let mut game: Vec<usize> = Vec::new();

    for n in numbers {
        game.push(n);
    }

    while game.len() < 30000000 {
        if game.len() % 10000 == 0 {
            println!("{}", game.len());
        }
        // Get the most recent spoken number
        let num = game[game.len() - 1];
        // Get everything before the most recent spoken number
        let before_game = &game[0..game.len()-2];
        if !before_game.contains(&num) {
            // If this is the first occurance of the number, say 0
            game.push(0);
            continue;
        } else {
            let mut closest_index = 0;
            // Find the closest index of the previous times the number has been spoken
            for (i,n) in game[0..game.len()-1].iter().enumerate() {
                if *n == num {
                    if i > closest_index {
                        closest_index = i;
                    }
                }
            }
            game.push((game.len() -1) - closest_index);
        }
    }

    println!("{}", game[game.len()-1]);
}