# What is this?

It is a simple web service that accepts XML from a Rainforest Eagle device and saves the instantaneous,
usage and summation data into a sqlite database. It's quite incomplete but I wanted something to
simply record everything so I can run SQL queries over it.

I use run this on a Raspberry PI 3 and it just keeps going.  

## Usage

- `go build server.go` ... compiles sqlite3 into it
- `./server <ip:port> <path/to/database>`

The server will create a new sqlite database if it doesn't exist.  You should be able
to query the database with any sqlite tool.

## Configuring the Eagle

* Settings > Cloud > Click the "+" to add a new recorder
* Name: rasp-pi-eagle-recorder
* url: http://<ip.address>:port/submit

It should take a few seconds for data to start flowing.  Don't worry if you see the red 429 status code.  There is some rate limiting done on the server side.

## Compiling a static binary

### Raspberry PI 3 B (arm-7)

```
xgo -targets "linux/arm-7" -ldflags='-w -extldflags "-static"' -out eagle-recorder .
```

## TODO

* Pretty much everything (including docs)
