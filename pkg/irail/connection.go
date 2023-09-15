package irail

type ConnectionContainer struct {
	Connections []Connection `json:"connection"`
}

type Stationinfo struct {
	ID           string `json:"id"`
	LocationX    string `json:"locationX"`
	LocationY    string `json:"locationY"`
	Standardname string `json:"standardname"`
	Name         string `json:"name"`
}

type Vehicleinfo struct {
	Name      string `json:"name"`
	Shortname string `json:"shortname"`
	ID        string `json:"@id"`
}

type Platforminfo struct {
	Name   string `json:"name"`
	Normal string `json:"normal"`
}

type Direction struct {
	Name string `json:"name"`
}

type Stop struct {
	ID                     string      `json:"id"`
	Station                string      `json:"station"`
	Stationinfo            Stationinfo `json:"stationinfo"`
	Time                   string      `json:"time"`
	Delay                  string      `json:"delay"`
	Canceled               string      `json:"canceled"`
	DepartureDelay         string      `json:"departureDelay"`
	DepartureCanceled      string      `json:"departureCanceled"`
	ScheduledDepartureTime string      `json:"scheduledDepartureTime"`
	ArrivalDelay           string      `json:"arrivalDelay"`
	ArrivalCanceled        string      `json:"arrivalCanceled"`
	IsExtraStop            string      `json:"isExtraStop"`
	ScheduledArrivalTime   string      `json:"scheduledArrivalTime"`
	DepartureConnection    string      `json:"departureConnection"`
}

type Stops struct {
	Number string `json:"number"`
	Stop   []Stop `json:"stop"`
}

type Alert struct {
	ID        string `json:"id"`
	Header    string `json:"header"`
	Lead      string `json:"lead"`
	Link      string `json:"link"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type Alerts struct {
	Number string  `json:"number"`
	Alert  []Alert `json:"alert"`
}

type Departure struct {
	Delay               string       `json:"delay"`
	Station             string       `json:"station"`
	Stationinfo         Stationinfo  `json:"stationinfo"`
	Time                string       `json:"time"`
	Vehicle             string       `json:"vehicle"`
	Vehicleinfo         Vehicleinfo  `json:"vehicleinfo"`
	Platform            string       `json:"platform"`
	Platforminfo        Platforminfo `json:"platforminfo"`
	Left                string       `json:"left"`
	Canceled            string       `json:"canceled"`
	Direction           Direction    `json:"direction"`
	Stops               Stops        `json:"stops"`
	Alerts              Alerts       `json:"alerts"`
	Walking             string       `json:"walking"`
	DepartureConnection string       `json:"departureConnection"`
}

type Arrival struct {
	Delay        string       `json:"delay"`
	Station      string       `json:"station"`
	Stationinfo  Stationinfo  `json:"stationinfo"`
	Time         string       `json:"time"`
	Vehicle      string       `json:"vehicle"`
	Vehicleinfo  Vehicleinfo  `json:"vehicleinfo"`
	Platform     string       `json:"platform"`
	Platforminfo Platforminfo `json:"platforminfo"`
	Arrived      string       `json:"arrived"`
	Canceled     string       `json:"canceled"`
	Walking      string       `json:"walking"`
	Direction    Direction    `json:"direction"`
}

type Via struct {
	ID          string      `json:"id"`
	Arrival     Arrival     `json:"arrival"`
	Departure   Departure   `json:"departure"`
	TimeBetween string      `json:"timeBetween"`
	Station     string      `json:"station"`
	Stationinfo Stationinfo `json:"stationinfo"`
	Vehicle     string      `json:"vehicle"`
	Direction   Direction   `json:"direction"`
}

type Vias struct {
	Number string `json:"number"`
	Via    []Via  `json:"via"`
}

type Connection struct {
	ID        string    `json:"id"`
	Departure Departure `json:"departure"`
	Arrival   Arrival   `json:"arrival"`
	Duration  string    `json:"duration"`
	Alerts    Alerts    `json:"alerts"`
	Vias      Vias      `json:"vias"`
}

func GetConnections(lang string, from string, to string) []Connection {
	req := Request[ConnectionContainer]{
		Path: "connections",
		Parameters: map[string]string{
			"lang": lang,
			"from": from,
			"to":   to,
		},
	}

	container := ConnectionContainer{}
	err := req.Do(&container)

	if err != nil {
		panic(err)
	}

	return container.Connections
}
