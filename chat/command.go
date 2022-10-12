package main

type commandID int

const (
	CmdNick commandID = iota
	CmdJoin
	CmdRooms
	CmdMsg
	CmdHelp
	CmdQuit
)

type command struct {
	id     commandID
	client *client
	args   []string
}
