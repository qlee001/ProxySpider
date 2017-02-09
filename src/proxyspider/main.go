package main

import (
	"encoding/json"
	"fmt"
	"getproxy"
	"net/http"
	"peer"
	"time"
)

var proxy_backup []peer.Peer
var proxy_available []peer.Peer

func getHandler(w http.ResponseWriter, req *http.Request) {
	data := proxy_available
	type json_struct struct {
		Ip    string
		Port  string
		Proto string
	}
	decode := make([]json_struct, 0, len(data))

	opt := req.FormValue("region")
	oversea := false
	if opt == "oversea" {
		oversea = true
	}

	for _, v := range data {
		if oversea && v.Status&2 == 0 {
			continue
		}
		one := json_struct{v.Ip, v.Port, v.Proto}
		decode = append(decode, one)
	}

	b, err := json.Marshal(decode)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func checkone(p *peer.Peer, ch chan bool) {
	status := p.Check()
	if status != 0 {
		ch <- true
	} else {
		ch <- false
	}
	return
}

func check_peers(peers []peer.Peer) []peer.Peer {

	count := len(peers)
	chs := make([]chan bool, count)
	results := make([]peer.Peer, 0, count)
	for i := 0; i < count; i += 100 {
		max := i + 100
		if max >= count {
			max = count
		}
		for j := i; j < max; j++ {
			ch := make(chan bool)
			chs[j] = ch
			one := &peers[j]
			go checkone(one, ch)
		}

		for j := i; j < max; j++ {
			ok := <-chs[j]
			if ok {
				fmt.Println(peers[j])
				results = append(results, peers[j])
			}
		}
	}

	return results
}

func check() {
	results := check_peers(proxy_backup)
	if len(results) != 0 {
		proxy_available = results
	}
}

func crawl() {
	peers := getproxy.Get()
	fmt.Println("from get: ", len(peers))
	/*
	   peers, err := storage.Get_backup()
	   if err != nil {
	       fmt.Println("get failed, err: ", err)
	   }

	*/
	results := check_peers(peers)
	proxy_backup = results
	proxy_available = results
	if len(proxy_backup) == 0 {
		fmt.Println("no proxy available")
		return
	}
	return
}

func cycle() {

	count := 100
	for {
		count++
		if count > 24 {
			crawl()
			count = 0
		} else {
			check()
			return
		}

		time.Sleep(time.Minute * 15)
	}

	return
}

func main() {
	go cycle()
	http.HandleFunc("/get", getHandler)
	http.ListenAndServe(":12345", nil)
}
