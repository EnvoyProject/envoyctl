package cmd

import (
	"time"
)

//arenaIndicator is an object in the arena
type arenaIndicator struct {
	//indicator
	ID            int64  `json:"id,string,omitempty"`     //the id of the indicator
	IndicatorType string `json:"indicatortype,omitempty"` //can be: url, hostname,uri. this is the TYPE(ex: IPv4)
	Indicator     string `json:"indicator,omitempty"`     //the actual value. this is the INDICATOR (ex:127.0.0.1)
	IndicatorInt  int64  `json:"indicatorint,omitempty"`  //this is to allow to store IPs also
	//subjective fields that are global
	ThreatType  string     `json:"threattype,omitempty"` //threattype calculated
	Score       float64    `json:"score,omitempty"`      //this is the confidence
	AddedOn     *time.Time `json:"addedon,omitempty"`    //time it was added or updated
	LastUpdated *time.Time `json:"lastupdated,omitempty"`
	ExpiredOn   *time.Time `json:"expiredon,omitempty"`
	Status      string     `json:"status,omitempty"`      //what is the status: UNKNOWN, NON_MALICIOUS, SUSPICIOUS, MALICIOUS
	Description string     `json:"description,omitempty"` //a general description
	Tags        string     `json:"tags,omitempty"`        //tags are labels different from the key. THREAT TYPES (ex:"MALICIOUS_IP")
}

//arenaDescriptor describes the attack and the indicator used
type arenaDescriptor struct {
	ID               int64      `json:"id,omitempty"` //unique identifier of the descriptor
	EventID          int64      `json:"eventid,omitempty"`
	UserID           int64      `json:"userid,omitempty"`           //the id of the user that submited
	ArenaIndicatorID int64      `json:"arenaindicatorid,omitempty"` //the id that is refering
	Source           string     `json:"source,omitempty"`           //who submited it -> from Provider
	SourceURI        string     `json:"sourceuri,omitempty"`        //where we can get some more context -> from Provider
	AttackType       string     `json:"attacktype,omitempty"`       //attack type that is associated with -> from Token
	ThreatType       string     `json:"threattype,omitempty"`       //the threat type that is associated with -> from Token
	AddedOn          *time.Time `json:"addedon,omitempty"`          //from processing
	LastUpdated      *time.Time `json:"lastupdated,omitempty"`      //from processing
	ExpiredOn        *time.Time `json:"expiredon,omitempty"`        // -> from Token
	Status           string     `json:"status,omitempty"`           //if it is malicious or not
	Description      string     `json:"description,omitempty"`      //a short description
	Score            float64    `json:"score,omitempty"`            //confidence in the value [0..10]
	ShareLevel       string     `json:"sharelevel,omitempty"`       //list of values
	ReviewStatus     string     `json:"reviewstatus,omitempty"`     //specifies the human or machine that has reviewed the information
	Tags             string     `json:"tags,omitempty"`             //a source can add some tags, to identify a specific attack(ex: MALICIOUS_IP)
	PrivacyMethod    string     `json:"privacymethod,omitempty"`
	PrivacyMembers   string     `json:"privacymembers,omitempty"`
}
