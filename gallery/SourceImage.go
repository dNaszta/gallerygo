package gallery

type SourceImage struct {
	Descriptions map[string]string	`json:"descriptions"`
	Source string					`json:"source"`
}
