# Junioraufgabe 2: St. Egano

## Lösungsidee

Für jedes Pixel, das zur Nachricht gehört, muss der RGB Wert ermittelt werden.
Der r Wert muss in das ASCII-Zeichen umgewandelt werden. Dieser Wert muss
an die bisherige Nachricht angehängt werden.
Dann muss die nächste horizontale Position ermittelt werden, indem der g Wert
zur aktuellen horizontalen Position addiert wird. Falls die neue horizontale Position
außerhalb des Bildes liegt, muss man die Breite des Bildes von der Position abziehen,
um die horizontale Position zu erhalten, die man erreichen würde, wenn man von der anderen Seite
des Bildes weiterzählt. Nach diesem Schritt kann die Position immer noch außerhalb des Bildes liegen,
daher wiederholt man diesen Schritt so lange, bis die Position innerhalb des Bildes liegt.
Dasselbe lässt sich jedoch schneller erreichen, indem man die horizontale Position modulo
der Breite des Bildes rechnet.
Für die neue vertikale Position tut man dasselbe mit dem b Wert und der Höhe des Bildes.
Diesen Prozess wiederholt man so lange wie der g oder b Wert des Pixels ungleich 0 ist.
Gestartet wird mit dem Pixel ganz oben links.

## Umsetzung

Die Datei, welche als Argument übergeben wurde, wird mit dem image Package
der Go-Standardbibliothek dekodiert. Dann wird ein StringBuilder zum Akkumulieren der Nachricht
erstellt. Die aktuelle Position wird mit zwei Variablen `x` und `y` gespeichert,
sie werden auf 0 initialisiert, was die Position ganz oben links im Bild ist.
Dann wird eine Schleife gestartet.
In der Schleife wird der RGB Wert des Pixels an der aktuellen Position ermittelt.
Der r Wert wird in eine Rune umgewandelt und an den StringBuilder angehängt.
x und y werden dann mit g und b addiert und modulo der Breite bzw. Höhe des Bildes gerechnet.
Wenn g und b 0 sind, wird die Schleife beendet.
Nach der Schleife wird der StringBuilder in einen String umgewandelt und ausgegeben.

## Beispiele

```shell
$ stegano Beispieleingaben/StEgano/bild01/bild01.png
Hallo Welt
```

```shell
$ stegano Beispieleingaben/StEgano/bild02/bild02.png
Hallo Gloria

Wie treffen uns am Freitag um 15:00 Uhr vor der Eisdiele am Markplatz.

Alle Liebe,
Juliane
```

```shell
$ stegano Beispieleingaben/StEgano/bild03/bild03.png
Hallo Juliane,

Super, ich werde da sein! Ich freue mich schon auf den riesen Eisbecher mit Erdbeeren. 

Bis bald,
Gloria
```

```shell
$ stegano Beispieleingaben/StEgano/bild04/bild04.png
Der Jugendwettbewerb Informatik ist ein Programmierwettbewerb für alle, die erste Programmiererfahrungen sammeln und vertiefen möchten. Programmiert wird mit Blockly, einer Bausteinorientierten Programmiersprache. Vorkenntnisse sind nicht nötig. Um sich mit den Aufgaben des Wettbewerbs vertraut zu machen, empfehlen wir unsere Trainingsseite . Er richtet sich an Schülerinnen und Schüler der Jahrgangsstufen 5 - 13, prinzipiell ist aber eine Teilnahme ab Jahrgangsstufe 3 möglich. Der Wettbewerb besteht aus drei Runden. Die ersten beiden Runden erfolgen online. In der 3. Runde werden zwei Aufgaben gestellt, diese gilt es mit eigenen Programmierwerkzeugen zuhause zu bearbeiten.
```

```shell
$ stegano Beispieleingaben/StEgano/bild05/bild05.png
Der Bundeswettbewerb Informatik richtet sich an Jugendliche bis 21 Jahre, vor dem Studium oder einer Berufstätigkeit. Der Wettbewerb beginnt am 1. September, dauert etwa ein Jahr und besteht aus drei Runden. Dabei können die Aufgaben der 1. Runde ohne größere Informatikkenntnisse gelöst werden; die Aufgaben der 2. Runde sind deutlich schwieriger.

Der Bundeswettbewerb ist fachlich so anspruchsvoll, dass die Gewinner i.d.R. in die Studienstiftung des deutschen Volkes aufgenommen werden. Aus den Besten werden die TeilnehmerInnen für die Internationale Informatik-Olympiade ermittelt. Der Bundeswettbewerb ermöglicht den Teilnehmenden, ihr Wissen zu vertiefen und ihre Begabung weiterzuentwickeln. So trägt der Wettbewerb dazu bei, Jugendliche mit besonderem fachlichen Potenzial zu erkennen.
```

```shell
$ stegano Beispieleingaben/StEgano/bild06/bild06.png
Bonn

Die Bundesstadt Bonn (im Latein der Humanisten Bonna) ist eine kreisfreie Großstadt im Regierungsbezirk Köln im Süden des Landes Nordrhein-Westfalen und Zweitregierungssitz der Bundesrepublik Deutschland. Mit 336.465 Einwohnern (31. Dezember 2022) zählt Bonn zu den zwanzig größten Städten Deut[...]wa 30 Mitarbeitern in Bonn.

Arbeitsmarktbehörden
Bonn ist außerdem Standort der Zentralen Auslands- und Fachvermittlung (ZAV) der Bundesagentur für Arbeit (BA). Im Stadtteil Duisdorf befindet sich der Hauptsitz der ZAV mit ihren bundesweit 18 Standorten.

Quelle: https://de.wikipedia.org/wiki/Bonn
```

```shell
$ stegano Beispieleingaben/StEgano/bild07/bild07.png
Es hatte ein Mann einen Esel, der schon lange Jahre die Säcke unverdrossen zur Mühle getragen hatte, dessen Kräfte aber nun zu Ende giengen, so daß er zur Arbeit immer untauglicher ward. Da dachte der Herr daran, ihn aus dem Futter zu schaffen, aber der Esel merkte daß kein guter Wind wehte, lief fo[...]zt der Richter, der rief bringt mir den Schelm her. Da machte ich daß ich fortkam." Von nun an getrauten sich die Räuber nicht weiter in das Haus, den vier Bremer Musikanten gefiels aber so wohl darin, daß sie nicht wieder heraus wollten. Und der das zuletzt erzählt hat, dem ist der Mund noch warm.
```

## Quellcode
```go
package main

import (
	"fmt"
	"image"
	"os"
	"strings"

	_ "image/jpeg"
	_ "image/png"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: stegano <filepath>")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	var sb strings.Builder
	bounds := img.Bounds()
	var x, y int

	for {
		r, g, b := GetRGB(img, x, y)

		x += int(g)
		y += int(b)

		x %= bounds.Max.X
		y %= bounds.Max.Y

		sb.WriteRune(rune(r))

		if g == 0 && b == 0 {
			break
		}
	}

	fmt.Println(sb.String())
}

func GetRGB(img image.Image, x, y int) (r, g, b uint32) {
	r, g, b, _ = img.At(x, y).RGBA()
	return r >> 8, g >> 8, b >> 8
}
```
