package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
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
		Timeout: 5 * time.Second,
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

type remoteS struct {
	ip       string
	port     int
	password string
	// info     struct {
	// 	id     int
	// 	ppp    string
	// 	status bool
	// }
}

func main() {
	/* 初始化变量 */
	remoteDic := remoteS{}
	remoteDic.ip = "182.16.117.186"
	remoteDic.port = 39393
	remoteDic.password = "BXT#8dkskdv2%%vgAjdinsA1R%3#rFdeg"
	// remoteDic.info.id = 1
	// remoteDic.info.ppp = "asdf"
	// var port int = 39393
	var cmdlist = [...]string{"ps -ef |grep 5000", "ifconfig", "exit"}

	// fmt.Println("远程ip地址: ", remoteS["ip"])
	// fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	session, err := SSHConnect("root", remoteDic.password, remoteDic.ip, remoteDic.port)
	if err != nil {
		// fmt.Println("test0!")
		// log.Println("test!")
		log.Fatal(err)
	}
	defer session.Close()

	// 定义命令的输入管道
	stdinBuf, err := session.StdinPipe()
	if err != nil {
		log.Fatal(err)
		return
	}

	// 获取 Shell 上的输出
	var outbt, errbt bytes.Buffer
	session.Stdout = &outbt
	session.Stderr = &errbt
	err = session.Shell()
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, c := range cmdlist {
		c = c + "\n"
		stdinBuf.Write([]byte(c))
	}
	session.Wait()
	log.Fatalln((outbt.String() + errbt.String()))

	// log.Fatal("exit.")
	// session.Stdout = os.Stdout
	// session.Stderr = os.Stderr
	// session.Run("ps -ef |grep 5000")
	// session.Run("ifconfig")
}
