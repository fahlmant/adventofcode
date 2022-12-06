use std::{fs};

use itertools::Itertools;

fn main() {
    
    // Use fs to get all input to a Result<String>, which is then unwrapped
    let contents = fs::read_to_string("../input").unwrap();

    let mut index = 0;
    
    for i in 0..contents.trim().len() {
        if contents[i..i + 14].chars().unique().count() == 14  {
            index = i + 14;
            break;
        }
    }

    println!("{}", index);
}

