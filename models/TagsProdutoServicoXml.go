package models

type TagsProdutoServicoXml struct {
	CProd    string                 `xml:"cProd"`
	CEAN     string                 `xml:"cEAN,omitempty"`
	XProd    string                 `xml:"xProd"`
	NCM      string                 `xml:"NCM"`
	CEST     string                 `xml:"CEST,omitempty"`
	CBenef   string                 `xml:"cBenef,omitempty"`
	CFOP     string                 `xml:"CFOP"`
	UCom     string                 `xml:"uCom"`
	QCom     string                 `xml:"qCom"`
	VUnCom   string                 `xml:"vUnCom"`
	VProd    string                 `xml:"vProd"`
	CEANTrib string                 `xml:"cEANTrib,omitempty"`
	UTrib    string                 `xml:"uTrib"`
	QTrib    string                 `xml:"qTrib"`
	VUnTrib  string                 `xml:"vUnTrib"`
	VFrete   string                 `xml:"vFrete,omitempty"`
	VSeg     string                 `xml:"vSeg,omitempty"`
	VDesc    string                 `xml:"vDesc,omitempty"`
	VOutro   string                 `xml:"vOutro,omitempty"`
	IndTot   ProdutoCompoeValorNota `xml:"indTot"`
	XPed     string                 `xml:"xPed,omitempty"`
	NItemPed int                    `xml:"nItemPed,omitempty"`
	NFCI     string                 `xml:"nFCI,omitempty"`
}
