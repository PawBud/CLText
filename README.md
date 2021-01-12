# About

CLText is a CLI messaging service built using Go-Lang. It uses a TCP connection.

The whole application is built upon:-

- client: current user and its connection
- room
- command: from the client to the server
- server: which manages all incoming commands, as well it stores rooms and clients
- TCP server itself to accept network connections

# Commands
Currently the Chat supports the following commands:-
- `/nick <name>` - get a name, otherwise user will stay anonymous.
- `/join <name>` - join a room, if room doesn't exist, the new room will be created. User can be only in one room at the same time.
- `/rooms` - show list of available rooms to join.
- `/msg	<msg>` - broadcast message to everyone in a room.
- `/quit` - disconnects from the chat serer.

# How Do I use It?
**These Instructions are for Linux Ubuntu**
1. Install the source code 
2. open the directory in your IDE and run `cd chat/`
3. run `go build .` which builds an executable in the chat directory itself
4. run `./chat` which runs the exectuable. Now the server should start

# Screenshots:-
