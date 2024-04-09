package models

import "errors"

type OrigemNotaFiscal struct {
	Codigo    string
	Descricao string
}

// Definições possíveis para OrigemNotaFiscal
var (
	Nacional                                                                  = OrigemNotaFiscal{"0", "Nacional"}
	EstrangeiraImportacaoDireta                                               = OrigemNotaFiscal{"1", "Estrangeira importação direta"}
	EstrangeiraAdquiridaMercadoInterno                                        = OrigemNotaFiscal{"2", "Estrangeira adquirida mercado interno"}
	NacionalMercadoriaOuBemConteudoImportacaoSuperior40P                      = OrigemNotaFiscal{"3", "Nacional mercadoria ou bem conteúdo importação superior 40 P"}
	NacionalProducaoEmConformidadeComProcessosProdutivosBasicos               = OrigemNotaFiscal{"4", "Nacional produção em conformidade com processos produtivos básicos"}
	NacionalMercadoriaOuBemConteudoImportacaoInferior40P                      = OrigemNotaFiscal{"5", "Nacional mercadoria ou bem conteúdo importação inferior 40 P"}
	EstrangeiraImportacaoDiretaSemSimilarNacionalConstanteEmListaCamex        = OrigemNotaFiscal{"6", "Estrangeira importação direta sem similar nacional constante em lista Camex"}
	EstrangeiraAdquiridaMercadoInternoSemSimilarNacionalConstanteEmListaCamex = OrigemNotaFiscal{"7", "Estrangeira adquirida mercado interno sem similar nacional constante em lista Camex"}
	NacionalMercadoriaOuBemComConteudoImportacaoSuperior70P                   = OrigemNotaFiscal{"8", "Nacional mercadoria ou bem conteúdo importação superior 70 P"}
)

// ValorDeCodigo retorna a OrigemNotaFiscal correspondente ao código fornecido
func ValorDeCodigoOrigem(codigo string) (OrigemNotaFiscal, error) {
	switch codigo {
	case Nacional.Codigo:
		return Nacional, nil
	case EstrangeiraImportacaoDireta.Codigo:
		return EstrangeiraImportacaoDireta, nil
	case EstrangeiraAdquiridaMercadoInterno.Codigo:
		return EstrangeiraAdquiridaMercadoInterno, nil
	case NacionalMercadoriaOuBemConteudoImportacaoSuperior40P.Codigo:
		return NacionalMercadoriaOuBemConteudoImportacaoSuperior40P, nil
	case NacionalProducaoEmConformidadeComProcessosProdutivosBasicos.Codigo:
		return NacionalProducaoEmConformidadeComProcessosProdutivosBasicos, nil
	case NacionalMercadoriaOuBemConteudoImportacaoInferior40P.Codigo:
		return NacionalMercadoriaOuBemConteudoImportacaoInferior40P, nil
	case EstrangeiraImportacaoDiretaSemSimilarNacionalConstanteEmListaCamex.Codigo:
		return EstrangeiraImportacaoDiretaSemSimilarNacionalConstanteEmListaCamex, nil
	case EstrangeiraAdquiridaMercadoInternoSemSimilarNacionalConstanteEmListaCamex.Codigo:
		return EstrangeiraAdquiridaMercadoInternoSemSimilarNacionalConstanteEmListaCamex, nil
	case NacionalMercadoriaOuBemComConteudoImportacaoSuperior70P.Codigo:
		return NacionalMercadoriaOuBemComConteudoImportacaoSuperior70P, nil
	default:
		return OrigemNotaFiscal{}, errors.New("Origem Nota Fiscal não encontrado!")
	}
}
