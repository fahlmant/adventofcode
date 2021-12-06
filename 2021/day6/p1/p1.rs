use std::fs;

fn main() {
    let mut v: Vec<i32> = fs::read_to_string("../input").unwrap().trim().split(',').map(|i| i.parse::<i32>().unwrap()).collect();

    for _ in 0..80 {
        for fish in 0..v.len() {
            v[fish] = v[fish] - 1;
            if v[fish] == -1 {
                v[fish] = 6;
                v.push(8);
            }
        }
    }

    println!("{}", v.len())
}