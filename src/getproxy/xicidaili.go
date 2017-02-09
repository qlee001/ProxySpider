package getproxy

import (
    "fmt"
    "strings"
    //"regexp"
    "github.com/PuerkitoBio/goquery"
)

func build_xicidaili() []string {
    strs := make([]string, 0, 10)
    for i:=1; i <= 2; i++ {
        strs = append(strs, fmt.Sprintf("http://www.xicidaili.com/wt/%d", i))
    }
    return strs
}

func parse_xicidaili(example string) [][]string {
    doc, err := goquery.NewDocumentFromReader(strings.NewReader(example))
    if err != nil {
        return nil
    }

    proxyaddrs := make([][]string, 0, 10)
    doc.Find("#ip_list tr").Each(func(id int, s *goquery.Selection) {
        s2 := s.Find("td")
        if s2.Length() != 10 {
            return
        }

        ip := s2.Eq(1).Text()
        port := s2.Eq(2).Text()
        proto := s2.Eq(5).Text()
        //fmt.Println(ip, port, proto)
        addr := []string{ip, port, proto}
        proxyaddrs = append(proxyaddrs, addr)
    })

    return proxyaddrs
}










