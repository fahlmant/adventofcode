use std::fs;

fn main() {

    let contents = fs::read_to_string("../input").unwrap();

    let mut full_seats: Vec<isize> = Vec::new();

    for line in contents.trim().split('\n') {
        let seat = &line.to_string().replace("F", "0").replace("B","1").replace("L", "0").replace("R","1");
        let row = isize::from_str_radix(&seat[..7], 2).unwrap();
        let column = isize::from_str_radix(&seat[7..], 2).unwrap();
    
        full_seats.push((row * 8) + column);
    }
    
    full_seats.sort();
    for i in 0..full_seats.len()-1 {
        if full_seats[i+1] != (full_seats[i] +1) {
            println!("{}", full_seats[i]+1);
            break;
        }
    }
}