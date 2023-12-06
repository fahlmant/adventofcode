
use std::fs;

fn main() {

    // Use fs to get all input to a Result<String>, which is then unwrapped
    let contents = fs::read_to_string("../input").unwrap();
    let mut total = 0;

    let lines:Vec<&str> = contents.trim().split('\n').collect();

    let times: Vec<i64> = lines[0].trim().split(":").skip(1).map(|x|->i64{x.replace(" ", "").parse().unwrap()}).collect();
    let distances: Vec<i64> = lines[1].trim().split(":").skip(1).map(|x|->i64{x.replace(" ", "").parse().unwrap()}).collect();
    
    for i in 0..times.len() {
        let highscore = distances[i];
        for j in 0..times[i] {
            if calculate_distance(j, times[i]) > highscore {
                total += 1;
            }
        }
    }

    println!("{}", total);
}

fn calculate_distance(time: i64, total: i64) -> i64 {

    let total_time_traveling = total - time;
    return total_time_traveling * time;
}