use std::fs;

fn main() {

    // Create a mutable vector of i32
    let mut v: Vec<i32> = Vec::new();

    // Use fs to get all input to a Result<String>, which is then unwrapped
    let contents = fs::read_to_string("../input").unwrap();

    // For every enumerated line (split \n)
    for number in contents.trim().split('\n').enumerate() {
        // enumerate here returns (usize, &str), so
        // taking number.1 gets the &str, then its parsed into
        // a Result<i32>, which is unwrapped and pushed into the vector v
        v.push(number.1.parse::<i32>().unwrap())
    }   

    let answer = v.windows(4).filter(|d| d[3] > d[0]).count();

    println!("{}", answer)
}