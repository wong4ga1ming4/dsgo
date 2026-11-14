package output

import (
	"encoding/json"
	"io"
	"strings"
)

// Write sends either indented JSON or plain text. Text must already end with a newline.
func Write(w io.Writer, asJSON bool, text string, value any) error {
	if !asJSON {
		_, err := io.WriteString(w, text)
		return err
	}
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	return enc.Encode(value)
}

func JoinPages(parts []string) string {
	clean := make([]string, 0, len(parts))
	for _, part := range parts {
		clean = append(clean, strings.TrimRight(part, "\n"))
	}
	return strings.Join(clean, "\n\n") + "\n"
}
