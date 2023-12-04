use std::fs;
use std::collections::HashMap;

fn main() {

    // Use fs to get all input to a Result<String>, which is then unwrapped
    let contents = fs::read_to_string("../input").unwrap();
    let mut total = 0;


    let mut grid = vec![vec![' ' as char; 140]; 140];
    for line in contents.trim().split('\n').enumerate() {
        let mut i = 0;
        for char in line.1.chars() {
            grid[line.0][i] = char;
            i += 1;
        }
    }

    let mut star_adjacency_list: Vec<(String, i32)> = Vec::new();

    let mut i = 0;
    let mut j = 0; 
    while i < 140 {
        while j < 140 {
            if grid[i][j].is_digit(10) {
                // Seek for end of number
                let mut k = j;
                while k < 139 && grid[i][k+1].is_digit(10) {
                    k += 1;
                }                
                let mut star_map = HashMap::new();
                let offsets = vec![-1, 0, 1 as i32];
                for l in j..k+1 {
                    for x_offset in &offsets {
                        let a = x_offset + i as i32;
                        for y_offset in &offsets { 
                            let b = y_offset + l as i32;
                            if a >= 0 && a <= 139 && b >=0 && b <= 139 {
                                let target = grid[a as usize][b as usize];
                                if target == '*' {
                                    star_map.insert(a,b);
                                }
                            }
                        }
                    }
                }

                let mut digit_string: String = String::new();
                for l in j..k+1 {
                    digit_string.push(grid[i][l]);
                }
                let digit = digit_string.parse::<i32>().unwrap();
                for (a,b) in star_map {
                    let coord_string = a.to_string() + " " + &b.to_string();
                    star_adjacency_list.push((coord_string, digit));
                }

                // Increment j to end of number
                j = k + 1;

            } else {
                j += 1;
            }

        }
        i += 1;
        j = 0;
    }

    let mut coords_vec:Vec<String> = Vec::new();


    for item in &star_adjacency_list {
        if !coords_vec.contains(&item.0) {
            coords_vec.push(item.0.clone());
            if star_adjacency_list.iter().filter(|i| i.0 == item.0).count() == 2 {
                let numbers: Vec<i32> = star_adjacency_list.clone().into_iter().filter(|i| i.0 == item.0).map(|i| i.1).collect();
                let product = numbers[0] * numbers[1];
                total += product;
            }
        }
    }
    println!("{:?}", total);
}
