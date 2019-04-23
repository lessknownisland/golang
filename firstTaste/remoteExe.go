package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

// SSHConnect for ssh connect
func SSHConnect(user, password, host string, port int) (*ssh.Session, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User: user,
		Auth: auth,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: 30 * time.Second,
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create session
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}

	return session, nil
}

func main() {
	/* 初始化变量 */
	remoteS := make(map[string]string)
	remoteS["ip"] = "182.16.117.186"
	remoteS["port"] = "39393"
	remoteS["password"] = "BXT#8dkskdv2%%$vgAjdinsA1R%3#rFdeg"

	// fmt.Println("远程ip地址: ", remoteS["ip"])
	// fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	session, err := SSHConnect("root", remoteS["password"], remoteS["ip"], 39393)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Run("ps -ef |grep 5000")
	session.Run("ifconfig")
}
