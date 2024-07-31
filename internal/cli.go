package cli

import (
	"fmt"

	"github.com/manifoldco/promptui"
    "github.com/toqueteos/webbrowser"
)

func cli() {
	prompt := promptui.Select{
		Label: "Select Source",
        HideHelp: true,
        Size: 10,
		Items: []string{"U.GG", "OP.GG", "LEAGUEOFGRAPHS.COM"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

    switch result{

    case"U.GG":
        webbrowser.Open("https://u.gg")
    case"OP.GG":
        webbrowser.Open("https://op.gg")
    case"LEAGUEOFGRAPHS.COM":
        webbrowser.Open("https://LEAGUEOFGRAPHS.COM")
    }

}