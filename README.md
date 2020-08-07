## Debounce Example

Start the server

```bash
$ go run main.go
```

Execute this multiple times

```bash
$ curl 'localhost:8000?msg=HelloWorld!'
$ curl 'localhost:8000?msg=HelloWorld!!'
$ curl 'localhost:8000?msg=HelloWorld!!!'
```

The `msg` is logged once after a second between requests has passed.

```
$ go run main.go
2020/08/06 23:25:21 HelloWorld!!!
```
