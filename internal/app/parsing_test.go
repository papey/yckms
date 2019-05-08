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
