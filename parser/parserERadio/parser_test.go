package parserERadio

import (
	"testing"
)

func TestGetLastPodcasts(t *testing.T) {
	pd, err := GetLastPodcasts()
	if err != nil {
		t.Fatal(err)
	}
	if len(pd) < 25 {
		t.Error("got podcasts < 25 must be more")
	}
}

//func TestParserERadio(t *testing.T) {
//	podcasts, err := parseLastPodcasts(ioutil.NopCloser(bytes.NewReader([]byte(rawHtml))))
//	if err != nil {
//		t.Error("Parse fall")
//	}
//	if len(podcastsFromRawHtml) != len(podcasts) {
//		t.Error("length not the same")
//	}
//	for i := 0; i < len(podcastsFromRawHtml); i++ {
//		if podcastsFromRawHtml[i].RawDate != podcasts[i].RawDate {
//			t.Error("mistake in parse data")
//		}
//		if podcastsFromRawHtml[i].ChannelName != podcasts[i].ChannelName {
//			t.Error("mistake in parse ChannelName")
//		}
//		if podcastsFromRawHtml[i].ChannelName != podcasts[i].ChannelName {
//			t.Error("mistake in parse ChannelName")
//		}
//		if podcastsFromRawHtml[i].Href != podcasts[i].Href {
//			t.Error("mistake in parse Href")
//		}
//	}
//}

//var podcastsFromRawHtml = []old_botdb.PodcastType{old_botdb.PodcastType{RawDate: "2017-06-03", ChannelName: "Muzaiko", Description: "Aleksks rakontas pri sia sperto kiel volontulo de TEJO Francesco Maurelli rakontas pri Kosmo (JES2013)", Href: "http://muzaiko.info/public/podkasto/podkasto-2017-06-02.mp3"}, old_botdb.PodcastType{RawDate: "2017-06-02", ChannelName: "Pola Retradio", Description: "La 655-a E_elsendo el la 02.06.2017 ĉe www.pola-retradio.org -Nia ĉefinformo lastmomenta koncernas membriĝon de Pollando kiel nekonstanta membro en la Sekureckonsilo de UN. – Nia ĉi-vendreda komenca kulturkroniko informas pri la Tagoj de la hungara komponisto, lingvisto, etnografo Zoltan Kodaly en Katowice kaj pri klopodoj repopularigi la Polan Muzeon en la svisa Rapperwil. – La […]", Href: "http://feedproxy.google.com/~r/retradio/~5/H2y7lfxNxOg/RetRadio_02.06.2017_vendredo.mp3"}, old_botdb.PodcastType{RawDate: "2017-06-02", ChannelName: "Ĉina Radio Internacia", Description: "Ĉina Radio Internacia. Fokusoj de la Semajno (Jun. 02).", Href: "http://mod.cri.cn/espe/r170602.m4a"}, old_botdb.PodcastType{RawDate: "2017-06-02", ChannelName: "Muzaiko", Description: "Raporto el la 4a JES - tria parto", Href: "http://muzaiko.info/public/podkasto/podkasto-2017-06-01.mp3"}, old_botdb.PodcastType{RawDate: "2017-06-01", ChannelName: "Muzaiko", Description: "Mirejo intervjuas Kingslim EDAH - 2 Elektitaj novaĵoj de El Popola Ĉinio por tiu ĉi semajno. Pliajn informojn kaj aliajn novaĵojn vi trovas en www.espero.com.cn Interalie - pozitivaj novaĵoj, elsendo 32", Href: "http://muzaiko.info/public/podkasto/podkasto-2017-05-31.mp3"}, old_botdb.PodcastType{RawDate: "2017-05-31", ChannelName: "Muzaiko", Description: "Probal rakontas pri kredo - dua parto Ari intervjuas Yordan Fonseca, unu el la ĉefaj organizantoj de la 2a Havana Festivalo Scenejo Aleks Kadar intervjuas Richard Delamore pri Esperanto TV", Href: "http://muzaiko.info/public/podkasto/podkasto-2017-05-30.mp3"}, old_botdb.PodcastType{RawDate: "2017-05-30", ChannelName: "Pola Retradio", Description: "La 654-a E-elsendo ĉe www.pola-retradio.org el la 30.05.2017: – Nia antaŭmikrofona gasto estas hodiaŭ nia redakcia kolego, d-ro Maciej Jaskot dividanta impresojn pri tri majaj konferencoj, kiujn li partoprenis en majo en la polaj Bjalistoko kaj Lublino kaj Sofio. Ili i.a. spegulas la problemojn de la lingvoinstruado en Pollando enhavante ankaŭ informojn interesajn por E-instruantoj. […]", Href: "http://feedproxy.google.com/~r/retradio/~5/I5zmo2IDjNs/RetRadio_30.05.2017_mardo.mp3"}, old_botdb.PodcastType{RawDate: "2017-05-30", ChannelName: "Muzaiko", Description: "Probal rakontas pri kredo - unua parto SATeH prezentas por vi: La prizonulo Junulara vespero en Bonaero (UK2014)", Href: "http://muzaiko.info/public/podkasto/podkasto-2017-05-29.mp3"}, old_botdb.PodcastType{RawDate: "2017-05-29", ChannelName: "3ZZZ en Esperanto", Description: "Legado: Laszlo el Sennaciulo “ Pri la 1a de majo en Francio” de Petro Levi Franciska:el la retejo de la Legio de Bona volo “ Intermiksiĝo en la mondo estas neevitebla” de Paiva Netto Parolado: Marcel Duoblaj prepozicioj. Legado: Laszlo el la revuo Esperanto de aprilo “ Ne materiaj heredaĵoj “ de Mireille Grosjean Kanto: […]", Href: "http://www.melburno.org.au/3ZZZradio/mp3/2017-05-29.3ZZZ.radio.mp3"}, old_botdb.PodcastType{RawDate: "2017-05-29", ChannelName: "Muzaiko", Description: "Pri siaj spertoj rilate al foraj rilatoj rakontas al ni István ERTL (dua parto) Aleks Kadar intervjuas Richard Delamore pri Esperanto TV", Href: "http://muzaiko.info/public/podkasto/podkasto-2017-05-28.mp3"}, old_botdb.PodcastType{RawDate: "2017-05-28", ChannelName: "Muzaiko", Description: "Pri siaj spertoj rilate al foraj rilatoj rakontas al ni István ERTL (unua parto) Mirejo pri geedziĝfesto Koreio Japanio", Href: "http://muzaiko.info/public/podkasto/podkasto-2017-05-27.mp3"}, old_botdb.PodcastType{RawDate: "2017-05-28", ChannelName: "Radio Havano Kubo", Description: "Radio Havano Kubo. Amika voĉo kiu trairas la mondon. Esperanto+28-05-17.", Href: "http://www.ameriko.org/eo/audio/download/879/Esperanto+28-05-17.mp3"}, old_botdb.PodcastType{RawDate: "2017-05-26", ChannelName: "Pola Retradio", Description: "La 653-a E_elsendo el la 26.05.2017 ĉe www.pola-retradio.org – Nian hodiaŭan felietonon pri koninda polo ni dediĉas al la pola komponisto, reprezentanto de Juna Pollando, Feliks Nowowiejski, kies 140a naskiĝdatreveno pasis nunjare. La felietonon akompanas muzikcitaĵoj el jutubo el liaj komponaĵoj: marŝo „Sub la standardo de paco”, la patriota, himneca kanto „Rota” kaj el la […]", Href: "http://feedproxy.google.com/~r/retradio/~5/UfiqqWb4zAA/RetRadio_26.05.2017_vendredo.mp3"}, old_botdb.PodcastType{RawDate: "2017-05-25", ChannelName: "Kernpunkto", Description: "Tiu ĉi epizodo estas enkonduka epizodo al la nuna stato de Eŭropa Unio. Ni planas okazigi scenejan debaton dum la Germana Esperanto Kongreso en Freiburg semajnon poste por diskuti la eblojn pretigi Eŭropan Union por la estonteco. En tiu ĉi epizodo ni parolas pri la komencoj de Eŭropa Unio kaj la ĝisnunaj atingoj.", Href: "https://kern.punkto.info/podlove/file/1458/s/feed/c/mp3/kp117-euxropa-unio.mp3"}, old_botdb.PodcastType{RawDate: "2017-05-23", ChannelName: "Pola Retradio", Description: "La 652-a E-elsendo ĉe www.pola-retradio.org el la 23.05.2017: – La komenca aktualaĵo rilatas al la laŭvica terorisma atenco en Eŭropo, ĉi-foje en Manĉestro, Britio kaŭzinte la pereon de almenaŭ 22 personoj, inkluzive infanojn kaj preskaŭ 60 vunditojn. La sekvaj informoj koncernas esperatan rikolton en la polaj vitejoj kaj bierproduktadon en Pollando. – La sciencbultenaj temoj […]", Href: "http://feedproxy.google.com/~r/retradio/~5/5Lo04OMjftA/RetRadio_23.05.2017_mardo.mp3"}, old_botdb.PodcastType{RawDate: "2017-05-22", ChannelName: "3ZZZ en Esperanto", Description: "Kanto:el la kompaktdisko Ni renkontiĝas de la plej Popularaj Kantoj de Sovetiaj kaj Rusaj Esperantistoj “ Sezonoj” Legado: Laszlo el Sennaciulo “ Fidel Castro Ruz” de Vilhelmo Lutermano Heather El Esperanto de aprilo 2017 “ La morto en Venecio” Recenzo de Christoph Klawe Kanto: el la kompaktdisko Urbano de Inicialoj DC “ La kontraŭsenco” Legado: […]", Href: "http://www.melburno.org.au/3ZZZradio/mp3/2017-05-22.3ZZZ.radio.mp3"}, old_botdb.PodcastType{RawDate: "2017-05-19", ChannelName: "Pola Retradio", Description: "La 651-a E_elsendo el la 19.05.2017 ĉe www.pola-retradio.org – Nia antaŭmikrofona gasto hodiaŭ estas la renkontita dum la aprila Malferma Tago en Roterdamo, Stela Besenyei-Merger, la redaktorino de la presita versio de la Pasporta Servo 2017. – La komenca ampleksa kulturkroniko informs pri Festivalo de Wajda-filmoj en Novjorko, pri la 110-jariĝo de la varsovia Publika […]", Href: "http://feedproxy.google.com/~r/retradio/~5/S7qdrp2dTso/RetRadio_19.05.2017_vendredo.mp3"}, old_botdb.PodcastType{RawDate: "2017-05-16", ChannelName: "Pola Retradio", Description: "La 650-a E-elsendo ĉe www.pola-retradio.org el la 16.05.2017: – En la hodiaŭa elsendo ni prezentas interparolon kun d-ro Vilmos Benczik marĝene de du liaj artikoloj el la eldonita pasintjare en Budapeŝto libro „Pri natureco kaj artefariteco de lingvoj kaj aliaj studoj”. La okazo por la antaŭmikrofona renkontiĝo estis la februaraj Interlingvistikaj Studoj en Poznano. – […]", Href: "http://feedproxy.google.com/~r/retradio/~5/93Joduf9gKw/RetRadio_16.05.2017_wtorek.mp3"}, old_botdb.PodcastType{RawDate: "2017-05-15", ChannelName: "Varsovia Vento", Description: "Dum ceremonio de la monumento-tombo de Ludoviko L. Zamenhof Ĉe Fejsbuko ni kreis diskutgrupon, kie aperas freŝaj informoj, fotoj kaj filmetoj - filmetojn vi trovos ankaŭ en laste kreita jutuba kanalo: podkasto viavento. La 123a elsendo – la 1a parto (tempo-daŭro 30:37): Elŝutu podkaston Ho mia kor’ ̵ [...]", Href: "http://www.podkasto.net/wp-content/uploads/2017/05/170512VVE123P1.mp3"}, old_botdb.PodcastType{RawDate: "2017-05-15", ChannelName: "3ZZZ en Esperanto", Description: "Legado: Heather el Retradio “ Paŝtaĵo por homoj aŭ nutraĵo” de Andreo Bach Kanto. el Bando Barok’ “ Heredo de la postmilit” de Rafael Milhomen Intervjuo. Franjo Martin pri komuna vivo Legado: Karlo el Monato “ Fremdlingvaj indikiloj provokas proteston” de Cristina Casella Heather : el Monato. “ Kial ne aboli la mortpunon” de Isikawa Takasi […]", Href: "http://www.melburno.org.au/3ZZZradio/mp3/2017-05-15.3ZZZ.radio.mp3"}, old_botdb.PodcastType{RawDate: "2017-05-12", ChannelName: "Kernpunkto", Description: "Ofte oni aŭdis jam pri la elbrula sindromo, kiu koncernas homojn tro laborantajn en la moderna labormondo. Sen gasto ni laike parolas pri la fenomeno, kiu ne estas konsiderata malsano. Tamen multaj homoj je ĝi suferas kaj eĉ memmortigas sin.", Href: "https://kern.punkto.info/podlove/file/1445/s/feed/c/mp3/kp116-elbrula-sindromo.mp3"}, old_botdb.PodcastType{RawDate: "2017-05-12", ChannelName: "Pola Retradio", Description: "La 649-a E_elsendo el la 12.05.2017 ĉe www.pola-retradio.org – Hodiaŭ antaŭ nia mikrofono ni gastigas Veronika Poór, la Ĝeneralan Direktoron de la CO en Rotterdamo, ĉi-foje por interparoli pri la nova retejo de UEA, pri kio kelkajn aktualaĵojn eksciis la partoprenantoj de la lasta Malferma Tago en Rotterdamo. – La komenca kulturkronika aktualaĵo rilatas al […]", Href: "http://feedproxy.google.com/~r/retradio/~5/oa60_k2-SA8/RetRadio_12.05.2017_-vendredo.mp3"}, old_botdb.PodcastType{RawDate: "2017-05-09", ChannelName: "Pola Retradio", Description: "La 648-a E-elsendo ĉe www.pola-retradio.org el la 09.05.2017: – Nia nunsemajna E-elsendo denove enhavas literaturan segmenton, en kiu ĉi-foje aŭdiĝas la jam 55-a fragmento de la romano „Quo Vadis” de Henryk Sienkiewicz en la E-traduko de Lidia Zamenhof. – La komencaj aktualaĵoj rilatas al solenaĵoj omaĝantaj la finon de la 2-a mondmilito en 1945 kaj […]", Href: "http://feedproxy.google.com/~r/retradio/~5/kVYUr8dQN6w/RetRadio_09.05.2017_mardo.mp3"}, old_botdb.PodcastType{RawDate: "2017-05-08", ChannelName: "La Lingvo de Frateco", Description: "La Lingvo de Frateco (Rio). http://www.radioriodejaneiro1400am.com/audiosrrj/ESPERANTOLALINGVOFRATECO/esperantolalingvo290417.", Href: "http://www.radioriodejaneiro1400am.com/audiosrrj/ESPERANTO%20LA%20LINGVO%20FRATECO/esperanto%20la%20lingvo%20290417.mp3"}, old_botdb.PodcastType{RawDate: "2017-05-08", ChannelName: "3ZZZ en Esperanto", Description: "Kanto: el la kompaktdisko Ribela sono de Dolĉamar “ Duope” Legado: Karlo el la revuo Monato “ Nova ĉampiono” de Zaltko Tiŝljar Heather el Retradio de Anton Obendorfer “ 50 plus: sporto tenas menson freŝan” Parolado: Matt “ Nenie pli bone ol hejme” Kanto: el la kompaktdisko Muzikpluvo de Akordo “ Damo, mia kara” Legado: […]", Href: "http://www.melburno.org.au/3ZZZradio/mp3/2017-05-08.3ZZZ.radio.mp3"}, old_botdb.PodcastType{RawDate: "2017-05-08", ChannelName: "Radio Frei", Description: "Radio Frei. 20170507 turint klein.", Href: "http://audio.radio-frei.de/podcast/Esperanto/20170507%20turint%20klein.mp3"}, old_botdb.PodcastType{RawDate: "2017-05-05", ChannelName: "Pola Retradio", Description: "La 647-a E_elsendo el la 05.05.2017 ĉe www.pola-retradio.org – En nia nunsemajna vendreda programo ni gastigas Konstantan Kongresan Sekretarion, Clay Magelhães kuntekste kun la 102-a UK en Seulo, Koreio, al kiu aliĝis jam 1042 personoj el 61 landoj. – La komencaj kulturkronikaj informoj koncernas la decidon pri tio, ke fama skulptaĵo de la pola skulptisto […]", Href: "http://feedproxy.google.com/~r/retradio/~5/FDusEmBKijU/RetRadio_05.05.2017_vendredo.mp3"}, old_botdb.PodcastType{RawDate: "2017-05-03", ChannelName: "Pola Retradio", Description: "La 646-a E-elsendo ĉe www.pola-retradio.org el la 03.05.2017: – Nia nunsemajna sondokumenta parto enhavas la prezenton de Ionel Onet dum la lastsabata (29.04.2017) Malferma Tago en Roterdamo pri la eldonejaj aktualaĵoj. – La komencaj novaĵoj rilatas al la hodiaŭa nacia festo en Pollando, la 226-a datreveno de la proklamo de la 3-Maja Konstitucio, al la […]", Href: "http://feedproxy.google.com/~r/retradio/~5/fZGA1hpMrZw/RetRadio_03.05.2017_merkredo.mp3"}, old_botdb.PodcastType{RawDate: "2017-05-01", ChannelName: "3ZZZ en Esperanto", Description: "Kanto:el la kompaktdisko Svedaj kantoj 2003 ”Amata vi ŝajnas ja roz’ “ Legado: Franciska el la retejo de Labourstart “ Internacia tago de la Romaoj” Parolado: Marcel “ Kiel mi ne intence sed bonŝance trompis kelkiujn” Legado: Laszlo el la revuo Esperanto Pri la historio de la libro de Ulrich Linz La danĝera lingvo Kanto: el […]", Href: "http://www.melburno.org.au/3ZZZradio/mp3/2017-05-01.3ZZZ.radio.mp3"}, old_botdb.PodcastType{RawDate: "2017-04-24", ChannelName: "3ZZZ en Esperanto", Description: "Kanto: el la kompaktdisko Martin kaj la Talpoj “Vivo duras, sed vi molas” Legado: Laszlo el la revuo Esperanto “Novaj radikoj , neologismoj en la Evia lingvo” de Apelete Agbolo. Heather : “Vojaĝo al suda Panĝabio” de Nadipedia el Monato Karlo :el Fervoja Mondo de Jindrich Tomiŝek “ Interesaĵoj pri monto Gotthard tunelo” Kanto : […]", Href: "http://www.melburno.org.au/3ZZZradio/mp3/2017-04-24.3ZZZ.radio.mp3"}}
//
//var rawHtml = `
//<!DOCTYPE html>
//<html lang="eo">
//<head>
//<meta charset="UTF-8">
//<title>ESPERANTO RADIO</title>
//<meta name="ROBOTS" content="NOINDEX, NOFOLLOW" />
//<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1" />
//<link rel="icon" type="image/png" href="/esperanto_radio_icon.png" sizes="192x192">
//<link rel="icon" type="image/x-icon" href="/favicon.ico">
//<meta name="theme-color" content="#009000">
//<style type="text/css">body{color:#FFF;background:black}a{text-decoration:none}a:link{color:#FFF}a:visited{color:#FFF}a:hover{color:yellow}a:active{color:red}</style>
//</head>
//<body>
//<hr />
//<h1><a title="Plena hejmo" href="/">Esperanto-Radio</a></h1>
//<hr />
//<div id="versio_1496513231">
//<a href='http://muzaiko.info/public/podkasto/podkasto-2017-06-02.mp3'>
//<strong>2017-06-03
//Muzaiko</strong>
//<br />
//Aleksks rakontas pri sia sperto kiel volontulo de TEJO Francesco Maurelli rakontas pri Kosmo (JES2013)
//</a><hr />
//
//<a href='http://feedproxy.google.com/~r/retradio/~5/H2y7lfxNxOg/RetRadio_02.06.2017_vendredo.mp3'>
//<strong>2017-06-02
//Pola Retradio</strong>
//<br />
//La 655-a E_elsendo el la 02.06.2017 ĉe www.pola-retradio.org -Nia ĉefinformo lastmomenta koncernas membriĝon de Pollando kiel nekonstanta membro en la Sekureckonsilo de UN. – Nia ĉi-vendreda komenca kulturkroniko informas pri la Tagoj de la hungara komponisto, lingvisto, etnografo Zoltan Kodaly en Katowice kaj pri klopodoj repopularigi la Polan Muzeon en la svisa Rapperwil. – La […]
//</a><hr />
//
//<a href='http://mod.cri.cn/espe/r170602.m4a'>
//<strong>2017-06-02
//Ĉina Radio Internacia</strong>
//<br />
//Ĉina Radio Internacia. Fokusoj de la Semajno (Jun. 02).
//</a><hr />
//
//<a href='http://muzaiko.info/public/podkasto/podkasto-2017-06-01.mp3'>
//<strong>2017-06-02
//Muzaiko</strong>
//<br />
//Raporto el la 4a JES - tria parto
//</a><hr />
//
//<a href='http://muzaiko.info/public/podkasto/podkasto-2017-05-31.mp3'>
//<strong>2017-06-01
//Muzaiko</strong>
//<br />
//Mirejo intervjuas Kingslim EDAH - 2 Elektitaj novaĵoj de El Popola Ĉinio por tiu ĉi semajno. Pliajn informojn kaj aliajn novaĵojn vi trovas en www.espero.com.cn Interalie - pozitivaj novaĵoj, elsendo 32
//</a><hr />
//
//<a href='http://muzaiko.info/public/podkasto/podkasto-2017-05-30.mp3'>
//<strong>2017-05-31
//Muzaiko</strong>
//<br />
//Probal rakontas pri kredo - dua parto Ari intervjuas Yordan Fonseca, unu el la ĉefaj organizantoj de la 2a Havana Festivalo Scenejo Aleks Kadar intervjuas Richard Delamore pri Esperanto TV
//</a><hr />
//
//<a href='http://feedproxy.google.com/~r/retradio/~5/I5zmo2IDjNs/RetRadio_30.05.2017_mardo.mp3'>
//<strong>2017-05-30
//Pola Retradio</strong>
//<br />
//La 654-a E-elsendo ĉe www.pola-retradio.org el la 30.05.2017: – Nia antaŭmikrofona gasto estas hodiaŭ nia redakcia kolego, d-ro Maciej Jaskot dividanta impresojn pri tri majaj konferencoj, kiujn li partoprenis en majo en la polaj Bjalistoko kaj Lublino kaj Sofio. Ili i.a. spegulas la problemojn de la lingvoinstruado en Pollando enhavante ankaŭ informojn interesajn por E-instruantoj. […]
//</a><hr />
//
//<a href='http://muzaiko.info/public/podkasto/podkasto-2017-05-29.mp3'>
//<strong>2017-05-30
//Muzaiko</strong>
//<br />
//Probal rakontas pri kredo - unua parto SATeH prezentas por vi: La prizonulo Junulara vespero en Bonaero (UK2014)
//</a><hr />
//
//<a href='http://www.melburno.org.au/3ZZZradio/mp3/2017-05-29.3ZZZ.radio.mp3'>
//<strong>2017-05-29
//3ZZZ en Esperanto</strong>
//<br />
//Legado: Laszlo el Sennaciulo “ Pri la 1a de majo en Francio” de Petro Levi Franciska:el la retejo de la Legio de Bona volo “ Intermiksiĝo en la mondo estas neevitebla” de Paiva Netto Parolado: Marcel Duoblaj prepozicioj. Legado: Laszlo el la revuo Esperanto de aprilo “ Ne materiaj heredaĵoj “ de Mireille Grosjean Kanto: […]
//</a><hr />
//
//<a href='http://muzaiko.info/public/podkasto/podkasto-2017-05-28.mp3'>
//<strong>2017-05-29
//Muzaiko</strong>
//<br />
//Pri siaj spertoj rilate al foraj rilatoj rakontas al ni István ERTL (dua parto) Aleks Kadar intervjuas Richard Delamore pri Esperanto TV
//</a><hr />
//
//<a href='http://muzaiko.info/public/podkasto/podkasto-2017-05-27.mp3'>
//<strong>2017-05-28
//Muzaiko</strong>
//<br />
//Pri siaj spertoj rilate al foraj rilatoj rakontas al ni István ERTL (unua parto) Mirejo pri geedziĝfesto Koreio Japanio
//</a><hr />
//
//<a href='http://www.ameriko.org/eo/audio/download/879/Esperanto+28-05-17.mp3'>
//<strong>2017-05-28
//Radio Havano Kubo</strong>
//<br />
//Radio Havano Kubo. Amika voĉo kiu trairas la mondon. Esperanto+28-05-17.
//</a><hr />
//
//<a href='http://feedproxy.google.com/~r/retradio/~5/UfiqqWb4zAA/RetRadio_26.05.2017_vendredo.mp3'>
//<strong>2017-05-26
//Pola Retradio</strong>
//<br />
//La 653-a E_elsendo el la 26.05.2017 ĉe www.pola-retradio.org – Nian hodiaŭan felietonon pri koninda polo ni dediĉas al la pola komponisto, reprezentanto de Juna Pollando, Feliks Nowowiejski, kies 140a naskiĝdatreveno pasis nunjare. La felietonon akompanas muzikcitaĵoj el jutubo el liaj komponaĵoj: marŝo „Sub la standardo de paco”, la patriota, himneca kanto „Rota” kaj el la […]
//</a><hr />
//
//<a href='https://kern.punkto.info/podlove/file/1458/s/feed/c/mp3/kp117-euxropa-unio.mp3'>
//<strong>2017-05-25
//Kernpunkto</strong>
//<br />
//Tiu ĉi epizodo estas enkonduka epizodo al la nuna stato de Eŭropa Unio. Ni planas okazigi scenejan debaton dum la Germana Esperanto Kongreso en Freiburg semajnon poste por diskuti la eblojn pretigi Eŭropan Union por la estonteco. En tiu ĉi epizodo ni parolas pri la komencoj de Eŭropa Unio kaj la ĝisnunaj atingoj.
//</a><hr />
//
//<a href='http://feedproxy.google.com/~r/retradio/~5/5Lo04OMjftA/RetRadio_23.05.2017_mardo.mp3'>
//<strong>2017-05-23
//Pola Retradio</strong>
//<br />
//La 652-a E-elsendo ĉe www.pola-retradio.org el la 23.05.2017: – La komenca aktualaĵo rilatas al la laŭvica terorisma atenco en Eŭropo, ĉi-foje en Manĉestro, Britio kaŭzinte la pereon de almenaŭ 22 personoj, inkluzive infanojn kaj preskaŭ 60 vunditojn. La sekvaj informoj koncernas esperatan rikolton en la polaj vitejoj kaj bierproduktadon en Pollando. – La sciencbultenaj temoj […]
//</a><hr />
//
//<a href='http://www.melburno.org.au/3ZZZradio/mp3/2017-05-22.3ZZZ.radio.mp3'>
//<strong>2017-05-22
//3ZZZ en Esperanto</strong>
//<br />
//Kanto:el la kompaktdisko Ni renkontiĝas de la plej Popularaj Kantoj de Sovetiaj kaj Rusaj Esperantistoj “ Sezonoj” Legado: Laszlo el Sennaciulo “ Fidel Castro Ruz” de Vilhelmo Lutermano Heather El Esperanto de aprilo 2017 “ La morto en Venecio” Recenzo de Christoph Klawe Kanto: el la kompaktdisko Urbano de Inicialoj DC “ La kontraŭsenco” Legado: […]
//</a><hr />
//
//<a href='http://feedproxy.google.com/~r/retradio/~5/S7qdrp2dTso/RetRadio_19.05.2017_vendredo.mp3'>
//<strong>2017-05-19
//Pola Retradio</strong>
//<br />
//La 651-a E_elsendo el la 19.05.2017 ĉe www.pola-retradio.org – Nia antaŭmikrofona gasto hodiaŭ estas la renkontita dum la aprila Malferma Tago en Roterdamo, Stela Besenyei-Merger, la redaktorino de la presita versio de la Pasporta Servo 2017. – La komenca ampleksa kulturkroniko informs pri Festivalo de Wajda-filmoj en Novjorko, pri la 110-jariĝo de la varsovia Publika […]
//</a><hr />
//
//<a href='http://feedproxy.google.com/~r/retradio/~5/93Joduf9gKw/RetRadio_16.05.2017_wtorek.mp3'>
//<strong>2017-05-16
//Pola Retradio</strong>
//<br />
//La 650-a E-elsendo ĉe www.pola-retradio.org el la 16.05.2017: – En la hodiaŭa elsendo ni prezentas interparolon kun d-ro Vilmos Benczik marĝene de du liaj artikoloj el la eldonita pasintjare en Budapeŝto libro „Pri natureco kaj artefariteco de lingvoj kaj aliaj studoj”. La okazo por la antaŭmikrofona renkontiĝo estis la februaraj Interlingvistikaj Studoj en Poznano. – […]
//</a><hr />
//
//<a href='http://www.podkasto.net/wp-content/uploads/2017/05/170512VVE123P1.mp3'>
//<strong>2017-05-15
//Varsovia Vento</strong>
//<br />
//Dum ceremonio de la monumento-tombo de Ludoviko L. Zamenhof Ĉe Fejsbuko ni kreis diskutgrupon, kie aperas freŝaj informoj, fotoj kaj filmetoj - filmetojn vi trovos ankaŭ en laste kreita jutuba kanalo: podkasto viavento. La 123a elsendo – la 1a parto (tempo-daŭro 30:37): Elŝutu podkaston Ho mia kor’ &#821 [...]
//</a><hr />
//
//<a href='http://www.melburno.org.au/3ZZZradio/mp3/2017-05-15.3ZZZ.radio.mp3'>
//<strong>2017-05-15
//3ZZZ en Esperanto</strong>
//<br />
//Legado: Heather el Retradio “ Paŝtaĵo por homoj aŭ nutraĵo” de Andreo Bach Kanto. el Bando Barok’ “ Heredo de la postmilit” de Rafael Milhomen Intervjuo. Franjo Martin pri komuna vivo Legado: Karlo el Monato “ Fremdlingvaj indikiloj provokas proteston” de Cristina Casella Heather : el Monato. “ Kial ne aboli la mortpunon” de Isikawa Takasi […]
//</a><hr />
//
//<a href='https://kern.punkto.info/podlove/file/1445/s/feed/c/mp3/kp116-elbrula-sindromo.mp3'>
//<strong>2017-05-12
//Kernpunkto</strong>
//<br />
//Ofte oni aŭdis jam pri la elbrula sindromo, kiu koncernas homojn tro laborantajn en la moderna labormondo. Sen gasto ni laike parolas pri la fenomeno, kiu ne estas konsiderata malsano. Tamen multaj homoj je ĝi suferas kaj eĉ memmortigas sin.
//</a><hr />
//
//<a href='http://feedproxy.google.com/~r/retradio/~5/oa60_k2-SA8/RetRadio_12.05.2017_-vendredo.mp3'>
//<strong>2017-05-12
//Pola Retradio</strong>
//<br />
//La 649-a E_elsendo el la 12.05.2017 ĉe www.pola-retradio.org – Hodiaŭ antaŭ nia mikrofono ni gastigas Veronika Poór, la Ĝeneralan Direktoron de la CO en Rotterdamo, ĉi-foje por interparoli pri la nova retejo de UEA, pri kio kelkajn aktualaĵojn eksciis la partoprenantoj de la lasta Malferma Tago en Rotterdamo. – La komenca kulturkronika aktualaĵo rilatas al […]
//</a><hr />
//
//<a href='http://feedproxy.google.com/~r/retradio/~5/kVYUr8dQN6w/RetRadio_09.05.2017_mardo.mp3'>
//<strong>2017-05-09
//Pola Retradio</strong>
//<br />
//La 648-a E-elsendo ĉe www.pola-retradio.org el la 09.05.2017: – Nia nunsemajna E-elsendo denove enhavas literaturan segmenton, en kiu ĉi-foje aŭdiĝas la jam 55-a fragmento de la romano „Quo Vadis” de Henryk Sienkiewicz en la E-traduko de Lidia Zamenhof. – La komencaj aktualaĵoj rilatas al solenaĵoj omaĝantaj la finon de la 2-a mondmilito en 1945 kaj […]
//</a><hr />
//
//<a href='http://www.radioriodejaneiro1400am.com/audiosrrj/ESPERANTO%20LA%20LINGVO%20FRATECO/esperanto%20la%20lingvo%20290417.mp3'>
//<strong>2017-05-08
//La Lingvo de Frateco</strong>
//<br />
//La Lingvo de Frateco (Rio). http://www.radioriodejaneiro1400am.com/audiosrrj/ESPERANTOLALINGVOFRATECO/esperantolalingvo290417.
//</a><hr />
//
//<a href='http://www.melburno.org.au/3ZZZradio/mp3/2017-05-08.3ZZZ.radio.mp3'>
//<strong>2017-05-08
//3ZZZ en Esperanto</strong>
//<br />
//Kanto: el la kompaktdisko Ribela sono de Dolĉamar “ Duope” Legado: Karlo el la revuo Monato “ Nova ĉampiono” de Zaltko Tiŝljar Heather el Retradio de Anton Obendorfer “ 50 plus: sporto tenas menson freŝan” Parolado: Matt “ Nenie pli bone ol hejme” Kanto: el la kompaktdisko Muzikpluvo de Akordo “ Damo, mia kara” Legado: […]
//</a><hr />
//
//<a href='http://audio.radio-frei.de/podcast/Esperanto/20170507%20turint%20klein.mp3'>
//<strong>2017-05-08
//Radio Frei</strong>
//<br />
//Radio Frei. 20170507 turint klein.
//</a><hr />
//
//<a href='http://feedproxy.google.com/~r/retradio/~5/FDusEmBKijU/RetRadio_05.05.2017_vendredo.mp3'>
//<strong>2017-05-05
//Pola Retradio</strong>
//<br />
//La 647-a E_elsendo el la 05.05.2017 ĉe www.pola-retradio.org – En nia nunsemajna vendreda programo ni gastigas Konstantan Kongresan Sekretarion, Clay Magelhães kuntekste kun la 102-a UK en Seulo, Koreio, al kiu aliĝis jam 1042 personoj el 61 landoj. – La komencaj kulturkronikaj informoj koncernas la decidon pri tio, ke fama skulptaĵo de la pola skulptisto […]
//</a><hr />
//
//<a href='http://feedproxy.google.com/~r/retradio/~5/fZGA1hpMrZw/RetRadio_03.05.2017_merkredo.mp3'>
//<strong>2017-05-03
//Pola Retradio</strong>
//<br />
//La 646-a E-elsendo ĉe www.pola-retradio.org el la 03.05.2017: – Nia nunsemajna sondokumenta parto enhavas la prezenton de Ionel Onet dum la lastsabata (29.04.2017) Malferma Tago en Roterdamo pri la eldonejaj aktualaĵoj. – La komencaj novaĵoj rilatas al la hodiaŭa nacia festo en Pollando, la 226-a datreveno de la proklamo de la 3-Maja Konstitucio, al la […]
//</a><hr />
//
//<a href='http://www.melburno.org.au/3ZZZradio/mp3/2017-05-01.3ZZZ.radio.mp3'>
//<strong>2017-05-01
//3ZZZ en Esperanto</strong>
//<br />
//Kanto:el la kompaktdisko Svedaj kantoj 2003 ”Amata vi ŝajnas ja roz’ “ Legado: Franciska el la retejo de Labourstart “ Internacia tago de la Romaoj” Parolado: Marcel “ Kiel mi ne intence sed bonŝance trompis kelkiujn” Legado: Laszlo el la revuo Esperanto Pri la historio de la libro de Ulrich Linz La danĝera lingvo Kanto: el […]
//</a><hr />
//
//<a href='http://www.melburno.org.au/3ZZZradio/mp3/2017-04-24.3ZZZ.radio.mp3'>
//<strong>2017-04-24
//3ZZZ en Esperanto</strong>
//<br />
//Kanto: el la kompaktdisko Martin kaj la Talpoj “Vivo duras, sed vi molas” Legado: Laszlo el la revuo Esperanto “Novaj radikoj , neologismoj en la Evia lingvo” de Apelete Agbolo. Heather : “Vojaĝo al suda Panĝabio” de Nadipedia el Monato Karlo :el Fervoja Mondo de Jindrich Tomiŝek “ Interesaĵoj pri monto Gotthard tunelo” Kanto : […]
//</a><hr />
//
//
//<hr>
//</div>
//</body>
//</html>
//`
