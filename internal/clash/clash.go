package clash

import (
	"github.com/LuMa2003/clash-scouter-app/pkg/lcu"
	"github.com/tidwall/gjson"
	"github.com/manifoldco/promptui"
	"fmt"
	"os"

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


	for { 
	data2, err := lcu.LCU(lcu.Request{
		Conn:     connInfo,
		Method:   "GET",
		Endpoint: "/lol-clash/v1/bracket/" + bracketId,
		Body:     nil,
	})
	
	if err == nil{

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
	//ERROR HÄR PGA Ingen opponenet
		data3, err := lcu.LCU(lcu.Request{
			Conn:     connInfo,
			Method:   "GET",
			Endpoint: "/lol-clash/v1/roster/" + opponent,
			Body:     nil,
		})

		if err == nil {	
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
		
		//Om Opponent inte hittas
		if err != nil{
			prompt := promptui.Select{
				Label:    "Opponent not found",
				HideHelp: true,
				Size:     10,
				Items:    []string{"Retry","Exit"},
			}
			_, result, err := prompt.Run()
		
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				continue
			}
			if result == "Exit" {
				fmt.Println("Exiting...")
				os.Exit(0)
			}
		}

	}
	//Om Clash Bracket inte hittas
	if err != nil{
		prompt := promptui.Select{
			Label:    "Clashbracket not found",
			HideHelp: true,
			Size:     10,
			Items:    []string{"Retry","Exit"},
		}
		_, result, err := prompt.Run()
	
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			continue
		}
		if result == "Exit" {
			fmt.Println("Exiting...")
			os.Exit(0)
		}
	}

	}

}










