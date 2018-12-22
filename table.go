package main

import (
	"fmt"

	"github.com/geistesk/procon/pc"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

const (
	columnPros int = 0
	columnCons int = 1
)

var (
	table            *tview.Table
	tblPros, tblCons []*pc.Entry
)

// entryRepresentation returns a representing string for the table.
func entryRepresentation(entry *pc.Entry) string {
	if entry.IsPro() {
		return fmt.Sprintf("[white]%s [green]%2d ", entry.Text, entry.AbsValue())
	} else {
		return fmt.Sprintf("[red] %-2d [white]%s", entry.AbsValue(), entry.Text)
	}
}

// setupTableHeader creats the table's header.
func setupTableHeader() {
	cols := []struct {
		no   int
		text string
	}{
		{columnPros, "Pros"},
		{columnCons, "Cons"},
	}

	for _, col := range cols {
		table.SetCell(0, col.no,
			tview.NewTableCell(col.text).
				SetSelectable(false).
				SetAlign(tview.AlignCenter).
				SetTextColor(tcell.ColorYellow))
	}
}

// syncListToTable draws all elements of the dataList to the table.
func syncListToTable() {
	if dataList == nil {
		return
	}

	tblPros, tblCons = dataList.ProsConsEntries()

	for i := 0; i < len(tblPros); i++ {
		table.SetCell(i+1, columnPros,
			tview.NewTableCell(entryRepresentation(tblPros[i])).
				SetAlign(tview.AlignRight))
	}

	for i := 0; i < len(tblCons); i++ {
		table.SetCell(i+1, columnCons,
			tview.NewTableCell(entryRepresentation(tblCons[i])).
				SetAlign(tview.AlignLeft))
	}
}

// redrawTable redraws the table.
func redrawTable() {
	table.Clear()

	setupTableHeader()
	syncListToTable()
}

// removeTableEntry removes the current entry if one is selected.
func removeTableEntry() {
	var entry *pc.Entry

	r, c := table.GetSelection()
	if c == columnPros {
		if r-1 >= len(tblPros) {
			return
		}
		entry = tblPros[r-1]
	} else {
		if r-1 >= len(tblCons) {
			return
		}
		entry = tblCons[r-1]
	}

	dataList.RemoveEntry(*entry)
	redrawTable()
}

// tableHandleKeyPress is called when the table is in focus and a key was pressed.
func tableHandleKeyPress(event *tcell.EventKey) {
	if event.Key() != tcell.KeyRune {
		return
	}

	switch event.Rune() {
	case 'x':
		removeTableEntry()
	}
}

// setupTable creates the pros and cons table.
func setupTable() {
	table = tview.NewTable().
		SetSeparator(tview.Borders.Vertical).
		SetFixed(1, 1).
		SetSelectable(true, true)

	setupTableHeader()
	syncListToTable()
}