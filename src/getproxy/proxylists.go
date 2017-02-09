package getproxy

import (
    "fmt"
    "regexp"
)

func build_proxylists() []string {
    strs := make([]string, 0, 1)
    strs = append(strs, "http://www.proxylists.net/http_highanon.txt")
    strs = append(strs, "http://www.proxylists.net/http.txt")
    return strs
}

func parse_proxylists(body string) [][]string {
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










