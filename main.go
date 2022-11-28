package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/wanderer1a/randgen/cassobj"
)

var trx_count int
var card_count int
var seeds string

func init() {
	fmt.Println("init")
	seeds = os.Getenv("CASSANDRA_SEEDS")
	rand.Seed(time.Now().UnixNano())
	ClusterInit()
	fmt.Println("init complete")
	var err error
	trx_count, err = strconv.Atoi(os.Getenv("NUMBER_TRX_TO_GENERATE"))
	card_count, err = strconv.Atoi(os.Getenv("NUMBER_CARDS_TO_GENERATE"))
	if err != nil {
		log.Fatal(err)
	}
}

func Seed() string {
	arg := ""
	return arg
}

func ClusterInit() {
	//	type arguments struct {
	//		key   string
	//		value string
	//	}

	//	a := arguments{
	//		s: Seed,
	//	}
	//	fmt.Println(a)
	cassobj.KeyspaceCreate(seeds)
	cassobj.PanTableCreate(seeds)
	cassobj.TrxTableCrate(seeds)
}

//func GenerateDictionaries() {
//	cassobj.TrxTableCrate()
//}

func PanHashGenerate(pan string) string {
	h := sha256.New()
	h.Write([]byte(pan))
	return hex.EncodeToString(h.Sum(nil))
}

var letterRunes = []rune("1234567890qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")
var digitsRunes = []rune("1234567890")

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RandInt(n int) int {
	b := rand.Intn(n)
	return b
}

func RandIntStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = digitsRunes[rand.Intn(len(digitsRunes))]
	}
	return string(b)
}

type stringn struct {
	str string
	n   int
}

func GenTrx(t int) string {
	var pan int
	var err error
	tm := rand.Intn(1000)
	time.Sleep(time.Duration(tm))
	pan, err = strconv.Atoi(RandIntStr(16))
	if err != nil {
		log.Fatal(err)
	}
	now := time.Now().UTC().String()
	card_pan_sha256 := PanHashGenerate(strconv.Itoa(pan))
	cassobj.PanInsert(t, now, pan, card_pan_sha256, seeds)

	for i := 0; i <= trx_count; i++ {
		fmt.Println(i)
		trx := transaction_record{card_pan_sha256: card_pan_sha256}
		trx.part_key = "27.11.2022"
		trx.tran_date = "27.11.2022"
		trx.pan = strconv.Itoa(pan)
		trx.acceptor_city = RandString(10)
		trx.acceptor_country = RandString(10)
		trx.acceptor_name = RandString(20)
		trx.acceptor_postal_code = RandInt(6)
		fmt.Println(trx)

	}

}

func main() {
	var wg sync.WaitGroup
	wg.Add(20)
	go func() {
		for t := 0; t <= card_count; t++ {
			go GenTrx(t)
		}
		wg.Done()
	}()
	wg.Wait()

}
