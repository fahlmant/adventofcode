use std::iter::Peekable;

struct Directory(
    Vec<Directory>, 
    u64
);

fn main() {
    
    let contents = include_bytes!("../input");

    let mut lines = contents.split(|b| b == &b'\n').peekable();

    let root_dir = sum_dir(&mut lines);
    sum_dir(&mut lines.peekable());

    let size = find_dir(&root_dir, root_dir.1 - 40_000_000);

    println!("{}", size.unwrap());
}

fn sum_dir(lines: &mut Peekable<impl Iterator<Item = &'static [u8]>>) -> Directory {
    let mut dir_size: u64 = 0;
    let mut subdirs = vec![];
    
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
            _ => subdirs.push(sum_dir(lines))
        }
    }
    dir_size += subdirs.iter().map(|d| d.1).sum::<u64>();
    return Directory(subdirs, dir_size)
}

fn find_dir(root: &Directory, min_size: u64) -> Option<u64> {
    root.0.iter().filter(|d| d.1 >= min_size).flat_map(|d| [Some(d.1), find_dir(d, min_size)]).flatten().min()
}
