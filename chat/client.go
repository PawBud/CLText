package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type client struct {
	conn     net.Conn
	nick     string
	room     *room
	commands chan<- command
}

func (c *client) readInput() {
	for {
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			return
		}
		msg = strings.Trim(msg, "\r\n")
		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])
		//make sure that the length of the input is at least 1
		if len(args) == 1 && args[0] != "/help" && args[0] != "/rooms" {
			c.err(fmt.Errorf("Cannot pass empty characters after the commands except the help command.\n" +
				"Example of a valid Command: /msg Once upon a time in a far far away galaxy" +
				"\n"))
		}
		switch cmd {
		case "/nick":
			c.commands <- command{
				id:     CmdNick,
				client: c,
				args:   args,
			}
		case "/join":
			c.commands <- command{
				id:     CmdJoin,
				client: c,
				args:   args,
			}
		case "/rooms":
			c.commands <- command{
				id:     CmdRooms,
				client: c,
			}
		case "/msg":
			c.commands <- command{
				id:     CmdMsg,
				client: c,
				args:   args,
			}
		case "/help":
			c.commands <- command{
				id:     CmdHelp,
				client: c,
			}
		case "/quit":
			c.commands <- command{
				id:     CmdQuit,
				client: c,
			}
		default:
			c.err(fmt.Errorf("unknown command: %s", cmd))
		}
	}
}

func (c *client) err(err error) {
	c.conn.Write([]byte("err: " + err.Error() + "\n"))
}

func (c *client) msg(msg string) {
	c.conn.Write([]byte("> " + msg + "\n"))
}
