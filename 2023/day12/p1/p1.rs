use std::fs;

fn main() {

    // Use fs to get all input to a Result<String>, which is then unwrapped
    let contents = fs::read_to_string("../input").unwrap();
    let mut total = 0;

    let lines: Vec<&str> = contents.trim().split('\n').collect(); 

    let mut springs: Vec<(&str, &str)> = Vec::new();
    for line in lines {
        let split: Vec<&str> = line.split(" ").collect();
        springs.push((split[0],split[1]))
    }

    for spring in springs {
        let numbers: Vec<i32> = spring.1.split(",").map(|x| x.parse::<i32>().unwrap()).collect();
        total += get_arrangements(spring.0.to_string(), numbers);        
    }

    println!("{}", total);
}

fn get_arrangements(springs_map: String, counts: Vec<i32>) -> i32 {

    let springs_chars = springs_map.chars().collect::<Vec<char>>();
    if springs_map.len() == 0 {
        // If we're at the end of the map and all counts are accounted for, valid solution
        if counts.len() == 0 {
            return 1;
        } else {
            // If there are counts left, then it's not a valid solution
            return 0;
        }
    }
    if springs_chars[0] == '.' {
        return get_arrangements(springs_map.clone()[1..].to_string(),counts);
    }
    if springs_chars[0] == '?' {
        let mut new_map_hash = springs_map.clone().to_string();
        let mut new_map_period = springs_map.clone().to_string();
        new_map_hash.replace_range(0..1, &"?");
        new_map_period.replace_range(0..1, &".");

        // Return sum of both options of ? being a # and a .
        return get_arrangements(new_map_hash, counts.clone()) + get_arrangements(new_map_period, counts.clone())
    }
    if springs_chars[0] == '#' {
        // If there are no more counts left, but a hash, not a valid solution
        if counts.len() == 0 {
            return 0;
        }
        // If the number of elements left are less then the number in the current count, not a valid solution
        if springs_map.len() < counts[0] as usize {
            return 0;
        }
        for i in 0..counts[0] as usize {
            // If any of the characters in the string of length counts[0] are ., then not a valid solution
            if springs_chars[i] == '.' {
                return 0;
            }
        }

        if counts.len() > 1 {
            // If the next character after the counts size is a #, not a valid solution
            if springs_chars[counts[0] as usize] == '#' {
                return 0;
            }
            return get_arrangements(springs_chars[1 + counts[0] as usize..].iter().collect::<String>(), counts[1..].to_vec())
        } else {
            return get_arrangements(springs_chars[counts[0] as usize..].iter().collect::<String>(), counts[1..].to_vec())

        }
    }

    return 0;
}