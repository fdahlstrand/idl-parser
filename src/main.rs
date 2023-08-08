#![allow(dead_code)]

#[derive(PartialEq, Debug)]
enum TokenType {
    SEMICOLON,
}

#[derive(PartialEq, Debug)]
struct Token {
    typ: TokenType,
    literal: String,
}

struct Lexer {
    input: String,
}

impl Token {
    fn new(t: TokenType, literal: &str) -> Self {
        return Token {
            typ: t,
            literal: literal.to_string(),
        };
    }
}

impl Lexer {
    fn new(input: &str) -> Self {
        return Lexer {
            input: input.to_string(),
        };
    }

    fn next_token(&self) -> Token {
        return Token::new(TokenType::SEMICOLON, ";");
    }
}

fn main() {}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_next_token() {
        let lexer = Lexer::new(";");
        let token = lexer.next_token();
        assert_eq!(token, Token::new(TokenType::SEMICOLON, ";"));
    }
}
