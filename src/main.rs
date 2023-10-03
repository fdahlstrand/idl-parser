use std::fmt;
use std::fmt::Formatter;

use crate::lexer::{Lexer, Token, TokenType};

mod lexer;

impl fmt::Display for Token {
    fn fmt(&self, f: &mut Formatter<'_>) -> fmt::Result {
        match &self.token_type {
            TokenType::Abstract => write!(f, "<Keyword|abstract>"),
            TokenType::Comma => write!(f, "<Comma>"),
            TokenType::EOF => write!(f, "<EOF>"),
            TokenType::Invalid(msg) => write!(f, "<ERROR: {}", msg),
            TokenType::Identifier(id) => write!(f, "<Identifier, {}>", id),
            TokenType::Integer(value) => write!(f, "<Integer | {}>", value),
        }
    }
}


fn main() {
    let mut lexer = Lexer::new("Hello, world");

    loop {
        let t = lexer.next();
        println!("{}", t);
        if t.token_type == TokenType::EOF {
            break;
        }
    }
}