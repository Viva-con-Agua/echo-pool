package crm

type (
	CrmData struct {
		CampaignId int    `json:"campaign_id" validate:"required"`
		DropsId    string `json:"drops_id" validate:"required"`
		Activity   string `json:"activity" validate:"required"`
		Country    string `json:"country,omitempty"`
		Created    int64  `json:"created"`
	}
	CrmEmail struct {
		Email string `json:"email"`
		Link  string `json:"link"`
	}
	CrmEmailBody struct {
		CrmData CrmData  `json:"crm_data"`
		Mail    CrmEmail `json:"mail"`
	}
)
