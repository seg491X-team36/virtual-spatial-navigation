package experiment

/* experiment ExperimentStatus struct
if an experiment has 3 rounds the state should go:

start experiment -> {"roundInProgress": false, "roundNumber": 0}
start round -> {"roundInProgress": true, "roundNumber": 1}
stop round -> {"roundInProgress": false, "roundNumber": 1}
start round -> {"roundInProgress": true, "roundNumber": 2}
stop round -> {"roundInProgress": false, "roundNumber": 2}
start round -> {"roundInProgress": true, "roundNumber": 3}
stop round -> {"roundInProgress": false, "roundNumber": 3}
*/
type ExperimentStatus struct {
	RoundInProgress bool `json:"roundInProgress"`
	RoundNumber     int  `json:"roundNumber"`
	RoundsTotal     int  `json:"roundsTotal"`
}

func (s ExperimentStatus) Done() bool {
	return s.RoundNumber == s.RoundsTotal && !s.RoundInProgress
}
