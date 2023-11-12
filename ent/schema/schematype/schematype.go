package schematype

type Set struct {
	Reps *int    `json:"reps,omitempty"`
	Kg   *int    `json:"kg,omitempty"`
	Time *string `json:"time,omitempty"`
	Km   *int    `json:"km,omitempty"`
	Set  *int    `json:"set,omitempty"`
}

type Image struct {
	Height      *float64 `json:"height"`
	AspectRatio *float64 `json:"aspectRatio"`
	Width       *int     `json:"width"`
	Layout      string   `json:"layout"`
	Priority    string   `json:"priority"`
	ObjectFit   string   `json:"objectFit"`
	Background  string   `json:"background"`
	Src         string   `json:"src"`
	BreakPoints []int    `json:"breakpoints"`
}
