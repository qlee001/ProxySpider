package getproxy

import (
    "fmt"
    //"strings"
    "regexp"
    //"github.com/PuerkitoBio/goquery"
)

func build_cn_proxy() []string {
    strs := make([]string, 0, 10)
    strs = append(strs, "http://cn-proxy.com/")
    return strs
}

func parse_cn_proxy(body string) [][]string {
    reg, err := regexp.Compile("<tr>[\\s\\S]*?<td>(.*?)</td>[\\s\\S]*?<td>(.*?)</td>[\\s\\S]*?<td>(.*?)</td>[\\s\\S]*?</tr>")
    if err != nil {
        fmt.Println(err)
        return nil
    }

    proxyaddrs := make([][]string, 0, 10)
    matchs := reg.FindAllStringSubmatch(body, -1)
    for _, match := range matchs {
        if len(match) != 4 {
            continue
        }
        reg2 := regexp.MustCompile("(?:\\d{1,3}\\.){1,3}\\d{1,3}")
        if reg2.MatchString(match[1]) {
            proxyaddrs = append(proxyaddrs, []string{match[1], match[2], "HTTP"})
        }
    }

    return proxyaddrs
}










