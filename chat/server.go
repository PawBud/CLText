package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type server struct {
	rooms    map[string]*room
	commands chan command
}

func newServer() *server {
	return &server{
		rooms:    make(map[string]*room),
		commands: make(chan command),
	}
}

func (s *server) run() {
	for cmd := range s.commands {
		switch cmd.id {
		case CmdNick:
			s.nick(cmd.client, cmd.args[1])
		case CmdJoin:
			s.join(cmd.client, cmd.args[1])
		case CmdRooms:
			s.listRooms(cmd.client)
		case CmdMsg:
			s.msg(cmd.client, cmd.args)
		case CmdHelp:
			s.help(cmd.client)
		case CmdQuit:
			s.quit(cmd.client)
		}
	}
}

func (s *server) newClient(conn net.Conn) {
	log.Printf("new client has joined: %s", conn.RemoteAddr().String())
	c := client{
		conn:     conn,
		nick:     "anonymous",
		commands: s.commands,
	}
	c.msg(fmt.Sprintf("type '/help' to see the list of commands.\n"))
	c.msg(fmt.Sprintf("type '/nick {nickname}' to assign a nickname to yourself.\n"))
	c.msg(fmt.Sprintf("If you want to pick a random nickname, type /randomnick"))
	
	// // when a user creates a session, a nickname should be picked before
	// // doing anything else
	// new_session := true
	c.readInput()
}

func (s *server) nick(c *client, nick string) {
	if nick != "" {
		c.nick = nick
	} else {
		log.Printf("Bc error hogga\n")
	}
	c.msg(fmt.Sprintf("all right, I will call you %s", nick))
}

func (s *server) join(c *client, roomName string) {
	r, ok := s.rooms[roomName]
	if !ok {
		r = &room{
			name:    roomName,
			members: make(map[net.Addr]*client),
		}
		s.rooms[roomName] = r
	}
	r.members[c.conn.RemoteAddr()] = c

	s.quitCurrentRoom(c)
	c.room = r

	r.broadcast(c, fmt.Sprintf("%s joined the room", c.nick))

	c.msg(fmt.Sprintf("welcome to %s", roomName))
}

func (s *server) listRooms(c *client) {
	var rooms []string
	for name := range s.rooms {
		rooms = append(rooms, name)
	}

	c.msg(fmt.Sprintf("available rooms: %s", strings.Join(rooms, ", ")))
}

func (_ *server) msg(c *client, args []string) {
	msg := strings.Join(args[1:], " ")
	c.room.broadcast(c, c.nick+": "+msg)
}

func (s *server) quit(c *client) {
	log.Printf("client has left the chat: %s", c.conn.RemoteAddr().String())

	s.quitCurrentRoom(c)

	c.msg("sad to see you go :(")
	err := c.conn.Close()
	if err != nil {
		return 
	}
}

func (s *server) help(c *client) {
	c.msg("the commands that can be used are: '/join', '/rooms', '/nick', '/msg' and '/quit'.")
}

func (s *server) quitCurrentRoom(c *client) {
	if c.room != nil {
		oldRoom := s.rooms[c.room.name]
		delete(s.rooms[c.room.name].members, c.conn.RemoteAddr())
		oldRoom.broadcast(c, fmt.Sprintf("%s has left the room", c.nick))
	}
}
