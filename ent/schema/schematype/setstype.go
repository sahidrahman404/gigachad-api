package schematype

type Set struct {
	Set  int     `json:"set,omitempty"`
	Reps *int    `json:"reps,omitempty"`
	Kg   *int    `json:"kg,omitempty"`
	Time *string `json:"time,omitempty"`
	Km   *int    `json:"km,omitempty"`
}

type Sets []Set
