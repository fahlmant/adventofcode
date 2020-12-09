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
        let value = ins[1][1..].to_string().parse::<usize>().unwrap();
        match ins[0] {
            "acc" => {if sign == "-"{self.accumulator -= value} else{self.accumulator+= value}; self.ic = self.ic + 1;}
            "jmp" => if sign == "-"{self.ic -= value} else{self.ic+= value},
            "nop" => self.ic = self.ic + 1,
            _ => return,
        }
    }

    fn run_program(&mut self) -> Result<usize, String> {
        let mut ic_tracker: HashSet<usize> = HashSet::new();
        loop {
            if self.ic >= self.instructions.len() {
                return Ok(self.accumulator);
            }
            if ic_tracker.contains(&self.ic) {
                return Err("Infinite Loop".to_string());
            }
            
            ic_tracker.insert(self.ic);
            &self.run_instruction();
        }
    }
}

fn main() {

    let contents = fs::read_to_string("../input").unwrap();

    let lines: Vec<String> = contents.trim().lines().map(|l| l.to_string()).collect();

    let game_console: GameConsole = GameConsole{
        instructions: lines,
        accumulator: 0,
        ic: 0,
    };

    for i in 0..game_console.instructions.len() {
        let mut game = GameConsole {
            instructions: game_console.instructions.clone(),
            accumulator: 0,
            ic: 0,
        };

        if game.instructions[i].contains("jmp") {
            game.instructions[i] = game.instructions[i].replace("jmp", "nop");
        } else if game.instructions[i].contains("nop") {
            game.instructions[i] = game.instructions[i].replace("nop", "jmp");
        }

        let result = game.run_program();
        match result {
            Ok(answer) => {println!("{}", answer); break;},
            Err(_) => continue,
        }
    }
}