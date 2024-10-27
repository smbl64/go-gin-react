[private]
default:
  @just --list

run-backend:
  watchexec -c -f *.go -r 'go run -C backend . '


run-frontend:
  npm run --prefix frontend dev

