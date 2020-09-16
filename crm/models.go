package crm

type (
	CrmData struct {
		CampaignId int    `json:"campaign_id" validate:"required"`
		DropsId    string `json:"drops_id" validate:"required"`
		Activity   string `json:"activity" validate:"required"`
		Created    int64  `json:"created"`
	}
)
