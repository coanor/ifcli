package ifcli

import (
	"encoding/hex"
	"log"
	"testing"
	"time"
)

func TestEncrypt(t *testing.T) {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	msg := `abchahahdffsafdsafdsafdsafdsafdhshjkalfhdslafhdshafhdjslka`
	pwd := "123456"
	en := encrypt([]byte(msg), pwd)
	enhex := hex.EncodeToString(en)
	log.Printf("en: %s", enhex)

	var err error
	en, err = hex.DecodeString(enhex)
	if err != nil {
		t.Fatal(err)
	}

	x := decrypt(en, pwd)
	if string(x) == msg {
		log.Printf("decrypt ok")
	} else {
		log.Printf("decrypt failed: %s <> %s", msg, string(x))
	}
}

func TestAddConn(t *testing.T) {
	c := &Conn{
		LastConn: time.Now(),
		Created:  time.Now(),

		Host:      "https://abc.com:3234",
		User:      "lucifer",
		Password:  "123456",
		DefaultDB: "",
	}

	LoadHist()
	err := AddConn(c)
	if err != nil {
		t.Fatal(err)
	}
}
