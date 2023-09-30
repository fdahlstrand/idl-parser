use std::fmt;
use std::fmt::Formatter;

pub(crate) struct Lexer {
    input: Vec<char>,
    pos: usize,
    ch: char,
}

#[derive(PartialEq, Debug)]
pub(crate) enum Token {
    Abstract,
    Comma,
    EOF,
    Identifier(String),
}

const KEYWORDS: [(&str, Token); 1] = [("abstract", Token::Abstract)];

impl fmt::Display for Token {
    fn fmt(&self, f: &mut Formatter<'_>) -> fmt::Result {
        match self {
            Token::Abstract => write!(f, "<Keyword|abstract>"),
            Token::Comma => write!(f, "<Comma>"),
            Token::EOF => write!(f, "<EOF>"),
            Token::Identifier(id) => write!(f, "<Identifier, {}>", id),
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
                self.consume();
                Some(Token::Comma)
            }
            _ if self.ch.is_ascii_alphanumeric() || self.ch == '_' => self.identifier(),
            _ => None
        }
    }

    fn consume(&mut self) {
        self.pos += 1;
        if self.pos >= self.input.len() {
            self.ch = '\0';
        } else {
            self.ch = self.input[self.pos];
        }
    }

    fn identifier(&mut self) -> Option<Token> {
        let match_keyword = self.ch != '_';
        let ident = self.read_identifier();

        if match_keyword {
            match self.lookup_keyword(&ident) {
                None => Some(Token::Identifier(ident)),
                t => t
            }
        } else {
            Some(Token::Identifier(ident))
        }
    }

    fn lookup_keyword(&mut self, ident: &str) -> Option<Token> {
        let lowercase_ident = ident.to_lowercase();
        for (id, token) in KEYWORDS {
            if id == lowercase_ident {
                return Some(token);
            }
        }

        None
    }


    fn read_identifier(&mut self) -> String {
        let mut ident = "".to_string();
        if self.ch == '_' {
            self.consume();
        }
        loop {
            ident.push(self.ch);
            self.consume();
            if !self.ch.is_ascii_alphanumeric() {
                break;
            }
        }

        ident
    }

    fn skip_whitespace(&mut self) {
        while self.ch.is_ascii_whitespace() {
            self.consume();
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn identifier() {
        let mut lexer = Lexer::new("identifier");

        let token = lexer.next();

        assert_eq!(Some(Token::Identifier("identifier".to_string())), token);
    }

    #[test]
    fn escaped_identifier() {
        let mut lexer = Lexer::new("_identifier");

        let token = lexer.next();

        assert_eq!(Some(Token::Identifier("identifier".to_string())), token);
    }

    #[test]
    fn keyword() {
        let mut lexer = Lexer::new("abstract");

        let token = lexer.next();

        assert_eq!(Some(Token::Abstract), token);
    }

    #[test]
    fn escaped_keyword() {
        let mut lexer = Lexer::new("_abstract");

        let token = lexer.next();

        assert_eq!(Some(Token::Identifier("abstract".to_string())), token);
    }
}