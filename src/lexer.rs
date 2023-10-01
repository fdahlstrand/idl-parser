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
    Error(String),
    Identifier(String),
}

const KEYWORDS: [(&str, Token); 1] = [("abstract", Token::Abstract)];


impl Lexer {
    pub fn new(src: &str) -> Self {
        let mut lexer = Lexer {
            input: src.chars().collect(),
            pos: 0,
            ch: '\0',
        };

        if lexer.input.len() > 0 {
            lexer.ch = lexer.input[0];
        }

        lexer
    }

    pub fn next(&mut self) -> Token {
        self.skip_whitespace();
        match self.ch {
            '\0' => Token::EOF,
            ',' => {
                self.consume();
                Token::Comma
            }
            _ if self.ch.is_ascii_alphabetic() || self.ch == '_' => self.identifier(),
            ch => Token::Error(format!("Unexpected character '{}'", ch)),
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

    fn identifier(&mut self) -> Token {
        let match_keyword = self.ch != '_';
        let ident = self.read_identifier();

        if ident.is_empty() {
            return Token::Error("Invalid Identifier".to_string());
        }

        if match_keyword {
            match self.lookup_keyword(&ident) {
                None => Token::Identifier(ident),
                Some(t) => t
            }
        } else {
            Token::Identifier(ident)
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

    fn skip_whitespace(&mut self) {
        while self.ch.is_ascii_whitespace() {
            self.consume();
        }
    }
}

#[cfg(test)]
mod tests {
    use rstest::rstest;

    use super::*;

    #[rstest]
    #[case::empty_input("")]
    #[case::explicit_nul("\0")]
    #[case::whitespace("\n\r\t ")]
    fn eof(#[case] input: &str) {
        let mut lexer = Lexer::new(input);

        let token = lexer.next();

        assert_eq!(Token::EOF, token);
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

        assert_eq!(Token::Identifier(input.to_string()), token);
    }

    #[rstest]
    #[case::starting_with_number("9identifier")]
    fn bad_identifier(#[case] input: &str) {
        let mut lexer = Lexer::new(input);

        let token = lexer.identifier();

        assert_ne!(Token::Identifier(input.to_string()), token);
    }

    #[rstest]
    #[case::basic_identifer("_identifier")]
    #[case::escaped_keyword("_abstract")]
    fn good_escaped_identifier(#[case] input: &str) {
        let mut lexer = Lexer::new(input);

        let token = lexer.identifier();

        assert_eq!(Token::Identifier(input[1..].to_string()), token);
    }

    #[rstest]
    #[case::double_underscore("__identifier")]
    #[case::number_after_escape("_9identifier")]
    #[case::only_escape("_")]
    fn bad_escaped_identifier(#[case] input: &str) {
        let mut lexer = Lexer::new(input);

        let token = lexer.identifier();

        assert_ne!(Token::Identifier(input[1..].to_string()), token);
    }

    #[rstest]
    fn keyword() {
        let mut lexer = Lexer::new("abstract");

        let token = lexer.identifier();

        assert_eq!(Token::Abstract, token);
    }

    #[rstest]
    #[case::comma(",", Token::Comma)]
    fn punctuation(#[case] input: &str, #[case] expected: Token) {
        let mut lexer = Lexer::new(input);

        let token = lexer.next();

        assert_eq!(expected, token);
    }

    #[rstest]
    #[case::unexpected_token("ä")]
    fn lexer_error(#[case] input: &str) {
        let mut lexer = Lexer::new(input);

        let token = lexer.next();

        match token {
            Token::Error(_) => {}
            _ => assert!(false, "Lexer did not return error, got {}", token)
        }
    }
}