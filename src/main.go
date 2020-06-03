package main

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"

	"./sha"
	"golang.org/x/crypto/md4"
)

type user struct {
	UserName               string
	UserAuthNtLmSecureHash string //generateNtPasswordHash(pass)
	UserAuthPassword       string //hashPassword(UserName, pass)
}

type params struct {
	DDnsClientKey        string //hashAdminPassword(random)
	ServerHashedPassword string //hashAdminPassword(random)
	HubHashedPassword    string //hashAdminPassword(hub_random)
	CreatedTime          int64
	HubSecurePassword    string //hashPassword("administrator", hub_random)

	UserList []user

	IsRelayNode      bool
	NextNodeName     string
	NextNodeHost     string
	NextNodePort     uint64
	NextNodeUserName string
	NextNodePassword string

	Subnet                string
	VirtualHostMacAddress string
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
func randLowerLetter(n int) string {
	var buf []byte = make([]byte, n)
	rand.Read(buf)
	for i, b := range buf {
		buf[i] = 'a' + (b % ('z' - 'a'))
	}
	return string(buf)
}
func splitBySemicolon(c rune) bool {
	return c == ';'
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
	p.Subnet = fmt.Sprintf("192.168.%d", (rand.Uint32() % 256))

	var userList = strings.FieldsFunc(os.Getenv("UserList"), splitBySemicolon)

	_, p.IsRelayNode = os.LookupEnv("RelayNode")
	if p.IsRelayNode {
		p.NextNodeHost = os.Getenv("NextNodeHost")
		p.NextNodePort, _ = strconv.ParseUint(os.Getenv("NextNodePort"), 10, 16)
		p.NextNodeUserName = os.Getenv("NextNodeUserName")
		p.NextNodePassword = hashPassword(p.NextNodeUserName, os.Getenv("NextNodePassword"))
		p.NextNodeName = randLowerLetter(6)
	}

	if len(userList) == 0 {
		p.UserList = make([]user, 1)
		p.UserList[0].UserName = randLowerLetter(6)
		pass = randString(8)
		p.UserList[0].UserAuthPassword = hashPassword(p.UserList[0].UserName, pass)
		p.UserList[0].UserAuthNtLmSecureHash = generateNtPasswordHash(pass)
		fmt.Printf("[INFO] Random generated user is : %s\nPassword is %s\n", p.UserList[0].UserName, pass)
	} else {
		p.UserList = make([]user, len(userList))
		for i, u := range userList {
			var temp = strings.SplitN(u, ":", 2)
			if len(temp) == 1 {
				p.UserList[i].UserName = temp[0]
				pass = randString(8)
				p.UserList[i].UserAuthPassword = hashPassword(temp[0], pass)
				p.UserList[i].UserAuthNtLmSecureHash = generateNtPasswordHash(pass)
				fmt.Printf("[INFO] Random generated password for user '%s' is: %s\n", temp[0], pass)
			} else {
				p.UserList[i].UserName = temp[0]
				p.UserList[i].UserAuthPassword = hashPassword(temp[0], temp[1])
				p.UserList[i].UserAuthNtLmSecureHash = generateNtPasswordHash(temp[1])
			}
		}
	}

	generateConfigFile(&p)
}
