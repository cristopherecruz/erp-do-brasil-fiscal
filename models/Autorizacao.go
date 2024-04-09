package models

import (
	"encoding/xml"
	"fmt"
	"github.com/hooklift/gowsdl/soap"
	"log"
)

type NfeDadosMsg struct {
	XMLName xml.Name `xml:"http://www.portalfiscal.inf.br/nfe/wsdl/NFeAutorizacao4 nfeDadosMsg"`
}

type NfeResultMsg struct {
	XMLName xml.Name `xml:"http://www.portalfiscal.inf.br/nfe/wsdl/NFeAutorizacao4 nfeResultMsg"`
	Content string   `xml:",innerxml"`
}

type NfeMonitoria struct {
	XMLName      xml.Name `xml:"http://www.portalfiscal.inf.br/nfe/wsdl/NFeAutorizacao4 nfeMonitoria"`
	NomeServidor string   `xml:"nomeServidor,omitempty"`
	DhServidor   string   `xml:"dhServidor,omitempty"`
}

type LayoutMensagemEntradaXml struct {
	XMLName           xml.Name              `xml:"enviNFe"`
	Versao            string                `xml:"versao,attr"`
	IdLote            string                `xml:"idLote"`
	IndicadorProcesso Processamento         `xml:"indSinc"`
	Notas             []LayoutNotaFiscalXml `xml:"NFe"`
}

type LayoutNotaFiscalXml struct {
	XMLName            xml.Name               `xml:"NFe"`
	IdentificadorLocal int64                  `xml:"identificadorLocal"`
	Info               TagsDadosNotaFiscalXml `xml:"infNFe"`
	InfoSuplementar    NFNotaInfoSuplementar  `xml:"infNFeSupl,omitempty"`
	Assinatura         NFSignature            `xml:"Signature,omitempty"`
}

type TagsDadosNotaFiscalXml struct {
	XMLName               xml.Name                        `xml:"infNFe"`
	Versao                string                          `xml:"versao,attr"`
	Identificador         string                          `xml:"Id,attr"`
	Identificacao         TagsIdentificacaoXml            `xml:"ide"`
	Emitente              TagsEmitenteXml                 `xml:"emit"`
	Destinatario          NFNotaInfoDestinatario          `xml:"dest,omitempty"`
	Itens                 []TagsDetalharProdutoServicoXml `xml:"det"`
	AutXML                []TagsDePessoasAutorizadas      `xml:"autXML,omitempty"`
	Total                 TagsTotalXml                    `xml:"total"`
	Transporte            TagsTransporteXml               `xml:"transp"`
	Cobr                  TagGrupoDeCobrancaXml           `xml:"cobr,omitempty"`
	Pagamentos            []TagsFormaPagamentoXml         `xml:"pag"`
	InformacoesAdicionais TagsInformacaoAdicionalXml      `xml:"infAdic,omitempty"`
	InfRespTec            NFInfRespTec                    `xml:"infRespTec,omitempty"`
}

type TagsIdentificacaoXml struct {
	UF                                UnidadeFederativa         `xml:"cUF"`
	CodigoRandomico                   string                    `xml:"cNF"`
	NaturezaOperacao                  string                    `xml:"natOp"`
	Modelo                            ModeloNF                  `xml:"mod"`
	Serie                             string                    `xml:"serie"`
	NumeroNota                        string                    `xml:"nNF"`
	DataHoraEmissao                   string                    `xml:"dhEmi"`
	DHSaiEnt                          string                    `xml:"dhSaiEnt,omitempty"`
	Tipo                              TipoNota                  `xml:"tpNF"`
	IdentificadorLocalDestinoOperacao LocalDestinoOperacao      `xml:"idDest"`
	CodigoMunicipio                   string                    `xml:"cMunFG"`
	TipoImpressao                     TipoImpressao             `xml:"tpImp"`
	TipoEmissao                       TipoEmissao               `xml:"tpEmis"`
	DigitoVerificador                 int                       `xml:"cDV"`
	Ambiente                          Ambiente                  `xml:"tpAmb"`
	Finalidade                        Finalidade                `xml:"finNFe"`
	OperacaoConsumidorFinal           ConsumidorFinal           `xml:"indFinal"`
	IndicadorPresencaComprador        PresencaComprador         `xml:"indPres"`
	IndicadorDoIntermediador          IndicadorDoIntermediador  `xml:"indIntermed,omitempty"`
	ProgramaEmissor                   ProcessoEmissor           `xml:"procEmi"`
	VersaoEmissor                     string                    `xml:"verProc"`
	Referenciadas                     []TagsFiscalReferenciaXml `xml:"NFref,omitempty"`
}

type NFeAutorizacao4Soap interface {
	NfeAutorizacaoLote(request *NfeDadosMsg) (*NfeResultMsg, error)

	NfeAutorizacaoLoteZip(request *string) (*NfeResultMsg, error)
}

type nFeAutorizacao4Soap struct {
	client *soap.Client
}

func NewNFeAutorizacao4Soap(client *soap.Client) NFeAutorizacao4Soap {
	return &nFeAutorizacao4Soap{
		client: client,
	}
}

func (service *nFeAutorizacao4Soap) NfeAutorizacaoLote(request *NfeDadosMsg) (*NfeResultMsg, error) {
	response := new(NfeResultMsg)

	err := service.client.Call("http://www.portalfiscal.inf.br/nfe/wsdl/NFeAutorizacao4/nfeAutorizacaoLote", request, response)
	if err != nil {
		log.Println(err)
	}

	formattedXML, err2 := xml.MarshalIndent(response, "", "    ")
	if err2 != nil {
		fmt.Printf("Erro ao formatar XML: %v\n", err2)
	}

	// Imprime o XML formatado no console
	fmt.Println(string(formattedXML))

	return response, nil
}

func (service *nFeAutorizacao4Soap) NfeAutorizacaoLoteZip(request *string) (*NfeResultMsg, error) {
	response := new(NfeResultMsg)
	err := service.client.Call("http://www.portalfiscal.inf.br/nfe/wsdl/NFeAutorizacao4/nfeAutorizacaoLoteZip", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
