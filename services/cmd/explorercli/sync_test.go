package main

import (
	"testing"
	"github.com/cosmos/cosmos-sdk/client"
	"time"
	"log"
	"github.com/ly0129ly/explorer/services/modules/db"
	"github.com/cosmos/cosmos-sdk"
)

func TestProcessSync(t *testing.T) {
	c := client.GetNode("tcp://47.104.155.125:46757")
	db.Mgo.Init("localhost:27017")
	//processSync(c)
	processWatch(c)

	log.Printf(" finish %s","ok")
	time.Sleep(35 * time.Minute)
}

func TestLoadTx(t *testing.T){
	bin := []byte{22,3,1,5,112,97,110,103,117,0,0,0,0,0,0,0,0,105,0,0,0,31,1,1,0,1,4,115,105,103,115,1,20,246,7,46,65,200,94,227,39,181,172,33,215,214,241,93,28,78,86,242,64,32,1,1,0,1,4,115,105,103,115,1,20,246,7,46,65,200,94,227,39,181,172,33,215,214,241,93,28,78,86,242,64,1,1,1,4,105,114,105,115,0,0,0,0,0,0,0}
	txb, _ := sdk.LoadTx(bin)
	log.Printf(" finish %s",txb)
}