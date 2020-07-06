package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/marcusolsson/tui-go"
	"log"
	"net"
	"time"
)

func main() {
	var text string
	sidebar := tui.NewVBox(tui.NewLabel("Users"),
		tui.NewLabel("______"),
		tui.NewLabel("user1"),tui.NewSpacer())
	sidebar.SetBorder(true)

	conn, err := net.Dial("tcp", ":4000")
	if err != nil {
		log.Fatalf("could not %v", err)
	}

	history := tui.NewVBox()

	historyScroll := tui.NewScrollArea(history)
	historyScroll.SetAutoscrollToBottom(true)

	historyBox := tui.NewVBox(historyScroll)
	historyBox.SetBorder(true)

	input := tui.NewEntry()
	input.SetFocused(true)
	input.SetSizePolicy(tui.Expanding, tui.Maximum)

	inputBox := tui.NewHBox(input)
	inputBox.SetBorder(true)
	inputBox.SetSizePolicy(tui.Expanding, tui.Maximum)

	chat := tui.NewVBox(historyBox, inputBox)
	chat.SetFocused(true)
	chat.SetSizePolicy(tui.Expanding, tui.Expanding)

	user := color.GreenString("user")

	input.OnSubmit(func(entry *tui.Entry) {
		history.Append(tui.NewHBox(tui.NewLabel(time.Now().Format("03:45") ),
			tui.NewPadder(1, 0, tui.NewLabel(user)),
			tui.NewLabel(entry.Text()),
			tui.NewLabel(""),
			tui.NewSpacer()))
		input.SetText(text)
		fmt.Fprint(conn, text)
	})

	root := tui.NewHBox(sidebar, chat)

	ui, err := tui.New(root)
	if err != nil {
		log.Fatal(err)
	}

	ui.SetKeybinding("Esc", func() {
		ui.Quit()
	})

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}
