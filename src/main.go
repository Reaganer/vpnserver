package main

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"text/template"
	"time"

	"./sha"
	"golang.org/x/crypto/md4"
)

type params struct {
	DDnsClientKey         string //hashAdminPassword(random)
	ServerHashedPassword  string //hashAdminPassword(random)
	HubHashedPassword     string //hashAdminPassword(hub_random)
	CreatedTime           int64
	HubSecurePassword     string //hashPassword("administrator", hub_random)
	VirtualHostMacAddress string //00-AC-F8-1C-11-FD

	UserName               string
	UserAuthNtLmSecureHash string //generateNtPasswordHash(pass)
	UserAuthPassword       string //hashPassword(UserName, pass)
}

func hashPassword(name, pass string) string {
	sha := sha.New()
	sha.Write([]byte(pass))
	sha.Write([]byte(strings.ToUpper(name)))
	return base64.StdEncoding.EncodeToString(sha.Sum(nil))
}

func generateNtPasswordHash(pass string) string {
	md4 := md4.New()
	var buf = make([]byte, len(pass)*2)
	for i, c := range pass {
		buf[i*2] = byte(c)
	}
	md4.Write(buf)
	return base64.StdEncoding.EncodeToString(md4.Sum(nil))
}

func hashAdminPassword(pass string) string {
	sha := sha.New()
	sha.Write([]byte(pass))
	return base64.StdEncoding.EncodeToString(sha.Sum(nil))
}

func generateConfigFile(p *params) {
	t := template.Must(template.New("config").Parse(configTemplate))
	out, err := os.Create("./vpn_server.config")
	if err != nil {
		fmt.Printf("[ERR] Create config file: %v\n", err)
		return
	}
	defer out.Close()
	t.Execute(out, p)
}

func randString(n int) string {
	var buf []byte = make([]byte, n)
	rand.Read(buf)
	for i, b := range buf {
		buf[i] = ' ' + (b % ('~' - ' '))
	}
	return string(buf)
}

func main() {
	var p params
	var mac []byte = []byte{5: 0}
	var pass string
	p.CreatedTime = time.Now().Unix()
	rand.Seed(p.CreatedTime)
	pass = randString(8)
	p.DDnsClientKey = hashAdminPassword(pass)

	pass = randString(16)
	p.ServerHashedPassword = hashAdminPassword(pass)

	pass = randString(12)
	p.HubHashedPassword = hashAdminPassword(pass)
	p.HubSecurePassword = hashPassword("administrator", pass)

	rand.Read(mac)
	mac[0] &= 0xFE
	p.VirtualHostMacAddress = fmt.Sprintf("%02X-%02X-%02X-%02X-%02X-%02X", mac[0], mac[1], mac[2], mac[3], mac[4], mac[5])

	var user = os.Getenv("USRNAME")
	pass = os.Getenv("PASSWORD")

	if len(user) == 0 {
		user = "alice"
	}
	if len(pass) == 0 {
		pass = randString(8)
		fmt.Printf("[INFO] Random generated password is : %s\n", pass)
	}

	p.UserName = user
	p.UserAuthPassword = hashPassword(user, pass)
	p.UserAuthNtLmSecureHash = generateNtPasswordHash(pass)

	generateConfigFile(&p)
}
