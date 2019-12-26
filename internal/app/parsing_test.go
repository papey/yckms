package app

import (
	"testing"
)

func TestParseYCKMPlaylist(t *testing.T) {
	input := `Au programme :
- Revue de presse : Matthieu
- Chronique Fidlar : Théo
- Chronique Waste Of Space Orchestra : Eline

Playlist : Bus / I Buried Paul, Nails / Endless Resistance, Sepultura / Territory, Venom / Evilution Devilution, All Pigs Must Die / The Whip, Fidlar / Too Real, Obituary / Slowly We Rot, Wayfarer / Catcher, Waste of Space Orchestra / Seeker's Reflection, Bat / Long Live the Lewd, Witchfinder / Ouija, Gadget /Choice of a Lost Generation`

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
	input := `Enfin le voilà !

Le futur du Bruit :
(0:04:09) Le cool
La saison 02 : Nouveau graphics et jingles ! yeaaah !
  Balek de Facebook, insta et Goodbye au site !
  Les lives sur twitch motherfucker !
  Tee-shirts ?
  Nouvelle émission Le Bruit tous les lundis sur TF1

(0:15:10) Le pas cool
Futur strike sur apple podcast ?
    Ce qu'on va faire et ce qu’on ne va pas faire
      Tour de france en camionnette, représentation en live chez l'autochtone


C’est l’heure du BILAN mon chat !
10 :
💀 (0:22:39) Flight of the Conchords - Live in London (Ecouter l’album (https://album.link/fr/i/1447320701)) (Ecouter l’épisode concerné (https://podcast.ausha.co/le-bruit/le-bruit-de-l-avant-2019-case-3))
🐻 (0:23:58) Last Train - The Big Picture (Ecouter l’album (https://album.link/fr/i/1472019197)) (Ecouter l’épisode concerné (https://podcast.ausha.co/le-bruit/le-bruit-volume-12-l-e-pisode-sans-nom))
9 :
💀 (0:26:27) Jon and Roy - Here (Ecouter l’album (https://album.link/fr/i/1447292371))
🐻 (0:28:45) Alcest - Spiritual Instinct (Ecouter l’album (https://album.link/fr/i/1476281691))
8 :
💀 (0:31:05) Dangerface - Get Loud! (Ecouter l’album (https://album.link/fr/i/1448576237)) (Ecouter l’épisode concerné (https://podcast.ausha.co/le-bruit/le-bruit-volume-03-grand-e-cart-facial-entre-deux-brouettes))
🐻 (0:32:10) The Offering - HOME (Ecouter l’album (https://album.link/fr/i/1464688715)) (Ecouter l’épisode concerné (https://podcast.ausha.co/le-bruit/le-bruit-volume-11-des-menhirs-et-des-crobots))
7 :
💀 (0:34:09) Frank Carter &amp; The Rattlesnakes - End of Suffering (Ecouter l’album (https://album.link/fr/i/1447497006)) (Ecouter l’épisode concerné (https://podcast.ausha.co/le-bruit/le-bruit-volume-08-la-baleine-de-pointe-a-pitre))
🐻 (0:35:33) Cult of Luna - A Dawn to Fear (Ecouter l’album (https://album.link/fr/i/1472437655)) (Ecouter l’épisode concerné (https://podcast.ausha.co/le-bruit/le-bruit-volume-12-l-e-pisode-sans-nom))
6 :
💀 (0:36:32) Dinosaur Pile-Up - Celebrity Mansions (Ecouter l’album (https://album.link/fr/i/1456626323)) (Ecouter l’épisode concerné (https://podcast.ausha.co/le-bruit/le-bruit-volume-09-allez-vous-faire-loutre-1))
🐻 (0:38:15) Shadow Of Intent - Melancholy (Ecouter l’album (https://album.link/fr/i/1466697634)) (Ecouter l’épisode concerné (https://podcast.ausha.co/le-bruit/le-bruit-de-l-avant-2019-case-4))
5 :
💀 (0:39:07) Rival Sons - Feral Roots (Ecouter l’album (https://album.link/fr/i/1440340685)) (Ecouter l’épisode concerné (https://podcast.ausha.co/le-bruit/le-bruit-volume-02-merci-pour-les-champignons))
🐻 (0:40:17) HATH - Of Rot and Ruin (Ecouter l’album (https://album.link/fr/i/1449749750)) (Ecouter l’épisode concerné (https://podcast.ausha.co/le-bruit/le-bruit-volume-07-salsifis-les-conneries))
4 :
💀 (0:41:15) Slipknot - We Are Not Your Kind (Ecouter l’album (https://album.link/fr/i/1463706038)) (Ecouter l’épisode concerné (https://podcast.ausha.co/le-bruit/le-bruit-volume-11-des-menhirs-et-des-crobots))
🐻 (0:42:47) Rival Sons - Feral Roots (Ecouter l’album (https://album.link/fr/i/1440340685)) (Ecouter l’épisode concerné (https://podcast.ausha.co/le-bruit/le-bruit-volume-02-merci-pour-les-champignons))
3 :
💀 (0:43:27) Billie Eilish - When We All Fall Asleep, Where Do We Go? (Ecouter l’album (https://album.link/fr/i/1450695723))
🐻 (0:47:20) Wheel - Moving Backwards (Ecouter l’album (https://album.link/fr/i/1442249647)) (Ecouter l’épisode concerné (https://podcast.ausha.co/le-bruit/le-bruit-volume-05-loups-flamboyants-vs-cyber-tarentules))
2 :
💀 (0:49:20) Norma Jean - All Hail (Ecouter l’album (https://album.link/fr/i/1475151504)) (Ecouter l’épisode concerné (https://podcast.ausha.co/le-bruit/le-bruit-de-l-avant-2019-case-7))
🐻 (0:50:45) Periphery - Periphery IV: HAIL STAN (Ecouter l’album (https://album.link/fr/i/1450856455)) (Ecouter l’épisode concerné (https://podcast.ausha.co/le-bruit/le-bruit-volume-06-la-cornebidouille-de-satan))
1 :
💀 (0:53:17) Clowns - Nature / Nurture (Ecouter l’album (https://album.link/fr/i/1451092249)) (Ecouter l’épisode concerné (https://podcast.ausha.co/le-bruit/le-bruit-volume-07-salsifis-les-conneries))
🐻 (0:55:18) Wilderun - Veil of Imagination (Ecouter l’album (https://album.link/fr/i/1479174395)) (Ecouter l’épisode concerné (https://podcast.ausha.co/le-bruit/le-bruit-de-l-avant-2019-case-8))

Top 100 (Topsters):
🐻 https://i.imgur.com/3b31NW0.png (https://i.imgur.com/3b31NW0.png)
💀 https://i.imgur.com/ARMW0lb.png (https://i.imgur.com/ARMW0lb.png)

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

	artistExpected := "Flight of the Conchords"
	albumExpected := "Spiritual Instinct"

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
