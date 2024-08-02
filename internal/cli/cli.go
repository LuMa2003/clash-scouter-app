package cli

import (
	"fmt"
	"strings"
	"github.com/LuMa2003/clash-scouter-app/internal/clash"
	"github.com/manifoldco/promptui"
	"github.com/toqueteos/webbrowser"
)

func Cli(summoners *[]clash.Summoner, region string) {
	var summoner_string strings.Builder
	
	var source_url string = ""
	prompt := promptui.Select{
		Label:    "Select Source",
		HideHelp: true,
		Size:     10,
		Items:    []string{"U.GG", "OP.GG", "LEAGUEOFGRAPHS.COM"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if result == "OP.GG" {
		for itteration, summoner := range *summoners{
			if itteration == 0 {
				summoner_string.WriteString(fmt.Sprint(summoner.Name, "-",summoner.Tag))	
			}
			summoner_string.WriteString(fmt.Sprint("%2C", summoner.Name, "%23",summoner.Tag))
		}
	}
	for itteration, summoner := range *summoners{
		if itteration == 0 {
			summoner_string.WriteString(fmt.Sprint(summoner.Name, "-",summoner.Tag))	
		}
		summoner_string.WriteString(fmt.Sprint("%2C", summoner.Name, "-",summoner.Tag))
	}
	switch result {
		//https://u.gg/multisearch?summoners=name-region%2Cname-region%2Cname-region&region={REGION FRÅN RIOT API SE DISCORD}
		//Har enbart EUW nu då jag inte har kollat upp hur man ska på bästa sätt lägga till siffra med riot api värdet e.g(EUW1, PH2)
	case "U.GG":
		source_url = fmt.Sprint("https://u.gg/multisearch?summoners=", summoner_string.String(), "&region=euw1",)

		//https://www.op.gg/multisearch/{REGION}?summoners=name#region%2Cname#region%2Cname#region
	case "OP.GG":
		source_url = fmt.Sprint("https://www.op.gg/multisearch/", region,"?summoners=", summoner_string.String())
		
		//HAR POROFESSOR
		//https://porofessor.gg/pregame/{REGION}name-region%2Cname-region%2Cname-region
	case "LEAGUEOFGRAPHS.COM":
		source_url = fmt.Sprint("https://porofessor.gg/pregame/", strings.ToLower(region), "/", summoner_string.String())
	}
	webbrowser.Open(source_url)

}
