package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	// "io/ioutil"
	"log"
	"net"
	"strings"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

// code take from the https://pkg.go.dev/golang.org/x/crypto/ssh#example-NewServerConn
func main() {
	// An SSH server is represented by a ServerConfig, which holds
	// certificate details and handles authentication of ServerConns.
	config := &ssh.ServerConfig{
		// Remove to disable password auth.
		PasswordCallback: func(c ssh.ConnMetadata, pass []byte) (*ssh.Permissions, error) {
			// Should use constant-time compare (or better, salt+hash) in
			// a production setting.
			if c.User() == "testuser" && string(pass) == "tiger" {
				return nil, nil
			}
			return nil, fmt.Errorf("password rejected for %q", c.User())
		},
	}

	// privateBytes, err := ioutil.ReadFile("/home/prem/.ssh/id_rsa")
	// if err != nil {
	// 	log.Fatal("Failed to load private key: ", err)
	// }
	privateKey, _, err := GenerateKey(2048)
	if err != nil {
		log.Fatal("failed to generate key")
	}
	privateBytes := EncodePrivateKey(privateKey)
	private, err := ssh.ParsePrivateKey(privateBytes)
	if err != nil {
		log.Fatal("Failed to parse private key: ", err)
	}

	config.AddHostKey(private)

	// Once a ServerConfig has been configured, connections can be
	// accepted.
	// listener, err := net.Listen("tcp", "0.0.0.0:2022")
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		log.Fatal("failed to listen for connection: ", err)
	}
	log.Printf("listener started on %s\n", listener.Addr().String())
	nConn, err := listener.Accept()
	if err != nil {
		log.Fatal("failed to accept incoming connection: ", err)
	}

	// Before use, a handshake must be performed on the incoming
	// net.Conn.
	_, chans, reqs, err := ssh.NewServerConn(nConn, config)
	if err != nil {
		log.Fatal("failed to handshake: ", err)
	}
	// log.Printf("logged in with key %s", conn.Permissions.Extensions["pubkey-fp"])

	// The incoming Request channel must be serviced.
	go ssh.DiscardRequests(reqs)

	// Service the incoming Channel channel.
	for newChannel := range chans {
		// Channels have a type, depending on the application level
		// protocol intended. In the case of a shell, the type is
		// "session" and ServerShell may be used to present a simple
		// terminal interface.
		if newChannel.ChannelType() != "session" {
			newChannel.Reject(ssh.UnknownChannelType, "unknown channel type")
			continue
		}
		channel, requests, err := newChannel.Accept()
		if err != nil {
			log.Fatalf("Could not accept channel: %v", err)
		}

		// Sessions have out-of-band requests such as "shell",
		// "pty-req" and "env".  Here we handle only the
		// "shell" request.
		go func(in <-chan *ssh.Request) {
			for req := range in {
				req.Reply(req.Type == "shell" || req.Type == "pty-req", nil)
			}
		}(requests)

		term := terminal.NewTerminal(channel, "> ")
		// term.SetPrompt(">>")

		go func() {
			defer channel.Close()
		out:
			for {
				line, err := term.ReadLine()
				if err != nil {
					break
				}
				line = strings.Trim(line, "")

				// fmt.Fprintf(channel, "%s", line)

				switch line {
				case "info":
					term.Write([]byte(fmt.Sprintf("this is terminal emulator\n")))
				case "quit":
					term.Write([]byte(fmt.Sprintf("quitting now ...\n")))
					break out
				}

			}
		}()
	}
}

func GenerateKey(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	private, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	return private, &private.PublicKey, nil
}

func EncodePrivateKey(private *rsa.PrivateKey) []byte {
	return pem.EncodeToMemory(&pem.Block{
		Bytes: x509.MarshalPKCS1PrivateKey(private),
		Type:  "RSA PRIVATE KEY",
	})
}
