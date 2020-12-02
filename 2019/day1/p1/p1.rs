use std::fs;

fn main() {
    let contents = fs::read_to_string("../input").unwrap();

    let masses: Vec<i32> = contents.trim().lines().map(|i| i.parse::<i32>().unwrap()).collect();
    let mut sum: i32 = 0;

    for mass in masses {
        sum += (mass/3) - 2;
    }

    println!("{}", sum);
}