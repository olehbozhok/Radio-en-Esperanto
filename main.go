package main

import (
	"github.com/Oleg-MBO/Radio-en-Esperanto-in-tg/botdb"
	_ "github.com/Oleg-MBO/Radio-en-Esperanto-in-tg/plagins/parserERadio"

	"github.com/Oleg-MBO/Radio-en-Esperanto-in-tg/tgApiHelper"

	tg "github.com/bot-api/telegram"
	"golang.org/x/net/context"
	"log"
	"fmt"
	"time"
)

func main() {

	db := botdb.Initdb("db.sqlite3")
	defer db.Close()
	db.LogMode(false)


	api := tgApiHelper.InitApi("")
	api.Debug(true)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()



	for n,p := range botdb.GetNewPodcasts() {
		//getAndSendPodcast(n, p,  api, ctx )
		fmt.Println("\n\n",n)
		//time.Sleep(time.Second * 10)
		if cs := p.CanSendTotelegram(); cs.Status==true {
			getAndSendPodcast(n, p, api, ctx)
		}
		time.Sleep(time.Second * 10)

	}

}


func getAndSendPodcast(n int,p botdb.PodcastType,api *tg.API, ctx context.Context ) error {

	//api.SendMessage(ctx, tg.MessageCfg{
	//	BaseMessage:tg.BaseMessage{
	//		BaseChat: tg.BaseChat{
	//			ChannelUsername: "@tqwerdtfghjkjmdasgfnhmgf",
	//		},
	//	},
	//	Text: fmt.Sprintf("%v\n%s\n%s", n, p.Href,p.Name()),
	//})


	IF, err := p.InputFile()
	log.Printf("IF OK:  %#v, nil: %#v", IF.OK,err)
	if  IF.OK { // тут какая то бага с err!= nil  не делает запрос хз почему
		log.Printf("IF OK:  %#v", IF.OK)

		m, err := api.SendAudio(ctx, tg.AudioCfg{
			BaseFile: tg.BaseFile{
				BaseMessage:tg.BaseMessage{
					BaseChat: tg.BaseChat{
						ChannelUsername: "@tqwerdtfghjkjmdasgfnhmgf",
					},
				},
				//FileID: p.Href,
				InputFile: IF,
			},
			Title: p.Name(),
		})
		if err != nil {
			log.Printf("\n\nERROR:  %#v", err)
			return err
		} else {
			log.Printf("\n\nmsg:  %#v", m.Audio.FileID)
		}
	}
	if err!= nil{
		log.Printf("error send to tg: %s",err.Error())
	}

	log.Print("exit")
	return nil
}