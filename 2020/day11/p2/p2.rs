use std::fs;
use std::cmp;

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
                let adjacent = count_occupied(old_grid.clone(), i, j);
                if adjacent == 0 && grid[i][j] == 'L' {
                    grid[i][j] = '#';
                    change += 1;
                }
                if adjacent >=5 && grid[i][j] == '#' {
                    grid[i][j] = 'L';
                    change += 1;
                }
                
            }
        }
        println!("Changes: {}", change);
        /*for i in 0..grid.len() {
            for j in 0..grid[i].len() {
                print!("{}", grid[i][j]);
            }
            println!(" ")
        }*/
    }

    let mut total_seats_taken = 0;
    for i in 0..grid.len() {
        total_seats_taken += grid[i].iter().filter(|&n| *n == '#').count();
    }

    println!("{}", total_seats_taken)
}


fn count_occupied(grid: Vec<Vec<char>>, i: usize, j: usize) -> usize {

    let mut total = 0;

    // Check to the left
    for y in(0..j).rev() { 
        match grid[i][y] {
            'L' => break,
            '#' => {total +=1; break;},
            _ => continue,
        }
    }

    // Check to the right
    for y in j+1..grid[i].len() { 
        match grid[i][y] {
            'L' => break,
            '#' => {total +=1; break;},
            _ => continue,
        }
    }

    // Check up
    for x in (0..i).rev() {
        match grid[x][j] {
            'L' => break,
            '#' => {total +=1; break;},
            _ => continue,
        }
    }

    // Check down
    for x in i+1..grid.len() { 
        match grid[x][j] {
            'L' => break,
            '#' => {total +=1; break;},
            _ => continue,
        }
    }

    // Check up + left
    let mut x = j;
    let mut y = i;
    while x >= 1 && y >= 1 {
        x -=1;
        y -=1;
        match grid[y][x] {
            'L' => break,
            '#' => {total +=1; break;},
            _ => continue,
        }
    }
    // Check up + right
    let mut x = j;
    let mut y = i;
    while x < grid[1].len() -1 && y >=1 {
        x+=1;
        y-=1;
        match grid[y][x] {
            'L' => break,
            '#' => {total +=1; break;},
            _ => continue,
        }
    }
    // Check down + left
    let mut x = j;
    let mut y = i;
    while x >= 1 && y < grid.len() - 1 {
        x -= 1;
        y +=1;
        match grid[y][x] {
            'L' => break,
            '#' => {total +=1; break;},
            _ => continue,
        }
    }
    // Check down + right
    let mut x = j;
    let mut y = i;
    while x < grid[1].len() -1 && y < grid.len() -1 {
        x += 1;
        y += 1;
        match grid[y][x] {
            'L' => break,
            '#' => {total +=1; break;},
            _ => continue,
        }
    } 

    return total
}
