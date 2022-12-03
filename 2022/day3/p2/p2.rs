use std::fs;

fn main() {

    // include_bytes opens a file as a bytes
    let priority = include_bytes!("../input")
    // Split byte array by \n
    .split(|b| *b == b'\n')
    // Place each line in a vector
    .collect::<Vec<_>>()
    // Create chunks of 3 lines
    .chunks(3)
    // For each chunk of 3 lines, iterate through the characters of the first set
    // and see if that character is in both the other sets
    .map(|set| set[0]
        .iter()
        .find(|b| set[1].contains(b) && set[2].contains(b))
        .unwrap())
    // For any character that's in all 3 sets, pass to calc_prio
    .map(|b| calc_prio(*b as char))
    // Sum all priorities
    .sum::<i32>();


    println!("{}",priority);  
}


fn calc_prio(item: char) -> i32 {
    if item.is_lowercase() {
        item as i32 - 96
    } else {
        item as i32 - 38
    }
}