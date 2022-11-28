package dto

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
