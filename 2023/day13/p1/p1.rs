
use std::fs;

fn main() {

    // Use fs to get all input to a Result<String>, which is then unwrapped
    let contents = fs::read_to_string("../input").unwrap();
    let mut total = 0;

    let groups: Vec<&str> = contents.trim().split("\n\n").collect();

    for g in groups {
        let lines: Vec<Vec<char>> = g.split('\n').collect::<Vec<&str>>().iter().map(|x| x.chars().collect()).collect();

        for l in lines.iter().enumerate() {

            for j in 0..lines[l.0].len() - 1 {
                let col1: Vec<char> = lines.clone().iter().map(|l| l[j]).collect();
                let col2: Vec<char> = lines.clone().iter().map(|l| l[j+1]).collect();
                if col1 == col2 {
                    // Vertical match
                    total += 1 +  j as i32;
                    break;
                }
            }

            if l.0 < lines.len() -1 {
                if l.1 == &lines[l.0 + 1] {
                    // Horizontal match
                    total += 100 * (1 + l.0 as i32);
                    break;
                }
            }
        }
    }

    println!("{}", total);
}

fn vertical_reflection(grid: Vec<Vec<char>>, index: usize) -> bool {


    
    return false;
}

fn horizontal_reflection() -> bool {

}