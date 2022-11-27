package cassobj

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

func SessionInit(seeds string) *gocql.Session {
	cluster := gocql.NewCluster(seeds)
	cluster.WriteTimeout = 10000
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	return session
}

func TrxInsert(trx []string) {

}

func PanTableCreate(seeds string) {
	session := SessionInit(seeds)
	fmt.Print("If not exists creating table 'card_pan_sha256'...")
	if err := session.Query(`CREATE TABLE IF NOT EXISTS generator.card_pan_sha256
								(
									number	                  int,
									card_pan	              text,
									card_pan_sha256           text,
								PRIMARY KEY (number)
								) `, "generator").Exec(); err != nil {
		fmt.Println("Error creating table")
		log.Fatal(err)
	}
	fmt.Print("ok")
	fmt.Println()

	session.Close()
}

func TrxTableCrate(seeds string) {
	session := SessionInit(seeds)
	fmt.Print("If not exists creating table 'trx'...")
	if err := session.Query(`CREATE TABLE IF NOT EXISTS generator.trx
								(
									part_key                  text,
									proc_date                 text,
									tran_date                 text,
									tran_datetime             text,
									tran_hour_num             int,
									tran_day_of_week          int,
									tran_day_of_month         int,
									tran_day_of_year          int,
									local_datetime            text,
									proc_code                 int,
									mti                       int,
									card_length               int,
									card_exp_date             text,
									card_product              text,
									card_pan_sha256           text,
									tran_any_first_date       text,
									tran_purchase_first_date  text,
									tran_financial_first_date text,
									tran_tokenized_first_date text,
									tran_amount               int,
									tran_currency             int,
									analyze_amount            int,
									pem                       int,
									mcc                       int,
									eci                       int,
									resp_code                 text,
									token_requestor           int,
									terminal_id               text,
									merchant_id               text,
									acceptor_name             text,
									acceptor_city             text,
									acceptor_country          text,
									acceptor_postal_code      int,
									mcc_set_one_rus           text,
									mcc_set_two_rus           text,
									mcc_set_three_rus         text,
									in_gate_postbox           int,
									out_gate_postbox          int,
									benefit_code              int,
									is_contactless            int,
									is_token                  int,
									is_zero_amount            int,
									is_tran_currency_non_rub  int,
									is_approved               int,
									token_requestor_code      int,
									token_requestor_name      text,
									tran_type                 text,
									is_funding                int,
									is_payment                int,
									is_3ds                    int,
									is_3ds_attempted          int,
									is_3ds_authenticated      int,
									is_3ds_not_authenticated  int,
									is_declined               int,
									is_suspicious_decline     int,
									is_pickup_decline         int,
									is_bad_pin_decline        int,
									is_oom_decline            int,
									is_tech_decline           int,
								PRIMARY KEY (tran_datetime, card_pan_sha256, tran_amount)
								) `, "generator").Exec(); err != nil {
		fmt.Println("Error creating table")
		log.Fatal(err)
	}
	fmt.Print("ok")
	fmt.Println()
	session.Close()
}

func KeyspaceCreate(seeds string) {
	session := SessionInit(seeds)
	fmt.Print("If not exists creating keyspace 'generator'...")
	if err := session.Query(`CREATE KEYSPACE IF NOT EXISTS generator
							WITH REPLICATION =
							{'class': 'NetworkTopologyStrategy',
							'dc1': 2, 'dc2': 2}`).Exec(); err != nil {
		fmt.Print("Error creating keyspace")
		log.Fatal(err)
	}
	fmt.Print("ok")
	fmt.Println()
	session.Close()
}
