package models

import "fmt"

type ProcessoEmissor struct {
	Codigo    string
	Descricao string
}

// Função de construtor para criar uma nova instância de ProcessoEmissor
func NewProcessoEmissor(codigo, descricao string) ProcessoEmissor {
	return ProcessoEmissor{Codigo: codigo, Descricao: descricao}
}

// Mapa para mapear códigos de processos emissores para suas instâncias correspondentes
var processoMap = map[string]ProcessoEmissor{
	"0": NewProcessoEmissor("0", "Aplicativo do contribuinte"),
	"1": NewProcessoEmissor("1", "Emissão de NF-e avulsa pelo Fisco"),
	"2": NewProcessoEmissor("2", "Emissão de NF-e avulsa, pelo contribuinte com certificado digital através do Fisco"),
	"3": NewProcessoEmissor("3", "Emissão de NF-e pelo contribuinte com aplicativo fornecido pelo Fisco"),
}

// Função para obter uma instância de ProcessoEmissor pelo código
func GetProcessoEmissorByCodigo(codigo string) (ProcessoEmissor, error) {
	processo, ok := processoMap[codigo]
	if !ok {
		return ProcessoEmissor{}, fmt.Errorf("ProcessoEmissor com código %s não encontrado", codigo)
	}
	return processo, nil
}
