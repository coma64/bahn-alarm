package models

type Connection struct {
	Id                       int
	TrackedById              int
	FromId                   int
	ToId                     int
	DepartureMarginMinutes   int
	DepartureInfoHistoryDays int
}
