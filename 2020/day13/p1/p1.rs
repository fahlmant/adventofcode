use std::fs;

fn main() {

    let contents = fs::read_to_string("../input").unwrap();

    let lines: Vec<String> = contents.trim().lines().map(|l| l.to_string()).collect();

    let timestamp = lines[0].parse::<usize>().unwrap();

    let mut correct_bus = 0;
    let mut min_time = usize::MAX;

    for bus in lines[1].split(",") {
        if bus == "x" {
            continue
        }
        let b = bus.parse::<usize>().unwrap();

        // This is needed to make sure the bus id has a closer timestamp over the minimum timestamp
        let time = (b - (timestamp%b)) % b;
        if time < min_time {
            min_time = time;
            correct_bus = b;
        }
    }

    println!("{}", min_time * correct_bus)

}