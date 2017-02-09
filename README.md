# ProxySpider
Crawl free proxy ip, provide HTTP api for user.

# How to use
* Requirement: golang development environment

* clone project.
```
$ git clone https://github.com/qlee001/ProxySpider.git
```

* goquery to parse html.
```
$ go get github.com/PuerkitoBio/goquery
```

* go build or go run directly
```
$ go build proxyspider
$ ./proxyspider
```
* use http api
```
$ curl "127.0.01:12345/get" -v
*   Trying 127.0.0.1...
* Connected to 127.0.01 (127.0.0.1) port 12345 (#0)
> GET /get HTTP/1.1
> Host: 127.0.01:12345
> User-Agent: curl/7.47.0
> Accept: */*
> 
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Thu, 09 Feb 2017 03:48:03 GMT
< Transfer-Encoding: chunked
< 
[{"Ip":"113.252.129.133","Port":"8380","Proto":"HTTP"}]
```

* get oversea ip
```
$ curl "127.0.01:12345/get?region=oversea" -v
*   Trying 127.0.0.1...
* Connected to 127.0.01 (127.0.0.1) port 12345 (#0)
> GET /get?region=oversea HTTP/1.1
> Host: 127.0.01:12345
> User-Agent: curl/7.47.0
> Accept: */*
> 
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Thu, 09 Feb 2017 03:49:14 GMT
< Content-Length: 757
< 
* Connection #0 to host 127.0.01 left intact
[{"Ip":"158.69.209.181","Port":"8888","Proto":"HTTP"},{"Ip":"40.138.64.36","Port":"8080","Proto":"HTTP"}]
```
