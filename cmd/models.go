// Copyright © 2017 Envoy Project <hello@envoyproject.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"time"
)

//arenaIndicator is an object in the arena
type arenaIndicator struct {
	ID            int64      `json:"id,string,omitempty"`
	IndicatorType string     `json:"indicatortype,omitempty"`
	Indicator     string     `json:"indicator,omitempty"`
	IndicatorInt  int64      `json:"indicatorint,omitempty"`
	ThreatType    string     `json:"threattype,omitempty"`
	Score         float64    `json:"score,omitempty"`
	AddedOn       *time.Time `json:"addedon,omitempty"`
	LastUpdated   *time.Time `json:"lastupdated,omitempty"`
	ExpiredOn     *time.Time `json:"expiredon,omitempty"`
	Status        string     `json:"status,omitempty"`
	Description   string     `json:"description,omitempty"`
	Tags          string     `json:"tags,omitempty"`
}

//arenaDescriptor describes the attack and the indicator used
type arenaDescriptor struct {
	ID               int64      `json:"id,omitempty"`
	EventID          int64      `json:"eventid,omitempty"`
	UserID           int64      `json:"userid,omitempty"`
	ArenaIndicatorID int64      `json:"arenaindicatorid,omitempty"`
	Source           string     `json:"source,omitempty"`
	SourceURI        string     `json:"sourceuri,omitempty"`
	AttackType       string     `json:"attacktype,omitempty"`
	ThreatType       string     `json:"threattype,omitempty"`
	AddedOn          *time.Time `json:"addedon,omitempty"`
	LastUpdated      *time.Time `json:"lastupdated,omitempty"`
	ExpiredOn        *time.Time `json:"expiredon,omitempty"`
	Status           string     `json:"status,omitempty"`
	Description      string     `json:"description,omitempty"`
	Score            float64    `json:"score,omitempty"`
	ShareLevel       string     `json:"sharelevel,omitempty"`
	ReviewStatus     string     `json:"reviewstatus,omitempty"`
	Tags             string     `json:"tags,omitempty"`
	PrivacyMethod    string     `json:"privacymethod,omitempty"`
	PrivacyMembers   string     `json:"privacymembers,omitempty"`
}
