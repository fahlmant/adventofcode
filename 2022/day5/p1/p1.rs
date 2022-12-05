use std::{fs, ascii::AsciiExt};

use scan_fmt::scan_fmt;

fn main() {
    
    // Use fs to get all input to a Result<String>, which is then unwrapped
    let contents = fs::read_to_string("../input").unwrap();

    let (stacks, instructions) = contents.split_once("\n\n").unwrap();

    let mut box_stacks = vec![vec![];9];

    for line in stacks.lines().rev().map(str::as_bytes){
        for stack in 0..box_stacks.len() {
            // Loops through each stack in the line and sees if it has a box
            let character = line[stack*4 + 1];
            // If the character is alphabetic, i.e. not blank or a bracket []
            // then it's a box. Push that box onto the vector for that stack
            if character.is_ascii_alphabetic() {
                box_stacks[stack].push(character as char);
            }
        }
    }

    for line in instructions.lines().filter_map(|line| scan_fmt!(line, "move {d} from {d} to {d}", usize, usize, usize).ok()) {
        for _ in 0..line.0 {
            let item = box_stacks[line.1 - 1].pop().unwrap();
            box_stacks[line.2 - 1].push(item);
        }
    }

    for stack in box_stacks {
        println!("{:?}", stack.last().unwrap());
    }
}

