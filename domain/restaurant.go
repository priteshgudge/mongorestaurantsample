package domain

type Restaurant struct {
	DBID         ID `json:"id" bson:"_id"`
	Name         string
	Address      string
	AddressLine2 string
	URL          string
	Outcode      string
	Postcode     string
	Rating       float32
	TypeOfFood   string
}
