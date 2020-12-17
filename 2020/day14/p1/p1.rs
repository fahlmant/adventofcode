use std::fs;
use std::collections::HashMap;

fn main() {

    let contents = fs::read_to_string("../input").unwrap();

    let lines: Vec<String> = contents.trim().lines().map(|l| l.to_string()).collect();

    let mut memory: HashMap<usize, usize> = HashMap::new();

    let mut mask: String  = String::new();
    for line in lines {
        if line.contains("mask") {
            mask = line[line.find("=").unwrap() + 2..].to_string();
        } else {
            // Get the index of the memory
            let index: usize = line[4..line.find("]").unwrap()].parse().unwrap();
            // Get the desired value to insert
            let value: u64 = line[line.find("=").unwrap() + 2..].parse().unwrap();
            // Convert the value to a string of the binary
            let partial_value: String = format!{"{:b}", value};

            // Fill in the rest of the number with 0s
            let num_of_zeros = 36 - partial_value.len();
            let mut value_string: String = "".to_string();
            for _ in 0..num_of_zeros {
                value_string.push('0');
            }

            // Append the value to the end of the zeros
            value_string.push_str(&partial_value);
            
            // Apply the mask to the value
            for (i,x) in mask.chars().enumerate() {
                match x {
                    // Replace i to i + 1
                    // For some reason i..i adds an extra character
                    '0' => value_string.replace_range(i..i+1, "0"),
                    '1' => value_string.replace_range(i..i+1, "1"),
                    _ => continue,
                }
            }

            // Insert the value as a binary with the memory address into the hash map
            memory.insert(index, usize::from_str_radix(&value_string, 2).unwrap());
        }
    }
    
    let mut total = 0;
    for (_,num) in memory {
        total += num;
    }

    println!("{}", total);
}