[private]
default:
  @just --list

run-backend:
  watchexec -c -f *.go -r 'go run -C backend . '
  #watchexec -c -f *.go -r 'go build -C backend . && ./backend/go-gin-react'
