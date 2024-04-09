package models

type TagGrupoDeCobrancaXml struct {
	Fat     TagGrupoDeFaturaXml   `xml:"fat,omitempty"`
	DupList []TagGrupoParcelasXml `xml:"dup,omitempty"`
}
