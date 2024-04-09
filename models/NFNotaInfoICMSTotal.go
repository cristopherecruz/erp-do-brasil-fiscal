package models

type NFNotaInfoICMSTotal struct {
	VBC          string `xml:"vBC"`
	VICMS        string `xml:"vICMS"`
	VICMSDeson   string `xml:"vICMSDeson"`
	VFCPUFDest   string `xml:"vFCPUFDest,omitempty"`
	VICMSUFDest  string `xml:"vICMSUFDest,omitempty"`
	VICMSUFRemet string `xml:"vICMSUFRemet,omitempty"`
	VFCP         string `xml:"vFCP"`
	VBCST        string `xml:"vBCST"`
	VST          string `xml:"vST"`
	VFCPST       string `xml:"vFCPST"`
	VFCPSTRet    string `xml:"vFCPSTRet"`
	VProd        string `xml:"vProd"`
	VFrete       string `xml:"vFrete"`
	VSeg         string `xml:"vSeg"`
	VDesc        string `xml:"vDesc"`
	VII          string `xml:"vII"`
	VIPI         string `xml:"vIPI"`
	VIPIDevol    string `xml:"vIPIDevol"`
	VPIS         string `xml:"vPIS"`
	VCOFINS      string `xml:"vCOFINS"`
	VOutro       string `xml:"vOutro"`
	VNF          string `xml:"vNF"`
	VTotTrib     string `xml:"vTotTrib,omitempty"`
}
