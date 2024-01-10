## Uvod u Golang

Go (Golang) je programski jezik razvijen od strane Google-a krajem 2000-ih godina. U početku se koristio kao interni alat samo unutar kompanije, da bi krajem 2009. jezik objavili u vidu _open-source_ projekta. Među kreatorima jezika nalazi se i **Ken Thompson**, tvorac _Unix_ operativnog sistema. 

### Osnovne karakteristike Go-a

- **Kompajliran** jezik, nalik na C i C++. Za pokretanje programa napisanih u Go-u nije potrebno ništa osim binarne datoteke koja se kreira prilikom kompajliranja. Nema potrebe za izvornim kodom i dodatnim alatima poput virtualnih mašina unutar kojih se kod izvršava (Java) ili interpretera (Python, JS).
- Koristi “**garbage collector**” kao i Python i Java, te korisnici ne moraju da vode računa o alociranju i dealociranju memorije prilikom pisanja koda.
- Go je razvijen u trenutku kada su multiprocesorske mašine i procesori sa više jezgara bili već uveliko u širokoj upotrebi. Kreatorima jezika je jedan od prioriteta bio omogućavanje jednostavnog **pisanja konkurentnog koda**, i njegovo efikasno izvršavanje. Go ima podršku za rad sa nitima (go-rutinama) ne u vidu eksternih biblioteka, već u vidu funkcionalnosti koje dolaze zajedno sa samim jezikom, poput kreiranja funkcija, struktura ili promenjivih.
- Jednostavno **testiranje koda**, te ni ovde nema potrebe za korišćenje eksternih biblioteka i/ili alata.


### Zašto Go?

Cilj predmetnog projekta je izrada _mikroservisne_ veb aplikacije. O ovoj temi ćete detaljno čuti u narednim terminima vežbi, ali za sada možemo da kažemo da ćemo pisati nekoliko "jednostavnijih" programa umesto jednog kompleksnog, od kojih će svaki obavljati deo funkcionalnosti naše veb aplikacije. 

Kada pišemo kompletnu veb aplikaciju u vidu jednog programa koji koristi gigabajte RAM memorije tokom svog izvršavanja, i koji će raditi na tradicionalnim veb serverima, _overhead_ koji nose dodatni alati neophodni za rad same aplikacije (poput virtualne mašine ili interpretera) ne predstavlja veliki problem (procentualno troše malo resursa u odnosu na samu aplikaciju). Međutim, ako hoćemo da tu istu veb aplikaciju razdelimo na 10 ili 15 jednostavnijih programa, od kojih će svaki trošiti svega nekoliko stotina MB RAM memorije, odjednom _overhead_ koji nose propratni alati postaje primetan. Ovaj problem dolazi još do većeg izražaja ako uzmemo u obzir da ćemo te iste servise verovatno pokretati na mašinama _cloud provider_-a poput AWS-a, koji svoje usluge naplaćuju, delom, na osnovu resursa koje naši servisi troše. Zbog ovoga je prilikom razvijanja mikroservisnih aplikacija poželjno korišćenje kompajliranih jezika, poput Go-a.

Jedna od glavnih mana kompajliranih jezika jeste distribucija programa. Moramo posebno kompajlirati program za svaki tip mašine na kojoj će on biti pokrenut. Međutim, ako uzmemo u obzir da se naša mikroservisna aplikacija neće pokretati na korisničkim računarima, već na serverima gde mi imamo kontrolu nad hardverom, ovo nam neće predstavljati problem.

Pored toga Go ima _out-of-the-box_ podršku za rad sa mrežnim protokolima poput **HTTP**-a i **gRPC**-a, što je još jedan plus s obzirom da pišemo veb servise.


## Prvi Go program

Da biste pratili vežbe, potrebno je da prvo [instalirate go](https://go.dev/doc/install).

Što se tiče Text Editora/IDE-a, preporučujemo VS-Code zajedno sa pluginom za Go, mada možete pokušati i sa GoLand-om (JetBrains) ili nekim još jednostavnijim alternativama poput Vim-a.

Go kompajler očekuje da svaki program sadrži `main` paket, koji u sebi ima `main` funkciju. Ovo je ulazna tačka programa, poput main-a u C-u ili Javi. Datoteke sa go kodom se završavaju `.go` ekstenzijom, a u praksi se datoteka koja u sebi sadrži main funkciju često naziva `main.go`.

Kreiranje prvog programa možete pogledati na sledećem <a href='https://youtu.be/yZMr0gj8Vjc'>snimku</a>.

### Hello World

```go
package main

import "fmt"

func main() {
  fmt.Println("Hello, World")
}
```

Kako pokrenuti program napisan u Go-u?

Kao što smo par puta do sada već pomenuli, go je kompajliran jezik, te je neophodno prvo kompajlirati izvršnu datoteku uz pomoć komande  **go build**  [-o _nazil_kompaljirane_binarne_datoteke_] _naziv_main_datoteke_   

npr. `go build -o prvi_go_program main.go`

Nakon kompajliranja pokrećemo program tako što pozovemo binarnu datoteku. U linuksu bi to bilo `./prvi_go_program`

Prilikom razvijanja programa češće se koristi komanda **go run** _naziv_main_datoteke_ koja istovremeno kompajlira i odmah pokreće program, ali ne kreira izvršnu datoteku.
