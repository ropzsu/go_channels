package main

import (
    "fmt"
    "golang.org/x/crypto/ssh"
    "log"
    "os"
)

func main() {

    if (len(os.Args) < 4 ) {
     fmt.Println("Usage: ./$0  <username> <host-ip> <cmd> \n  Before use, must set TOKEN env")
     return  
   }
    user := os.Args[1] 
    host := os.Args[2] 
    cmd := os.Args[3] 
    passwd :=  os.Getenv("TOKEN") 

    // 建立SSH客户端连接
    client, err := ssh.Dial("tcp", host + ":36000", &ssh.ClientConfig{
        User:            user,
        Auth:            []ssh.AuthMethod{ssh.Password(passwd)},
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    })
    if err != nil {
        log.Fatalf("SSH dial error: %s", err.Error())
    }

    // 建立新会话
    session, err := client.NewSession()
    defer session.Close()
    if err != nil {
        log.Fatalf("new session error: %s", err.Error())
    }

    result, err := session.Output(cmd)
    if err != nil {
        fmt.Fprintf(os.Stdout, "Failed to run command, Err:%s", err.Error())
        os.Exit(0)
    }
    fmt.Println(string(result))
}

