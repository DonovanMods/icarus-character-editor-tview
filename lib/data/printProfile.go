package data

import (
	"fmt"
	"log"
	"strconv"

	"github.com/donovanmods/icarus-player-data/profile"
	"github.com/rivo/tview"
)

func PrintProfile(p *profile.ProfileData) tview.Primitive {
	saveCount := func(field string, text string) {
		if text == "" {
			return
		}

		count, err := strconv.Atoi(text)
		if err != nil {
			log.Print(fmt.Errorf("unable to convert %s to int: %w", text, err))
		}

		p.SetCountFor(field, count)
		p.Dirty = true
	}

	form := tview.NewForm()
	form.SetBorder(false).SetBorderPadding(1, 1, 1, 1)

	form.AddTextView("UserID", p.Profile.UserID, 40, 2, true, false)
	form.AddInputField("Credits", p.GetCountFor(profile.Credits), 10, nil, func(text string) {
		saveCount(profile.Credits, text)
	})
	form.AddInputField("Refund", p.GetCountFor(profile.Refund), 10, nil, func(text string) {
		saveCount(profile.Refund, text)
	})
	form.AddInputField("Purple Exotics", p.GetCountFor(profile.PurpleExotics), 10, nil, func(text string) {
		saveCount(profile.PurpleExotics, text)
	})
	form.AddInputField("Red Exotics", p.GetCountFor(profile.RedExotics), 10, nil, func(text string) {
		saveCount(profile.RedExotics, text)
	})

	// table := tview.NewTable().SetSelectable(true, true).SetBorders(false)
	// table.SetBorderPadding(1, 1, 1, 1)

	// table.SetCell(0, 0, tview.NewTableCell("UserID:").SetTextColor(tcell.ColorGreen).SetSelectable(false))
	// table.SetCell(0, 1, tview.NewTableCell(C.Profile.UserID).SetTextColor(tcell.ColorWhite).SetSelectable(false))

	// table.SetCell(2, 0, tview.NewTableCell("Credits:").SetTextColor(tcell.ColorGreen).SetSelectable(false))
	// table.SetCell(2, 1, tview.NewTableCell(C.getMetaCountFor(Credits)).SetTextColor(tcell.ColorYellow).SetSelectable(true))
	// table.SetCell(2, 2, tview.NewTableCell("Refund:").SetTextColor(tcell.ColorGreen).SetSelectable(false))
	// table.SetCell(2, 3, tview.NewTableCell(C.getMetaCountFor(Refund)).SetTextColor(tcell.ColorYellow).SetSelectable(true))

	// table.SetCell(4, 1, tview.NewTableCell("Purple").SetTextColor(tcell.ColorPurple).SetAlign(tview.AlignRight).SetSelectable(false))
	// table.SetCell(4, 2, tview.NewTableCell("Red").SetTextColor(tcell.ColorRed).SetAlign(tview.AlignRight).SetSelectable(false))

	// table.SetCell(5, 0, tview.NewTableCell("Exotics:").SetTextColor(tcell.ColorBlue).SetSelectable(false))
	// table.SetCell(5, 1, tview.NewTableCell(C.getMetaCountFor(Exotics)).SetTextColor(tcell.ColorPurple).SetAlign(tview.AlignRight).SetSelectable(true))
	// table.SetCell(5, 2, tview.NewTableCell(C.getMetaCountFor(ExoticsRed)).SetTextColor(tcell.ColorRed).SetAlign(tview.AlignRight).SetSelectable(true))

	return form
}
