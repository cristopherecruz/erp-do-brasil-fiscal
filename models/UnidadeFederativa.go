package models

type UnidadeFederativa struct {
	Codigo                         string
	Descricao                      string
	CodigoIbge                     string
	ZoneID                         string
	QRCodeHomologacao              string
	QRCodeProducao                 string
	ConsultaChaveAcessoHomologacao string
	ConsultaChaveAcessoProducao    string
}

// Mapa de unidades federativas
var Estados = map[string]UnidadeFederativa{
	"AC":       {"AC", "Acre", "12", "CST", "http://hml.sefaznet.ac.gov.br/nfce/qrcode", "http://sefaznet.ac.gov.br/nfce/qrcode", "http://sefaznet.ac.gov.br/nfce/consulta", "http://sefaznet.ac.gov.br/nfce/consulta"},
	"AL":       {"AL", "Alagoas", "27", "BET", "http://nfce.sefaz.al.gov.br/QRCode/consultarNFCe.jsp", "http://nfce.sefaz.al.gov.br/QRCode/consultarNFCe.jsp", "http://nfce.sefaz.al.gov.br/consultaNFCe.htm", "http://nfce.sefaz.al.gov.br/consultaNFCe.htm"},
	"AP":       {"AP", "Amapá", "16", "BET", "https://www.sefaz.ap.gov.br/nfcehml/nfce.php", "https://www.sefaz.ap.gov.br/nfce/nfce.php", "https://www.sefaz.ap.gov.br/sate1/seg/SEGf_AcessarFuncao.jsp?cdFuncao=FIS_1261", "https://www.sefaz.ap.gov.br/sate/seg/SEGf_AcessarFuncao.jsp?cdFuncao=FIS_1261"},
	"AM":       {"AM", "Amazonas", "13", "IET", "http://homnfce.sefaz.am.gov.br/nfceweb/consultarNFCe.jsp", "http://sistemas.sefaz.am.gov.br/nfceweb/consultarNFCe.jsp", "homnfce.sefaz.am.gov.br/nfceweb/formConsulta.do", "sistemas.sefaz.am.gov.br/nfceweb/formConsulta.do"},
	"BA":       {"BA", "Bahia", "29", "BET", "http://hnfe.sefaz.ba.gov.br/servicos/nfce/modulos/geral/NFCEC_consulta_chave_acesso.aspx", "http://nfe.sefaz.ba.gov.br/servicos/nfce/modulos/geral/NFCEC_consulta_chave_acesso.aspx", "http://hnfe.sefaz.ba.gov.br/servicos/nfce/default.aspx", "nfe.sefaz.ba.gov.br/servicos/nfce/default.aspx"},
	"CE":       {"CE", "Ceará", "23", "BET", "http://nfceh.sefaz.ce.gov.br/pages/ShowNFCe.html", "http://nfce.sefaz.ce.gov.br/pages/ShowNFCe.html", "https://nfeh.sefaz.ce.gov.br/pages/consultaChaveAcesso.jsf", "https://nfe.sefaz.ce.gov.br/pages/consultaChaveAcesso.jsf"},
	"DF":       {"DF", "Distrito Federal", "53", "BET", "http://dec.fazenda.df.gov.br/ConsultarNFCe.aspx", "http://dec.fazenda.df.gov.br/ConsultarNFCe.aspx", "http://dec.fazenda.df.gov.br/NFCE/", "http://dec.fazenda.df.gov.br/NFCE/"},
	"ES":       {"ES", "Espírito Santo", "32", "BET", "http://homologacao.sefaz.es.gov.br/ConsultaNFCe/qrcode.aspx", "http://app.sefaz.es.gov.br/ConsultaNFCe/qrcode.aspx", "http://homologacao.sefaz.es.gov.br/ConsultaNFCe", "http://app.sefaz.es.gov.br/ConsultaNFCe"},
	"GO":       {"GO", "Goiás", "52", "BET", "http://homolog.sefaz.go.gov.br/nfeweb/sites/nfce/danfeNFCe", "http://nfe.sefaz.go.gov.br/nfeweb/sites/nfce/danfeNFCe", "http://homolog.sefaz.go.gov.br/nfeweb/sites/nfce/danfeNFCe", "http://nfe.sefaz.go.gov.br/nfeweb/sites/nfce/danfeNFCe"},
	"MA":       {"MA", "Maranhão", "21", "BET", "http://www.hom.nfce.sefaz.ma.gov.br/portal/consultarNFCe.jsp", "http://www.nfce.sefaz.ma.gov.br/portal/consultarNFCe.jsp", "http://www.hom.nfce.sefaz.ma.gov.br/portal/consultarNFCe.jsp", "http://www.sefaz.ma.gov.br/nfce/consulta/"},
	"MT":       {"MT", "Mato Grosso", "51", "IET", "http://homologacao.sefaz.mt.gov.br/nfce/consultanfce", "http://www.sefaz.mt.gov.br/nfce/consultanfce", "http://homologacao.sefaz.mt.gov.br/nfce/consultanfce", "http://www.sefaz.mt.gov.br/nfce/consultanfce"},
	"MS":       {"MS", "Mato Grosso do Sul", "50", "IET", "http://www.dfe.ms.gov.br/nfce/qrcode", "http://www.dfe.ms.gov.br/nfce/qrcode", "http://www.dfe.ms.gov.br/nfce", "http://www.dfe.ms.gov.br/nfce"},
	"MG":       {"MG", "Minas Gerais", "31", "BET", "https://portalsped.fazenda.mg.gov.br/portalnfce/sistema/qrcode.xhtml", "https://portalsped.fazenda.mg.gov.br/portalnfce/sistema/qrcode.xhtml", "https://hportalsped.fazenda.mg.gov.br/portalnfce", "https://portalsped.fazenda.mg.gov.br/portalnfce"},
	"PA":       {"PA", "Pará", "15", "BET", "https://appnfc.sefa.pa.gov.br/portal-homologacao/view/consultas/nfce/nfceForm.seam", "https://appnfc.sefa.pa.gov.br/portal/view/consultas/nfce/nfceForm.seam", "https://appnfc.sefa.pa.gov.br/portal/view/consultas/nfce/consultanfce.seam", "https://appnfc.sefa.pa.gov.br/portal/view/consultas/nfce/consultanfce.seam"},
	"PB":       {"PB", "Paraíba", "25", "BET", "http://www.receita.pb.gov.br/nfcehom", "http://www.receita.pb.gov.br/nfce", "www.receita.pb.gov.br/nfcehom", "www.receita.pb.gov.br/nfce/consulta"},
	"PR":       {"PR", "Paraná", "41", "BET", "http://www.fazenda.pr.gov.br/nfce/qrcode", "http://www.fazenda.pr.gov.br/nfce/qrcode", "http://www.fazenda.pr.gov.br", "http://www.fazenda.pr.gov.br"},
	"PE":       {"PE", "Pernambuco", "26", "BET", "http://nfcehomolog.sefaz.pe.gov.br/nfce-web/consultarNFCe", "http://nfce.sefaz.pe.gov.br/nfce-web/consultarNFCe", "nfce.sefaz.pe.gov.br/nfce/consulta", "nfce.sefaz.pe.gov.br/nfce/consulta"},
	"PI":       {"PI", "Piauí", "22", "BET", "http://www.sefaz.pi.gov.br/nfce/qrcode", "http://www.sefaz.pi.gov.br/nfce/qrcode", "http://www.sefaz.pi.gov.br/nfce/consulta", "http://www.sefaz.pi.gov.br/nfce/consulta"},
	"RJ":       {"RJ", "Rio de Janeiro", "33", "BET", "http://www4.fazenda.rj.gov.br/consultaNFCe/QRCode", "http://www4.fazenda.rj.gov.br/consultaNFCe/QRCode", "www.nfce.fazenda.rj.gov.br/consulta", "www.nfce.fazenda.rj.gov.br/consulta"},
	"RN":       {"RN", "Rio Grande do Norte", "24", "BET", "http://hom.nfce.set.rn.gov.br/consultarNFCe.aspx", "http://nfce.set.rn.gov.br/consultarNFCe.aspx", "http://nfce.set.rn.gov.br/consultarNFCe.aspx", "http://nfce.set.rn.gov.br/consultarNFCe.aspx"},
	"RS":       {"RS", "Rio Grande do Sul", "43", "BET", "https://www.sefaz.rs.gov.br/NFCE/NFCE-COM.aspx", "https://www.sefaz.rs.gov.br/NFCE/NFCE-COM.aspx", "https://www.sefaz.rs.gov.br/NFCE/NFCE-COM.aspx", "https://www.sefaz.rs.gov.br/NFCE/NFCE-COM.aspx"},
	"RO":       {"RO", "Rondônia", "11", "IET", "http://www.nfce.sefin.ro.gov.br/consultanfce/consulta.jsp", "http://www.nfce.sefin.ro.gov.br/consultanfce/consulta.jsp", "http://www.nfce.sefin.ro.gov.br", "http://www.nfce.sefin.ro.gov.br"},
	"RR":       {"RR", "Roraima", "14", "BET", "http://200.174.88.103:8080/nfce/servlet/qrcode", "https://www.sefaz.rr.gov.br/nfce/servlet/qrcode", "http://200.174.88.103:8080/nfce/servlet/wp_consulta_nfce", "https://www.sefaz.rr.gov.br/nfce/servlet/wp_consulta_nfce"},
	"SC":       {"SC", "Santa Catarina", "42", "BET", "https://hom.sat.sef.sc.gov.br/nfce/consulta", "https://sat.sef.sc.gov.br/nfce/consulta", "https://hom.sat.sef.sc.gov.br/nfce/consulta", "https://sat.sef.sc.gov.br/nfce/consulta"},
	"SE":       {"SE", "Sergipe", "28", "BET", "http://www.hom.nfe.se.gov.br/portal/consultarNFCe.jsp", "http://www.nfce.se.gov.br/portal/consultarNFCe.jsp", "http://www.hom.nfe.se.gov.br/portal/portalNoticias.jsp", "http://www.nfce.se.gov.br/portal/portalNoticias.jsp"},
	"TO":       {"TO", "Tocantins", "17", "BET", "http://homologacao.sefaz.to.gov.br/nfce/qrcode", "http://www.sefaz.to.gov.br/nfce/qrcode", "http://homologacao.sefaz.to.gov.br/nfce/consulta.jsf", "http://www.sefaz.to.gov.br/nfce/consulta.jsf"},
	"NACIONAL": {"NC", "Nacional", "90", "", "", "", "", ""},
	"RFB":      {"RFB", "RFB", "91", "", "", "", "", ""},
	"EX":       {"EX", "Exterior", "99", "", "", "", "", ""},
}

// RecuperarUFPorCodigo retorna a unidade federativa com base no código
func RecuperarUFPorCodigo(codigo string) (UnidadeFederativa, bool) {
	uf, encontrado := Estados[codigo]
	return uf, encontrado
}

// RecuperarUFPorCodigoIBGE retorna a unidade federativa com base no código IBGE
func RecuperarUFPorCodigoIBGE(codigo string) (UnidadeFederativa, bool) {
	for _, uf := range Estados {
		if uf.CodigoIbge == codigo {
			return uf, true
		}
	}
	return UnidadeFederativa{}, false
}
