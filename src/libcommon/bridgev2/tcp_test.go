package bridgev2

import (
	"app"
	"fmt"
	json "github.com/json-iterator/go"
	"libservicev2"
	"testing"
	"util/db"
	"util/logger"
)

func init() {
	logger.SetLogLevel(1)
	app.Secret = "123456"
	app.BasePath = "E:\\godfs-storage\\storage1"
	libservicev2.SetPool(db.NewPool(1))
}

func PrintResult(result ...interface{}) {
	fmt.Println("\n\n+++~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~+++")
	if result != nil {
		for i := range result {
			obj := result[i]
			bs, _ := json.Marshal(obj)
			fmt.Println(string(bs))
		}
	}
	fmt.Println("+++~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~+++")
}

func TestServer(t *testing.T) {
	app.UUID = "tracker01"
	server := NewServer("", 1022)
	server.Listen()
}

func TestClient(t *testing.T) {
	app.UUID = "storage01"
	storage := &app.StorageDO{
		Host:          "localhost",
		Port:          1022,
		AdvertiseAddr: "192.168.1.142",
		AdvertisePort: 1022,
		AccessFlag:    app.AccessFlagNone,
	}
	server := &app.ServerInfo{}
	server.FromStorage(storage)
	index := 0
	for {
		client := NewTcpClient(server)
		if err := client.Connect(); err != nil {
			panic(err)
		}
		index++
		PrintResult(client.Validate())
		client.GetConnManager().Close()
		fmt.Println(index)
	}

}
