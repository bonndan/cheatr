package main

import (
	"fmt"
	"os"
)

func main() {

	topics := initTopics()

	if len(os.Args) == 1 {
		fmt.Println("No topic given. Select one below:")
		for key, _ := range topics {
			fmt.Printf("%s \n", key)
		}
		return
	}

	topicname := os.Args[1]
	topic, exists := topics[topicname]
	if !exists {
		fmt.Println("Topic does not exist:", topicname)
		return
	}

	topic.print()
}

type Topic struct {
	title  string
	cheats map[string]string
}

func (topic Topic) print() {
	fmt.Println(topic.title)
	fmt.Println("--------")

	for key, desc := range topic.cheats {
		fmt.Printf("%-12s", key)
		fmt.Printf("%s\n", desc)
	}
}

func initTopics() map[string]Topic {

	topics := make(map[string]Topic)

	vim := Topic{
		title: "vim",
		cheats: map[string]string {
			",e": "explore",
			"^p": "open file",
			"^w q": "close pane",
			"yyp": "copy line",
			">>": "indent",
			"<<": "unindent",
			"0": "line start",
			"/": "search",
			"n": "search next",
			"*": "hilite",
			":%s/_/_/g": "replace _ -> _",
		},
	}
	topics["vim"] = vim

	w3m := Topic{
		title: "w3m",
		cheats: map[string]string {
			"B": "back",
			"U": "open url",
			"H": "help",
			"T": "new Tab",
			"^t": "list tabs",
			"ESC-a": "Bookmark",
		},
	}
	topics["w3m"] = w3m

	tmux := Topic {
		title: "tmux",
		cheats: map[string]string {
			"^b c": "new",
			"^b w": "list",
			"^b \"": "split horizontal",
			"^b %": "split vertical",
			"^b o": "to next split",
			"^b n": "next",
			"^b p": "previous",
			"^b hjkl": "move in direction to split",
            "^b &": "kill window",
		},
	}
	topics["tmux"] = tmux

	mutt := Topic {
		title: "mutt",
		cheats: map[string]string {
			"c": "change box",
			"m": "new message",
            "y": "send",
            "i": "quit edit mode",
            "r": "reply",
            "p": "menu for gpg",
            "$": "sync",
            "g": "reply all",
            "v": "view message attachments, s to save",
		},
	}
	topics["mutt"] = mutt

	various := Topic {
		title: "various",
		cheats: map[string]string {
			"tig": "git (d=diff)",
			"wego": "weather (github.com/schachmat/wego)",
			"gcalcli": "calendar",
		},
	}
	topics["var"] = various

	cmus := Topic {
		title: "cmus",
		cheats: map[string]string {
			"3": "playlist",
			"-/+": "volume",
			"c": "pause",
			"x": "play",
		},
	}
	topics["cmus"] = cmus

	return topics
}
