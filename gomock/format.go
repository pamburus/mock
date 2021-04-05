package gomock

import "strings"

/*
IndentString takes a string and indents each line by the specified amount.
*/
func IndentString(s string, indentation uint) string {
	components := strings.Split(s, "\n")
	result := ""
	indent := strings.Repeat(indentString, int(indentation))
	for i, component := range components {
		result += indent + component
		if i < len(components)-1 {
			result += "\n"
		}
	}

	return result
}

const indentString = "    "
