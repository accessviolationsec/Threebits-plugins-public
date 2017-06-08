package Threebits_plugins_public

import (
	"net"
	"net/http"
	"errors"
	"strconv"
	"github.com/g-clef/Threebits/structures"
)

type HTTPBanner struct{
}


// Note: this plugin is completely ignoring the socket provided by the framework (which
// is a bit of a shame. In theory I could have it write to the socket & basically build a
// mini http client, but that seemed overkill. Forcing the golang http client to use the
// existing socket mostly worked, but blew up when sites like google tried to shift to
// http/2 . So, this just makes its own socket. Sue me.
func (h HTTPBanner) Handle(socket net.Conn, test structures.Test) (bool, string, error){
	var err error

	client := &http.Client{}
	response, err := client.Get("http://" + test.Target + ":" + strconv.Itoa(test.Port))

	if err != nil{
		return false, "", err
	}
	server := response.Header.Get("Server")
	response.Body.Close()
	if server != ""{
		return true, server, nil
	} else {
		return false, "", errors.New("Server string not found")
	}
}

func (h HTTPBanner) Protocol()(string){
	return "tcp"
}

func (h HTTPBanner) DefineArguments() {
}

func (h HTTPBanner) Initialize() (error){
    return nil
}

