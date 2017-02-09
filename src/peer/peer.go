package peer

import (
    //"fmt"
    "net/http"
    "net/url"
    "time"
    "strings"
    "io/ioutil"
    )

type Peer struct {
    Ip      string
    Port    string
    Proto   string
    Status  int
}

func (p *Peer) GetByHTTPProxy(request_url string, proxy_addr string) ([]byte, error) {
	proxy, err := url.Parse(proxy_addr)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
            DisableKeepAlives : true,
		},
        Timeout :   3 * time.Second,
	}

	request, _ := http.NewRequest("GET", request_url, nil)

    resp, err := client.Do(request)
    if err != nil {
        return nil, err
    }
	defer resp.Body.Close()
	body,_ := ioutil.ReadAll(resp.Body)
    return body, nil
}


func (p *Peer)Check() int {
    var status int = 0
    var mainland int  = 1
    var abroad int = 2

    if p.Proto != "HTTP" && p.Proto != "HTTPS" {
        return status
    }

    url := "http://www.baidu.com/js/bdsug.js?v=1.0.3.0"
    substring := "(function(){var M=navigator.userAgent.indexOf("
    if p.Proto == "HTTPS" {
        url = "https://ss2.bdstatic.com/70cFsjip0QIZ8tyhnq/js/tangram-1.3.5.js"
        substring = "var T,baidu=T=baidu||{version:\"1.3.5\"};"
    }
    for i := 0; i < 3; i++ {
        if p.checkproxy(url, substring) {
            status |= mainland
            break
        }
    }

    if status == 0 {
        return status
    }

    url = "http://webcache.googleusercontent.com/search?q=cache:www.zhihu.com"
    substring = "<div id=\"google-cache-hdr\""
    if p.Proto == "HTTPS" {
        url = "https://apis.google.com/js/rpc:shindig_random.js?onload=init"
        substring = "var gapi=window.gapi=window.gapi||{};gapi._bs=new Date().getTime();("
    }
    for i := 0; i < 3; i++ {
        if p.checkproxy(url, substring) {
            status |= abroad
            break
        }
    }
    p.Status = status
    return status
}


func (p *Peer)checkproxy(url string, substring string) bool {
    proxy := "http://" + p.Ip + ":" + p.Port + "/"
    data, err := p.GetByHTTPProxy(url, proxy)
    if err != nil || strings.Contains(string(data), substring) == false {
        return false
    }
    return true
}


