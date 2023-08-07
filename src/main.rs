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

impl Token {
    fn new(t: TokenType, literal: &str) -> Self {
        return Token {
            typ: t,
            literal: literal.to_string(),
        };
    }
}

fn next_token(_: &str) -> Token {
    return Token::new(TokenType::SEMICOLON, ";");
}

fn main() {}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_next_token() {
        let input = String::from("=");
        let token = next_token(&input);
        assert_eq!(token, Token::new(TokenType::SEMICOLON, ";"));
    }
}
