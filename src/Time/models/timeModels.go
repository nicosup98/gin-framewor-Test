package timemodels

type TimeByUtcDTO struct {
	Date     string `json:"date"`
	TimeZone int64  `json:"timeZone"`
}
