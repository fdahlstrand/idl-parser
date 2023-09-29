use std::fmt;
use std::fmt::Formatter;

use crate::lexer::Token::Identifier;

pub(crate) struct Lexer {
    input: Vec<char>,
    pos: usize,
    ch: char,
}

#[derive(PartialEq)]
pub(crate) enum Token {
    Comma,
    Identifier(String),
    EOF,
}

impl fmt::Display for Token {
    fn fmt(&self, f: &mut Formatter<'_>) -> fmt::Result {
        match self {
            Token::Comma => write!(f, "<Comma>"),
            Identifier(id) => write!(f, "<Identifier, {}>", id),
            Token::EOF => write!(f, "<EOF>"),
        }
    }
}

impl Lexer {
    pub fn new(src: &str) -> Self {
        let mut lexer = Lexer {
            input: src.chars().collect(),
            pos: 0,
            ch: '\0',
        };
        lexer.ch = lexer.input[0];

        lexer
    }

    pub fn next(&mut self) -> Option<Token> {
        self.skip_whitespace();
        match self.ch {
            '\0' => Some(Token::EOF),
            ',' => {
                self.advance();
                Some(Token::Comma)
            }
            _ if self.ch.is_ascii_alphanumeric() => self.identifier(),
            _ => None
        }
    }

    fn advance(&mut self) {
        self.pos += 1;
        if self.pos >= self.input.len() {
            self.ch = '\0';
        } else {
            self.ch = self.input[self.pos];
        }
    }

    fn identifier(&mut self) -> Option<Token> {
        let mut ident = "".to_string();
        loop {
            ident.push(self.ch);
            self.advance();
            if !self.ch.is_ascii_alphanumeric() {
                break;
            }
        }

        Some(Identifier(ident))
    }

    fn skip_whitespace(&mut self) {
        while self.ch.is_ascii_whitespace() {
            self.advance();
        }
    }
}