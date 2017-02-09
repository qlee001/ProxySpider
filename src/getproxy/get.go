package getproxy

import (
    "fmt"
    "time"
    "peer"
    "io/ioutil"
    "net/http"
)


type target struct {
    Build       func () []string
    Parse       func (string) [][]string
}

var targets = []target{
    target{build_cn_proxy, parse_cn_proxy},
    target{build_cnproxy, parse_cnproxy},
    target{build_proxylists, parse_proxylists},
    target{build_cybersyndrome, parse_cybersyndrome},
    target{build_cn_proxy, parse_cn_proxy},
    target{build_cz88, parse_cz88},
    target{build_kxdaili, parse_kxdaili},
    target{build_xicidaili, parse_xicidaili},
    target{build_ip181, parse_ip181},
    target{build_httpdaili, parse_httpdaili},

    target{build_pachong, parse_pachong},
}

var proxyaddrs [][]string

func get_proxy_from_url(url string, f func (string) [][]string, ch chan bool) {
    client := http.Client{
        Timeout : time.Second * 10,
    }

    request, err := http.NewRequest("GET", url, nil)
    if err != nil {
        ch<- false
        return
    }

    request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.87 Safari/537.36")

    resp, err := client.Do(request)
    if err != nil {
        fmt.Println("get failed: ", err)
        ch <- false
        return
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    addrs := f(string(body))
    fmt.Println("url: ", url, ", addrs num: ", len(addrs))
    for _, addr := range addrs {
        proxyaddrs = append(proxyaddrs, addr)
    }

    ch<- true
    return
}

func Get() []peer.Peer {
    chs := make([]chan bool, 0)

    for _, t := range targets {
        urls := t.Build()
        //fmt.Println(urls)
        for _, url := range urls {
            ch := make(chan bool)
            chs = append(chs, ch)
            go get_proxy_from_url(url, t.Parse, ch)
        }
    }

    for _, ch := range chs {
        <-ch
    }

    m := make(map[string]bool)
    addrs := make([]peer.Peer, 0, 100)
    for _, addr := range proxyaddrs {
        key := addr[0] + addr[1] + addr[2]
        if _, ok := m[key]; !ok {
            s := peer.Peer{addr[0], addr[1], addr[2], 0}
            addrs = append(addrs, s)
            m[key] = true
        }
    }

    return addrs
}








