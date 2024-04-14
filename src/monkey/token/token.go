package token 

type TokenType string

type Token struct {
    Type TokenType 
    Literal string
}

const (
  ILLEGAL = "ILLEGAL"
  EOF = "EOF"


  // identifiers + literals

  IDENT = "IDENT" // add, foobar, xy etc.
  INT = "INT" //Applications

  // operators 
  ASSIGN = "="
  PLUS = "+"

  //delimiter
  COMMA = ","
  SEMICOLON = ";"

  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  RBRACE = "}"

  // keywords 
  FUNCTION = "FUNCTION"
  LET = "LET"

)











