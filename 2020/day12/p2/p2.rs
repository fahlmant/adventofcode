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
    waypoint_x: isize,
    waypoint_y: isize,
}

impl Ferry {
    fn new() -> Ferry {
        Ferry {
            x: 0,
            y: 0, 
            waypoint_x: 10,
            waypoint_y: -1, 
        }
    }

    fn move_ferry(&mut self, times: isize) {
        assert_eq!(times > 0, true);
        for _ in 0..times {
            self.x += self.waypoint_x;
            self.y += self.waypoint_y;
        }
    }

    fn move_waypoint(&mut self, dir: Direction, length: isize) {
        match dir {
            Direction::N => self.waypoint_y -= length,
            Direction::S => self.waypoint_y += length,
            Direction::E => self.waypoint_x += length,
            Direction::W => self.waypoint_x -= length,
        }
    }

    fn turn_waypoint_left(&mut self) {
        let x = self.waypoint_x;
        let y = self.waypoint_y;
        self.waypoint_x = y;
        self.waypoint_y = -x;
    }

    fn turn_waypoint_right(&mut self) {
        let x = self.waypoint_x;
        let y = self.waypoint_y;
        self.waypoint_x = -y;
        self.waypoint_y = x;
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
            'N' => ferry.move_waypoint(Direction::N, value),
            'E' => ferry.move_waypoint(Direction::E, value),
            'S' => ferry.move_waypoint(Direction::S, value),
            'W' => ferry.move_waypoint(Direction::W, value),
            'L' => {let turns = value / 90; for _ in 0..turns {ferry.turn_waypoint_left()}},
            'R' => {let turns = value / 90; for _ in 0..turns {ferry.turn_waypoint_right()}},
            'F' => {ferry.move_ferry(value)},
            _ => continue,
        }
    }
    println!("{}", ferry.x.abs() + ferry.y.abs());

}