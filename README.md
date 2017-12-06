# What is this?

It is a simple web service that accepts XML from a Rainforest Eagle device and saves the instantaneous,
usage and summation data into a sqlite database. It's quite incomplete but I wanted something to
simply record everything so I can run SQL queries over it.

# Usage

- `go build server.go` ... compiles sqlite3 into it
- `./server -path <path/to/sqlite-db.file> -listen <ip:port>`
