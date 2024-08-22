// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

package main

import (
	"caga/websocket/random"
	"flag"
	"log"
	"math"
	"net/http"

	"math/big"

	"github.com/gorilla/websocket"
)

// We can set a higher number for big.Int, but in this code, we use MaxUint64 for practice
// To increase speed and save memory
var maxNumber = new(big.Int).SetUint64(math.MaxUint64)

// To be able to generate a non-repeating random number, we need to store the used numbers, we can use third parties like redis.
// But in this test, we will store it in a map for faster access, which may take up a lot of memory - DON't set maxNumber too big
var generated = map[string]bool{}

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options

func echo(w http.ResponseWriter, r *http.Request) {
	// accept any connection from any origin
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Upgrade error: ", err)
		return
	}
	defer c.Close()
	for {
		mt, _, err := c.ReadMessage()
		if err != nil {
			log.Println("Failed to read message:", err)
			break
		}

		num, err := random.Random(generated, maxNumber)
		if err != nil {
			log.Println("Failed to generate new number:", err)
			break
		}

		log.Printf("Generate new number %s \n", num.String())
		err = c.WriteMessage(mt, []byte(num.String()))
		if err != nil {
			log.Println("Failed to send message:", err)
			break
		}
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/", echo)

	log.Printf("Ws server is listening at %s \n", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
