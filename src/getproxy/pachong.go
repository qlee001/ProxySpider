package getproxy

import (
    "strings"
    "regexp"
    "strconv"
    "github.com/PuerkitoBio/goquery"
)

func build_pachong() []string {
    strs := make([]string, 0, 10)
    strs = append(strs, "http://pachong.org/")
    strs = append(strs, "http://pachong.org/anonymous.html")
    strs = append(strs, "http://pachong.org/transparent.html")
    strs = append(strs, "http://pachong.org/area/short/name/cn.html")
    strs = append(strs, "http://pachong.org/area/short/name/br.html")
    strs = append(strs, "http://pachong.org/area/short/name/us.html")
    strs = append(strs, "http://pachong.org/area/short/name/ve.html")
    strs = append(strs, "http://pachong.org/area/short/name/in.html")
    return strs
}

func parse_pachong(body string) [][]string {
    doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
    if err != nil {
        return nil
    }
    
    reg1 := regexp.MustCompile("^var\\s+(\\w+)=(\\d+)\\+(\\d+)$")
    reg2 := regexp.MustCompile("^var\\s+(\\w+)=(\\d+)\\+(\\d+)\\^(\\w+)$")
    reg_port := regexp.MustCompile("document\\.write\\(\\((\\d+)\\^(\\w+)\\)\\+(\\d+)\\)")

    m := make(map[string]int)
    doc.Find("head script").Each(func(i int, s *goquery.Selection) {
        if i != 2 {
            return
        }
        script := s.Text()


        for _, s := range strings.Split(script, ";") {
            match := reg1.FindStringSubmatch(s)
            if match != nil {
                arg1, _ := strconv.Atoi(match[2])
                arg2, _ := strconv.Atoi(match[3])
                m[match[1]] = arg1 + arg2
                continue
            }

            match = reg2.FindStringSubmatch(s)
            if match != nil {
                arg1, _ := strconv.Atoi(match[2])
                arg2, _ := strconv.Atoi(match[3])
                arg3, _ := strconv.Atoi(match[4])
                m[match[1]] = arg1 + arg2 ^ arg3
                continue
            }
        }
        return
    })

    if len(m) == 0 {
        return nil
    }

    addrs := make([][]string, 0, 10)
    doc.Find("tbody tr").Each(func(i int, s *goquery.Selection) {
        if s.Find("td").Length() != 7 {
            return
        }
        var ip, port, proto string
        s.Find("td").Each(func(i2 int, s2 *goquery.Selection) {
            switch i2 {
            case 1:
                ip = s2.Text()
            case 2:
                portstr := s2.Text()
                match := reg_port.FindStringSubmatch(portstr)
                if match == nil {
                    return
                }
                arg1, _ := strconv.Atoi(match[1])
                arg2, _ := m[match[2]]
                arg3, _ := strconv.Atoi(match[3])
                port = strconv.Itoa(arg1 ^ arg2 + arg3)
            case 4:
                proto = s2.Text()
                if proto == "socks4" {
                    proto = "SOCKS4"
                } else if proto == "socks5" {
                    proto = "SOCKS5"
                } else {
                    proto = "HTTP"
                }

            default:
                return
            }
        })
        if len(ip) == 0 || len(port) == 0 || len(proto) == 0 {
            return
        }

        addr := make([]string, 3)
        addr[0] = ip
        addr[1] = port
        addr[2] = proto
        addrs = append(addrs, addr)
    })

    return addrs
}










