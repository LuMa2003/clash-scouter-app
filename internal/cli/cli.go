package cli

import (
	"fmt"
	"github.com/LuMa2003/clash-scouter-app/internal/clash"
	"github.com/manifoldco/promptui"
	"github.com/toqueteos/webbrowser"
	"strings"
)

func Cli(summoners *[]clash.Summoner, region string) {
	var sourceUrl string
	var summonerString strings.Builder

	uggMap := map[string]string{
		"BR":   "BR1",
		"EUNE": "EUN1",
		"EUW":  "EUW1",
		"JP":   "JP1",
		"KR":   "KR",
		"LAN":  "LA1",
		"LAS":  "LA2",
		"NA":   "NA1",
		"OCE":  "OC1",
		"TR":   "TR1",
		"RU":   "RU",
		"PH":   "PH2",
		"SG":   "SG2",
		"TH":   "TH2",
		"TW":   "TW2",
		"VN":   "VN2",
	}

	prompt := promptui.Select{
		Label:    "Select Source",
		HideHelp: true,
		Items:    []string{"U.GG", "OP.GG", "LEAGUEOFGRAPHS.COM"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if result == "OP.GG" {
		for i, summoner := range *summoners {
			if i == 0 {
				summonerString.WriteString(fmt.Sprint(summoner.Name, "-", summoner.Tag))
			}
			summonerString.WriteString(fmt.Sprint("%2C", summoner.Name, "%23", summoner.Tag))
		}
	}
	for i, summoner := range *summoners {
		if i == 0 {
			summonerString.WriteString(fmt.Sprint(summoner.Name, "-", summoner.Tag))
		}
		summonerString.WriteString(fmt.Sprint("%2C", summoner.Name, "-", summoner.Tag))
	}
	switch result {
	//https://u.gg/multisearch?summoners=name-region%2Cname-region%2Cname-region&region={REGION FRÅN RIOT API SE DISCORD}
	//Har enbart EUW nu då jag inte har kollat upp hur man ska på bästa sätt lägga till siffra med riot api värdet e.g(EUW1, PH2)
	case "U.GG":
		sourceUrl = fmt.Sprintf("https://u.gg/multisearch?summoners=%v&region=%v", summonerString.String(), uggMap[region])

		//https://www.op.gg/multisearch/{REGION}?summoners=name#region%2Cname#region%2Cname#region
	case "OP.GG":
		sourceUrl = fmt.Sprintf("https://www.op.gg/multisearch/%v?summoners=%v", region, summonerString.String())

		//HAR POROFESSOR
		//https://porofessor.gg/pregame/{REGION}name-region%2Cname-region%2Cname-region
	case "LEAGUEOFGRAPHS.COM":
		sourceUrl = fmt.Sprintf("https://porofessor.gg/pregame/%v/%v", strings.ToLower(region), summonerString.String())
	}
	webbrowser.Open(sourceUrl)

}
