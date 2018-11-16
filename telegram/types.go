package telegram

type Message struct {
	MessageID int                    `json:"message_id"`
	From      map[string]interface{} `json:"from"`
	Chat      map[string]interface{} `json:"chat"`
	Date      int                    `json:"date"`
	Text      string                 `json:"text"`
}

type TelegramResult struct {
	UpdateID int       `json:"update_id"`
	Messages []Message `json:"message"`
}

type TelegramResponse struct {
	Success bool           `json:"ok"`
	Result  TelegramResult `json:"result"`
}
