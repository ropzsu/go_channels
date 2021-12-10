package main

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"os"
)

type Message struct {
	Host   string
	Result string
	Err    error
}

// Get username, password by environment variable.
func connect_ssh_cmd(host, port, command string, ch chan Message) {
	user := os.Getenv("USER_ID")
	password := os.Getenv("TOKEN")

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
			ssh.KeyboardInteractive(SshInteractive),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	addr := host + ":" + port

	client, err := ssh.Dial("tcp", addr, config)
	dial_failed := 0

	if err != nil {
		// fmt.Println("\n\n== ssh [Dia host] failed. err = ", err)
		// set notify. will try again.
		dial_failed = 1
	}

	// try dia again
	if dial_failed == 1 {
		client, err = ssh.Dial("tcp", addr, config)
		if err != nil {
			ch <- Message{host, fmt.Sprintf("== ssh [Dia host] final failed. err = %s", err), err}
			return
		}
	}

	session, err := client.NewSession()
	if err != nil {
		ch <- Message{host, fmt.Sprintf("== ssh [NewSession] failed. err = %s", err), err}
		return
	}

	// run command and get results.
	var b bytes.Buffer
	session.Stdout = &b
	session.Run(command)

	msg := fmt.Sprintf(b.String())
	client.Close()

	ch <- Message{host, msg, nil}
	return
}

// Get password by environment variable.
func SshInteractive(user, instruction string, questions []string, echos []bool) (answers []string, err error) {
	answers = make([]string, len(questions))
	// The second parameter is unused
	for n, _ := range questions {
		answers[n] = os.Getenv("TOKEN")
	}
	return answers, nil
}

func batch_ssh_run(cmd string, port string, ips []string) chan Message {
	ch := make(chan Message)

	for i := 0; i < len(ips); i++ {
		h := ips[i]
		go func() {
			connect_ssh_cmd(h, port, cmd, ch)
		}()
	}
	return ch
}

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Must SET environ : USER_ID  + TOKEN")
		fmt.Printf("Usage: %s <CMD>  <HOST1> <HOST2> <HOST3> ...\n", os.Args[0])
		return
	}

	cmd := os.Args[1]
	ips := os.Args[2:]
	port := "36000"

	jobs := batch_ssh_run(cmd, port, ips)

	for i := 0; i < len(ips); i++ {
		m := <-jobs
		fmt.Println(">> Host: ", m.Host, " msg = ", m.Result)
	}

}
