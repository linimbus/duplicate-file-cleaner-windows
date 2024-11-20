package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

var search, delete, move, cancel *walk.PushButton

func ButtonEnable() {
	search.SetEnabled(true)
	delete.SetEnabled(true)
	move.SetEnabled(true)
	cancel.SetEnabled(true)
	ProcessUpdate(0)
}

func ButtonDisable() {
	search.SetEnabled(false)
	delete.SetEnabled(false)
	move.SetEnabled(false)
	cancel.SetEnabled(false)
}

func ActiveWidget() []Widget {
	return []Widget{
		PushButton{
			AssignTo: &search,
			Text:     "Search",
			OnClicked: func() {
				ButtonDisable()
				go func() {
					SearchFileActive()
					ButtonEnable()
				}()
			},
		},
		PushButton{
			AssignTo: &delete,
			Text:     "Delete",
			OnClicked: func() {
				ButtonDisable()
				DeleteAction(mainWindow, func(isNew bool) {
					DeleteFileActive(isNew)
					ButtonEnable()
				}, func() {
					ButtonEnable()
				})
			},
		},
		PushButton{
			AssignTo: &move,
			Text:     "Move",
			OnClicked: func() {
				ButtonDisable()
				MoveAction(mainWindow, func(isNew bool) {
					// MoveAction(isNew)
					ButtonEnable()
				}, func() {
					ButtonEnable()
				})

			},
		},
		PushButton{
			AssignTo: &cancel,
			Text:     "Cancel",
			OnClicked: func() {
				CloseWindows()
			},
		},
		HSpacer{},
	}
}
