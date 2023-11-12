# Aufgabe 3: Zauberschule

## Lösungsidee

Der Plan der Zauberschule kann als gewichteter Graph dargestellt werden, wobei die Knoten
je einem Feld des Plans entsprechen und die Kanten die möglichen Bewegungen zwischen den Feldern.
Das Gewicht einer Kante ist für alle Kanten, außer einem Stockwerkwechsel, 1. Für einen Stockwerkwechsel
beträgt es 3.
Um den schnellsten Weg zu finden, muss der Weg mit den geringsten Kosten gefunden werden,
wozu der ~~Djikstra-Algorithmus~~ Dijkstra-Algorithmus verwendet werden kann.
Für jeden Knoten wird die Distanz zum Startknoten gespeichert, die anfangs unendlich ist.
Die Distanz des Startknotens wird auf 0 gesetzt.
Dann wird der Knoten mit der geringsten Distanz, welcher mit der aktuellen Distanz noch nicht
besucht wurde, ausgewählt und die Distanz zu allen Nachbarn des Knotens wird aktualisiert,
falls die Distanz über den aktuellen Knoten kürzer ist. Dieser Prozess wird
so lange wiederholt, bis der Zielknoten erreicht wird. Dann ist der Weg mit den geringsten Kosten
gefunden.

## Umsetzung

Die Datei, welche als Argument übergeben wurde, wird in eine Datenstruktur umgewandelt,
welche den Plan als 3d-Array darstellt. Die erste Dimension ist die Stockwerknummer,
die zweite und dritte Dimension sind die x- und y-Koordinate des Feldes.
Dann 3d-Array enthält booleans die angeben, ob das Feld begehbar ist oder nicht.
Während des Einlesens wird auch die Start- und Zielposition gespeichert.
Der Plan wird nicht als Graph dargestellt, da die Kanten nicht gespeichert werden müssen.
Sie sind immer die gleiche, links, rechts, oben, unten und Stockwerkwechsel.
Es muss nur überprüft werden, ob das Feld begehbar ist.
Um die Distanz zu jedem Feld zu speichern, wird eine map `distances` erstellt.
Um später den Weg zu rekonstruieren, wird ein map `previous` erstellt, welche für jeden Knoten
den Vorgängerknoten speichert. Um den Knoten mit der geringsten Distanz zu finden, wird ein
Min-Queue `queue` erstellt, welcher nur die Knoten enthält, die noch nicht besucht mit ihrer
aktuellen Distanz besucht wurden.
Dann wird eine Schleife gestartet, welche so lange wiederholt wird,
bis der Zielknoten erreicht wurde oder die Queue leer ist.
In der Schleife wird der Knoten mit der geringsten Distanz aus der Queue entfernt.
Dann wird überprüft, ob der Knoten der Zielknoten ist, wenn ja, wird die Schleife beendet.
Dann wird für jeden Nachbarn die Distanz über den aktuellen Knoten berechnet.
Wenn die Distanz kürzer ist, als die aktuelle Distanz des Nachbarn,
oder es noch keine Distanz gibt, wird die Distanz des Nachbarn
auf die neue Distanz gesetzt und der Nachbar mit der neuen Distanz in die Queue eingefügt.
Der aktuelle Knoten wird als Vorgänger des Nachbarn gespeichert.
Nach dem Ende der Schleife wird der Weg rekonstruiert, indem der Vorgänger des Zielknotens
zum Weg hinzugefügt wird und dann sein Vorgänger usw., bis der Startknoten erreicht wurde.
Der Weg ist dann in umgekehrter Reihenfolge gespeichert, daher wird er umgedreht.

## Beispiele

## Quellcode
