use std::fs;

fn main() {

    // Use fs to get all input to a Result<String>, which is then unwrapped
    let contents = fs::read_to_string("../input").unwrap();
    let mut total = 0;

    let lines: Vec<&str> = contents.trim().split('\n').collect(); 

    let grid: Vec<Vec<char>> = create_grid(lines); 

    let path = get_pipe_path(&grid);
    
    total = path.len() / 2;
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

fn get_start(grid: &Vec<Vec<char>>) -> Option<(usize, usize)> {
    
    for j in 0..grid.len() {
        for i in 0..grid[0].len() {
            if grid[j][i] == 'S' {
                return Some((i,j))
            }
        }
    }
    None
}

fn get_next_locations_from_pipe(grid: &Vec<Vec<char>>, (x,y): (usize, usize)) -> Option<((usize, usize), (usize,usize))> {

    if y >= grid.len() || x >= grid[0].len() {
        return None;
    }

    match grid[y][x] {
        // | connects North and South (y-1 and y +1)
        '|' => Some(((x, y.wrapping_sub(1)), (x, y+1))),
        // - connects East and West (x-1 and x+1)
        '-' => Some(((x.wrapping_sub(1), y), (x+1, y))),
        // L connects North and East (y-1 and x+11)
        'L' => Some(((x, y.wrapping_sub(1)), (x + 1, y))),
        // J connects North and West (y-1 and x-1)
        'J' => Some(((x, y.wrapping_sub(1)), (x.wrapping_sub(1), y))),
        // 7 connects South and West (y+1 and x-1)
        '7' => Some(((x.wrapping_sub(1), y), (x, y + 1))),
        // F connects South and East (y+1 and x+1)
        'F' => Some(((x, y + 1), (x + 1, y))),
        // . is ground and doesn't connect to anything
        '.' => None,
        // S is the start and doesn't tell us anything about how it connects
        'S' => None,
        _ => None,
    }   
}

fn get_pipe_path(grid: &Vec<Vec<char>>) -> Vec<(usize, usize)> {

    let mut pipe_path: Vec<(usize, usize)> = Vec::new();

    let start = get_start(grid).unwrap();

    let mut current_position = start;
    // Neightbors are always left (x-1), right(x+1), up (y-1) and down (y+1)
    let neighbors = vec![
        (current_position.0.wrapping_sub(1), current_position.1),
        (current_position.0 + 1, current_position.1),
        (current_position.0, current_position.1.wrapping_sub(1)),
        (current_position.0, current_position.1 + 1),
        ];

    // Find a connecting pipe
    for n in neighbors {
        // If the coordinate for S matches a connecting coordinate for a neighbor, we know that that neightbor connects to S
        if n == current_position {
            continue
        }
        let next_pipe = get_next_locations_from_pipe(grid, n);
        match next_pipe {
            Some((connection1 ,connection2)) => {
                if connection1 == current_position || connection2 == current_position {
                    // We've found a pipe that's connected to S
                    current_position = n;
                    break;
                }
            }
            None => {},
        }
    }

    pipe_path.push(start);

    // Follow the first pipe until we hit S again
    while grid[current_position.1][current_position.0] != 'S' {
        let (connection1, connection2) = get_next_locations_from_pipe(grid, current_position).unwrap();
        let previous_position = pipe_path.last().unwrap();
        let next_pipe: (usize, usize);
        if connection1 == *previous_position {
            next_pipe = connection2;
        } else {
            next_pipe = connection1;
        }
        pipe_path.push(current_position);
        current_position = next_pipe;

    }

    return pipe_path;
}