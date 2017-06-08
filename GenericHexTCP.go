package Threebits_plugins_public

import (
    "encoding/hex"
    "flag"
	"net"
	"github.com/g-clef/Threebits/structures"
	"reflect"
)




type GenericHexTCP struct{
    sendArg []byte
    receiveArg []byte
}

func (s GenericHexTCP)Handle(socket net.Conn, test structures.Test) (bool, string, error) {
    var err error
    var numBytes int

    s.sendArg, err = hex.DecodeString(test.Args.Generic_TCP.Send)
    if err != nil {
        return false, "error reading config", err
    }
    s.receiveArg, err = hex.DecodeString(test.Args.Generic_TCP.Receive)
    if err != nil {
        return false, "error reading config", err
    }

	_, err = socket.Write(s.sendArg)
	if err != nil{
		return false, "error writing to socket", err
	}
	var response =make([]byte, 1024)
	numBytes, err = socket.Read(response)
	if err != nil{
		return false, "error reading response", err
	}
	if numBytes != len(s.receiveArg){
		return false, "response different size than test response string", nil
	}
	if reflect.DeepEqual(response, s.receiveArg) {
	    return true, "", nil
	} else {
	    return false, "handshake did not match", nil
	}
}


func (s GenericHexTCP) Protocol()(string){
	return "tcp"
}


func (s GenericHexTCP) DefineArguments() {
    
    flag.StringVar(&structures.AllArgs.Generic_TCP.Send, "hexsend", "", "(for generic hex plugin)Hex string to send to target")
    flag.StringVar(&structures.AllArgs.Generic_TCP.Receive, "hexreceive", "", "(for generic hex plugin)Hex string to check responses against.")
}


