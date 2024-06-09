# chat-app-go

A chat app written in go

## How to run

### From web site

Not yet available.

### From source code

Prerequisites:

- Git (https://git-scm.com/)
- Go (https://go.dev/)
- Node (https://nodejs.org/en)

Minimal steps:

1. Clone the repository. `git clone https://github.com/TeenageMutantCoder/chat-app-go.git` or `git clone git@github.com:TeenageMutantCoder/chat-app-go.git`
2. Change directories into the repository. `cd chat-app-go`
3. Run the server. `go run .`
4. Visit the web page at http://localhost:8080/

Steps for development:

1. Clone the repository. `git clone https://github.com/TeenageMutantCoder/chat-app-go.git`
2. Change directories into the repository. `cd chat-app-go`
3. Install node dependencies. `npm install`
4. Run the development server. `npm run dev`
5. In another terminal, allow tailwind to watch the template files. `npm run tw:watch`
6. Visit the web page at http://localhost:8080/

## Ideas for future development (not in any particular order)

- [ ] Store messages and users in a database (or don't, if a more privacy-focused app)
- [ ] Use websockets instead of polling to retrieve new messages
- [ ] Add ability to create/join/remove different chat rooms
- [ ] Add ability to send multimedia messages
- [ ] Deploy this app online
- [ ] Add better auth
- [ ] Make it easier to change nickname
- [ ] Add ability to mute/block users
- [ ] Add rate limiting, timeouts, bans to prevent spam (possibly optional per chat room)
- [ ] Add notifications when there are new messages
- [ ] Add roles and permissions per chat room
