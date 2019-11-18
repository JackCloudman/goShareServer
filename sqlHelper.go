package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var apis = make(map[string]int)
var cur *sql.DB

type File struct {
	ID     int
	Nombre string `json:"nombre"`
	Hash   string `json:"hash"`
	Size   int    `json:"size"`
	Type   string `json:"type"`
}
type Peer struct {
	ID   int    `json:"id"`
	IP   string `json:"ip"`
	Port string `json:"port"`
}

func startSql() {
	db, err := sql.Open("mysql", "root@/GoShare")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	cur = db
}
func searchByHash(hash string) *File {
	rows, err := cur.Query("SELECT * FROM Files WHERE hash = ?", hash)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	for rows.Next() {
		result := &File{}
		// Scan the value to []byte
		err = rows.Scan(&result.ID, &result.Nombre, &result.Hash, &result.Size, &result.Type)

		if err != nil {
			panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
		}
		return result
	}
	return nil
}
func updateLastSeen(peerID, fileID int) {
	stmtIns, err := cur.Query("UPDATE Peer_File set last_seen = now() WHERE id_peer = ? AND id_file = ?", peerID, fileID)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close()
}
func insertPeer(p *Peer) {
	stmtIns, err := cur.Prepare("INSERT INTO Peers VALUES( ?, ? ,?)") // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close()
	res, err := stmtIns.Exec(nil, p.IP, p.Port) // Insert tuples (i, i^2)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	id, err := res.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	p.ID = int(id)
}
func getPeer(ip, port string) *Peer {
	rows, err := cur.Query("SELECT * FROM Peers WHERE ip = ? AND port = ?", ip, port)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	result := &Peer{}
	result.IP = ip
	result.Port = port
	for rows.Next() {
		log.Println("Se encontro el Peer")
		// Scan the value to []byte
		err = rows.Scan(&result.ID, &result.IP, &result.Port)
		if err != nil {
			panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
		}
		return result
	}
	log.Println("No se encontro el peer")
	insertPeer(result)
	return result
}
func existPeerFile(peerID, fileID int) bool {
	log.Println("PeerID : ", peerID, " FileID:", fileID)
	rows, err := cur.Query("SELECT 1 FROM Peer_File WHERE id_file = ? AND id_peer = ?", fileID, peerID)
	if err != nil {
		log.Println("Error al obtener si existe el PeerID")
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	for rows.Next() {
		return true
	}
	return false
}
func insertPeerFile(peerID, fileID int) {
	stmtIns, err := cur.Prepare("INSERT INTO Peer_File VALUES( ?, ?, ? ,?)") // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close()
	_, err = stmtIns.Exec(nil, fileID, peerID, nil) // Insert tuples (i, i^2)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}
func insertFile(file *File) {
	stmtIns, err := cur.Prepare("INSERT INTO Files VALUES( ?, ? , ? , ? , ?)") // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close()
	res, err := stmtIns.Exec(nil, file.Nombre, file.Hash, file.Size, file.Type) // Insert tuples (i, i^2)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	id, err := res.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	file.ID = int(id)
}
func addPeerFile(file_id, peer_id int) {
	if existPeerFile(peer_id, file_id) {
		log.Println("Existe el Peerfile")
		updateLastSeen(peer_id, file_id)
	} else {
		log.Println("NO existe el Peerfile")
		insertPeerFile(peer_id, file_id)
	}

}
func uploadfile(file *File, p *Peer) {
	file_tmp := searchByHash(file.Hash)
	p = getPeer(p.IP, p.Port)
	log.Println("Se obtuvo el hash y el peer")
	if file_tmp != nil {
		log.Println("Ya se encontro el archivo")
		addPeerFile(file_tmp.ID, p.ID)
	} else {
		log.Println("No se encontro el archivo")
		insertFile(file)
		addPeerFile(file.ID, p.ID)
	}

}
