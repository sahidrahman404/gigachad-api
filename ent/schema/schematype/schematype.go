package schematype

type Set struct {
	Reps *int    `json:"reps,omitempty"`
	Kg   *int    `json:"kg,omitempty"`
	Time *string `json:"time,omitempty"`
	Km   *int    `json:"km,omitempty"`
}

type Style struct {
	Width       *string `json:"width"`
	Height      *string `json:"height"`
	MaxWidth    *string `json:"max-width"`
	MaxHeight   *string `json:"max-height"`
	AspectRatio *string `json:"aspectRatio"`
}

type Image struct {
	Style
	Height        *float64 `json:"height"`
	AspectRatio   *float64 `json:"aspectRatio"`
	Width         *int     `json:"width"`
	Layout        string   `json:"layout"`
	ObjectFit     string   `json:"objectFit"`
	Src           string   `json:"src"`
	SrcSet        string   `json:"srcset"`
	Filename      string   `json:"filename"`
	Loading       *string  `json:"loading"`
	FetchPriority *string  `json:"fetchpriority"`
	Decoding      *string  `json:"decoding"`
	Alt           *string  `json:"alt"`
	Role          *string  `json:"role"`
	Sizes         string   `json:"sizes"`
	BreakPoints   []int    `json:"breakpoints"`
	Priority      bool     `json:"priority"`
}
