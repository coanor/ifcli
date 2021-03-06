package ifcli

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path"
	"time"

	"github.com/BurntSushi/toml"
	_ "github.com/influxdata/influxdb1-client" // this is important because of the bug in go mod
	client "github.com/influxdata/influxdb1-client/v2"
)

var (
	curConnections       = map[string]*Conn{}
	curConn        *Conn = nil

	ifclirc = ""
)

func SetIfCliRC(s string) error {
	if _, err := os.Stat(s); err != nil {
		return err
	}

	ifclirc = s
	fmt.Printf("use %s as ifclirc\n", ifclirc)
	return nil
}

func LoadHist() {

	if ifclirc == "" {
		u, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}

		ifclirc = path.Join(u.HomeDir, `.ifclirc`)
		fmt.Printf("use %s as ifclirc\n", ifclirc)
	}

	data, err := ioutil.ReadFile(ifclirc)
	if err != nil {
		return // ignore
	}

	histConn := map[string]*Conn{}
	if err := toml.Unmarshal(data, &histConn); err != nil {
		log.Fatal(err)
	}

	for k, v := range histConn {

		enPwd, err := hex.DecodeString(v.PasswordEncrypted)
		if err != nil {
			fmt.Printf("[error] invalid password on %s: %s, ignored\n", k, err.Error())
			continue
		}

		pwd, err := decrypt(enPwd, pwd)
		if err != nil {
			fmt.Printf("invalid InfluxDB password: %s, ignored\n", err.Error())
			continue
		}

		v.Password = string(pwd)

		if err := AddConn(v); err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("load connection %s\n", k)
		}
	}
}

type Conn struct {
	LastConn time.Time `toml:"last_conn"`
	LastExit time.Time `toml:"last_exit"`
	Created  time.Time `toml:"created"`

	Host              string `toml:"host"`
	User              string `toml:"user"`
	Password          string `toml:"-"`
	PasswordEncrypted string `toml:"password"`
	DefaultDB         string `toml:"default_db"`
	Prompt            string `toml:"prompt"`

	curDB string        `toml:"-"`
	cli   client.Client `toml:"-"`
}

func (c *Conn) Connect() error {

	cli, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:      c.Host,
		Username:  c.User,
		Password:  c.Password,
		UserAgent: "ifcli",
	})

	if err != nil {
		return err
	}

	c.cli = cli

	if c.DefaultDB != "" && c.curDB == "" {
		c.curDB = c.DefaultDB
	}

	c.LastConn = time.Now()

	if err := AddConn(c); err != nil {
		return err
	}

	// update current working connection
	curConn = c

	fmt.Printf("conect to %s ok\n", c.Key())
	return nil
}

func (c *Conn) Close() error {
	c.LastExit = time.Now()
	if err := AddConn(c); err != nil {
		return err
	}

	if err := c.cli.Close(); err != nil {
		return err
	}

	AddConn(c)

	return nil
}

func (c *Conn) Key() string {
	return c.User + "::" + c.Prompt
}

func StatQuery(t string, n int64) {

	if curConn == nil {
		fmt.Printf("not connected :(\n")
		return
	}

	q := client.NewQuery(t, curConn.curDB, ``)
	start := time.Now()
	var maxCost, totalCost time.Duration

	for i := 0; i < int(n); i++ {
		start_ := time.Now()
		if resp, err := curConn.cli.Query(q); err == nil && resp.Error() == nil {
			cost := time.Since(start_)
			if cost > maxCost {
				maxCost = cost
			}
		} else {
			if err == nil {
				fmt.Printf("[error] resp Err: %s\n", resp.Error())
			} else {
				fmt.Printf("[error] %s, resp Err: %s\n", err.Error(), resp.Error())
			}
			return
		}
	}

	totalCost = time.Since(start)

	fmt.Printf("time: total: %v, avg: %v, max: %v\n", totalCost, totalCost/time.Duration(n), maxCost)
}

func DoQuery(t string) {

	if curConn == nil {
		fmt.Printf("not connected :(\n")
		return
	}

	q := client.NewQuery(t, curConn.curDB, ``)
	start := time.Now()

	if resp, err := curConn.cli.Query(q); err == nil && resp.Error() == nil {
		n := ShowResp(resp)
		fmt.Printf("\n%d rows in set\n", n)
		fmt.Printf("time: %v\n", time.Since(start))
	} else {
		if err == nil {
			fmt.Printf("[error] resp Err: %s\n", resp.Error())
		} else {
			fmt.Printf("[error] %s, resp Err: %s\n", err.Error(), resp.Error())
		}
	}
}

// for exists connection, update it, or add it
func AddConn(c *Conn) error {

	var err error

	enPwd, err := encrypt([]byte(c.Password), pwd)
	if err != nil {
		return err
	}

	c.PasswordEncrypted = hex.EncodeToString(enPwd)

	if c.Created.Unix() == 0 {
		c.Created = time.Now()
	}

	// add new/update exits connectionss
	curConnections[c.Key()] = c

	AddSug(`CONN ` + c.Key()) // add suggestion

	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(curConnections); err != nil {
		return err
	}

	f, err := os.OpenFile(ifclirc, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	defer f.Close()
	if _, err := f.Write(buf.Bytes()); err != nil {
		return err
	}

	return nil
}
