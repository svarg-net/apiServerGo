package model

type TypeMenu struct {
	Href string `json:"href"`
	Text string `json:"text"`
	//"title": "About",
	//"target": "",
	//"type": "_blank"
}
type TypeMeta struct {
	Description string `json:"description"`
	Title       string `json:"title"`
}
type TypeBLock struct {
	Position int
	Tag      string
	Text     string
}
type TypePage struct {
	Link   string      `json:"link"`
	Meta   TypeMeta    `json:"meta"`
	Blocks []TypeBLock `json:"block"`
}
type TypeData struct {
	Menu  []TypeMenu `json:"menu"`
	Pages []TypePage `json:"pages"`
}
