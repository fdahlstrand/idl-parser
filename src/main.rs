#![allow(dead_code)]
use std::iter::Peekable;
use std::str::Chars;

#[derive(PartialEq, Debug)]
enum TokenType {
    SEMICOLON,
    LCURLY,
    RCURLY,
}

#[derive(PartialEq, Debug)]
struct Token {
    typ: TokenType,
    literal: String,
}

struct Lexer<'a> {
    input: Peekable<Chars<'a>>,
}

impl Token {
    fn new(t: TokenType, literal: &str) -> Self {
        return Token {
            typ: t,
            literal: literal.to_string(),
        };
    }
}

impl<'a> Lexer<'a> {
    fn new(input: Peekable<Chars<'a>>) -> Self {
        Lexer { input }
    }

    fn next_token(&mut self) -> Token {
        let ch = self.input.next();
        match ch {
            Some(c) => match c {
                ';' => Token::new(TokenType::SEMICOLON, ";"),
                '{' => Token::new(TokenType::LCURLY, "{"),
                '}' => Token::new(TokenType::RCURLY, "}"),
                _ => Token::new(TokenType::SEMICOLON, ";"),
            },
            None => Token::new(TokenType::SEMICOLON, ";"),
        }
    }
}

fn main() {}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_next_token() {
        let mut lexer = Lexer::new(";{}".chars().peekable());
        assert_eq!(lexer.next_token(), Token::new(TokenType::SEMICOLON, ";"));
        assert_eq!(lexer.next_token(), Token::new(TokenType::LCURLY, "{"));
        assert_eq!(lexer.next_token(), Token::new(TokenType::RCURLY, "}"));
    }
}
