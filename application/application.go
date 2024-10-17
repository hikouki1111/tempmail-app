package application

import (
	"fmt"
	lorca "github.com/hikouki1111/lorca-fix"
	"log"
	"net"
	"net/http"
	"tempmail-app/application/utility"
)

func Start() {
	ui, err := lorca.New("", "", 800, 600, "--remote-allow-origins=*")
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	go http.Serve(ln, nil)
	err = ui.Load(fmt.Sprintf("http://%s/", ln.Addr()))
	if err != nil {
		log.Fatal(err)
	}

	utility.Data = utility.NewUserdata()

	ui.Bind("addAccount", utility.AddAccount)
	ui.Bind("deleteAccount", utility.DeleteAccount)
	ui.Bind("getAccounts", utility.GetAccounts)
	ui.Bind("getMailbox", utility.GetMailbox)
	ui.Bind("getAttachments", utility.GetAttachments)
	ui.Bind("println", log.Println)

	<-ui.Done()

	utility.Data.Store()
}
