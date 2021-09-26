package simpleTcpMessage

func typeSupported(t byte) bool {
	return t <= TypeMask
}

type Message struct {
	fields map[byte][]byte
}

func (msg *Message) AppendField(_type byte, p []byte) bool {
	if typeSupported(_type) && len(p) > 0 {
		if _, ok := msg.fields[_type]; !ok {
			msg.fields[_type] = make([]byte, 0, len(p))
		}
		msg.fields[_type] = append(msg.fields[_type], p...)
		return true
	}
	return false
}

func NewMessage() *Message {
	return &Message{
		fields: make(map[byte][]byte),
	}
}

func (msg *Message) GetField(_type byte) ([]byte, bool) {
	if typeSupported(_type) {
		if _, ok := msg.fields[_type]; ok {
			return msg.fields[_type], true
		}
	}
	return nil, false
}

func (msg *Message) RemoveFieldIfExist(_type byte) bool {
	if typeSupported(_type) {
		if _, ok := msg.fields[_type]; ok {
			delete(msg.fields, _type)
			return true
		}
	}
	return false
}
