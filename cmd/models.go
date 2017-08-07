package cmd

//arenaIndicator is an object in the arena
type arenaIndicator struct {
	//indicator
	ID            int64  `json:"id,string,omitempty"`                                   //the id of the indicator
	IndicatorType string `json:"indicatortype,omitempty"` //can be: url, hostname,uri. this is the TYPE(ex: IPv4)
	Indicator     string `json:"indicator,omitempty"`    //the actual value. this is the INDICATOR (ex:127.0.0.1)
	IndicatorInt  int64  `json:"indicatorint,omitempty"`                                //this is to allow to store IPs also
	//subjective fields that are global
	ThreatType  string     `json:"threattype,omitempty"` //threattype calculated
	Score       float64    `json:"score,omitempty"`         //this is the confidence
	AddedOn     *time.Time `json:"addedon,omitempty"`                                  //time it was added or updated
	LastUpdated *time.Time `json:"lastupdated,omitempty"`
	ExpiredOn   *time.Time `json:"expiredon,omitempty"`
	Status      string     `json:"status,omitempty"`       //what is the status: UNKNOWN, NON_MALICIOUS, SUSPICIOUS, MALICIOUS
	Description string     `json:"description,omitempty"` //a general description
	Tags        string     `json:"tags,omitempty"`        //tags are labels different from the key. THREAT TYPES (ex:"MALICIOUS_IP")
}

