package getproxy

import (
    "fmt"
    "strings"
    "regexp"
    "github.com/PuerkitoBio/goquery"
)

func build_cnproxy() []string {
    strs := make([]string, 0, 10)
    for i:=1; i <= 10; i++ {
        strs = append(strs, fmt.Sprintf("http://www.cnproxy.com/proxy%d.html", i))
    }
    return strs
}

func parse_cnproxy(example string) [][]string {
    doc, err := goquery.NewDocumentFromReader(strings.NewReader(example))
    if err != nil {
        return nil
    }

    proxyaddrs := make([][]string, 0, 10)
    doc.Find("#proxylisttb table tr").Each(func(id int, s *goquery.Selection) {
        html, err := s.Html()
        if err != nil {
            return
        }

        reg, err := regexp.Compile("<td>(.+?)<script.*?>document.write\\(\":\"\\+(.+?)\\)</script></td><td>(.*?)</td><td>.*?</td><td>(.*?)</td>")
        if err != nil {
            return
        }

        captures := reg.FindStringSubmatch(html)
        if len(captures) != 5 {
            return
        }

        plist := strings.Split(captures[2], "+")
        portstr := ""
        for _, char := range plist {
            switch char {
            case "v":
                portstr += "3"
            case "m":
                portstr += "4"
            case "a":
                portstr += "2"
            case "l":
                portstr += "9"
            case "q":
                portstr += "0"
            case "b":
                portstr += "5"
            case "i":
                portstr += "7"
            case "w":
                portstr += "6"
            case "r":
                portstr += "8"
            case "c":
                portstr += "1"
            default:
                break
            }
        }
        addr := []string{captures[1], portstr, captures[3]}
        proxyaddrs = append(proxyaddrs, addr)

    })

    return proxyaddrs
}










