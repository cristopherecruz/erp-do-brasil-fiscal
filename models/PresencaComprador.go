package models

import "fmt"

type PresencaComprador struct {
	Codigo    string
	Descricao string
}

// Valores pré-definidos
var (
	NAO_APLICA                               = PresencaComprador{"0", "Não se aplica"}
	OPERACAO_PRESENCIAL                      = PresencaComprador{"1", "Operação presencial"}
	OPERACAO_NAO_PRESENCIAL_INTERNET         = PresencaComprador{"2", "Operação não presencial - Internet"}
	OPERACAO_NAO_PRESENCIAL_TELEATENDIMENTO  = PresencaComprador{"3", "Operação não presencial - Teleatendimento"}
	NFCE_EM_OPERACAO_COM_ENTREGA_DOMICILIO   = PresencaComprador{"4", "NFC-e em operação com entrega a domicílio"}
	OPERACAO_PRESENCIAL_FORA_ESTABELECIMENTO = PresencaComprador{"5", "Operação presencial, fora do estabelecimento"}
	OPERACAO_NAO_PRESENCIAL_OUTROS           = PresencaComprador{"9", "Operação não presencial - Outros"}
)

// Função para recuperar a presença do comprador com base no código
func RecuperarPresencaCompradorPorCodigo(codigo string) (PresencaComprador, error) {
	presencas := map[string]PresencaComprador{
		NAO_APLICA.Codigo:                               NAO_APLICA,
		OPERACAO_PRESENCIAL.Codigo:                      OPERACAO_PRESENCIAL,
		OPERACAO_NAO_PRESENCIAL_INTERNET.Codigo:         OPERACAO_NAO_PRESENCIAL_INTERNET,
		OPERACAO_NAO_PRESENCIAL_TELEATENDIMENTO.Codigo:  OPERACAO_NAO_PRESENCIAL_TELEATENDIMENTO,
		NFCE_EM_OPERACAO_COM_ENTREGA_DOMICILIO.Codigo:   NFCE_EM_OPERACAO_COM_ENTREGA_DOMICILIO,
		OPERACAO_PRESENCIAL_FORA_ESTABELECIMENTO.Codigo: OPERACAO_PRESENCIAL_FORA_ESTABELECIMENTO,
		OPERACAO_NAO_PRESENCIAL_OUTROS.Codigo:           OPERACAO_NAO_PRESENCIAL_OUTROS,
	}

	if presenca, ok := presencas[codigo]; ok {
		return presenca, nil
	}
	return PresencaComprador{}, fmt.Errorf("Presença do comprador não encontrada para o código %s", codigo)
}
