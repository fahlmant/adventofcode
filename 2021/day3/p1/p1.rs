use std::fs;

fn main() {

    // Create a mutable vector of Strings
    let v: Vec<Vec<char>> = fs::read_to_string("../input").unwrap().trim().split('\n').map(|line| line.chars().collect()).collect();

    // Base 2 for conversion to actual nubmers later
    let base: i32 = 2;

    // The number of bits in a given data point
    let number_len = v[0].len();
    // The frequence of 1s for each bit
    let mut frequencies = Vec::new();

    // 
    for i in 0..number_len {
        let mut frequency = 0;

        for d in v.iter() {
            if d[i] == '1' {
                frequency += 1;
            }
        }

        frequencies.push(frequency);
    }


    let mut gamma_rate = 0;
    let mut epsilon_rate = 0;

    for (i,freq) in frequencies.iter().enumerate() {
        let power = number_len as i32 - i as i32 - 1;
        let value = match power {
            -1 => 1,
            _ => base.pow(power as u32)
        };
        
        if freq < &(v.len()/2) {
            epsilon_rate += value;
        } else {
            gamma_rate += value;
        }
    }
    
    println!("{}", gamma_rate * epsilon_rate);
}