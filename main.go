package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
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

type transaction_record struct {
	part_key                  string
	proc_date                 string
	tran_date                 string
	tran_datetime             string
	tran_hour_num             int
	tran_day_of_week          int
	tran_day_of_month         int
	tran_day_of_year          int
	local_datetime            string
	proc_code                 int
	mti                       int
	card_length               int
	card_exp_date             string
	card_product              string
	card_pan_sha256           string
	tran_any_first_date       string
	tran_purchase_first_date  string
	tran_financial_first_date string
	tran_tokenized_first_date string
	tran_amount               int
	tran_currency             int
	analyze_amount            int
	pem                       int
	mcc                       int
	eci                       int
	resp_code                 string
	token_requestor           int
	terminal_id               string
	merchant_id               string
	acceptor_name             string
	acceptor_city             string
	acceptor_country          string
	acceptor_postal_code      int
	mcc_set_one_rus           string
	mcc_set_two_rus           string
	mcc_set_three_rus         string
	in_gate_postbox           int
	out_gate_postbox          int
	benefit_code              int
	is_contactless            int
	is_token                  int
	is_zero_amount            int
	is_tran_currency_non_rub  int
	is_approved               int
	token_requestor_code      int
	token_requestor_name      string
	tran_type                 string
	is_funding                int
	is_payment                int
	is_3ds                    int
	is_3ds_attempted          int
	is_3ds_authenticated      int
	is_3ds_not_authenticated  int
	is_declined               int
	is_suspicious_decline     int
	is_pickup_decline         int
	is_bad_pin_decline        int
	is_oom_decline            int
	is_tech_decline           int
	pan                       string
}

func main() {
	println(RandInt(10))
	println(RandString(10))

	//	tr := transaction_record{}
	now := time.Now().UTC().UnixNano()
	fmt.Println("Current datetime:", now)
	fmt.Println("Random Int length of 5:", RandInt(5))

	for t := 0; t <= card_count; t++ {
		var pan int
		var err error
		pan, err = strconv.Atoi(RandIntStr(16))
		if err != nil {
			log.Fatal(err)
		}

		card_pan_sha256 := PanHashGenerate(strconv.Itoa(pan))
		cassobj.PanInsert(t, pan, card_pan_sha256, seeds)

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
}
