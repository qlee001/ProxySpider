package getproxy

import (
    "fmt"
    "regexp"
)

func build_my_proxy() []string {
    strs := make([]string, 0, 1)
    for i := 1; i <= 10; i++ {
        strs = append(strs, fmt.Sprintf("http://www.my-proxy.com/free-proxy-list-%d.html", i))
    }
    strs = append(strs, "http://www.my-proxy.com/free-proxy-list-s1.html")
    strs = append(strs, "http://www.my-proxy.com/free-proxy-list-s2.html")
    strs = append(strs, "http://www.my-proxy.com/free-proxy-list-s3.html")
    return strs
}

func parse_my_proxy(body string) [][]string {
    reg, err := regexp.Compile("(\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3})\\:(\\d{2,5})")
    if err != nil {
        fmt.Println(err)
        return nil
    }

    proxyaddrs := make([][]string, 0, 10)
    matchs := reg.FindAllStringSubmatch(body, -1)
    for _, match := range matchs {
        if len(match) == 3 {
            proxyaddrs = append(proxyaddrs, []string{match[1], match[2], "HTTP"})
        }
    }

    return proxyaddrs
}










