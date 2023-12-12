extern crate itertools;

use std::fs;
use std::cmp;
use itertools::Itertools;

fn main() {

    // Use fs to get all input to a Result<String>, which is then unwrapped
    let contents = fs::read_to_string("../input").unwrap();
    let mut total = 0;

    let lines: Vec<&str> = contents.trim().split('\n').collect(); 

    let grid: Vec<Vec<char>> = create_grid(lines); 

    let mut galaxies: Vec<(usize,usize)> = Vec::new();

    let mut extra_rows: Vec<usize> = Vec::new();
    let mut extra_cols: Vec<usize> = Vec::new();

    for y in 0..grid.len() {
        // If the row has a galaxy, then it needs to be expanded
        if !grid[y].contains(&'#') {
            extra_rows.push(y);
        }

        // Build the column, and if it has a galaxy, then it needs to be expanded
        let column: Vec<char> = grid.clone().iter().map(|l| l[y]).collect();
        if !column.contains(&'#') {
            extra_cols.push(y);
        }

        for x in 0..grid[0].len() {
            if grid[y][x] == '#' {
                galaxies.push((x,y));
                continue;
            }
        }
    }

    for g in galaxies.into_iter().combinations(2) {

        let manhattan_dist = g[0].0.abs_diff(g[1].0) + g[0].1.abs_diff(g[1].1);
        let mut row_expansion = 0;
        let mut col_expansion = 0;


        for i in cmp::min(g[0].0,g[1].0)..cmp::max(g[0].0,g[1].0) {
            if extra_cols.contains(&i) {
                col_expansion += 1;
            }
        }

        for j in cmp::min(g[0].1,g[1].1)..cmp::max(g[0].1,g[1].1){
            if extra_rows.contains(&j) {
                row_expansion += 1;
            }
        }

        total += manhattan_dist + (row_expansion* (1000000-1)) + (col_expansion* (1000000-1));
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