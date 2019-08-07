package app

import (
	"testing"
)

func TestParseYCKMPlaylist(t *testing.T) {
	input := `Au programme :
- Revue de presse : Matthieu
- Chronique Fidlar : Théo
- Chronique Waste Of Space Orchestra : Eline

Playlist : Bus / I Buried Paul, Nails / Endless Resistance, Sepultura / Territory, Venom / Evilution Devilution, All Pigs Must Die / The Whip, Fidlar / Too Real, Obituary / Slowly We Rot, Wayfarer / Catcher, Waste of Space Orchestra / Seeker's Reflection, Bat / Long Live the Lewd, Witchfinder / Ouija, Gadget /Choice of a Lost Generationi`

	artistExpected := "Bus"
	songExpected := "Territory"

	s := parseYCKMPlaylist(input)
	if s == nil {
		t.FailNow()
	}

	if s[0].artist != artistExpected {
		t.Errorf("Expected : %s, Get : %s", artistExpected, s[0].artist)
	}

	if s[2].title != songExpected {
		t.Errorf("Expected : %s, Get : %s", songExpected, s[2].title)
	}

}

func TestParseLeBruitPlaylist(t *testing.T) {
	input := `On en a gros (c'est faux) ! Au programme : Bouli prends goût à se faire casser les dents, Germain dors avec un Irlandais, nous faisons des sous-entendus graveleux avec une élégance sans égale et en plus y a du rab !

Le pod :
Le Bruit sur son 31

Les chroniques du bruit :
💀 0:06:07 | Clowns - Nature / Nurture (Ecouter (https://song.link/album/fr/i/1451092249))
🐻 0:15:25 | Belzebubs - Pantheon Of The Nightside Gods (Ecouter (https://song.link/album/fr/i/1453386311))
💀 0:25:26 | Boundaries - Turning Point (Ecouter (https://song.link/album/fr/i/1453296200))
🐻 0:34:50 | Hath - Of Rot And Ruin (Ecouter (https://song.link/album/fr/i/1449749750))
💀 0:47:28 | Danko Jones - A Rock Supreme (Ecouter (https://song.link/album/fr/i/1448249776))
🐻 0:58:04 | Glen Hansard - This Wild Willing (Ecouter (https://song.link/album/fr/i/1449024496))

Dans notre poche :
🐻 1:09:16 | Suldusk - Lunar Falls (Ecouter (https://song.link/album/fr/i/1455319108))
🐻 1:10:10 | Extortionist - Sever the Cord (Ecouter (https://song.link/album/fr/i/1451678823))
💀 1:11:10 | Billie Eilish – when we all fall asleep, where do we go? (Ecouter (https://song.link/album/fr/i/1450695723))
💀 1:12:05 | Jon and Roy – Here (Ecouter (https://song.link/album/fr/i/1447292371))
💩 1:15:44 | Bêtisier

La playlist Spotify rassemblant les morceaux des épisodes : La playlist de l'année 2019 (https://open.spotify.com/user/coloneltabasco/playlist/5m4bJu85sXYO6XybLtEprb?si=0gxIYytsQoCba9pZgGCmfw)

Réseaux Sociaux :
Twitter : @LeBruitPodcast (https://twitter.com/LeBruitPodcast)
  Twitter Bouli : @Boulinosaure (https://twitter.com/Boulinosaure)
  Twitter Germain : @Bearded__Bear (https://twitter.com/Bearded__Bear)
  Instagram : @lebruitpodcast (https://www.instagram.com/lebruitpodcast/)
  Facebook : lebruitpodcast (https://www.facebook.com/lebruitpodcast)
  Twitch : lebruitpodcast (https://www.twitch.tv/lebruitpodcast)
  Discord : https://discord.gg/qQ2XGkk (https://discord.gg/qQ2XGkk)

Plateforme d’écoutes :
Ausha (https://podcast.ausha.co/le-bruit), Youtube (https://www.youtube.com/channel/UCjIrHhD3HXBZFDd4cR5Mouw), Spotify (https://open.spotify.com/show/31ZkKfw71Dp6uGhxmB7joR?si=JRoQXdO6TJqlVa-brimyeg), Feedburner (http://feeds.feedburner.com/lebruitpodcast), Tunein Radio (https://tunein.com/podcasts/Music-Podcasts/Le-Bruit-p1183237/), Itunes (https://itunes.apple.com/fr/podcast/le-bruit/id1448164973?l=en), Stitcher (https://www.stitcher.com/podcast/guillaume-delacroix/le-bruit?refid=stpr), PocketCast (https://pca.st/VKP0), Podcloud (https://podcloud.fr/podcast/le-bruit), Deezer (https://www.deezer.com/fr/show/70571), Google podcast (https://www.google.com/podcasts?feed=aHR0cDovL2ZlZWRzLmZlZWRidXJuZXIuY29tL2xlYnJ1aXRwb2RjYXN0)`

	artistExpected := "Clowns"
	albumExpected := "Of Rot And Ruin"

	s := parseLeBruitPlaylist(input)
	if s == nil {
		t.FailNow()
	}

	if s[0].artist != artistExpected {
		t.Errorf("Expected : %s, Get : %s", artistExpected, s[0].artist)
	}

	if s[3].album != albumExpected {
		t.Errorf("Expected : %s, Get : %s", albumExpected, s[3].album)
	}

}

func TestParseHarryCoverPlaylist(t *testing.T) {
	input := `<p>J'inaugure aujourd'hui une nouvelle catégorie d'épisodes, les "<strong>Fausses Bonnes Idées</strong>" entre hommages ratés, plaisirs coupables et autres plantages.</p>
<p>Pour ce premier épisode, retour sur l'intemporel "Wicked Game" de Chris Isaak. Non, la reprise évoquée ne sera ni celle de James Vincent Mc Morrow, ni celle de London Grammar, pas meilleures que l'originale, mais pas assez décalées pour être considérées comme ratées.</p>
<p>J'attends vos retours, commentaires et étoiles Itunes.</p>
<p><a href="https://open.spotify.com/track/390AWnOn2rfe9FzQjYmxIH?si=iRpCp690Q2C5jkF6vzs2Sw">La version originale Spotify</a></p>
<p><a href="https://www.youtube.com/watch?v=dlJew-Dw87I">La version originale Youtube</a></p>
<p><a href="https://open.spotify.com/track/5XenUjG7cRnTkUe8AVuuMX?si=xmc2-xfpSK2ENCJlHppLrw">La reprise Spotify</a></p>
<p><a href="https://www.youtube.com/watch?v=8oYodfK4DkE">La reprise Youtube</a></p>
<p><a href="https://open.spotify.com/playlist/2MEGtxKmjFC8vSxBj45Wi3?si=CzXMQUCiR_6yKSmkIyOnNg">La playlist Spotify de l'émission</a></p>
<p><a href="https://discord.gg/FjeJpx">Le discord des streetcasteurs</a></p>
<p>Me contacter sur Twitter : @HCoverpodcast</p>
<p><br></p>
<p><strong>Pour aller plus loin :</strong></p>
<p><a href="https://open.spotify.com/album/36tz5XLSdscEvXvzWQwSXv?si=eCGtZ6MhR5CXxBIQXAOC9A">-"Razorblade Romance" sur Spotify</a></p>
<p><a href="https://www.youtube.com/watch?v=jK4IBAZGkdE">-"Razorblade Romance" sur Youtube</a></p>
<p><a href="https://open.spotify.com/track/3V1H6liHwCDcWeqdPJabOM?si=XV2EIGWtTAyIXex42UeVDA">-"Wicked Game" par Stone Sour sur Spotify</a></p>
<p><a href="https://www.youtube.com/watch?v=cncoJB_C-m0">-"Wicked Game" par Stone Sour sur Youtube</a></p>
<p><br></p>
<p><strong>Ecoutez ce podcast sur :</strong></p>
<p><a href="https://anchor.fm/leotot8">Anchor</a></p>
<p><a href="https://podcasts.apple.com/us/podcast/harry-cover-reprises-fra%C3%AEches-et-d%C3%A9s%C3%A9quilibr%C3%A9es/id1463761176?uo=4">Apple Podcasts</a></p>
<p><a href="https://open.spotify.com/show/6usEP3XdEwNcotYdQ8MtWy">Spotify</a></p>
<p><a href="https://anchor.fm/s/b3b7468/podcast/rss">Flux rss</a></p>
<p><br></p>
<p>Générique d'intro : "I will survive" par Cake</p>
<p><br></p>`

	idExpected := "390AWnOn2rfe9FzQjYmxIH"

	s := parseHarryCoverPlaylist(input)
	if s == nil {
		t.FailNow()
	}

	if s[0].id != idExpected {
		t.Errorf("Expected : %s, Get : %s", idExpected, s[0].id)
	}

}
