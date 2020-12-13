use std::fs;

fn main() {

    let contents = fs::read_to_string("../input").unwrap();

    let lines: Vec<String> = contents.trim().lines().map(|l| l.to_string()).map(|l| l.chars().collect()).collect();

    let mut grid: Vec<Vec<char>> = vec![vec!['a'; lines[0].len()]; lines.len() ];

    for i in 0..lines.len() {
        for j in 0..lines[i].len() {
            grid[i][j] = lines[i].chars().nth(j).unwrap();
        }
    }

    let mut change = 1;
    while change > 0 {
        change = 0;
        let old_grid = grid.clone();
        for i in 0..grid.len() {
            for j in 0..grid[i].len() {
                if grid[i][j] == '.' {
                    continue
                }
                let adjacent = count_adjacent(old_grid.clone(), i, j);
                if adjacent == 0 && grid[i][j] == 'L' {
                    grid[i][j] = '#';
                    change += 1;
                }
                if adjacent >=4 && grid[i][j] == '#' {
                    grid[i][j] = 'L';
                    change += 1;
                }
                
            }
        }
        println!("Changes: {}", change);
    }

    let mut total_seats_taken = 0;
    for i in 0..grid.len() {
        total_seats_taken += grid[i].iter().filter(|&n| *n == '#').count();
    }

    println!("{}", total_seats_taken)
}


fn count_adjacent(grid: Vec<Vec<char>>, i: usize, j: usize) -> usize {

    let mut total = 0;

    if i > 0 {
        if j > 0 {if grid[i-1][j-1] == '#' {total += 1;}}
        if j < grid[i].len() - 1 {if grid[i-1][j+1] == '#' {total+=1;}}
        if grid[i-1][j] == '#' {total+=1;}
    }

    if j >  0 {
        if grid[i][j-1] == '#'{total+=1;}
    }

    if j < grid[i].len() - 1 {
        if grid[i][j+1] == '#'{total+=1;}
    }

    if i < grid.len() - 1{
        if j > 0 {if grid[i + 1][j - 1] == '#' {total += 1;}}
        if j < grid[i].len() - 1 {if grid[i+1][j+1] == '#' {total+=1;}}
        if grid[i+1][j] == '#' {total+=1;}
    }

    return total
}
