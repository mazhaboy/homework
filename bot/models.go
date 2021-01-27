package main

type Update struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	Chat Chat   `json:"chat"`
	From User   `json:"from"`
	Text string `json:"text"`
	Poll Poll   `json:"poll"`
}

type Chat struct {
	ChatID int `json:"id"`
}

type User struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"username"`
}

type RestResponse struct {
	Result []Update `json:"result"`
}

type BotMessage struct {
	ChatID int    `json:"chat_id"`
	Text   string `json:"text"`
}
type Poll struct {
	PollID              int             `json:"poll_id"`
	TotalVoterCount     int             `json:"total_voter_count"`
	ExplanationEntities []MessageEntity `json:"explanation_entities"`
	options             []PollOption    `json:"options"`
	PollAnswer          PollAnswer      `json:"poll_answer"`
}
type MessageEntity struct {
	Type     string `json:"type"`
	User     User   `json:"text_mention"`
	Language string `json:"language"`
}
type PollOption struct {
	Text string `json:"text"`
}
type PollAnswer struct {
	PollID int  `json:"poll_id"`
	User   User `json:"user"`
}
