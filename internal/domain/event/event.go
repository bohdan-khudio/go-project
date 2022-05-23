package event

type Event struct {
	Id        int64   `json:"id" db:"id,omitempty"`
	Title     string  `json:"title" db:"title"`
	ShortDesc string  `json:"short_desc" db:"short_desc"`
	Desc      string  `json:"desc" db:"description"`
	Long      float64 `json:"long" db:"long"`
	Lat       float64 `json:"lat" db:"lat"`
}
