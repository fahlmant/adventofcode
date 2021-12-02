use std::fs;

struct Instruction {
    direction: String,
    magnitude: i32
}

fn main() {

    // Use fs to get all input to a Result<String>, which is then unwrapped
    let contents = fs::read_to_string("../input").unwrap();


    let instructions: Vec<Instruction> = contents.trim().lines().map(|i| parse_line(i)).collect();

    let mut x: i32 = 0;
    let mut y: i32 = 0;

    for i in instructions {
        if i.direction == "forward" {
            x += i.magnitude;
        } else if i.direction == "up" {
            y -= i.magnitude
        } else if i.direction == "down" {
            y += i.magnitude
        }
    }

    println!("{}", x*y);

}

fn parse_line(content: &str) -> Instruction {

    let split: Vec<&str> = content.split(" ").collect();
    let direction = String::from(split[0]);
    let magnitude = split[1].parse::<i32>().unwrap();

    Instruction {
        direction: direction,
        magnitude: magnitude
    }
}
