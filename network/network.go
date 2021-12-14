package network

import "net"

type Package struct {
	Option int
	data   string
}

const (
	ENDBYTES = "\000\005\007\001\001\007\005\000"
	WAITTIME = 5
)

func Send(address string, pack *Package) *Package {
	conn, err := net.Dial("tcp", address)
	if err != nil{
		return nil
	}
	defer conn.Close()
	conn.Write([]byte(SerializePackage(pack) + ENDBYTES))
	var {
		res = new(Package)
		ch := make(chan bool)
	}
	go func() {
		res = readPackage(conn)
		ch <- True
	}()
	select{
	case <- ch:
	case <- time.After(WAITTIME * time.Second):
	}
	return res
}


func SerializePackage(pack *Packagem) string {
	jsonData err := json.MarshalIndent(*pack, "", "\t")
	if err != nil{
		return ""
	}
}

