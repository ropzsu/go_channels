
package main
import (
    "bytes"
    "golang.org/x/crypto/ssh"
    "fmt"
    "os"
)

func connectViaSsh(user, host string, password string) (*ssh.Client, *ssh.Session) {
     config := &ssh.ClientConfig{
        User: "brianguo",
        Auth: []ssh.AuthMethod{
             ssh.Password( password ),
             ssh.KeyboardInteractive(SshInteractive),
        },
         HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }

    client, err := ssh.Dial("tcp", host, config)
     
    if err != nil {
       fmt.Println("\n\n== ssh [Dia host] failed. err = ", err)
       os.Exit(-1)
    }

    session, err := client.NewSession()
     if err != nil {
       fmt.Println("== ssh [NewSession] failed. err = ", err)
       os.Exit(-1)
    }

    return client, session
}


func SshInteractive(user, instruction string, questions []string, echos []bool) (answers []string, err error) {
    answers = make([]string, len(questions))
    // The second parameter is unused
    for n, _ := range questions {
        answers[n] = os.Getenv("TOKEN")
    }

    return answers, nil
}

 

func main() {

    if (len(os.Args) < 3 ) {
        fmt.Println("Must SET environ : USER_ID  + TOKEN")
        fmt.Printf("Usage: %s  <HOST> <CMD>\n", os.Args[0])
        return 
    }

    user := os.Getenv("USER_ID")
    passwd := os.Getenv("TOKEN")
    host := os.Args[1]
    cmd := os.Args[2]


    fmt.Print(">>>  Connect HOST: ", host, ", msg = ")
    client, session := connectViaSsh( user , host + ":36000" , passwd)

    var b bytes.Buffer   
    session.Stdout = &b
    session.Run(cmd)

    fmt.Println(b.String())
    client.Close()
}


