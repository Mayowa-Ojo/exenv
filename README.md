## Exenv

> Zero dependency .env.example generator (is that fancy enough?)

Generates an `env.example` file in the current directory from the specified `.env` file.

#### Usage
Build the binary
```bash
$ git clone https://github.com/Mayowa-Ojo/exenv.git
$ cd exenv
$ go build -o bin/
$ go install
```
Commands

```bash
$ exenv -path .env -sep =
```