package main

type ItemOwner struct {
	Reputation   int    `json:"reputation"`
	UserId       int    `json:"user_id"`
	UserType     string `json:"user_type"`
	ProfileImage string `json:"profile_image"`
	DisplayName  string `json:"display_name"`
	Link         string `json:"link"`
}

type SearchItem struct {
	Tags             []string  `json:"tags"`
	Owner            ItemOwner `json:"owner"`
	IsAnswered       bool      `json:"is_answered"`
	ViewCount        int       `json:"view_count"`
	ClosedDate       int       `json:"closed_date"`
	AnswerCount      int       `json:"answer_count"`
	Score            int       `json:"score"`
	LastActivityDate int       `json:"last_activity_date"`
	CreationDate     int       `json:"creation_date"`
	LastEditDate     int       `json:"last_edit_date"`
	QuestionId       int       `json:"question_id"`
	Link             string    `json:"link"`
	ClosedReason     string    `json:"closed_reason"`
	Title            string    `json:"title"`
}

type SearchAPIResponse struct {
	Items          []SearchItem `json:"items"`
	HasMore        bool         `json:"has_more"`
	QuotaMax       int          `json:"quota_max"`
	QuotaRemaining int          `json:"quoata_remaining"`
}
