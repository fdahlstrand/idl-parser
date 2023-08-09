#![allow(dead_code)]
use std::iter::Peekable;
use std::str::Chars;

#[derive(PartialEq, Debug)]
enum TokenType {
    EOF,
    SEMICOLON,
    LBRACE,
    RBRACE,
    ILLEGAL,
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
                '{' => Token::new(TokenType::LBRACE, "{"),
                '}' => Token::new(TokenType::RBRACE, "}"),
                _ => Token::new(TokenType::ILLEGAL, ""),
            },
            None => Token::new(TokenType::EOF, ""),
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
        let expected_tokens = [
            Token::new(TokenType::SEMICOLON, ";"),
            Token::new(TokenType::LBRACE, "{"),
            Token::new(TokenType::RBRACE, "}"),
            Token::new(TokenType::EOF, ""),
        ];
        for expected in expected_tokens {
            assert_eq!(lexer.next_token(), expected);
        }
    }

    #[test]
    fn test_bad_token() {
        let mut lexer = Lexer::new("ö".chars().peekable());
        assert_eq!(lexer.next_token(), Token::new(TokenType::ILLEGAL, ""));
    }
}
