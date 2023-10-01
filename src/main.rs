use std::fmt;
use std::fmt::Formatter;

use crate::lexer::{Lexer, Token};

mod lexer;

impl fmt::Display for Token {
    fn fmt(&self, f: &mut Formatter<'_>) -> fmt::Result {
        match self {
            Token::Abstract => write!(f, "<Keyword|abstract>"),
            Token::Comma => write!(f, "<Comma>"),
            Token::EOF => write!(f, "<EOF>"),
            Token::Error(msg) => write!(f, "<ERROR: {}", msg),
            Token::Identifier(id) => write!(f, "<Identifier, {}>", id),
        }
    }
}


fn main() {
    let mut lexer = Lexer::new("Hello, world");

    loop {
        let t = lexer.next();
        println!("{}", t);
        if t == Token::EOF {
            break;
        }
    }
}