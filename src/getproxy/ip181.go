package getproxy

import (
    "fmt"
    "strings"
    //"regexp"
    "github.com/PuerkitoBio/goquery"
)

func build_ip181() []string {
    strs := make([]string, 0, 10)
    for i:=1; i <= 2; i++ {
        strs = append(strs, fmt.Sprintf("http://www.ip181.com/daili/%d.html", i))
    }
    return strs
}

func parse_ip181(example string) [][]string {
    doc, err := goquery.NewDocumentFromReader(strings.NewReader(example))
    if err != nil {
        return nil
    }

    proxyaddrs := make([][]string, 0, 10)
    doc.Find(`table[class^="table table-hover panel-default panel ctable"] tbody tr`).Each(func(id int, s *goquery.Selection) {
        s2 := s.Find("td")
        if s2.Length() != 7 {
            return
        }

        ip := s2.Eq(0).Text()
        port := s2.Eq(1).Text()
        proto := s2.Eq(3).Text()
        //fmt.Println(ip, port, proto)
        addr := []string{ip, port, proto}
        proxyaddrs = append(proxyaddrs, addr)
    })

    return proxyaddrs
}










