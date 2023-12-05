use std::fmt::DebugSet;
use std::fs;
use std::ops::Range;

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

    let mut seeds: Vec<i64> = parts.next().unwrap().split_whitespace().skip(1).map(|x| x.parse().unwrap()).collect();

    let maps: Vec<Vec<AlmanacMap>> = parts.map(|pt| generate_map(pt)).collect();

    for (i,_) in seeds.clone().iter().enumerate() {
        for m in &maps {
            seeds[i] = apply_transformation(m, seeds[i]);
        }
    }

    let min_value = seeds.iter().min();
    match min_value {
        Some(min) => min_location = *min,
        None      => println!( "Vector is empty" ),
    }
    

    println!("{}", min_location);
}

fn apply_transformation(m: &Vec<AlmanacMap>, seed: i64) -> i64{
    let mut result: i64 = 0;
    for r in m {
        if (r.source_range_start..r.source_range_start+r.range_count).contains(&seed) {
            let offset = seed - r.source_range_start;
            result = r.dest_range_start + offset;
            return result;
        } else {
            result = seed;
        }
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
