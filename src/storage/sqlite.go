package storage

import (
    "peer"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

var sqlite_path = "/home/ubuntu/ProxySpider/src/proxyspider.db"


func sqlite_select(tname string) ([]peer.Peer, error) {
	db, err := sql.Open("sqlite3", sqlite_path)
    if err != nil {
        return nil, err
    }

    rows, err := db.Query("SELECT ip, port, protocol FROM " + tname + ";")
    if err != nil {
        return nil, err
    }

    peers := make([]peer.Peer, 0, 4096)
	for rows.Next() {
        var ip string
        var port string
        var protocol string
        err = rows.Scan(&ip, &port, &protocol)
        if err != nil {
            continue
        }
        one := peer.Peer{ip, port, protocol, 0}
        peers = append(peers, one)
    }

    return peers, nil
}


func get_available_from_sqlite() ([]peer.Peer, error) {
    return sqlite_select("available")
}

func get_backup_from_sqlite() ([]peer.Peer, error) {
    return sqlite_select("backup")
}

func sqlite_append(tname string, peers []peer.Peer) error {
	db, err := sql.Open("sqlite3", sqlite_path)
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("insert into ? (ip, port, protocol) values(?, ?, ?);")
	if err != nil {
		return err
	}
	defer stmt.Close()

    for _,one := range peers {
		_, err = stmt.Exec(tname, one.Ip, one.Port, one.Proto)
		if err != nil {
			return err
		}
	}

	return nil
}


func sqlite_replace(tname string, peers []peer.Peer) error {
	db, err := sql.Open("sqlite3", sqlite_path)
	if err != nil {
		return err
	}
	defer db.Close()

    _, err = db.Exec("DELETE FROM " + tname + ";")
    if err != nil {
        return err
    }

	stmt, err := db.Prepare("insert into " + tname + "(ip, port, protocol, status) values(?, ?, ?, ?);")
	if err != nil {
		return err
	}
	defer stmt.Close()

    for _,one := range peers {
		_, err = stmt.Exec(one.Ip, one.Port, one.Proto, one.Status)
		if err != nil {
			return err
		}
	}

	return nil
}

func update_available_to_sqlite(peers []peer.Peer) error {
    return sqlite_replace("available", peers)
}

func update_backup_to_sqlite(peers []peer.Peer) error {
    return sqlite_replace("backup", peers)
}

func record_available_to_sqlite(peers []peer.Peer) error {
    return sqlite_append("record", peers)
}

