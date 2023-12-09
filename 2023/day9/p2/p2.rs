use std::fs;
use std::collections::HashMap;


fn main() {

    // Use fs to get all input to a Result<String>, which is then unwrapped
    let contents = fs::read_to_string("../input").unwrap();
    let mut total = 0;

    let lines: Vec<&str> = contents.trim().split('\n').collect(); 

    let mut line_answers: Vec<i32> = Vec::new();

    for l in lines {
        let mut value_history: Vec<i32> = l.split(" ").map(|n| n.parse::<i32>().unwrap()).collect();
        value_history.reverse();
        let mut history_diff_lasts: Vec<i32> = vec![value_history.last().unwrap().clone()];

        loop {
            let diffs: Vec<i32> = value_history.windows(2).map(|n| n[1] - n[0]).collect();
            value_history = diffs.clone();

            history_diff_lasts.push(value_history.last().unwrap().clone());

            let mut zero_count = 0;
            for i in diffs.clone() {
                if i == 0 {
                    zero_count += 1;
                }
            }
            if zero_count == diffs.len() {
                break;
            }
        }

        let mut line_answer = 0;
        for i in history_diff_lasts {
            line_answer += i;
        }

        line_answers.push(line_answer);
    }
    
    total = line_answers.iter().sum();
    
    println!("{}", total);
}
