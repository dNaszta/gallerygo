package gallery

type Image struct {
	Descriptions map[string]string	`json:"descriptions"`
	Original ImageProperty			`json:"original"`
	Instances []ImageProperty		`json:"instances"`
}
