package models

import "errors"

// TributacaoICMS representa a tributação do ICMS
type TributacaoICMS struct {
	Codigo    string
	Descricao string
}

// Definições possíveis para TributacaoICMS
var (
	TributacaoIntegralmente                                                              = TributacaoICMS{"00", "Tributada integralmente"}
	TributadaComCobrancaICMSPorSubstituicaoTributaria                                    = TributacaoICMS{"10", "Tributada com cobrança de ICMS por ST"}
	ComReducaoBaseCalculo                                                                = TributacaoICMS{"20", "Com redução da base de cálculo"}
	IsentaOuNaoTributadaComCobrancaICMSPorSubstituicaoTributaria                         = TributacaoICMS{"30", "Isenta ou não tributada com cobrança de ICMS por ST"}
	Isenta                                                                               = TributacaoICMS{"40", "Isenta"}
	NaoTributado                                                                         = TributacaoICMS{"41", "Não tributada"}
	Suspensao                                                                            = TributacaoICMS{"50", "Suspensão"}
	Diferimento                                                                          = TributacaoICMS{"51", "Diferimento"}
	ICMSCobradoAnteriormentePorSubstituicaoTributaria                                    = TributacaoICMS{"60", "ICMS cobrado anteriormente por ST"}
	ComReducaoBaseCalculoCobrancaICMSPorSubstituicaoTributariaICMSSubstituicaoTributaria = TributacaoICMS{"70", "Com redução da base de cálculo/Cobrança ICMS por ST/ICMS ST"}
	Outros                                                                               = TributacaoICMS{"90", "Outros"}

	// CSOSN
	TributadaPeloSimplesNacionalComPermissaoDeCreditoDeICMS                   = TributacaoICMS{"101", "Tributada pelo Simples Nacional com permissão de crédito de ICMS"}
	TributadaPeloSimplesNacionalSemPermissaoDeCredito                         = TributacaoICMS{"102", "Tributada pelo Simples Nacional sem permissão de crédito"}
	IsencaoICMSNoSimplesNacionalNaFaixaDeReceitaBruta                         = TributacaoICMS{"103", "Isenção de ICMS no Simples Nacional na faixa de receita bruta"}
	TributadaPeloSimplesNacionalComPermissaoDeCreditoECobrancaDoICMSPorST     = TributacaoICMS{"201", "Tributada pelo Simples Nacional com permissão de crédito e cobrança do ICMS por ST"}
	TributadaPeloSimplesNacionalSemPermissaoDeCreditoEComCobrancaDoICMSPorST  = TributacaoICMS{"202", "Tributada pelo Simples Nacional sem permissão de crédito e com cobrança do ICMS por ST"}
	IsencaoDoICMSNoSimplesNacionalParaFaixaDeReceitaBrutaECobrancaDeICMSPorST = TributacaoICMS{"203", "Isenção do ICMS no Simples Nacional para faixa de receita bruta e cobrança de ICMS por ST"}
	ImuneDeICMS                                                               = TributacaoICMS{"300", "Imune de ICMS"}
	NaoTributadaPeloSimplesNacional                                           = TributacaoICMS{"400", "Não tributada pelo Simples Nacional"}
	ICMSCobradoAnteriormentePorSubstituicaoTributariaOuPorAntecipacao         = TributacaoICMS{"500", "ICMS cobrado anteriormente por ST ou por antecipação"}
	OutrosCSOSN                                                               = TributacaoICMS{"900", "Outros (operações que não se enquadram nos códigos anteriores)"}
)

// ValorDeCodigo retorna a TributacaoICMS correspondente ao código fornecido
func ValorDeCodigoTributacao(codigo string) (TributacaoICMS, error) {
	switch codigo {
	case TributacaoIntegralmente.Codigo:
		return TributacaoIntegralmente, nil
	case TributadaComCobrancaICMSPorSubstituicaoTributaria.Codigo:
		return TributadaComCobrancaICMSPorSubstituicaoTributaria, nil
	case ComReducaoBaseCalculo.Codigo:
		return ComReducaoBaseCalculo, nil
	case IsentaOuNaoTributadaComCobrancaICMSPorSubstituicaoTributaria.Codigo:
		return IsentaOuNaoTributadaComCobrancaICMSPorSubstituicaoTributaria, nil
	case Isenta.Codigo:
		return Isenta, nil
	case NaoTributado.Codigo:
		return NaoTributado, nil
	case Suspensao.Codigo:
		return Suspensao, nil
	case Diferimento.Codigo:
		return Diferimento, nil
	case ICMSCobradoAnteriormentePorSubstituicaoTributaria.Codigo:
		return ICMSCobradoAnteriormentePorSubstituicaoTributaria, nil
	case ComReducaoBaseCalculoCobrancaICMSPorSubstituicaoTributariaICMSSubstituicaoTributaria.Codigo:
		return ComReducaoBaseCalculoCobrancaICMSPorSubstituicaoTributariaICMSSubstituicaoTributaria, nil
	case Outros.Codigo:
		return Outros, nil
	case TributadaPeloSimplesNacionalComPermissaoDeCreditoDeICMS.Codigo:
		return TributadaPeloSimplesNacionalComPermissaoDeCreditoDeICMS, nil
	case TributadaPeloSimplesNacionalSemPermissaoDeCredito.Codigo:
		return TributadaPeloSimplesNacionalSemPermissaoDeCredito, nil
	case IsencaoICMSNoSimplesNacionalNaFaixaDeReceitaBruta.Codigo:
		return IsencaoICMSNoSimplesNacionalNaFaixaDeReceitaBruta, nil
	case TributadaPeloSimplesNacionalComPermissaoDeCreditoECobrancaDoICMSPorST.Codigo:
		return TributadaPeloSimplesNacionalComPermissaoDeCreditoECobrancaDoICMSPorST, nil
	case TributadaPeloSimplesNacionalSemPermissaoDeCreditoEComCobrancaDoICMSPorST.Codigo:
		return TributadaPeloSimplesNacionalSemPermissaoDeCreditoEComCobrancaDoICMSPorST, nil
	case IsencaoDoICMSNoSimplesNacionalParaFaixaDeReceitaBrutaECobrancaDeICMSPorST.Codigo:
		return IsencaoDoICMSNoSimplesNacionalParaFaixaDeReceitaBrutaECobrancaDeICMSPorST, nil
	case ImuneDeICMS.Codigo:
		return ImuneDeICMS, nil
	case NaoTributadaPeloSimplesNacional.Codigo:
		return NaoTributadaPeloSimplesNacional, nil
	case ICMSCobradoAnteriormentePorSubstituicaoTributariaOuPorAntecipacao.Codigo:
		return ICMSCobradoAnteriormentePorSubstituicaoTributariaOuPorAntecipacao, nil
	case OutrosCSOSN.Codigo:
		return OutrosCSOSN, nil
	default:
		return TributacaoICMS{}, errors.New("Tributação ICMS não encontrada")
	}
}
