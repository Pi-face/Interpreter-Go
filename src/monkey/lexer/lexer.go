package lexer

type Lexer struct {
      input     string
      position  int // current positions in input or pointer(points to current char)
      readposition int // current reading position in input (after current char)
      ch          byte //current character under observation.
      
}

func New(input string) *Lexer {
  l := &Lexer{input:input}
  return l
}
