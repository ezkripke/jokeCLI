package model

// Twopart struct represents jokeAPI response when joke is a "two parter"
type Twopart struct {
	Category string `json:"category"`
	Type     string `json:"type"`
	Setup    string `json:"setup"`
	Delivery string `json:"delivery"`
	Flags    struct {
		Nsfw      bool `json:"nsfw"`
		Religious bool `json:"religious"`
		Political bool `json:"political"`
		Racist    bool `json:"racist"`
		Sexist    bool `json:"sexist"`
	} `json:"flags"`
	ID    int  `json:"id"`
	Error bool `json:"error"`
}

// Single struct represent jokeAPI response when joke is a "one parter"
type Single struct {
	Category string `json:"category"`
	Type     string `json:"type"`
	Joke     string `json:"joke"`
	Flags    struct {
		Nsfw      bool `json:"nsfw"`
		Religious bool `json:"religious"`
		Political bool `json:"political"`
		Racist    bool `json:"racist"`
		Sexist    bool `json:"sexist"`
	} `json:"flags"`
	ID    int  `json:"id"`
	Error bool `json:"error"`
}

// Output represents a joke as it will be output to user
type Output struct {
	Fulljoke string `json:"fulljoke"`
	ID       int    `json:"id"`
}

// Output (overloaded) creates an Output struct from an API response
func (s Single) Output() Output {
	return Output{
		Fulljoke: s.Joke,
		ID:       s.ID,
	}
}

// Output (overloaded) creates an Output struct from an API response
func (t Twopart) Output() Output {
	return Output{
		Fulljoke: t.Setup + "\n" + t.Delivery,
		ID:       t.ID,
	}
}
