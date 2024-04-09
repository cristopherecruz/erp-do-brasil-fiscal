package models

type FormaDePagamento string

const (
	DINHEIRO            FormaDePagamento = "01" // Dinheiro
	CHEQUE              FormaDePagamento = "02" // Cheque
	CARTAO_CREDITO      FormaDePagamento = "03" // Cartão de crédito
	CARTAO_DEBITO       FormaDePagamento = "04" // Cartão de débito
	CREDITO_LOJA        FormaDePagamento = "05" // Crédito Loja
	VALE_ALIMENTACAO    FormaDePagamento = "10" // Vale alimentação
	VALE_REFEICAO       FormaDePagamento = "11" // Vale refeição
	VALE_PRESENTE       FormaDePagamento = "12" // Vale presente
	VALE_COMBUSTIVEL    FormaDePagamento = "13" // Vale combustível
	CARTAO_LOJA         FormaDePagamento = "14" // Cartão da loja
	CREDIARIO           FormaDePagamento = "15" // Crediário
	CARTAO              FormaDePagamento = "16" // Cartão
	OUTRO               FormaDePagamento = "99" // Outros
	VOUCHER             FormaDePagamento = "17" // Voucher
	SEM_PAGAMENTO_FORMA FormaDePagamento = "90" // Sem Pagamento
	TEF_CREDITO         FormaDePagamento = "21" // TEF Crédito
	TEF_DEBITO          FormaDePagamento = "22" // TEF Débito
	TEF_PIX             FormaDePagamento = "25" // TEF Pix
)
