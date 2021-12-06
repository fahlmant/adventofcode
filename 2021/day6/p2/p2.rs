use std::fs;

fn main() {
    let mut v: Vec<usize> = fs::read_to_string("../input").unwrap().trim().split(',').map(|i| i.parse::<usize>().unwrap()).collect();

    let mut fishes: [usize; 9] = [0,0,0,0,0,0,0,0,0];

    for fish in v {
        fishes[fish] += 1;
    }

    for day in 0..256 {

        let new_fish = fishes[0];
        fishes[0] = fishes[1];
        fishes[1] = fishes[2];
        fishes[2] = fishes[3];
        fishes[3] = fishes[4];
        fishes[4] = fishes[5];
        fishes[5] = fishes[6];
        fishes[6] = fishes[7] + new_fish;
        fishes[7] = fishes[8];
        fishes[8] = new_fish;
    }

    let mut answer = 0;
    for i in 0..fishes.len() {
        answer += fishes[i];
    }
    
    println!("{}", answer)
}