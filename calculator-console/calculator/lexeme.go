package calculator

import "math"

type LexemeType string
type LexemeClassification string

// types of lexeme - operator (e.g. + - * / etc) operand (number) or internal function (exp, log, ln, etc.)
const (
	LEXEME_CLASSIFICATION_OPERATOR          LexemeClassification = "OPERATOR"
	LEXEME_CLASSIFICATION_OPERAND           LexemeClassification = "OPERAND"
	LEXEME_CLASSIFICATION_INTERNAL_FUNCTION LexemeClassification = "INTERNAL_FUNCTION"
	LEXEME_CLASSIFICATION_INVALID           LexemeClassification = "INVALID"
)

// valid operator lexemes
const (
	LEXEME_INVALID     LexemeType = "INVALID"
	LEXEME_OPEN_PAREN  LexemeType = "("
	LEXEME_CLOSE_PAREN LexemeType = ")"
	LEXEME_PLUS        LexemeType = "+"
	LEXEME_MINUS       LexemeType = "-"
	LEXEME_MULTIPLY    LexemeType = "*"
	LEXEME_DIVIDE      LexemeType = "/"
	LEXEME_EQUAL       LexemeType = "="
	LEXEME_SPACE       LexemeType = " "
)

// valid internal function lexemes
const (
	LEXEME_FUNCTION_EXPONENT    LexemeType = "exp"
	LEXEME_FUNCTION_LOG         LexemeType = "log"
	LEXEME_FUNCTION_NATURAL_LOG LexemeType = "ln"
)

type Lexeme struct {
	Runes          []rune
	Type           LexemeType
	Classification LexemeClassification
}

func IsLexeme(runeSeq []rune) (isLexeme bool, lexemeType LexemeType, classification LexemeClassification) {

	if len(runeSeq) == 0 { // empty sequence is not a lexeme
		return false, LEXEME_INVALID, LEXEME_CLASSIFICATION_INVALID
	} else if len(runeSeq) == 1 { // single character - might be an operator

		switch LexemeType(runeSeq[0]) {
		case LEXEME_OPEN_PAREN:
			return true, LEXEME_OPEN_PAREN, LEXEME_CLASSIFICATION_OPERATOR
		case LEXEME_CLOSE_PAREN:
			return true, LEXEME_CLOSE_PAREN, LEXEME_CLASSIFICATION_OPERATOR
		case LEXEME_PLUS:
			return true, LEXEME_PLUS, LEXEME_CLASSIFICATION_OPERATOR
		case LEXEME_MINUS:
			return true, LEXEME_MINUS, LEXEME_CLASSIFICATION_OPERATOR
		case LEXEME_MULTIPLY:
			return true, LEXEME_MULTIPLY, LEXEME_CLASSIFICATION_OPERATOR
		case LEXEME_DIVIDE:
			return true, LEXEME_DIVIDE, LEXEME_CLASSIFICATION_OPERATOR
		case LEXEME_EQUAL:
			return true, LEXEME_EQUAL, LEXEME_CLASSIFICATION_OPERATOR
		case LEXEME_SPACE:
			return true, LEXEME_SPACE, LEXEME_CLASSIFICATION_OPERATOR
		default:
			return false, LEXEME_INVALID, LEXEME_CLASSIFICATION_INVALID
		}

	}

	// runeSequence contains 2 or more runes - check for internal functions
	switch LexemeType(runeSeq) {
	case LEXEME_FUNCTION_EXPONENT:
		return true, LEXEME_FUNCTION_EXPONENT, LEXEME_CLASSIFICATION_OPERATOR
	case LEXEME_FUNCTION_LOG:
		return true, LEXEME_FUNCTION_LOG, LEXEME_CLASSIFICATION_OPERATOR
	case LEXEME_FUNCTION_NATURAL_LOG:
		return true, LEXEME_FUNCTION_NATURAL_LOG, LEXEME_CLASSIFICATION_INTERNAL_FUNCTION
	}

	// runSequence contains 2 or more runes and is not an internal function - check for numbers
	positive := true
	foundDecimal := false
	fitsInFloat := true
	var floatRepresentation float64
	var foundDigit float64
	for charPos, char := range runeSeq {

		// check for a sign if at the first character of the potential number
		if charPos == 0 && char == '+' {
			positive = true
			continue
		} else if charPos == 0 && char == '-' {
			positive = false
			continue
		}

		// check for a decimal, if one exists
		if char == '.' {
			if foundDecimal { // invalid to have multiple decimal places - break out and return invalid lexeme
				break
			} else {
				foundDecimal = true
				continue
			}
		}

		// check for a valid digit
		if char >= '0' && char <= '9' {

			foundDigit = float64(char - '0')

			// check if adding this new digit would fit in a float64
			if fitsInFloat && positive && (math.MaxFloat64-foundDigit)/10 >= floatRepresentation {
				floatRepresentation = (10 * floatRepresentation) + foundDigit
			} else if fitsInFloat && !positive && (math.MaxFloat64-foundDigit)/10 >= floatRepresentation {

			} else {
				fitsInFloat = false
			}

		} else { // invalid character - break out and return invalid lexeme
			break
		}

	}

	return false, LEXEME_INVALID, LEXEME_CLASSIFICATION_INVALID

}

func (l *Lexeme) String() string {
	return string(l.Runes)
}

func LexemeSliceToString(lexemes []Lexeme) string {

	result := ""

	for i := 0; i < len(lexemes); i++ {

		result += lexemes[i].String()
		if i < len(lexemes)-1 {
			result += ", "
		}

	}

	return result

}
