use std::fs;

fn main() {


    
    // Use fs to get all input to a Result<String>, which is then unwrapped
    let contents = fs::read_to_string("../input").unwrap();

    let mut priority = 0;

    for line in contents.trim().split('\n').enumerate() {
 
        let length = line.1.chars().count();
        // Split the rucksuck contents in half, since both pokcets have the same number of items
        let first = &line.1[0..(length/2)];
        let second = &line.1[length/2..length];

        let mut items: Vec<char> = Vec::new();
        // Place all unique items from first into a vector
        for item in first.chars() {
            if !items.contains(&item) {
                items.push(item)
            }
        }    

        // Find the shared item (is assuming there's only one shared item per rucksack ok?)
        for item in second.chars() {
            if items.contains(&item) {
                println!("{}: {}", line.0, item);
                // Add priority to running count
                priority += calc_prio(item);
                break;
            }
        }
    }

    println!("{}", priority);   
}


fn calc_prio(item: char) -> i32 {
    if item.is_lowercase() {
        item as i32 - 96
    } else {
        item as i32 - 38
    }
}