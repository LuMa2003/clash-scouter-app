package cli

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/toqueteos/webbrowser"
)


func cli(summoner string) {
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

	switch result {

	case "U.GG":
		source_url = "https://u.gg/lol/profile/euw1/" + summoner +"-euw/overview"
	case "OP.GG":
		source_url = "https://www.op.gg/summoners/euw/" + summoner + "-EUW"
	case "LEAGUEOFGRAPHS.COM":
		source_url = "https://www.leagueofgraphs.com/summoner/euw/" + summoner + "-EUW"	
	}
	webbrowser.Open(source_url)
	
	
}
