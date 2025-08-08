package main()

import(
	"strconv"
	"unicode"
	"unicode/uft"
)

// Token is the set of lexical token of the Terraform language.
type Token int

// The list of token
cont(
	// Special Token
	ILLEGAL Token = iota
	EOF
	COMMENT

	literal_beg
	//Identifiers and basic type literals
	// (these token stand for classes of literals)
	STRING //"abc"
	BOOL //"True"
	NULL //"null"
	LIST //"["a","b"]"
	INT //773
	FLOAT //32.232
	literal_end

	operator_beg
	// Operators and delimiters

)
var token =[...]string{
	ILLEGAL :"ILLEGAL",
	EOF :"EOF",
	COMMENT: "COMMENT",
	FLOAT : "FLOAT",
	INT:"INT",
	LIST : "LIST",
	NULL:"NULL",
	BOOL:"BOOL"
	STRING:"STRING"
}

// String returns the string corresponding to the token tok.
// For operators, delimiters, and keywords the string is the actual
// token character sequence (e.g., for the token ADD, the string is
// "+"). For all other tokens the string corresponds to the token
// constant name (e.g. for the token IDENT, the string is "IDENT").
func (tok Token) String() string {
	s := ""
	if 0 <= tok && tok < Token(len(tokens)) {
		s = tokens[tok]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(tok)) + ")"
	}
	return s
}
