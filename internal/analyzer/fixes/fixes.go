package fixes

import (
	"go/ast"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

const (
	msgFixDisallowSpecialChars = "Remove special characters or emojis from log message"
	msgFixLowercaseStart       = "Convert log message to start with lowercase letter"
)

func FixDisallowSpecialChars(lit *ast.BasicLit) analysis.SuggestedFix {
	unquoted, _ := strconv.Unquote(lit.Value)
	var builder strings.Builder

	for _, r := range unquoted {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r) {
			builder.WriteRune(r)
		}
	}

	fixed := builder.String()
	return analysis.SuggestedFix{
		Message: msgFixDisallowSpecialChars,
		TextEdits: []analysis.TextEdit{
			{
				Pos:     lit.Pos(),
				End:     lit.End(),
				NewText: []byte(strconv.Quote(fixed)),
			},
		},
	}
}

func FixLowercaseStart(lit *ast.BasicLit) analysis.SuggestedFix {
	unquoted, _ := strconv.Unquote(lit.Value)

	var fixed string
	for i, r := range unquoted {
		if unicode.IsLetter(r) {
			fixed = string(unicode.ToLower(r)) + unquoted[i+1:]
			break
		}
	}

	return analysis.SuggestedFix{
		Message: msgFixLowercaseStart,
		TextEdits: []analysis.TextEdit{
			{
				Pos:     lit.Pos(),
				End:     lit.End(),
				NewText: []byte(strconv.Quote(fixed)),
			},
		},
	}
}
