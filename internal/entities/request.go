package entities

type Request struct {
	PeriodStart         string `json:"period_start"`
	PeriodEnd           string `json:"period_end"`
	PeriodKey           string `json:"period_key"`
	IndicatorToMoID     string `json:"indicator_to_mo_id"`
	IndicatorToMoFactID string `json:"indicator_to_mo_fact_id"`
	Value               string `json:"value"`
	FactTime            string `json:"fact_time"`
	IsPlan              string `json:"is_plan"`
	AuthUserID          string `json:"auth_user_id"`
	Comment             string `json:"comment"`
}

func NewRequest(
	periodStart string,
	periodEnd string,
	periodKey string,
	indicatorToMoID string,
	indicatorToMoFactID string,
	value string,
	factTime string,
	isPlan string,
	authUserID string,
	comment string) *Request {
	return &Request{
		PeriodStart:         periodStart,
		PeriodEnd:           periodEnd,
		PeriodKey:           periodKey,
		IndicatorToMoID:     indicatorToMoID,
		IndicatorToMoFactID: indicatorToMoFactID,
		Value:               value,
		FactTime:            factTime,
		IsPlan:              isPlan,
		AuthUserID:          authUserID,
		Comment:             comment,
	}
}
