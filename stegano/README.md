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

## Quellcode
