[ ] **Aufgabe B1: Potenzfunktion**

Schreiben Sie eine Funktion:

``` GO
func potenziere(basis, exponent int) int {...}
```

welche die 𝑛‑te Potenz einer Zahl berechnet.
Beachten Sie dabei:
- Der Exponent 𝑛 ist eine nicht‑negative ganze Zahl. Dies muss Ihre Funktion jedoch nicht überprüfen.
- Wir definieren 00 als 1.
- Implementieren Sie Ihre Funktion mit einer Schleife, d.h. benutzen Sie nicht die Bibliotheksfunktion math.pow().
- Sie dürfen davon ausgehen, dass Ihrer Funktion nur Argumente übergeben werden, die einen
Potenzwert im zulässigen Integer‑Wertebereich ergeben.

  
***Aufgabe B1: Potenzfunktion***

Im Ordner dieser Aufgabe finden Sie eine Datei potenziere.go, die Sie entsprechend der Aufgabenstellung
abändern sollen.

Weiterhin stellen wir in der Datei main.go eine main‑Funktion zur Verfügung, damit Sie Ihre Funktion
in einem beispielhaften Kontext kompilieren und ausführen können. Nachdem Sie in den Aufgabenordner
gewechselt sind, geben Sie dazu folgenden Befehl ein:

``` GO
go run .
```

Die Datei potenziere_test.go stellt Tests bereit, die Sie mit folgendem Befehl durchführen können:

``` GO
go test
```