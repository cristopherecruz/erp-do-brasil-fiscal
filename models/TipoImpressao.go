package models

import "fmt"

type TipoImpressao struct {
	Codigo    string
	Descricao string
}

// Valores pré-definidos
var (
	SEM_GERACAO_DANFE              = TipoImpressao{"0", "Sem geração de DANFe"}
	DANFE_NORMAL_RETRATO           = TipoImpressao{"1", "DANFe normal retrato"}
	DANFE_NORMAL_PAISAGEM          = TipoImpressao{"2", "DANFe normal paisagem"}
	DANFE_SIMPLIFICADO             = TipoImpressao{"3", "DANFe simplificado"}
	DANFE_NFCE                     = TipoImpressao{"4", "DANFe NFCe"}
	DANFE_NFCE_MENSAGEM_ELETRONICA = TipoImpressao{"5", "DANFe NFCe mensagem eletrônica"}
)

// Função para recuperar o TipoImpressao com base no código
func RecuperarTipoImpressaoPorCodigo(codigo string) (TipoImpressao, error) {
	tipos := map[string]TipoImpressao{
		SEM_GERACAO_DANFE.Codigo:              SEM_GERACAO_DANFE,
		DANFE_NORMAL_RETRATO.Codigo:           DANFE_NORMAL_RETRATO,
		DANFE_NORMAL_PAISAGEM.Codigo:          DANFE_NORMAL_PAISAGEM,
		DANFE_SIMPLIFICADO.Codigo:             DANFE_SIMPLIFICADO,
		DANFE_NFCE.Codigo:                     DANFE_NFCE,
		DANFE_NFCE_MENSAGEM_ELETRONICA.Codigo: DANFE_NFCE_MENSAGEM_ELETRONICA,
	}

	if ti, ok := tipos[codigo]; ok {
		return ti, nil
	}
	return TipoImpressao{}, fmt.Errorf("Tipo Impressão não encontrado para o código %s", codigo)
}
