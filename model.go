package models

//Create Struct
type Meeting struct {
	Id                string `json:"id"`
	Title             string `json:"title"`
	Participants      string `json:"participants"`
	StartTime         string `json:"starttime"`
	EndTime           string `json:"endtime"`
	CreationTimestamp string `json:"timestamp"`
}

type Participant struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	RSVP  string `json:"RSVP"`
}
