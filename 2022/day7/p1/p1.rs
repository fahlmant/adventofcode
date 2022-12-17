use std::iter::Peekable;

fn main() {
    
    let contents = include_bytes!("../input");

    let lines = contents.split(|b| b == &b'\n');

    let mut sum = 0;

    sum_dir(&mut lines.peekable(), &mut sum);

    println!("{}", sum)
}

fn sum_dir(lines: &mut Peekable<impl Iterator<Item = &'static [u8]>>, sum: &mut u64) -> u64 {
    let mut dir_size: u64 = 0;
    
    while let Some(x) = lines.next() {
        match x {
            b"$ cd .." => break,
            _ if &x[0..4] == b"$ ls" => {
                loop {
                    let line = lines.peek();
                    match line {
                        None => break,
                        Some(line) => {
                            if line[0] == b'$' {
                                break;
                            }
                            if line[0] != b'd' {
                                dir_size += atoi::atoi::<u64>(line.split(|b| b == &b' ').next().unwrap()).unwrap();
                            }
                            lines.next();
                        }
                    }
 
                }    
            }
            _ => dir_size += sum_dir(lines, sum)
        }
    }

    if dir_size < 100_000 {
        *sum += dir_size
    }

    return dir_size
}

