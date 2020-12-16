use std::fs;

fn main() {

    let contents = fs::read_to_string("../input").unwrap();

    let lines: Vec<String> = contents.trim().lines().map(|l| l.to_string()).collect();

    let mut constraints: Vec<(usize, usize)> = Vec::new();

    // Bus b with index i departs at time T + i, which means T + i is a multiple of k given the input rules
    // This means that T + i % b == 0
    // This also means that T % b == -i
    // To ensure T%b is between 0 and b, this means T % b == (b-(i%b)) % b
    // Chinese remainder theorem: Let N be the product of the bus IDs in our input.lines
    // There exists one T < N that satisfies the above rules
    let mut N = 1;
    for (i, bus) in lines[1].split(",").enumerate() {
        if bus == "x" {
            continue
        }
        let b = bus.parse::<usize>().unwrap();
        let j = i % b;
        constraints.push(((b-j)%b,b));
        N *= b;
    }

    let mut total = 0;
    for x in constraints {
        let i = x.0;
        let b = x.1;
        let NI = N/x.1;
        let mi = mod_inverse(NI, b);
        // Total is the sum of all index * mod inverse * Product of all other bus IDs
        // Yay Chinese remainder theorem
        total += i*mi*NI;
    }

    println!("{}", total % N);

}

fn mod_inverse(x: usize, y: usize) -> usize {
    let a = x % y;
    for x in 1..y {
        if (a*x) % y == 1 {
            return x;
        }
    }

    return 1;
}