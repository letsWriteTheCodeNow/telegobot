package keyboard

type IncomingMessage struct {
	Ok     bool `json:"ok"`
	Result []struct {
		Update_id int `json:"update_id"`
		Message   struct {
			Message_id int `json:"message_id"`
			From       struct {
				Id int `json:"id"`
			} `json:"from"`
			Text    string `json:"text"`
			Contact struct {
				Phone_number string `json:"phone_number"`
			} `json:"contact"`
		} `json:"message"`
		HandlerFunction struct {
			Name string `json:"data"`
		} `json:"callback_query"`
	} `json:"result"`
}

type KeyboardButton struct {
	Text            string `json:"text"`
	Request_contact bool   `json:"request_contact"`
}

type InlineKeyboardButton struct {
	Text          string `json:"text"`
	Callback_data string `json:"callback_data"`
}

type Keyboard struct {
	KeyboardButtonArray [][]KeyboardButton `json:"keyboard"`
	Resize_keyboard     bool               `json:"resize_keyboard"`
	One_time_keyboard   bool               `json:"one_time_keyboard"`
}

type InlineKeyboard struct {
	ButtonType [][]InlineKeyboardButton `json:"inline_keyboard"`
}
