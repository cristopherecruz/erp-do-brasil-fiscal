package models

type OperadoraCartao string

const (
	VISA            OperadoraCartao = "01"
	MASTERCARD      OperadoraCartao = "02"
	AMERICANEXPRESS OperadoraCartao = "03"
	SOROCRED        OperadoraCartao = "04"
	DINERSCLUB      OperadoraCartao = "05"
	ELO             OperadoraCartao = "06"
	HIPERCARD       OperadoraCartao = "07"
	AURA            OperadoraCartao = "08"
	CABAL           OperadoraCartao = "09"
	OUTROS          OperadoraCartao = "99"
)
