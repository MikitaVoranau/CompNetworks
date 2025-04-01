package internal

const (
	MessageTypeText = 1
	MessageTypeFile = 2
)

type Message struct {
	Type    int
	Content []byte
	Name    string
}
