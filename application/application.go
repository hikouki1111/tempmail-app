package application

import (
    "fmt"
    lorca "github.com/hikouki1111/lorca-fix"
    tempmail "github.com/hikouki1111/tempmail-wrapper"
    "log"
    "net"
    "net/http"
)

var (
    accounts []tempmail.Account
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

    ui.Bind("addAccount", func() {
        acc, err := tempmail.NewAccount()
        if err != nil {
            log.Println(err)
        }
        accounts = append(accounts, *acc)
    })

    ui.Bind("getAccounts", func() []tempmail.Account {
        return accounts
    })

    <-ui.Done()
}
