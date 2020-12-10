use std::fs;

fn main() {

    let contents = fs::read_to_string("../input").unwrap();

    let lines: Vec<String> = contents.trim().lines().map(|l| l.to_string()).collect();

    let mut answer: usize = 0;
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
        answer = *current_num;
        break;
    }

    for i in 0..lines.len()-1 {
        let mut sum: usize = lines[i].parse::<usize>().unwrap();
        let mut min: usize = sum;
        let mut max: usize = sum;
        for j in i+1..lines.len() {
            let num2 = lines[j].parse::<usize>().unwrap();
            min = usize::min(min, num2);
            max = usize::max(max, num2);
            sum += num2;
            if sum == answer{
                println!("{}", min + max);
                return;
            }
            if sum > answer {
                break;
            }
        }
    }
}