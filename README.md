## Debounce Example

Start the server

```bash
$ go run main.go
```

Execute the following

```bash
$ curl 'localhost:8000?msg=HelloWorld!'
curl 'localhost:8000?msg=HelloWorld!!'
curl 'localhost:8000?msg=HelloWorld!!!'
```

The `msg` query parameter is logged after a second between requests has passed.

```
$ go run main.go
2020/08/06 23:25:21 HelloWorld!!!
```
