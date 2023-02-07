# backend

## Pre requisites

- Go installed
- Environment setup (GOPATH and GOROOT) propperly set
- Link on how to https://go.dev/doc/install
---
## Folder structure
The logic about the folder structure is by domain
The only domain in this application is `notifications` here the global structs and interfaces are hold to minimize circular dependency among packages
the folders presents are:
- cmd: contains the main.go file, which initializes the logger, notifiers, storage and server
- internal: holds all the logic related to notifications domain
- logs: contains the txt file to store the messages recorded
