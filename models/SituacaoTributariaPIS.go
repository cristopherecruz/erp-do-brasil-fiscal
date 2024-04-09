package models

import "errors"

type SituacaoTributariaPIS struct {
	Codigo    string
	Descricao string
}

var SituacoesTributariasPIS = []SituacaoTributariaPIS{
	{"01", "Operação tributável cumulativo/não cumulativo"},
	{"02", "Operação tributável alíquota diferenciada"},
	{"03", "Operação tributável quantidade vendida por alíquota por unidade produto"},
	{"04", "Operação tributável monofásica alíquota zero"},
	{"05", "Operação tributável ST"},
	{"06", "Operação tributável alíquota zero"},
	{"07", "Operação isenta de contribuição"},
	{"08", "Operação sem incidência de contribuição"},
	{"09", "Operação com suspensão de contribuição"},
	{"49", "Outras operações de saída"},
	{"50", "Operação direito crédito vinculada exclusivamente receita tributada mercado interno"},
	{"51", "Operação direito crédito vinculada exclusivamente à receita não tributada mercado interno"},
	{"52", "Operação direito crédito vinculada exclusivamente receita exportação"},
	{"53", "Operação direito crédito vinculada receitas tributada e não tributada mercado interno"},
	{"54", "Operação direito crédito vinculada receitas tributadas no mercado interno exportação"},
	{"55", "Operação direito crédito vinculada receitas não tributada no mercado interno exportação"},
	{"56", "Operação direito crédito vinculada receitas tributadas e não tributadas mercado interno exportação"},
	{"60", "Crédito presumido operação aquisição vinculada exclusivamente receita tributada mercado interno"},
	{"61", "Crédito presumido operação aquisição vinculada exclusivamente à receita não tributada mercado interno"},
	{"62", "Crédito presumido operação aquisição vinculada exclusivamente receita exportação"},
	{"63", "Crédito presumido operação aquisição vinculada receitas tributadas mercado interno"},
	{"64", "Crédito presumido operação aquisição vinculada receitas tributadas mercado interno exportação"},
	{"65", "Crédito presumido operação aquisição vinculada receitas não tributadas mercado interno exportação"},
	{"66", "Crédito presumido operação aquisição vinculada receitas tributadas e não tributadas mercado interno exportação"},
	{"67", "Crédito presumido outras operações"},
	{"70", "Operação aquisição sem direito crédito"},
	{"71", "Operação aquisição com isenção"},
	{"72", "Operação aquisição com suspensão"},
	{"73", "Operação aquisição alíquota zero"},
	{"74", "Operação aquisição sem incidência de contribuição"},
	{"75", "Operação aquisição por substituição tributária"},
	{"98", "Outras operações de entrada"},
	{"99", "Outras operações"},
}

func ValueOfCodigo(codigo string) (SituacaoTributariaPIS, error) {
	for _, situacao := range SituacoesTributariasPIS {
		if situacao.Codigo == codigo {
			return situacao, nil
		}
	}
	return SituacaoTributariaPIS{}, errors.New("Situacao Tributaria PIS não encontrado!")
}
