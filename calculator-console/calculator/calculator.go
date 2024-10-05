package calculator

import (
	"fmt"
)

const (
	ERROR_INVALID_COMMAND string = "INVALID COMMAND"
)

// Calculator defines a calculator which accepts a string representation of a command, processes it, and outputs
// results in string format. Under the hood, the Command is immediately cast to []rune.
type Calculator struct {
	Command      []rune
	lexemes      []Lexeme
	hasProcessed bool
	Result       string
	Error        error
}

func NewCalculator() *Calculator {
	return &Calculator{}
}

// UpdateCommandString updates the calculator instance with a new command string to process
func (c *Calculator) UpdateCommandString(newCommand string) { // update the command string
	c.Command = []rune(newCommand)
	c.hasProcessed = false
}

// GetResult processes the command string if it has not already been processed since the last update and returns
// the result as a string or an error if a problem occured
func (c *Calculator) GetResult() (result string, err error) { // get the result of calculating the current command string
	if !c.hasProcessed { // process the current command string if it hasn't already been handled

		// loop over command rune sequence and attempt to consume lexemes.
		// since lexemes can be of length 1-N, we must keep consuming chars until the current potential lexeme
		// becomes invalid to ensure full consumption of each lexeme and prevent errors.
		var isLexeme bool
		var lexemeType LexemeType
		var classification LexemeClassification
		left := 0
		right := 1
		for left < len(c.Command) { // manually adjust pointers and keep looping until the entire sequence is consumed

			if left == len(c.Command)-1 { // final char reached

				isLexeme, lexemeType, classification = IsLexeme(c.Command[left : left+1])

				if isLexeme { // last char is a lexeme - add it to the list and we're done

					c.lexemes = append(c.lexemes, Lexeme{
						Runes:          c.Command[left : left+1],
						Type:           lexemeType,
						Classification: classification,
					})

					left++ // technically unneeded since we're explicitly calling break after this line
					break

				} else { // last char is not a lexeme - command string is invalid

					c.Result = ""
					c.Error = fmt.Errorf(
						"%s: unable to consume final character as a valid lexeme. Command string: %q; consumed lexemes: %s; invalid final char: %v",
						ERROR_INVALID_COMMAND, string(c.Command), LexemeSliceToString(c.lexemes), c.Command[left:left+1])
					c.hasProcessed = true
					return c.Result, c.Error // set processed flag to true and return the error string

				}

			}

			// if we reach this point, there is at least one more character to the right of the "left" char -
			// attempt to build the longest valid lexeme from that position before advancing
			for right = len(c.Command) - 1; right > left; right-- {

				isLexeme, lexemeType, classification = IsLexeme(c.Command[left : right+1])

				if isLexeme { // longest valid lexeme found!
					c.lexemes = append(c.lexemes, Lexeme{
						Runes:          c.Command[left : right+1],
						Type:           lexemeType,
						Classification: classification,
					})

					// advance left index
					left = right + 1

					break
				}

			}

			if right <= left { // right will be <= left if no valid lexeme remains between left and the end of the command string

				c.Result = ""
				c.Error = fmt.Errorf(
					"%s: unable to consume the next lexeme. Command string: %q; consumed lexemes: %s",
					ERROR_INVALID_COMMAND, string(c.Command), LexemeSliceToString(c.lexemes))
				c.hasProcessed = true
				return c.Result, c.Error // set processed flag to true and return the error string

			}

		}

		c.hasProcessed = true
	}

	// process lexeme list and return result
	c.Result = c.processLexemes()
	return c.Result, nil
}

// processLexemes is a helper function to perform the actual calculations on a list of valid lexemes
func (c *Calculator) processLexemes() (result string) {

	result = fmt.Sprintf("made it into c.processLexemes! list of consumed lexemes: %s", LexemeSliceToString(c.lexemes))

	return result

}
