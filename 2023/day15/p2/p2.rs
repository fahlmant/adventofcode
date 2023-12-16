use std::{fs, collections::HashMap};

fn main() {

    // Use fs to get all input to a Result<String>, which is then unwrapped
    let contents = fs::read_to_string("../input").unwrap();
    let mut total = 0;

    let sequence: Vec<String> = contents.trim().split(',').map(|x| x.to_string()).collect(); 

    let mut boxes = HashMap::new();
    for i in 0..256{
        boxes.insert(i, vec![]);
    }

    for s in sequence {
        let instruction = &s.chars().collect::<Vec<char>>();

        if instruction.contains(&'=') {
            // Split the string by the =
            let split = s.split_once('=').unwrap();
    
            // The label comes before the split and gets hashed into the box number
            let label = split.0.to_string();
            let box_number = hash_string(label.clone());

            // The length is after the split
            let length = split.1;

            // There should only be one lens with each unique label in the box
            let mut found = false;

            // Build out the lens string
            let lens_string: String = label.clone() + " " + length;

            let mut lenses: Vec<String> = boxes.get(&box_number).unwrap().clone();
            for l in lenses.clone().iter().enumerate() {
                // If the box contains a lens with that label already, update the entry with the new focal length
                if l.1.contains(&label) {
                    let mut new_lenses = lenses.clone();
                    new_lenses[l.0] = lens_string.clone();
                    boxes.insert(box_number, new_lenses);
                    found = true;
                    break;
                }
            }

            // If the box doesn't contain that label, add a new lens to the end
            if !found {
                lenses.push(lens_string);
                boxes.insert(box_number, lenses);
            }
        }

        // If there is a - in the instruction
        if instruction.contains(&'-') {
            // Get the label from everything before the -, which is the last character
            let label: String = instruction[..instruction.len() - 1].iter().collect();

            // Get the box number by hasing the label
            let box_number = hash_string(label.clone());

            // Look through existing lense in that box to see if it contains the one we're removing
            let mut lenses: Vec<String> = boxes.get(&box_number).unwrap().clone();
            for l in lenses.clone().iter().enumerate() {
                // If the box contains the target lens, remove it
                if l.1.contains(&label) {
                    lenses.remove(l.0);
                    boxes.insert(box_number, lenses.clone());
                    break;
                }
            }
        }
    }

    for b in boxes {
        let empty: Vec<String> = vec![];
        if b.1 != empty {
            for c in b.1.iter().enumerate() {
                let Some((_,length)) = c.1.split_once(" ") else {todo!()};
                let focal_length = length.parse::<i32>().unwrap();
                total += (b.0 as i32 + 1) * (c.0 as i32 + 1) * focal_length;
            }
        }
    }


    println!("{}", total);
}

fn hash_string(s: String) -> i32 {
    let mut result: i32 = 0;
    for c in s.chars() {
                result += c.clone() as i32;
                result *= 17;
                result %= 256;
    }
    return result;
}