package models

import (
	"errors"
	"strings"
)

type ModalidadeFrete string

const (
	ModalidadeFreteContratacaoPorContaDoRemetente    ModalidadeFrete = "0" // Por Conta do Emitente (CIF)
	ModalidadeFreteContratacaoPorContaDoDestinatario ModalidadeFrete = "1" // Por Conta do Destinatário (FOB)
	ModalidadeFreteContratacaoPorContaDeTerceiros    ModalidadeFrete = "2" // Por Conta de Terceiros
	ModalidadeFreteProprioPorContaDoRemetente        ModalidadeFrete = "3" // Transporte próprio por conta do remetente
	ModalidadeFreteProprioPorContaDoDestinatario     ModalidadeFrete = "4" // Transporte próprio por conta do destinatário
	ModalidadeFreteSemOcorrenciaTransporte           ModalidadeFrete = "9" // Sem Frete
)

var modalidadeFreteDescriptions = map[ModalidadeFrete]string{
	ModalidadeFreteContratacaoPorContaDoRemetente:    "Por Conta do Emitente (CIF)",
	ModalidadeFreteContratacaoPorContaDoDestinatario: "Por Conta do Destinatário (FOB)",
	ModalidadeFreteContratacaoPorContaDeTerceiros:    "Por Conta de Terceiros",
	ModalidadeFreteProprioPorContaDoRemetente:        "Transporte próprio por conta do remetente",
	ModalidadeFreteProprioPorContaDoDestinatario:     "Transporte próprio por conta do destinatário",
	ModalidadeFreteSemOcorrenciaTransporte:           "Sem Frete",
}

func (m ModalidadeFrete) Descricao() string {
	if desc, ok := modalidadeFreteDescriptions[m]; ok {
		return desc
	}
	return ""
}

func ModalidadeFreteFromDescricao(descricao string) (ModalidadeFrete, error) {
	for modalidade, desc := range modalidadeFreteDescriptions {
		if strings.EqualFold(desc, descricao) {
			return modalidade, nil
		}
	}
	return "", errors.New("ModalidadeFrete não encontrada para a descrição fornecida")
}

func ModalidadeFreteFromCodigo(codigo string) (ModalidadeFrete, error) {
	switch codigo {
	case "0":
		return ModalidadeFreteContratacaoPorContaDoRemetente, nil
	case "1":
		return ModalidadeFreteContratacaoPorContaDoDestinatario, nil
	case "2":
		return ModalidadeFreteContratacaoPorContaDeTerceiros, nil
	case "3":
		return ModalidadeFreteProprioPorContaDoRemetente, nil
	case "4":
		return ModalidadeFreteProprioPorContaDoDestinatario, nil
	case "9":
		return ModalidadeFreteSemOcorrenciaTransporte, nil
	default:
		return "", errors.New("ModalidadeFrete não encontrada para o código fornecido")
	}
}
