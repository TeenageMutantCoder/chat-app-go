package resources

type Message struct {
	Content string
	Author  string
}

var messages = []Message{}

func GetMessages() []Message {
	return messages
}

func CreateMessage(content string, author string) {
	messages = append(messages, Message{Content: content, Author: author})
}
