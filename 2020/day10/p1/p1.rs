use std::fs;

fn main() {

    let contents = fs::read_to_string("../input").unwrap();

    let mut lines: Vec<usize> = contents.trim().lines().map(|l| l.to_string()).map(|l| l.parse::<usize>().unwrap()).collect();

    lines.sort();

    let mut one_count: usize = 1;
    let mut three_count: usize = 1;

    for i in 1..lines.len() {
        if lines[i] - lines[i-1] == 1 {
            one_count +=1;
        } else if lines[i] - lines[i-1] == 3{
            three_count +=1;
        }
    }

    println!("{}", one_count*three_count);
}   