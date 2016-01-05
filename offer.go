package gopaymill

type Offer struct {
	ID                string `json:id`
	AppId             string `json:app_id`
	CreatedAt         int    `json:created_at`
	UpdatedAt         int    `json:updated_at`
	Name              string `json:name`
	Amount            int    `json:amount`
	Interval          string `json:interval`
	TrialPeriodDays   int    `json:trial_period_days`
	SubscriptionCount count  `json:subscription_count`
}
