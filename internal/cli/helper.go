package cli

import (
	"fmt"
	"strings"
)

func renderMainMenu() string {
	return fmt.Sprintf(
		`[%d] Add
[%d] List
[%d] View
[%d] Edit
[%d] Delete
`, optAdd, optList, optView, optEdit, optDelete)
}

func renderHeader(header string) string {
	boxWidth := 40
	header = " " + header + " "

	// NOTE: should be 2?
	innerWidth := boxWidth - 3

	// Calculate padding for centering
	paddingLeft := (innerWidth - len(header)) / 2
	paddingRight := innerWidth - len(header) - paddingLeft

	line := "|" + strings.Repeat("-", paddingLeft) + header + strings.Repeat("-", paddingRight) + "|"

	return fmt.Sprintf(`===== Employee Management Program =====

+-------------------------------------+
%s
+-------------------------------------+

`, line)
}

func renderFooter(footer string) string {
	hint := "\n\n-------------------------------------------------------\n" + footer
	return hint
}

