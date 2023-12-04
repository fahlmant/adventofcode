use std::fs;

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
                // Check each number for adjacent symbol until one is found
                let mut adjacent = false; 
                let offsets = vec![-1, 0, 1 as i32];
                for l in j..k+1 {
                        for x_offset in &offsets {
                                let a = x_offset + i as i32;
                                for y_offset in &offsets { 
                                    let b = y_offset + l as i32;
                                    if a >= 0 && a <= 139 && b >=0 && b <= 139 {
                                        let target = grid[a as usize][b as usize];
                                        if !target.is_digit(10) && target != '.' {
                                            adjacent = true;
                                            break;
                                        }
                                    }
                                }
                    }
                }

                if adjacent {
                    let mut digit_string: String = String::new();
                    for l in j..k+1 {
                        digit_string.push(grid[i][l]);
                    }
                    let digit = digit_string.parse::<i32>().unwrap();
                    total += digit
    
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

    println!("{}", total);
}
