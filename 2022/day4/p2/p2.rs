use std::fs;

use scan_fmt::scan_fmt;

fn main() {
    
    // Use fs to get all input to a Result<String>, which is then unwrapped
    let contents = fs::read_to_string("../input").unwrap();

    let answer = &contents.lines()
        .map(|l| scan_fmt!(l, "{d}-{d},{d}-{d}", i32, i32, i32, i32).unwrap())
        .filter(|(a, b, c, d)| ((a <= c && c <= b) || (c <= a && a <= d)))
        .count();

    println!("{}", answer)
}

