package Threebits_plugins_public

import (
	"net"
	"github.com/accessviolationsec/Threebits/structures"
	"bufio"
)


type SSHBanner struct{}

func (s SSHBanner)Handle(socket net.Conn, test structures.Test) (bool, string, error) {
	connbuf := bufio.NewReader(socket)
	version, err := connbuf.ReadString('\n')
	if err != nil {
		return false, "", err
	}
	return true, version, nil
}


func (s SSHBanner) Protocol()(string){
	return "tcp"
}