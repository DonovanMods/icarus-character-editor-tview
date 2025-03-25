package main

import (
	"fmt"
	"log"

	"github.com/donovanmods/icarus-character-editor/lib/data"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var App *tview.Application

func main() {
	startApp()
}

func saveAll(view *tview.Flex) {
	var errors int

	view.Clear().SetTitle("[ Writing Player Data ]")

	if err := data.ProfileData.Write(); err != nil {
		view.AddItem(textView(fmt.Sprintf("[red::b]error writing profile data: %s[-::-]", err)), 0, 1, false)
		errors++
	}

	if err := data.CharacterData.Write(); err != nil {
		view.AddItem(textView(fmt.Sprintf("[red::b]error writing character data: %s[-::-]", err)), 0, 1, false)
		errors++
	}

	if errors == 0 {
		view.AddItem(textView("[green]Player Data Saved Successfully[-]"), 0, 1, false)
	}
}

func startApp() {
	if err := data.Read(); err != nil {
		log.Fatal("error reading player data: %w", err)
	}

	// Create a new TUI application
	App = tview.NewApplication()
	App.EnableMouse(true)

	// Create a TextView that will display the character list in the TUI
	mainMenu := tview.NewList().SetHighlightFullLine(true).SetWrapAround(false)
	mainMenu.
		SetBorder(true).
		SetBorderPadding(1, 1, 1, 1).
		SetTitle("[ Characters ]")
	mainMenu.SetCurrentItem(0) // Set the first item as the current item

	dataView := tview.NewFlex()
	dataView.SetBorder(true).SetBorderPadding(0, 0, 0, 0)

	// Iterate through characters and add each item to the character list
	for i, item := range data.CharacterData.Characters {
		mainMenu.AddItem(item.Name, "", rune(i+49), nil)
	}

	// Add a profile option
	mainMenu.AddItem("Player Profile", "", 'p', nil)

	// Add a quit option
	mainMenu.AddItem("Exit the Program", "", 'q', func() {
		App.Stop()
	})

	mainMenu.AddItem("Write Player Data", "", 'w', func() {
		dataView.Clear().SetTitle("[ Write Player Data ]")
		saveAll(dataView)
	})

	mainMenu.SetSelectedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		switch shortcut {
		case 'p':
			App.SetFocus(dataView)
		case 'q':
			App.Stop()
		case 'w':
			saveAll(dataView)
		}
	})

	// Set the function to be called when a character is selected
	mainMenu.SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		if shortcut == rune('p') {
			dataView.Clear().SetTitle("[ Player Profile ]")
			dataView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
				switch event.Key() {
				case tcell.KeyEsc:
					App.SetFocus(mainMenu)
					return nil
				}
				return event
			})
			dataView.AddItem(data.PrintProfile(data.ProfileData, App), 0, 1, true)

			return
		}

		if shortcut == rune('q') {
			dataView.Clear().SetTitle("[ Quit ]")
			dataView.AddItem(textView("[green]Exit the Character Editor without Saving[-]"), 0, 1, false)

			return
		}

		if shortcut == rune('w') {
			dataView.Clear().SetTitle("[ Write Player Data]")
			dataView.AddItem(textView("[green]Save/Write the Player Data[-]"), 0, 1, false)

			return
		}

		// Print the selected character data
		dataView.Clear().SetTitle("[ Character Data ]")
		dataView.AddItem(data.PrintCharacter(data.CharacterData, index), 0, 1, false)
	})

	// Print the first character data by default
	dataView.Clear().SetTitle("[ Character Data ]")
	dataView.AddItem(data.PrintCharacter(data.CharacterData, 0), 0, 1, false)

	// Create a layout using Flex to display the character list and the form side by side
	flex := tview.NewFlex()
	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// Only if the mainMenu is in focus
		if App.GetFocus() == dataView {
			return event
		}

		switch event.Key() {
		case tcell.KeyRune:
			// VIM style navigation
			switch event.Rune() {
			case 'k':
				App.SetFocus(mainMenu.SetCurrentItem(mainMenu.GetCurrentItem() - 1))
			case 'j':
				App.SetFocus(mainMenu.SetCurrentItem(mainMenu.GetCurrentItem() + 1))
			default:
				return event
			}
		case tcell.KeyCtrlR:
			mainMenu.Clear()
			dataView.Clear()
			App.SetFocus(mainMenu.SetCurrentItem(0)).ForceDraw()
		default:
			return event
		}
		return nil
	})

	flex.AddItem(mainMenu, 0, 1, true)  // Left side
	flex.AddItem(dataView, 0, 4, false) // Right side

	// Start the TUI application
	if err := App.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}

func textView(text string) tview.Primitive {
	view := tview.NewTextView()
	view.
		SetDynamicColors(true).
		SetBorderPadding(1, 1, 1, 1)

	fmt.Fprintf(view, "[green]%s[-]", text)

	return view
}
