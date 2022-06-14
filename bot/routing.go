package bot

import (
	"strings"

	"github.com/zemlyak-l/vkcringe/object"
)

const ChatPeerID = 2000000000

type HandlerFunc func(object.NewMessage)
type HandlerFuncMap map[string]HandlerFunc

func EmptyHandlerFunc(object.NewMessage) {}

type TextRoutes struct {
	sent HandlerFunc

	AllRoutes     map[string]HandlerFunc
	PrivateRoutes map[string]HandlerFunc
	ChatRoutes    map[string]HandlerFunc
	AllHandler    HandlerFunc
}

func NewRoutes() *TextRoutes {
	return &TextRoutes{
		AllRoutes:     make(HandlerFuncMap),
		PrivateRoutes: make(HandlerFuncMap),
		ChatRoutes:    make(HandlerFuncMap),
		AllHandler:    EmptyHandlerFunc,
	}
}

func (t *TextRoutes) checkRoute(routeType string, cmd string) {
	var f HandlerFunc
	var ok bool

	switch routeType {
	case "all":
		f, ok = t.AllRoutes[cmd]
	case "private":
		f, ok = t.PrivateRoutes[cmd]
	case "chat":
		f, ok = t.ChatRoutes[cmd]
	}

	if ok && t.sent == nil {
		t.sent = f
	}
}

func (bot *Bot) OnAll(f HandlerFunc) {
	bot.TextRoutes.AllHandler = f
}

func (bot *Bot) OnMessage(text string, f HandlerFunc) {
	bot.TextRoutes.AllRoutes[text] = f
}

func (bot *Bot) OnPrivateMessage(text string, f HandlerFunc) {
	bot.TextRoutes.PrivateRoutes[text] = f
}

func (bot *Bot) OnChatMessage(text string, f HandlerFunc) {
	bot.TextRoutes.ChatRoutes[text] = f
}

func (bot *Bot) messageHandler(message object.NewMessage) {
	cmdArgs := strings.Split(message.Text, " ")
	cmdName := cmdArgs[0]
	if len(cmdArgs) != 1 {
		message.CmdArgs = cmdArgs[1:]
	}

	bot.TextRoutes.checkRoute("all", cmdName)
	if message.PeerID < ChatPeerID {
		bot.TextRoutes.checkRoute("private", cmdName)
	} else {
		bot.TextRoutes.checkRoute("chat", cmdName)
	}

	if bot.TextRoutes.sent != nil {
		go bot.TextRoutes.sent(message)
	} else {
		go bot.TextRoutes.AllHandler(message)
	}
}
