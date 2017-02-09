package getproxy

import (
    "fmt"
    "strings"
    "regexp"
    "github.com/PuerkitoBio/goquery"
)

func build_cz88() []string {
    strs := make([]string, 0, 10)
    strs = append(strs, "http://www.cz88.net/proxy/index.shtml")
    for i:=1; i<=10; i++ {
        strs = append(strs, fmt.Sprintf("http://www.cz88.net/proxy/http_%d.shtml", i))
    }
    return strs
}

func parse_cz88(body string) [][]string {
    doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
    if err != nil {
        return nil
    }

    addrs := make([][]string, 0, 10)
    reg := regexp.MustCompile("(?:\\d{1,3}\\.){3}\\d{3}")
    doc.Find("#boxright div ul li").Each(func(i int, s *goquery.Selection) {
        s2 := s.Find("div")
        if s2.Length() != 4 {
            return
        }

        ip := s2.Eq(0).Text()
        if !reg.MatchString(ip) {
            return
        }
        port := s2.Eq(1).Text()
        addr := make([]string, 3)
        addr[0] = ip
        addr[1] = port
        addr[2] = "HTTP"
        addrs = append(addrs, addr)
    })

    return addrs
}










