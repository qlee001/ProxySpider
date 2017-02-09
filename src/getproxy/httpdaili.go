package getproxy

import (
    //"fmt"
    "strings"
    //"regexp"
    "github.com/PuerkitoBio/goquery"
)

func build_httpdaili() []string {
    strs := make([]string, 0, 10)
    strs = append(strs, "http://www.httpdaili.com/mfdl/")
    return strs
}

func parse_httpdaili(example string) [][]string {
    doc, err := goquery.NewDocumentFromReader(strings.NewReader(example))
    if err != nil {
        return nil
    }

    proxyaddrs := make([][]string, 0, 10)
    doc.Find(`div[class^="kb-item-wrap11"] table`).Each(func(protoid int, s *goquery.Selection) {
        s.Find("tr").Each(func(idx int, s2 *goquery.Selection) {
            s3 := s2.Find("td")
            if s3.Length() != 5 {
                return
            }

            ip := s3.Eq(0).Text()
            port := s3.Eq(1).Text()
            var proto string
            if protoid == 2 {
                proto = "HTTPS"
            } else {
                proto = "HTTP"
            }

            addr := []string{ip, port, proto}
            proxyaddrs = append(proxyaddrs, addr)
         })
    })

    return proxyaddrs
}










