package main

import (

	"github.com/Oleg-MBO/Radio-en-Esperanto-in-tg/plagins/parserERadio"
	"fmt"
	"github.com/Oleg-MBO/Radio-en-Esperanto-in-tg/botdb"
)



func main()  {

	db := botdb.Initdb("db.sqlite3")
	defer db.Close()



	lastpodkasts := parserERadio.GetLastPodcasts()
	fmt.Println("got lastpodkasts ", len(lastpodkasts))
	for _, q := range lastpodkasts {
		//fmt.Println("newchannel")
		//
		//fmt.Println(botdb.Debug().NewRecord(*q.Channel))// => returns `true` as primary key is blank
		//
		//fmt.Println(botdb.Debug().Create(q.Channel))
		//
		//fmt.Println(botdb.Debug().NewRecord(q.Channel))// => return `false` after `user` created


		fmt.Println("new ", q.ChannelName)
		fmt.Println(db.NewRecord(q) )// => returns `true` as primary key is blank

		fmt.Println(db.Create(&q) )


		fmt.Println( db.NewRecord(q) )// => return `false` after `user` created
	}




}