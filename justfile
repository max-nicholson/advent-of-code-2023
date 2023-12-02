run day:
    go run pkg/$(printf "%02.0f" {{day}})/main.go

test day:
    go test ./pkg/$(printf "%02.0f" {{day}})

fetch day:
    go run cmd/fetch.go --day {{day}}
