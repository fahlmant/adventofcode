use std::fs;

fn main() {

    let contents = fs::read_to_string("../input").unwrap();

    let sections: Vec<String> = contents.trim().split("\n\n").map(|l| l.to_string()).collect();

    let mut ranges: Vec<(usize, usize)> = Vec::new();
    for line in sections[0].clone().lines().map(|l| l.to_string()) {
        let range_rules: Vec<&str> = line.split(":").collect();
        let range_groupings: Vec<&str> = range_rules[1].split("or").collect();
        for r in range_groupings {
            let bounds: Vec<&str> = r.split("-").collect();
            ranges.push((bounds[0].trim().parse::<usize>().unwrap(), bounds[1].trim().parse::<usize>().unwrap()));
        }
    }

    let mut total = 0;

    let nearby_tickets: Vec<String> = sections[2].split("\n").map(|l| l.to_string()).collect();
    for i in 1..nearby_tickets.len() {
        let numbers: Vec<usize> = nearby_tickets[i].split(",").map(|n| n.parse::<usize>().unwrap()).collect();
        'outer: for n in numbers {
            for r in ranges.clone() {
                if n >= r.0 && n <= r.1 {
                    continue 'outer;
                }
            }
            total += n;
        }
    }

    println!("{}", total);
}