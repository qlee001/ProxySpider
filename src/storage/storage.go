package storage

import (
    "fmt"
    "peer"
)


func Get_available() ([]peer.Peer, error) {
    return  get_available_from_sqlite()
}

func Get_backup() ([]peer.Peer, error) {
    result, err := get_backup_from_sqlite()
    if err != nil {
        return nil, err
    }

    return result, nil
}


func Update_available(s[]peer.Peer) {
    update_available_to_sqlite(s)
    record_available_to_sqlite(s)
	return
}

func Update_backup(s []peer.Peer) {
    err := update_backup_to_sqlite(s)
    if err != nil {
        fmt.Println("update backup error: ", err)
    }
    return
}
