package resources

type (
	// ResourcePredicateData is a container used to query predicate information
	ResourcePredicateData struct {
		Schema []struct {
			Name      string `json:"predicate"`
			Type      string
			Index     bool
			Reverse   string
			Tokenizer []string
			List      bool
			Count     bool
			Upsert    string
			Lang      bool
		}
	}

	ResourceTypeData struct {
		Types []struct {
			Name   string
			Fields []struct {
				Name string
			}
		}
	}
)
