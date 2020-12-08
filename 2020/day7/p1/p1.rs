use std::fs;
use std::collections::HashSet;
use std::collections::HashMap;

type BagRules = HashMap<String, HashMap<String, usize>>;

fn main() {

    let contents = fs::read_to_string("../input").unwrap();

    let lines: Vec<&str> = contents.trim().lines().collect();

    let bag_rules: BagRules = build_bag_rules(lines);

    // HashSet to keep a list of all unique bags that can hold a shiny gold bag, directly or indirectly
    let mut can_contain_shiny_gold: HashSet<String> = HashSet::new();

    // Loop through each parent bag, and find which ones can directly hold a shiny gold bag
    for (k,v) in bag_rules.clone() {
        for (k2, _) in v {
            if k2.contains("shiny gold") {
                can_contain_shiny_gold.insert(k.clone());
            }
        }
    }

    let mut change: usize = 1;
    // Loop through all other bags to see if they can contain the bags that contain shiny gold
    while change != 0 {
        let bag_count = can_contain_shiny_gold.len();
        for bag in can_contain_shiny_gold.clone() {
            for (k,v) in bag_rules.clone() {
                for (k2, _) in v {
                    // If the parent bag can contain a bag that contains a shiny gold (or further), keep track of it
                    if k2.contains(&bag) {
                        can_contain_shiny_gold.insert(k.clone());
                    }
                }
            }
        }
        // If the length of the HashSet has changed, we know a new bag has been added to the list
        change = can_contain_shiny_gold.len() - bag_count;
    }

    println!("{}", can_contain_shiny_gold.len());

}

fn build_bag_rules(lines: Vec<&str>) -> BagRules {

    let mut bag_rules: BagRules = HashMap::new();
    for line in lines {
        // Split on contain, so 0 is the parent bag, and 1 is all the child bags
        let split: Vec<&str> = line.split(" contain ").collect();
            // Clean up parent bag to just be the color
            let parent_bag = split[0].to_string().replace(" bags", "");
            // Child bags is a list of all the possible child bags and how many of each
            let mut child_bags: HashMap<String, usize> = HashMap::new();
            // Remove spaces after commas for easier parsing, then split child bags by comma
            for bag in split[1].replace(", ", ",").split(",") {
                // Split the number from the bag, i.e. 5 bright orange bags
                let split_child_bag: Vec<&str> = bag.splitn(2, " ").collect();
                // Special case, no means 0 bags total.
                if split_child_bag[0] == "no" {
                    continue;
                }
                // Parse the bag count into a usize
                let bag_count: usize = split_child_bag[0].parse::<usize>().unwrap();
                // Strip the period and use the color as the type
                let bag_type: String = split_child_bag[1].replace(".", "").to_string();
                child_bags.insert(bag_type, bag_count);
            }
            // Parent bag is the key, list of child bags and their counts are the values
            bag_rules.insert(parent_bag, child_bags);
    }

     return bag_rules;
}