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

    let mut start_strings: Vec<String> = Vec::new();
    for (k, _) in nodes.clone() {
        if k.chars().collect::<Vec<char>>()[2] == 'A' {
            start_strings.push(k);
        }
    }

    let mut steps_in_path: Vec<i32> = vec![0 as i32; start_strings.len()];

    for i in 0..start_strings.len() {
        let mut total_steps = 0;
        let mut done = false;

        while !done {
            for j in instructions.clone() {
                total_steps += 1;
                if j == 'L' {
                    start_strings[i] = nodes.get(&start_strings[i]).unwrap().0.clone();
                } else if j == 'R' {
                    start_strings[i] = nodes.get(&start_strings[i]).unwrap().1.clone();
                }
                if start_strings[i].chars().collect::<Vec<char>>()[2] == 'Z'{
                    done = true;
                    break;
                }
            }
        }
        steps_in_path[i] = total_steps;
    }
    
    total = find_lcm(&steps_in_path);

    println!("{}", total);
}

fn find_lcm(numbers: &Vec<i32>) -> i64 {
    let mut answer: i64 = numbers[0] as i64;

    for i in 1..numbers.len() {
        let a = numbers[i] as i64;
        answer = (a * answer) / (gcd(a, answer));
    }

    return answer;
}

fn gcd(a: i64, b: i64) -> i64 {
    if  b == 0 {
        return a;
    }
    return gcd(b, a%b);
}