package data

import (
	"fmt"
	"log"
	"strconv"

	"github.com/donovanmods/icarus-player-data/profile"
	"github.com/rivo/tview"
)

var statusField *tview.TextView

func PrintProfile(p *profile.ProfileData, app *tview.Application) tview.Primitive {
	saveCount := func(field string, text string) {
		if text == "" {
			return
		}

		count, err := strconv.Atoi(text)
		if err != nil {
			if statusField == nil {
				log.Fatal("unable to get status field")
			}
			statusField.SetText(fmt.Errorf("[red::b]unable to convert %s to int: %w[-::-]", text, err).Error())

			return
		}

		statusField.SetText("")

		p.SetCountFor(field, count)
		// saveAll()
	}

	form := tview.NewForm()
	form.SetBorder(false).SetBorderPadding(1, 1, 1, 1)

	// form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
	// 	index, btn := form.GetFocusedItemIndex()
	// 	if btn > 0 {
	// 		index = btn
	// 	}

	// 	switch event.Key() {
	// 	case tcell.KeyUp, tcell.KeyBacktab:
	// 		form.SetFocus(index - 1)
	// 	case tcell.KeyDown, tcell.KeyTab:
	// 		form.SetFocus(index + 1)
	// 	case tcell.KeyRune:
	// 		// VIM style navigation
	// 		switch event.Rune() {
	// 		case 'k':
	// 			form.SetFocus(index - 1)
	// 		case 'j':
	// 			form.SetFocus(index + 1)
	// 		default:
	// 			return event
	// 		}
	// 	default:
	// 		return event
	// 	}

	// 	return nil
	// })

	form.AddTextView("", "", 40, 2, true, false)
	statusField = form.GetFormItem(0).(*tview.TextView)
	if statusField == nil {
		log.Fatal("unable to get status field")
	}

	form.AddTextView("UserID", p.UserID, 40, 2, true, false)
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

	return form
}
