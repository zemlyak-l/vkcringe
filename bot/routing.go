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
	sentFunc HandlerFunc
	isSent   bool

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
		isSent:        false,
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

	if ok && !t.isSent {
		t.isSent = true
		t.sentFunc = f
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

	if bot.TextRoutes.isSent {
		bot.TextRoutes.isSent = false
		go bot.TextRoutes.sentFunc(message)
	} else {
		go bot.TextRoutes.AllHandler(message)
	}
}
