package events

type OutputSolutionEvent struct {
	Created     float64 `json:"created"`
	Duration    float64 `json:"duration"`
	Exitcode    int     `json:"exitcode"`
	Root        string  `json:"root"`
	Environment struct {
	} `json:"environment"`
	Summary struct {
		Passed    int `json:"passed"`
		Total     int `json:"total"`
		Collected int `json:"collected"`
	} `json:"summary"`
	Collectors []struct {
		Nodeid  string `json:"nodeid"`
		Outcome string `json:"outcome"`
		Result  []struct {
			Nodeid string `json:"nodeid"`
			Type   string `json:"type"`
		} `json:"result"`
	} `json:"collectors"`
	Tests []struct {
		Nodeid   string   `json:"nodeid"`
		Lineno   int      `json:"lineno"`
		Outcome  string   `json:"outcome"`
		Keywords []string `json:"keywords"`
		Setup    struct {
			Duration float64 `json:"duration"`
			Outcome  string  `json:"outcome"`
		} `json:"setup"`
		Call struct {
			Duration float64 `json:"duration"`
			Outcome  string  `json:"outcome"`
		} `json:"call"`
		Teardown struct {
			Duration float64 `json:"duration"`
			Outcome  string  `json:"outcome"`
		} `json:"teardown"`
	} `json:"tests"`
}
