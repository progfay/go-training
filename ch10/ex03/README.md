# ex10/ch03

## Log

```sh
❯ make
go build -o fetch.out gopl.io/ch1/fetch
./fetch.out 'http://gopl.io/ch1/helloworld?go-get=1' | grep 'go-import'
<meta name="go-import" content="gopl.io git https://github.com/adonovan/gopl.io">
```

## Using `curl`

```sh
❯ curl -L -D - 'http://gopl.io/ch1/helloworld?go-get=1'
HTTP/1.1 302 Found
Content-Type: text/html; charset=utf-8
Location: https://gopl.io/ch1/helloworld?go-get=1
Date: Fri, 01 Jan 2021 12:58:58 GMT
Content-Length: 62

HTTP/2 200
content-type: text/html; charset=utf-8
content-length: 212
date: Fri, 01 Jan 2021 12:58:59 GMT

<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="go-import" content="gopl.io git https://github.com/adonovan/gopl.io">
</head>
<body>
</body>
</html>
```
