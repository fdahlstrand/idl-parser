use std::fmt;
use std::fmt::Formatter;

use crate::lexer::{Lexer, Position, Token, TokenType};

mod lexer;

impl fmt::Display for Position {
    fn fmt(&self, f: &mut Formatter<'_>) -> fmt::Result {
        write!(f, "@{}:{}", self.line, self.column)
    }
}

impl fmt::Display for Token {
    fn fmt(&self, f: &mut Formatter<'_>) -> fmt::Result {
        let name = match &self.token_type {
            TokenType::Abstract => "abstract".to_string(),
            TokenType::Comma => "comma".to_string(),
            TokenType::EOF => "EOF".to_string(),
            TokenType::Invalid(msg) => format!("INVALID {}", msg),
            TokenType::Identifier(id) => format!("identifier({})>", id),
            TokenType::Integer(value) => format!("integer({})", value),
        };

        write!(f, "<{} {}-{} = \"{}\">", name, self.start, self.end, self.literal)
    }
}


fn main() {
    let mut lexer = Lexer::new("Hello, world\nAgain is abstract");

    loop {
        let t = lexer.next();
        println!("{}", t);
        if t.token_type == TokenType::EOF {
            break;
        }
    }
}