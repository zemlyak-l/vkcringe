package bot

import "github.com/zemlyak-l/vkgottle/object"

type Message object.NewMessage

func (m *Message) Answer(send *object.Message) error {
	return nil
}

func (m *Message) Reply(send *object.Message) error {
	return nil
}
