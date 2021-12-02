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

    let mut answer = 0;
    // Iterate through each element of the vector with iter and enumerate
    for (i, number) in v.iter().enumerate() {
        // Nested iteration for comparison
        for (j, number2) in v.iter().enumerate() {
            // Don't check the same index
            if i+1 == j {
                if number < number2 {
                    answer += 1;
                }
            }

        }
    }

    println!("{}", answer)
}