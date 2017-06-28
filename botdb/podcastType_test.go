package botdb

import (
	"testing"
	"os"


)

func TestPodcastTypeSave(t *testing.T)  {
	db := Initdb(os.Getenv("dbString"))
	defer db.Close()
	var n int64 = 0
	p := PodcastType{
		RawDate: "2000-01-01",
		ChannelName: "test",
		Description: "testo messago",
		Href: "http://example.com/podkasto-2017-06-02.mp3",
		TgMsgIdDescription: &n,
		TgFromIdChat: &n,
		TgMsgFileId: &n,

	}

	p.Delete()
	defer func() {
		if err :=p.Delete(); err!= nil {
			t.Errorf("Error delete podcast: %!v",err)
		}
	}()
	if err :=p.Save(); err!= nil {
		t.Errorf("Error save podcast: %!v",err)
	}
	if p.IsUnique() {
		t.Errorf("Error now is unique")
		t.Fail()
	}
	p.ID = 0
	if err :=p.Save(); err== nil {
		t.Errorf("Error save podcast 2, must be unsaved: %!v",err)
	}
}