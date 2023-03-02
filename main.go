package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("vim-go")

	// init config
	//
	var cfgFile string
	flag.StringVar(&cfgFile, "cfg", "app.cfg", "config filepath for this service")

	f, err := os.Open(cfgFile)
	must(err)
	cfg := NewConfig(f)
	f.Close()

	// init database
	//
	db, err := NewMySQL(cfg.MySQL)
	must(err)
	defer db.Close()

	// init server
	//
	svc := Service{db}
	h := &Handler{svc}
	// route
	http.Handle("/piglatins", h)
	// start
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Printf("Server exited with error: %v", err)
	}
}

func must(e error) {
	if e != nil {
		log.Fatalf("Start server failed: %v", e)
	}
}

// Config represents the entire config for this app.
type Config struct {
	Server ServerConfig `json:"server"`
	MySQL  MySQLConfig  `json:"mysql"`
}

type ServerConfig struct {
	Port int `json:"port"`
}

// NewConfig parses text to a Config.
func NewConfig(r io.Reader) Config {
	var cfg Config
	d := json.NewDecoder(r)
	must(d.Decode(&cfg))
	return cfg
}
