use std::fs;

fn main() {

    let contents = fs::read_to_string("../input").unwrap();

    let grid: Vec<&str> = contents.trim().split("\n").collect();

    let total: usize = count_trees(grid.clone(), 1,3);
    println!("{}",total);
    let total1 = count_trees(grid.clone(), 1, 1);
    println!("{}",total1);
    let total2 = count_trees(grid.clone(), 1, 5);
    println!("{}",total2);
    let total3 = count_trees(grid.clone(), 1, 7);
    println!("{}",total3);
    let total4 = count_trees(grid.clone(), 2, 1);
    println!("{}",total4);

   println!("{}", total* total1 * total2 * total3 * total4);
}

fn count_trees(grid: Vec<&str>, x: usize, y: usize) -> usize {

    let mut total: usize = 0;
    let rows = grid.len();

    let mut i = 0;
    let mut j = 0;
    while i < rows {
        let y_index =j*y % grid[i].len();
        let cell = grid[i].chars().nth(y_index).unwrap();

        if cell == "#".chars().nth(0).unwrap() {
            total += 1;
        }
        i += x;
        j +=1;
    }

    return total;
}