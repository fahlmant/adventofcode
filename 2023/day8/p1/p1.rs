use std::fs;
use std::collections::HashMap;


fn main() {

    // Use fs to get all input to a Result<String>, which is then unwrapped
    let contents = fs::read_to_string("../input").unwrap();
    let mut total = 0;

    let lines:Vec<&str> = contents.trim().split('\n').collect(); 

    let instructions: Vec<char> = lines[0].chars().collect();

    let mut nodes: HashMap<String, (String, String)> = HashMap::new();
    // Skip instruction line and blank line
    for i in 2..lines.len() {
        let split:Vec<&str> = lines[i].trim().split("=").collect();
        let node_mapping_string: String = split[1].replace("(", "").replace(")", "").replace(" ", "");
        let node_maping_vec: Vec<&str> = node_mapping_string.split(",").collect();
        let node_map: (String, String) = (node_maping_vec[0].to_string(), node_maping_vec[1].to_string());
        nodes.insert(split[0].trim().to_string(), node_map);
    }

    let mut done = false;
    let mut node_key: String = "AAA".to_string();
    while !done {
        for i in instructions.clone() {
            total += 1;
            if i == 'L' {
                node_key = nodes.get(&node_key).unwrap().0.clone();
            } else if i == 'R' {
                node_key = nodes.get(&node_key).unwrap().1.clone();
            }
            if node_key == "ZZZ".to_string() {
                done = true;
                break;
            }
        }
    } 
    

    
    println!("{}", total);
}
