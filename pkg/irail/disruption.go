package irail

type DisruptionContainer struct {
	Timestamp   string       `json:"timestamp"`
	Disruptions []Disruption `json:"disturbance"`
}

type Disruption struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Type        string `json:"type"`
	Timestamp   string `json:"timestamp"`
}

func GetDisruptions(lang string) []Disruption {
	req := Request[DisruptionContainer]{
		Path: "disturbances",
		Parameters: map[string]string{
			"lang": lang,
		},
	}

	disrutions := DisruptionContainer{}
	err := req.Do(&disrutions)

	if err != nil {
		panic(err)
	}

	return disrutions.Disruptions
}
