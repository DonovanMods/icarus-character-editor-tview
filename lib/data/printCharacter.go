package data

import (
	"fmt"
	"strings"

	"github.com/donovanmods/icarus-player-data/character"
	"github.com/rivo/tview"
)

func PrintCharacter(c *character.CharacterData, index int) tview.Primitive {
	subView := tview.NewTextView()
	subView.SetDynamicColors(true).SetBorderPadding(1, 1, 1, 1)

	if index < 0 || index >= len(c.Characters) {
		fmt.Fprintln(subView, "Invalid Character")
		return subView
	}

	char := &c.Characters[index]

	// Iterate through characters and print each item to the TextView
	fmt.Fprint(subView, nameString(char))
	fmt.Fprint(subView, xpString(char))
	fmt.Fprintf(subView, "Known Talents: %d\n", len(char.Talents))

	return subView
}

func nameString(c *character.Character) string {
	status := make([]string, 0, 2)
	statusString := ""

	if c.IsDead {
		status = append(status, "[red::bi]DEAD[-::-]")
	}

	if c.IsAbandoned {
		status = append(status, "[purple::bi]Abandoned[-::-]")
	}

	if len(status) > 0 {
		statusString = fmt.Sprintf("(%s)", strings.Join(status, " & "))
	}

	return fmt.Sprintf("[yellow::b]%s[-::-] %s\n\n", c.Name, statusString)
}

func xpString(c *character.Character) string {
	return fmt.Sprintf("Level: %-3d (XP: %d%s)\n\n", c.Level(), c.XP, xpDebtString(c))
}

func xpDebtString(c *character.Character) string {
	if c.XP_Debt > 0 {
		return fmt.Sprintf("; Debt: %d", c.XP_Debt)
	}

	return ""
}
