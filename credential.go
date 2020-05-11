package main

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

const credentialFile = "./credential.toml"

/*
[ID]

app = A0136PQFA6A
client = 3465148148.1108806520214

[Secret]

client = d8ffd2b8affd777215e45263e2cd078c
Signing = dc6f1c7aefaddb92fc00ed0d58b8a3f0

[Token]

verification = pcoFWuRnnWoTebl2FsUbISwM
*/
type credentialStruct struct {
	ID     credentialID
	Secret credentialSecret
	Token  credentialToken
}

type credentialID struct {
	App    string
	Client string
}

type credentialSecret struct {
	Client  string
	Signing string
}

type credentialToken struct {
	Verification string
}

func parseCredentials() *credentialStruct {
	f, err := os.Open(credentialFile)
	if err != nil {
		log.Fatalln(err)
	}

	var credential credentialStruct
	metaData, err := toml.DecodeReader(f, &credential)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(metaData)

	return &credential
}
