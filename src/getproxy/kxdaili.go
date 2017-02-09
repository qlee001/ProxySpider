package getproxy

import (
    "fmt"
    "strings"
    "regexp"
    "github.com/PuerkitoBio/goquery"
)

func build_kxdaili() []string {
    strs := make([]string, 0, 10)
    for i:=1; i<=10; i++ {
        strs = append(strs, fmt.Sprintf("http://www.kxdaili.com/ipList/%d.html#ip", i))
    }
    return strs
}

func parse_kxdaili(body string) [][]string {
    doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
    if err != nil {
        return nil
    }

    addrs := make([][]string, 0, 10)
    reg := regexp.MustCompile("(?:\\d{1,3}\\.){3}\\d{1,3}")
    doc.Find("table[class=\"ui table segment\"] tbody tr").Each(func(i int, s *goquery.Selection) {
        s2 := s.Find("td")
        if s2.Length() != 7 {
            return
        }

        ip := s2.Eq(0).Text()
        if !reg.MatchString(ip) {
            return
        }

        port := s2.Eq(1).Text()
        proto := s2.Eq(3).Text()
        if proto == "HTTP" || proto == "HTTP,HTTPS" {
            addr := make([]string, 3)
            addr[0] = ip
            addr[1] = port
            addr[2] = "HTTP"
            addrs = append(addrs, addr)
        }

        if proto == "HTTP,HTTPS" {
            addr := make([]string, 3)
            addr[0] = ip
            addr[1] = port
            addr[2] = "HTTPS"
            addrs = append(addrs, addr)
        }
    })

    return addrs
}










