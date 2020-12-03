use std::fs;

fn main() {

    let contents = fs::read_to_string("../input").unwrap();

    // Create a Vector with entry for each line in input
    let grid: Vec<&str> = contents.trim().split("\n").collect();

    let total: usize = count_trees(grid, 1,3);

    println!("{}", total);
}

fn count_trees(grid: Vec<&str>, x: usize, y: usize) -> usize {

    let mut total: usize = 0;
    // Rows is the length of the Vector
    let rows = grid.len();

    for i in 0..rows {
        // row * steps in the x direction of the slope
        let x_index =i*x;
        // column * steps in the y direction, mod by the length as each row repeats.
        let y_index =i*y % grid[i].len();

        // Get char at the x,y position
        let cell = grid[x_index].chars().nth(y_index).unwrap();

        // If the cell is a tree, add to the total count
        if cell == "#".chars().nth(0).unwrap() {
            total += 1;
        }
        
    }

    return total;
}