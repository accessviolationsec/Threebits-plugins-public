package Threebits_plugins_public

import (
	"net"
	"github.com/g-clef/Threebits/structures"
	"reflect"
	"encoding/hex"
	"bytes"
	"flag"
)




type GenericHexTCP struct{
    sendArg []byte
    receiveArg []byte
    receiveBeginArg []byte
}

func (s GenericHexTCP)Handle(socket net.Conn, test structures.Test) (bool, string, error) {
    var err error
    var numBytes int

    s.sendArg, err = hex.DecodeString(test.Args.Generic_TCP.Send)
    if err != nil {
        return false, "error reading config", err
    }
    if test.Args.Generic_TCP.Receive != "" {
        s.receiveArg, err = hex.DecodeString(test.Args.Generic_TCP.Receive)
        if err != nil {
            return false, "error reading config", err
        }
    }
    if test.Args.Generic_TCP.ReceiveStart != "" {
            s.receiveBeginArg, err = hex.DecodeString(test.Args.Generic_TCP.ReceiveStart)
        if err != nil {
            return false, "error reading config", err
        }
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
	if s.receiveArg != nil {
	    if numBytes != len(s.receiveArg){
		    return false, "response different size than test response string", nil
	    }
	    if reflect.DeepEqual(response, s.receiveArg) {
	        return true, "", nil
	    } else {
	        return false, "handshake did not match", nil
	    }
	} else {
	    if numBytes < len(s.receiveBeginArg) {
	        return false, "response shorter than test comparison string", nil
	    }
	    if bytes.HasPrefix(response, s.receiveBeginArg){
	        return true, "", nil
	    } else {
	        return false, "handshake did not match", nil
	    }
	}
	    
}


func (s GenericHexTCP) Protocol()(string){
	return "tcp"
}


func (s GenericHexTCP) DefineArguments() {
    
    flag.StringVar(&structures.AllArgs.Generic_TCP.Send, "hexsend", "", "(for generic hex plugin)Hex string to send to target")
    flag.StringVar(&structures.AllArgs.Generic_TCP.Receive, "hexreceive", "", "(for generic hex plugin)Hex string to check responses against.")
    flag.StringVar(&structures.AllArgs.Generic_TCP.ReceiveStart, "hexreceiveatstart", "", "(for generic hex plugin)Hex string to check just the beginning of the response against.")
}


