use std::fs;

fn main() {

    let contents = fs::read_to_string("../input").unwrap();

    let mut highest_seat_id: isize = 0;

    for line in contents.trim().split('\n') {
        let seat = &line.to_string().replace("F", "0").replace("B","1").replace("L", "0").replace("R","1");
        let row = isize::from_str_radix(&seat[..7], 2).unwrap();
        let column = isize::from_str_radix(&seat[7..], 2).unwrap();

        let seat_id = (row * 8) + column;
        if seat_id > highest_seat_id {
            highest_seat_id = seat_id;
        }
    }

    println!("Highest Seat ID: {}", highest_seat_id)
}