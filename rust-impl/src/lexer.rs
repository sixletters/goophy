use std::string;



pub struct Lexer{
    input: String,
    pos: u64, // Current position in input (points to current char)
    read_pos: u64,  // Current read pos in input(after current char)
    ch: u8,  // Current read pos in input(after current char)
}

impl Lexer {
    pub fn new(input: String) -> Self {
        let lexer = Self{
            input,
            pos: 0,
            read_pos: 0,
            ch: 0,
        };
        lexer.read_char();
        lexer
    }

    pub fn read_char(&self) {
        if self.read_pos >= self.input.len() {
            
        }
    }
}