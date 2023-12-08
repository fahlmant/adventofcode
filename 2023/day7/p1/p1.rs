use std::fs;
use std::collections::HashMap;

#[derive(Clone)]
pub struct hand {
    cards: Vec<String>,
    bid: i32,
    strength: i32
}

fn main() {

    // Use fs to get all input to a Result<String>, which is then unwrapped
    let contents = fs::read_to_string("../input").unwrap();
    let mut total = 0;

    let lines:Vec<&str> = contents.trim().split('\n').collect(); 

    let mut hands: Vec<hand> = Vec::new();

    for line in lines {
        let parts: Vec<&str> = line.split_whitespace().collect();
        let hand_chars: Vec<char> = parts[0].chars().collect();
        let mut hand_str: Vec<String> = Vec::new();
        for c in hand_chars {
            hand_str.push(c.to_string());
        }
        let bid = parts[1].parse().unwrap();
        let strength = determine_strength(&hand_str);
        let h = hand{cards: hand_str, bid: bid, strength: strength};
        hands.push(h);
    }

    hands.sort_by_key(|s| s.strength);

    let mut hand_groups: Vec<Vec<hand>> = vec![vec![]; 7];
    for mut h in hands {
        let mut new_cards: Vec<String> = Vec::new();
        for c in h.cards.iter().enumerate() {
            let new_char = c.1.replace("A", "14").replace("K", "13").replace("Q", "12").replace("J", "11").replace("T", "10");
            new_cards.push(new_char);
        }
        h.cards = new_cards;
        hand_groups[h.strength as usize - 1].push(h);
    }
    
    for hg in &mut hand_groups {
        hg.sort_by_key(|s: &hand| s.cards[0].parse::<i32>().unwrap());
        let len = hg.len();

        let mut swapped;
        loop {
            swapped = false;

            if len <= 0 {
                break;
            }
            for i in 0..len - 1 {
                if needs_swap(&hg[i].cards, &hg[i+1].cards) {
                    hg.swap(i, i+1);
                    swapped = true;
                    break;
                }
            }

            if !swapped {
                break;
            }
        }
    }

    let mut rank = 1;
    for hg in hand_groups {
        for h in hg {
            total += rank * h.bid;
            rank += 1;
        }
    }    
    
    println!("{}", total);
}

fn determine_strength(hand: &Vec<String>) -> i32 {
    let mut counts = HashMap::new(); 
    hand.iter().for_each( |val| {counts.entry(val) .and_modify(|count| { *count += 1 }) .or_insert(1);});

    if counts.values().any(|&x| x == 5) {
        return 7
    } else if counts.values().any(|&x| x == 4 ) {
        return 6
    } else if counts.values().any(|&x| x == 3) {
        if counts.values().any(|&x| x == 2) {
            return 5
        } else {
            return 4
        }
    } else if counts.values().any(|&x| x == 2) {
        let mut i = 0;
        for (_,v) in counts {
            if v == 2 {
                i += 1;
            }
        }
        if i == 2 {
            return 3
        } else {
            return 2
        }
    } else {
        return 1;
    }
}

fn needs_swap(cards1: &Vec<String>, cards2: &Vec<String>) -> bool {
    let mut swap = false;
    let mut i = 0;
    for j in 0..cards1.len() {
        let a = cards1[j].parse::<i32>().unwrap();
        let b = cards2[j].parse::<i32>().unwrap();
        if a == b {
            continue;
        } else {
            i = j;
            break
        }
    }

    let a = cards1[i].parse::<i32>().unwrap();
    let b = cards2[i].parse::<i32>().unwrap();
    if a > b {
        swap = true;
    }
    return swap;
}