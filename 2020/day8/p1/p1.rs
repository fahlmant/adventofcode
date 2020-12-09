use std::fs;
use std::collections::HashSet;

struct GameConsole {
    instructions: Vec<String>,
    accumulator: usize,
    ic: usize,
}

impl GameConsole {
    fn run_instruction(&mut self) {
        
        let ins: Vec<&str> = self.instructions[self.ic].split(" ").collect();

        let sign = ins[1].chars().nth(0).unwrap().to_string();
        let mut value = ins[1][1..].to_string().parse::<usize>().unwrap();
        match ins[0] {
            "acc" => {if sign == "-"{self.accumulator -= value} else{self.accumulator+= value}; self.ic = self.ic + 1;}
            "jmp" => if sign == "-"{self.ic -= value} else{self.ic+= value},
            "nop" => self.ic = self.ic + 1,
            _ => return,
        }
    }
}

fn main() {

    let contents = fs::read_to_string("../input").unwrap();

    let lines: Vec<String> = contents.trim().lines().map(|l| l.to_string()).collect();

    let mut game_console: GameConsole = GameConsole{
        instructions: lines,
        accumulator: 0,
        ic: 0,
    };

    let mut ic_tracker: HashSet<usize> = HashSet::new();

    loop {
        if ic_tracker.contains(&game_console.ic) {
            break;
        }
        ic_tracker.insert(game_console.ic);
        game_console.run_instruction();
    }
    
    println!("{}", game_console.accumulator);
}