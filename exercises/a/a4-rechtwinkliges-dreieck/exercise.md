[ ] **Aufgabe A4: Rechtwinkliges Dreieck**

Schreiben Sie eine Funktion:

``` GO
func istRechtwinklig(x, y, z int) bool
```

Der Funktion werden die Seitenlängen eines Dreiecks übergeben. Sie soll genau dann true zurückgeben,
wenn ein Dreieck mit diesen Seitenlängen rechtwinklig ist.

Anmerkung: Sie müssen nicht überprüfen, ob die übergebenen Seitenlängen ein Dreieck bilden. Dies
dürfen Sie voraussetzen.

Tipp: Ein Dreieck ist genau dann rechtwinklig, wenn für seine Seitenlängen der Satz des Pythagoras
gilt.

***Ressourcen***

Im Ordner dieser Aufgabe finden Sie eine Datei istRechtwinklig.go, die Sie entsprechend der
Aufgabenstellung abändern sollen. Weiterhin stellen wir in der Datei main.go eine main-Funktion
zur Verfügung, damit Sie Ihre Funktion in einem beispielhaftem Kontext kompilieren und ausführen
können. Nachdem Sie in den Aufgabenordner gewechselt sind, geschieht das mit dem Befehl:

``` GO
go run .
```

Die DateiistRechtwinklig_test.gostellt Tests bereit, die Sie mit folgendemBefehl durchführen
können:

``` GO
go test
```

Wenn alle Tests erfolgreich sind, erhalten Sie eine Meldung wie:

``` GO
PASS
ok aufgaben/istRechtwinklig 0.125s
```

Wenn hingegen einige Tests fehlschlagen, gibt es eine Meldung der Form: