use std::fs;
use std::collections::HashMap;

type BagRules = HashMap<String, HashMap<String, usize>>;

fn main() {

    let contents = fs::read_to_string("../input").unwrap();

    let lines: Vec<&str> = contents.trim().lines().collect();

    let bag_rules: BagRules = build_bag_rules(lines);

    let total_bags: usize = sum_bag("shiny gold".to_string(), bag_rules);

    println!("{}", total_bags);

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
                let bag_type: String = split_child_bag[1].replace(".", "").replace(" bags", "").replace(" bag", "").to_string();
                child_bags.insert(bag_type, bag_count);
            }
            // Parent bag is the key, list of child bags and their counts are the values
            bag_rules.insert(parent_bag, child_bags);
    }

     return bag_rules;
}

fn sum_bag(color: String, map: BagRules) -> usize {

    let mut sum: usize = 0;
    // If the bag is not a parent of any others, it only counts for 1 bag
    if !map.contains_key(&color) {
        return 1;
    }
    let child_bags = map.get(&color).unwrap();
    for (k,v) in child_bags.iter() {
        // Recursively sum all child bags
        sum += v + v*sum_bag(k.to_string() ,map.clone());
    }

    return sum;
}