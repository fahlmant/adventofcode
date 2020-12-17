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
            let value: usize = line[line.find("=").unwrap() + 2..].parse().unwrap();
            // Convert the index to a string of the binary
            let partial_index: String = format!{"{:b}", index};

            // Fill in the rest of the number with 0s
            let num_of_zeros = 36 - partial_index.len();
            let mut index_string: String = "".to_string();
            for _ in 0..num_of_zeros {
               index_string.push('0');
            }

            // Append the value to the end of the zeros
            index_string.push_str(&partial_index);
            
            // Apply the mask to the index
            for (i,x) in mask.chars().enumerate() {
                match x {
                    // Replace i to i + 1
                    // For some reason i..i adds an extra character
                    'X' => index_string.replace_range(i..i+1, "X"),
                    '1' => index_string.replace_range(i..i+1, "1"),
                    _ => continue,
                }
            }

            // Count the number of Xs and generate all possible combinatations of 1s and 0s that will replace them
            let x_count: u32 = index_string.matches("X").count() as u32;
            let binary_strings = generate_binary_combinations(x_count);
           
            // For each combination of 0s and 1s
            for bs in binary_strings {
                let mut bs_index = 0;
                let mut temp_index = index_string.clone();
                // Create a clone of the index string and replace all Xs with that combos 0s and 1s
                for (i,c) in temp_index.clone().chars().enumerate() {
                    if c == 'X' {
                        temp_index.replace_range(i..i+1, bs.chars().nth(bs_index).unwrap().to_string().as_str());
                        bs_index += 1;
                    }
                }
                // Insert the value as a binary with the memory address into the hash map
                memory.insert(usize::from_str_radix(&temp_index, 2).unwrap(), value);
            }

            

        }
    }
    let mut total = 0;
    for (_,num) in memory {
        total += num;
    }

    println!("{}", total);
    
}


// Generates all binary numbers as strings between 0 and length
// For example, for a length 2, the following values will be returned:
// 00, 01, 10, 11
fn generate_binary_combinations(length: u32) -> Vec<String>{

    let mut binary_string_vec: Vec<String> = Vec::new();

    let mut total = 0;
    let base: usize = 2;
    // Calculate the highest number that the binary
    // string can represent by adding
    // 2^0 + 2^1 .. 2^length
    for i in 0..length {
        total += base.pow(i);
    }

    // For each number from 0 to highest number (inclusive)
    for x in 0..total+1 {
        // Convert number to a string of binary represention
        let mut value = format!{"{:b}", x};
        // Pad the start with 0s if the string is not long enough
        if value.len () < length as usize {
            let mut zeros: String = "".to_string();
            for _ in 0..(length as usize - value.len()) {
                zeros.push('0');
             }
             zeros.push_str(&value);
             value = zeros;
        }
        binary_string_vec.push(value);
    }

    return binary_string_vec;
}