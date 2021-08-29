package keyboard

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

func (k *Keyboard) ByDefault() {
	k.Resize_keyboard = true
	k.One_time_keyboard = true
}

func (k *Keyboard) AddButtonRequestContact(text string) {

	var newKeyboardButton KeyboardButton
	var newKeyboardButtonArray []KeyboardButton

	newKeyboardButton.Text = text
	newKeyboardButton.Request_contact = true

	newKeyboardButtonArray = append(newKeyboardButtonArray, newKeyboardButton)

	k.KeyboardButtonArray = append(k.KeyboardButtonArray, newKeyboardButtonArray)
	// 		var keyboardButton keyboard.KeyboardButton
	// 		var keyboardButtonArray []keyboard.KeyboardButton

	// 		keyboardButton.Text = "Отправить номер"
	// 		keyboardButton.Request_contact = true

	// 		keyboardButtonArray = append(keyboardButtonArray, keyboardButton)
}

type InlineKeyboard struct {
	ButtonType [][]InlineKeyboardButton `json:"inline_keyboard"`
}

func AddKeyboardButton(text string, request_contact bool) {

}
