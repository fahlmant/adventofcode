use std::fs;

fn main() {

    // Use fs to get all input to a Result<String>, which is then unwrapped
    let contents = fs::read_to_string("../input").unwrap();
    let mut total = 0;
    let mut total_ways_to_beat: Vec<i32> = Vec::new();

    let lines:Vec<&str> = contents.trim().split('\n').collect();

    let times: Vec<i32> = lines[0].split_whitespace().skip(1).map(|x|->i32{x.parse().unwrap()}).collect();
    let distances: Vec<i32> = lines[1].split_whitespace().skip(1).map(|x|->i32{x.parse().unwrap()}).collect();
    
    for i in 0..times.len() {
        let highscore = distances[i];
        let mut ways_to_beat = 0;
        for j in 0..times[i] {
            if calculate_distance(j, times[i]) > highscore {
                ways_to_beat += 1;
            }
        }
        total_ways_to_beat.push(ways_to_beat);
    }

    total = total_ways_to_beat[0];
    for i in 1..total_ways_to_beat.len() {
        total *= total_ways_to_beat[i];
    }
    println!("{}", total);
}

fn calculate_distance(time: i32, total: i32) -> i32 {

    let total_time_traveling = total - time;
    return total_time_traveling * time;
}