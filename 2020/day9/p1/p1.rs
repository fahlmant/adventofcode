use std::fs;

fn main() {

    let contents = fs::read_to_string("../input").unwrap();

    let lines: Vec<String> = contents.trim().lines().map(|l| l.to_string()).collect();

    'outer: for i in 25..lines.len() {
        let current_num = &lines[i].parse::<usize>().unwrap();
        for j in i-25..i {
            for k in i-25..i {
                if j != k {
                    let num1 = &lines[j].parse::<usize>().unwrap();
                    let num2 = &lines[k].parse::<usize>().unwrap();
                    if (num1 + num2) == *current_num{
                        continue 'outer;
                    }
                }
            }
        }
        println!("{}", current_num);
        break;
    }
}