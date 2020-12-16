use std::fs;

fn main() {

    let contents = fs::read_to_string("../input").unwrap();

    let lines: Vec<String> = contents.trim().lines().map(|l| l.to_string()).collect();

    let timestamp: usize = 0;

    for _ in 0..usize::MAX {
        for (i, bus) in lines[1].split(",").enumerate() {
            if bus == "x" {
                continue
            }
            let b = bus.parse::<usize>().unwrap();
            
        }
    }


    println!("{}", timestamp)

}