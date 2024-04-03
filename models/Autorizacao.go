package models

import (
	"encoding/xml"
	"fmt"
	"github.com/hooklift/gowsdl/soap"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type NfeDadosMsg struct {
	XMLName xml.Name `xml:"http://www.portalfiscal.inf.br/nfe/wsdl/NFeAutorizacao4 nfeDadosMsg"`
}

type NfeResultMsg struct {
	XMLName xml.Name `xml:"http://www.portalfiscal.inf.br/nfe/wsdl/NFeAutorizacao4 nfeResultMsg"`
}

type NfeMonitoria struct {
	XMLName xml.Name `xml:"http://www.portalfiscal.inf.br/nfe/wsdl/NFeAutorizacao4 nfeMonitoria"`

	NomeServidor string `xml:"nomeServidor,omitempty"`

	DhServidor string `xml:"dhServidor,omitempty"`
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

	formattedXML, err := xml.MarshalIndent(request, "", "    ")
	if err != nil {
		fmt.Printf("Erro ao formatar XML: %v\n", err)
	}

	// Imprime o XML formatado no console
	fmt.Println(string(formattedXML))
	/*err := service.client.Call("http://www.portalfiscal.inf.br/nfe/wsdl/NFeAutorizacao4/nfeAutorizacaoLote", request, response)
	if err != nil {
		return nil, err
	}*/

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
