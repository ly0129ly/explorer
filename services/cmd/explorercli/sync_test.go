package main

import (
	"testing"
	"github.com/cosmos/cosmos-sdk/client"
	"time"
	"log"
	"github.com/ly0129ly/explorer/services/modules/db"
)

func TestProcessSync(t *testing.T) {
	c := client.GetNode("tcp://47.104.155.125:46757")
	db.Mgo.Init("localhost:27017")
	processSync(c)
	processWatch(c)

	log.Printf(" finish %s","ok")
	time.Sleep(35 * time.Minute)
}