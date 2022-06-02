package polling

import (
	"log"
	"net/url"
	"strconv"

	"github.com/zemlyak-l/vkgottle/api"
	"github.com/zemlyak-l/vkgottle/object"
)

const (
	baseLongpollWait = 15
)

type Longpoll struct {
	Api    *api.Api
	Params url.Values
	Server string
	Routes *Routes
}

func NewLongpoll(api *api.Api, groupID int) (*Longpoll, error) {
	// Getting longpoll server and configs
	r := object.ResponseInit{}
	message := object.GetServerMessage{GroupID: groupID}
	if err := api.GroupsGetLongPollServer(message, &r); err != nil {
		return nil, err
	}

	// Convert LongpollWait value to string
	strWait := strconv.Itoa(baseLongpollWait)

	// Create url values
	urlParams := url.Values{}
	urlParams.Add("act", "a_check")
	urlParams.Add("ts", r.Response.Ts)
	urlParams.Add("key", r.Response.Key)
	urlParams.Add("wait", strWait)

	// Init Longpoll struct
	return &Longpoll{
		Api:    api,
		Params: urlParams,
		Server: r.Response.Server,
		Routes: &Routes{},
	}, nil
}

func (lp *Longpoll) Request() (*object.LongpollResponse, error) {
	// Create request to longpoll
	r := &object.LongpollResponse{}
	err := lp.Api.Post(lp.Server, []byte(lp.Params.Encode()), &r)
	if err != nil {
		return nil, err
	}

	if r.Ts == "" {
		return nil, nil
	}

	lp.Params.Set("ts", r.Ts)

	return r, nil
}

func (lp *Longpoll) ListenNewEvents() {
	// Listen new longpoll events
	for {
		// Create request
		event, err := lp.Request()
		if err != nil {
			log.Fatal(err)
		}
		if event == nil {
			continue
		}
		// Check updates
		for _, update := range event.Updates {
			if err := lp.CheckEvent(update); err != nil {
				log.Fatal(err)
			}
		}
	}
}
