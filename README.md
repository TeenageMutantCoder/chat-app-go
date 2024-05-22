# chat-app-go

## A chat server written in go

I am working on this to help myself learn Golang. This was inspired by [this video](https://youtu.be/qmmQAAJzM54?si=XpcClNx5z1JBYm3z) where Tsoding did the same thing.

### Current commands

At the moment, each action needs to be prefixed by a command. Otherwise, the server will respond with an error message.

#### send

`send Hello World` sends a message that says "Hello World"

#### close

`close` closes the connection to the server.

### How to run

Prerequisites:

- Git (https://git-scm.com/)
- Go (https://go.dev/)
- `nc` util (https://en.wikipedia.org/wiki/Netcat) or an alternative way to connect to the tcp server

Steps:

1. Clone the repository. `git clone https://github.com/TeenageMutantCoder/chat-app-go.git` or `git clone git@github.com:TeenageMutantCoder/chat-app-go.git`
2. Change directories into the repository. `cd chat-app-go`
3. Run the server. `go run .`
4. Connect to the server. `nc localhost 8080`
