package models

type IndicadorFormaPagamento string

const (
	A_VISTA       IndicadorFormaPagamento = "0" // Pagamento à vista
	A_PRAZO       IndicadorFormaPagamento = "1" // Pagamento a prazo
	SEM_PAGAMENTO IndicadorFormaPagamento = "2" // Sem Pagamento
)
