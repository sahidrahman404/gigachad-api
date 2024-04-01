package schematype

import "strconv"

type Set struct {
	Reps     *int    `json:"reps,omitempty"`
	Kg       *int    `json:"kg,omitempty"`
	Duration *string `json:"duration,omitempty"`
	Km       *int    `json:"km,omitempty"`
}

type SetDisplay struct {
	Kg       string `json:"kg,omitempty"`
	Reps     string `json:"reps,omitempty"`
	Duration string `json:"duration,omitempty"`
	Km       string `json:"km,omitempty"`
}

func (s Set) Display() SetDisplay {
	sd := SetDisplay{}
	empty := "-"
	if s.Reps == nil {
		sd.Reps = empty
	}
	if s.Kg == nil {
		sd.Kg = empty
	}
	if s.Duration == nil {
		sd.Duration = empty
	}
	if s.Km == nil {
		sd.Km = empty
	}

	if s.Reps != nil {
		sd.Reps = strconv.Itoa(*s.Reps)
	}
	if s.Kg != nil {
		sd.Kg = strconv.Itoa(*s.Kg)
	}
	if s.Duration != nil {
		sd.Duration = *s.Duration
	}
	if s.Km != nil {
		sd.Km = strconv.Itoa(*s.Km)
	}
	return sd
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

type Reminder struct {
	Day    int
	Hour   int
	Minute int
	Second int
}
