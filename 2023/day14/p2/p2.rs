
use std::fs;

fn main() {

    // Use fs to get all input to a Result<String>, which is then unwrapped
    let contents = fs::read_to_string("../input").unwrap();
    let mut total = 0;

    let lines: Vec<&str> = contents.trim().split('\n').collect(); 

    let mut grid: Vec<Vec<char>> = create_grid(lines);
    
    let mut moved = true;
    while moved {
        moved = false;
        let mut new_grid: Vec<Vec<char>> = grid.clone();
        new_grid[0] = grid[0].clone();

        for g in grid.iter().enumerate() {
        // Skip the first line
            if g.0 != 0 {
                for i in 0..g.1.len() {
                   if g.1[i] == 'O' {
                        if grid[g.0 - 1][i] == '.' {
                            //println!("Coordinate {} and {} is {}", g.0-1, i, grid[g.0-1][i]);
                            new_grid[g.0 - 1][i] = 'O';
                            new_grid[g.0][i] = '.';
                            moved = true;
                        }   
                    }
                }
            }
        }
        grid = new_grid;
    }


    for g in grid.clone().iter().enumerate() {
        for i in 0..g.1.len() {
            if g.1[i] == 'O' {
                println!("Total += {}", grid.len() - g.0);
                total += grid.len() - g.0;
            }
        }
    }

    println!("{}", total);
}


fn create_grid(input: Vec<&str>) -> Vec<Vec<char>> {

    let mut grid: Vec<Vec<char>> = Vec::new();
    for l in input {
        let line: Vec<char> = l.chars().collect();
        grid.push(line);
    }

    return grid;
}