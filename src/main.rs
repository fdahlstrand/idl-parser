use crate::lexer::{Lexer, Token};

mod lexer;


fn main() {
    let mut lexer = Lexer::new("Hello, world");

    while let Some(t) = lexer.next() {
        println!("{}", t);
        if t == Token::EOF {
            break;
        }
    }
}