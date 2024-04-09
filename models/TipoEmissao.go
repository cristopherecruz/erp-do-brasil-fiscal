package models

import "fmt"

type TipoEmissao struct {
	Codigo    string
	Descricao string
}

// Valores pré-definidos
var (
	EMISSAO_NORMAL      = TipoEmissao{"1", "Normal"}
	CONTINGENCIA_FS_IA  = TipoEmissao{"2", "Contingência FS"}
	CONTINGENCIA_SCAN   = TipoEmissao{"3", "Contingência SCAN"}
	CONTINGENCIA_DPEC   = TipoEmissao{"4", "Contingência DPEC"}
	CONTINGENCIA_FSDA   = TipoEmissao{"5", "Contingência FSDA"}
	CONTINGENCIA_SVCAN  = TipoEmissao{"6", "Contingência SVCAN"}
	CONTINGENCIA_SVCRS  = TipoEmissao{"7", "Contingência SVCRS"}
	CONTIGENCIA_OFFLINE = TipoEmissao{"9", "Contingência offline NFC-e"}
)

// Função para recuperar o TipoEmissao com base no código
func RecuperarTipoEmissaoPorCodigo(codigo string) (TipoEmissao, error) {
	tipos := map[string]TipoEmissao{
		EMISSAO_NORMAL.Codigo:      EMISSAO_NORMAL,
		CONTINGENCIA_FS_IA.Codigo:  CONTINGENCIA_FS_IA,
		CONTINGENCIA_SCAN.Codigo:   CONTINGENCIA_SCAN,
		CONTINGENCIA_DPEC.Codigo:   CONTINGENCIA_DPEC,
		CONTINGENCIA_FSDA.Codigo:   CONTINGENCIA_FSDA,
		CONTINGENCIA_SVCAN.Codigo:  CONTINGENCIA_SVCAN,
		CONTINGENCIA_SVCRS.Codigo:  CONTINGENCIA_SVCRS,
		CONTIGENCIA_OFFLINE.Codigo: CONTIGENCIA_OFFLINE,
	}

	if te, ok := tipos[codigo]; ok {
		return te, nil
	}
	return TipoEmissao{}, fmt.Errorf("Tipo Emissão não encontrado para o código %s", codigo)
}
