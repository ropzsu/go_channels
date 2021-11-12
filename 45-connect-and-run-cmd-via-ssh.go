package main

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"os"
)

func connectViaSsh(user, host string, password string) (*ssh.Client, *ssh.Session) {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.KeyboardInteractive(SshInteractive),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", host, config)
	fmt.Println(err)
	session, err := client.NewSession()
	fmt.Println(err)

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
	var b bytes.Buffer
	if len(os.Args) < 3 {
		fmt.Println("Must SET environ : USER_ID  + TOKEN")
		fmt.Printf("Usage: %s  <HOST> <CMD>\n", os.Args[0])
		return
	}

	user := os.Getenv("USER_ID")
	host := os.Args[1]
	cmd := os.Args[2]

	fmt.Println(">>> Step 1: Connect HOST: ", host)
	client, session := connectViaSsh(user, host+":22", "Password:")

	session.Stdout = &b
	session.Run(cmd)
	fmt.Println(">>> Step 2: Print Result: ")
	fmt.Println(b.String())
	client.Close()
}
