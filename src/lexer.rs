pub(crate) struct Lexer {
    input: Vec<char>,
    pos: usize,
    ch: char,
    literal: String,
    start: Position,
    current: Position,
}

#[derive(PartialEq, Debug)]
pub(crate) enum TokenType {
    Abstract,
    Comma,
    EOF,
    Invalid(String),
    Identifier(String),
    Integer(i64),
}

#[derive(Debug, PartialEq)]
pub(crate) struct Position {
    pub(crate) line: usize,
    pub(crate) column: usize,
}

impl Clone for Position {
    fn clone(&self) -> Self {
        Position {
            line: self.line,
            column: self.column,
        }
    }
}

pub(crate) struct Token {
    pub(crate) token_type: TokenType,
    pub(crate) literal: String,
    pub(crate) start: Position,
    pub(crate) end: Position,
}

const KEYWORDS: [(&str, TokenType); 1] = [("abstract", TokenType::Abstract)];


impl Lexer {
    pub fn new(src: &str) -> Self {
        let mut lexer = Lexer {
            input: src.chars().collect(),
            pos: 0,
            ch: '\0',
            literal: "".to_string(),
            start: Position {
                line: 1,
                column: 1,
            },
            current: Position {
                line: 1,
                column: 1,
            },
        };

        if lexer.input.len() > 0 {
            lexer.ch = lexer.input[0];
        }

        lexer
    }

    pub fn next(&mut self) -> Token {
        self.skip_whitespace();
        match self.ch {
            '\0' => {
                self.mark_start();
                self.advance();
                self.emit(TokenType::EOF)
            }
            ',' => {
                self.mark_start();
                self.consume();
                self.emit(TokenType::Comma)
            }
            _ if self.ch.is_ascii_digit() => self.number(),
            _ if self.ch.is_ascii_alphabetic() || self.ch == '_' => self.identifier(),
            ch => {
                self.mark_start();
                self.emit(TokenType::Invalid(format!("Unexpected character '{}'", ch)))
            }
        }
    }

    fn emit(&mut self, t: TokenType) -> Token {
        let token = Token {
            token_type: t,
            literal: self.literal.clone(),
            start: self.start.clone(),
            end: Position {
                line: self.current.line,
                column: self.current.column - 1,
            },
        };

        self.literal = "".to_string();

        token
    }

    fn mark_start(&mut self) {
        self.start = self.current.clone();
    }

    fn advance(&mut self) {
        if self.ch == '\n' {
            self.current.line += 1;
            self.current.column = 1;
        } else {
            self.current.column += 1;
        }
        self.pos += 1;
        if self.pos >= self.input.len() {
            self.ch = '\0';
        } else {
            self.ch = self.input[self.pos];
        }
    }

    fn consume(&mut self) {
        self.literal.push(self.ch);
        self.advance();
    }

    fn peek(&mut self) -> Option<char> {
        return if self.pos + 1 >= self.input.len() {
            None
        } else {
            Some(self.input[self.pos + 1])
        };
    }

    fn identifier(&mut self) -> Token {
        let match_keyword = self.ch != '_';
        let ident = self.read_identifier();

        if ident.is_empty() {
            return self.emit(TokenType::Invalid("Invalid Identifier".to_string()));
        }

        if match_keyword {
            match self.lookup_keyword(&ident) {
                None => self.emit(TokenType::Identifier(ident)),
                Some(t) => self.emit(t)
            }
        } else {
            self.emit(TokenType::Identifier(ident))
        }
    }

    fn lookup_keyword(&mut self, ident: &str) -> Option<TokenType> {
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

        self.mark_start();

        if self.ch == '_' {
            self.consume();
        }

        if self.ch.is_ascii_alphabetic() {
            loop {
                ident.push(self.ch);
                self.consume();
                if !(self.ch.is_ascii_alphanumeric() || self.ch == '_') {
                    break;
                }
            }
        }

        ident
    }

    fn number(&mut self) -> Token {
        let mut value: i64 = 0;
        if self.ch == '0' && (self.peek() == Some('x') || self.peek() == Some('X')) {
            self.mark_start();
            self.consume();
            self.consume();

            if !self.ch.is_ascii_hexdigit() {
                return self.emit(TokenType::Invalid("Invalid hexadecimal literal".to_string()));
            }

            loop {
                value = value * 16 + i64::from(self.ch.to_digit(16).unwrap());
                self.consume();

                if !self.ch.is_ascii_hexdigit() {
                    break;
                }
            }
        } else if self.ch == '0' {
            self.mark_start();
            loop {
                value = value * 8 + i64::from(self.ch.to_digit(10).unwrap());
                self.consume();
                if !('0' <= self.ch && self.ch <= '7') {
                    break;
                }
            }
        } else {
            self.mark_start();
            loop {
                value = value * 10 + i64::from(self.ch.to_digit(10).unwrap());
                self.consume();
                if !self.ch.is_ascii_digit() {
                    break;
                }
            }
        }

        self.emit(TokenType::Integer(value))
    }

    fn skip_whitespace(&mut self) {
        while self.ch.is_ascii_whitespace() {
            self.advance();
        }
    }
}

#[cfg(test)]
mod tests {
    use rstest::rstest;

    use super::*;

    #[rstest]
    #[case::empty_input("", Position {line: 1, column: 1}, Position {line: 1, column: 1})]
    #[case::single_character("a", Position {line: 1, column: 1}, Position {line: 1, column: 1})]
    #[case::multiple_characters("identifier", Position {line: 1, column: 1}, Position {line: 1, column: 10})]
    #[case::multiple_lines("\n\nidentifier", Position {line: 3, column: 1}, Position {line: 3, column: 10})]
    #[case::whitespace_only("  \t", Position {line: 1, column: 4}, Position {line: 1, column: 4})]
    fn position(#[case] input: &str, #[case] start: Position, #[case] end: Position) {
        let mut lexer = Lexer::new(input);

        let token = lexer.next();

        assert_eq!(token.start, start, "unexpected start of token");
        assert_eq!(token.end, end, "unexpected end of token");
    }

    #[rstest]
    #[case::empty_input("")]
    #[case::explicit_nul("\0")]
    #[case::whitespace("\n\r\t ")]
    fn eof(#[case] input: &str) {
        let mut lexer = Lexer::new(input);

        let token = lexer.next();

        assert_eq!(TokenType::EOF, token.token_type);
    }

    #[rstest]
    #[case::basic_identifier("identifier")]
    #[case::short_identifer("i")]
    #[case::mixed_case("IdEnTiFiEr")]
    #[case::with_numbers("Id9nT2FiEr")]
    #[case::with_underscore("Id9n_T2FiEr_")]
    fn good_identifier(#[case] input: &str) {
        let mut lexer = Lexer::new(input);

        let token = lexer.identifier();

        assert_eq!(TokenType::Identifier(input.to_string()), token.token_type);
    }

    #[rstest]
    #[case::starting_with_number("9identifier")]
    fn bad_identifier(#[case] input: &str) {
        let mut lexer = Lexer::new(input);

        let token = lexer.identifier();

        assert_ne!(TokenType::Identifier(input.to_string()), token.token_type);
    }

    #[rstest]
    #[case::basic_identifer("_identifier")]
    #[case::escaped_keyword("_abstract")]
    fn good_escaped_identifier(#[case] input: &str) {
        let mut lexer = Lexer::new(input);

        let token = lexer.identifier();

        assert_eq!(TokenType::Identifier(input[1..].to_string()), token.token_type);
    }

    #[rstest]
    #[case::double_underscore("__identifier")]
    #[case::number_after_escape("_9identifier")]
    #[case::only_escape("_")]
    fn bad_escaped_identifier(#[case] input: &str) {
        let mut lexer = Lexer::new(input);

        let token = lexer.identifier();

        assert_ne!(TokenType::Identifier(input[1..].to_string()), token.token_type);
    }

    #[rstest]
    fn keyword() {
        let mut lexer = Lexer::new("abstract");

        let token = lexer.identifier();

        assert_eq!(TokenType::Abstract, token.token_type);
    }

    #[rstest]
    #[case::comma(",", TokenType::Comma)]
    fn punctuation(#[case] input: &str, #[case] expected: TokenType) {
        let mut lexer = Lexer::new(input);

        let token = lexer.next();

        assert_eq!(expected, token.token_type);
    }

    #[rstest]
    #[case::unexpected_token("ä")]
    fn lexer_error(#[case] input: &str) {
        let mut lexer = Lexer::new(input);

        let token = lexer.next();

        match token.token_type {
            TokenType::Invalid(_) => {}
            _ => assert!(false, "Lexer did not return error, got {}", token)
        }
    }

    #[rstest]
    #[case::simple_integer("1", 1)]
    #[case::long_integer("1133388990", 1133388990)]
    #[case::simple_octal("01", 1)]
    #[case::octal_double_digit("011", 9)]
    #[case::simple_hexadecimal("0xA", 10)]
    #[case::long_hexadecimal("0Xffff", 65535)]
    fn integer(#[case] input: &str, #[case] value: i64) {
        let mut lexer = Lexer::new(input);

        let token = lexer.next();

        assert_eq!(TokenType::Integer(value), token.token_type);
    }

    #[rstest]
    #[case::truncated_hexadecimal("0x")]
    #[case::truncated_hexadecimal("0X")]
    fn bad_integer(#[case] input: &str) {
        let mut lexer = Lexer::new(input);

        let token = lexer.next();

        match token.token_type {
            TokenType::Invalid(_) => {}
            _ => assert!(false, "Lexer did not return error, got {}", token)
        }
    }
}