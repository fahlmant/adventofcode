use std::fmt::DebugSet;
use std::fs;
use std::ops::Range;

#[derive(Debug)]
pub struct AlmanacMap {
    source_range_start: i64,
    dest_range_start: i64,
    range_count: i64
}

fn main() {

    // Use fs to get all input to a Result<String>, which is then unwrapped
    let contents = fs::read_to_string("../input").unwrap();
    let mut min_location = 0;

    let mut parts = contents.as_str().split("\n\n");

    let seed_ranges: Vec<(i64, i64)> = parts.next().unwrap().split_whitespace().skip(1).map(|x| x.parse().unwrap()).collect::<Vec<_>>().chunks(2).map(|s| (s[0], s[1])).collect();

    let mut maps: Vec<Vec<AlmanacMap>> = parts.map(|pt| generate_map(pt)).collect();
    // Reverse the order of transforms. That allows starting from a location, transforming to seed, and checking if the seed is valid
    maps.reverse();

    let mut valid_seed = false;
    let mut loc: i64 = 0;
    while !valid_seed {
        let mut result = loc;
        for m in &maps {
            result = apply_transformation(m, result);
        }
        for r in seed_ranges.clone() {
            if (r.0..r.0+r.1).contains(&result) {
                valid_seed = true;
                min_location = loc;
                break;
            }
        }
        loc += 1;
    }

    println!("{}", min_location);
}

fn apply_transformation(m: &Vec<AlmanacMap>, seed: i64) -> i64{
    let mut result: i64 = 0;
    for r in m {
        if (r.dest_range_start..r.dest_range_start+r.range_count).contains(&seed) {
            let offset = seed - r.dest_range_start;
            result = r.source_range_start + offset;
            return result;
        } else {
            result = seed;        }
    }
    return result;
}


fn generate_map(input: &str) -> Vec<AlmanacMap>  {
    let mut ranges:Vec<AlmanacMap> = Vec::new();
    for line in input.lines().skip(1) {
        let transform: Vec<i64> = line.split_whitespace().map(|x| x.parse().unwrap()).collect();
        let range = AlmanacMap {
            dest_range_start: transform[0],
            source_range_start: transform[1],
            range_count: transform[2],
        };
        ranges.push(range);
    }
    ranges
}
