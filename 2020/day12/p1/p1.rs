use std::fs;

#[derive(Clone)]
enum Direction {
    N,
    E,
    S,
    W,
}

struct Ferry {
    x: isize,
    y: isize,
    direction: Direction,
}

impl Ferry {
    fn new() -> Ferry {
        Ferry {
            x: 0,
            y: 0, 
            direction: Direction::E,
        }
    }

    fn move_ferry(&mut self, dir: Direction, length: isize) {
        match dir {
            Direction::N => self.y -= length,
            Direction::S => self.y += length,
            Direction::E => self.x += length,
            Direction::W => self.x -= length,
        }
    }

    fn turn_left(&mut self) {
        match self.direction {
            Direction::N => self.direction = Direction::W,
            Direction::S => self.direction = Direction::E,
            Direction::E => self.direction = Direction::N,
            Direction::W => self.direction = Direction::S,  
        }
    }

    fn turn_right(&mut self) {
        match self.direction {
            Direction::N => self.direction = Direction::E,
            Direction::S => self.direction = Direction::W,
            Direction::E => self.direction = Direction::S,
            Direction::W => self.direction = Direction::N,  
        }
    }
}
 
fn main() {

    let contents = fs::read_to_string("../input").unwrap();

    let lines: Vec<String> = contents.trim().lines().map(|l| l.to_string()).collect();


    let mut ferry: Ferry = Ferry::new();
    for l in lines {
        let action = l.chars().nth(0).unwrap();
        let value = l[1..].to_string().parse::<isize>().unwrap();
        match action  {
            'N' => ferry.move_ferry(Direction::N, value),
            'E' => ferry.move_ferry(Direction::E, value),
            'S' => ferry.move_ferry(Direction::S, value),
            'W' => ferry.move_ferry(Direction::W, value),
            'L' => {let turns = value / 90; for _ in 0..turns {ferry.turn_left()}},
            'R' => {let turns = value / 90; for _ in 0..turns {ferry.turn_right()}},
            'F' => {ferry.move_ferry(ferry.direction.clone(), value)},
            _ => continue,
        }
    }
    println!("{}", ferry.x.abs() + ferry.y.abs());

}