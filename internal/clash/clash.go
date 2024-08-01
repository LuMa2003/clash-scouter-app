package clash

import (
	"github.com/LuMa2003/clash-scouter-app/pkg/lcu"
	"github.com/tidwall/gjson"
)

type Summoner struct {
	Name string
	Tag  string
}

func ClashOpponent(connInfo *lcu.ConnInfo) ([]Summoner, error) {

	data, err := lcu.LCU(lcu.Request{
		Conn:     connInfo,
		Method:   "GET",
		Endpoint: "/lol-clash/v1/tournament-summary",
		Body:     nil,
	})
	if err != nil {
		return nil, err
	}
	summary := gjson.GetManyBytes(data, "0.rosterId", "0.bracketId")
	me := summary[0].Int()
	bracketId := summary[1].String()

	data2, err := lcu.LCU(lcu.Request{
		Conn:     connInfo,
		Method:   "GET",
		Endpoint: "/lol-clash/v1/bracket/" + bracketId,
		Body:     nil,
	})
	if err != nil {
		return nil, err
	}

	bracket := gjson.GetManyBytes(data2, `matches.#(status=="UPCOMING")#.rosterId1`, `matches.#(status=="UPCOMING")#.rosterId2`)
	roster1 := &bracket[0]
	roster2 := &bracket[1]

	var opponent string

	for index := range roster1.Indexes {
		if roster1.Array()[index].Int() == me {
			opponent = roster2.Array()[index].String()
		} else if roster2.Array()[index].Int() == me {
			opponent = roster1.Array()[index].String()
		}
	}

	data3, err := lcu.LCU(lcu.Request{
		Conn:     connInfo,
		Method:   "GET",
		Endpoint: "/lol-clash/v1/roster/" + opponent,
		Body:     nil,
	})
	if err != nil {
		return nil, err
	}
	roster := gjson.GetBytes(data3, "members.#.summonerId").String()

	data4, err := lcu.LCU(lcu.Request{
		Conn:     connInfo,
		Method:   "GET",
		Endpoint: "/lol-summoner/v2/summoners/?ids=" + roster,
		Body:     nil,
	})
	if err != nil {
		return nil, err
	}
	summonersJson := gjson.GetManyBytes(data4, "#.gameName", "#.tagLine")
	name := &summonersJson[0]
	tag := &summonersJson[1]

	var summoners []Summoner

	for index := range name.Indexes {
		summoners = append(summoners, Summoner{
			Name: name.Array()[index].String(),
			Tag:  tag.Array()[index].String(),
		})
	}
	return summoners, nil
}
