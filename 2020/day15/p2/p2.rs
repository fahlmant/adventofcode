use std::fs;
use std::collections::HashMap;

fn main() {

    let contents = fs::read_to_string("../input").unwrap();

    let numbers: Vec<usize> = contents.trim().split(",").map(|l| l.to_string()).map(|l| l.parse::<usize>().unwrap()).collect();

    let mut game: Vec<usize> = Vec::new();

    let mut seen: HashMap<usize, Vec<usize>> = HashMap::new();
    for n in numbers.clone() {
        game.push(n);
        seen.insert(n, vec![game.len()-1]);
    }
        
    let mut temp_vec = seen.get(&0).unwrap().clone();
    game.push(0);
    temp_vec.push(game.len()-1);    
    seen.insert(0,temp_vec);

    while game.len() < 30000000 {
        if game.len() %100000 == 0{
            println!("{}", game.len());
        }
        // Get the most recent spoken number
        let num = game[game.len() - 1];
        // Get everything before the most recent spoken number
        let before_game = &game[0..game.len()-2];
        if !before_game.contains(&num) {
            // If this is the first occurance of the number, say 0
            let mut temp_vec = seen.get(&0).unwrap().clone();
            game.push(0);
            temp_vec.push(game.len()-1);
            seen.insert(0, temp_vec);
        } else {
            let seen_vec = seen.get(&num).unwrap().clone();
            let most_recent_seen = seen_vec[seen_vec.len()-1];
            let spoken_word = most_recent_seen - seen_vec[seen_vec.len()-2];
            game.push(spoken_word);

            if seen.contains_key(&spoken_word) {
                let mut temp_vec = seen.get(&spoken_word).unwrap().clone();
                temp_vec.push(game.len()-1);
                seen.insert(spoken_word, temp_vec);
            } else {
                let mut temp_vec: Vec<usize> = Vec::new();
                temp_vec.push(game.len()-1);
                seen.insert(spoken_word, temp_vec);
            }

        }
    }

    println!("{}", game[game.len()-1]);
}