package db

type TimeSelection string

const (
	TimeSelectionLatest TimeSelection = "latest"
	TimeSelectionBefore TimeSelection = "before"
	TimeSelectionAfter  TimeSelection = "after"
	TimeSelectionAt     TimeSelection = "at"
)
