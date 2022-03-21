# gochat

Realtime Chat API using WebSockets through [gorilla/websocket](https://github.com/gorilla/websocket). It can handle multiple connections at once, with all 
connections sending/receiving data between each other. (1 "room")

The connection pool is built in a way that it is in it's own goroutine,  communicating to the main thread via channels.

Also follows the unofficial but distinguished [project-layout](https://github.com/golang-standards/project-layout) file structure to practice writing production-grade code.

## Features
- Handling multiple connections via connection pool
- Dedicated goroutine for connection pools
- Familiar file structure
