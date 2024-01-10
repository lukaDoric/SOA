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

## Paketi

Osnovna gradivna komponenta svakog Go programa jesu paketi. Program mora da ima minimum jedan paket - `main`. Svaka funkcija ili struktura koju napišemo u Go-u će pripadati nekom paketu (koji je definisan u prvoj liniji svake datoteke). Ne moramo sve funkcionalnosti paketa definisati unutar jedne datoteke, već ih možemo razdeliti na proizvoljno mnogo, ali je tada neophodno da se one sve nalaze u istom direktorijumu.

Pored toga što dele našu aplikaciju na zasebne celine i daju kodu strukturu, paketi se takođe koriste umesto biblioteka. Svaki put kada u kodu naiđemo na liniju **import _naziv_paketa_**, tada u svoj program uvozimo sve funkcionalnosti definisane u okviru tog paketa. Na primeru gore možete primetiti da smo import-ovali paket `fmt`, koji u sebi sadrži, između ostalog, funkciju za ispis teksta na standardni izlaz.

## Dependency management

Uz instalaciju go jezika ste dobili i kolekciju paketa sa funkcionalnostima koje se najčešće koriste prilikom pisanja programa. U [standardnoj biblioteci](https://pkg.go.dev/std) se, između ostalog, nalaze i paketi za rad sa stringovima, bufferima, http protokolom, operativnim sistemom itd.

Kako međutim da importujemo pakete koji nisu deo standardne biblioteke? Go poseduje i alat za dependency management: `go mod`. Uz pomoć ovog alata možemo na jednostavan način da skinemo sve pakete koji su  potrebni za rad programa, kao i da vodimo računa o verzijama istih, nalik maven-u ili npm-u. Da bi koristili go mod, prvo treba da inicijalizujemo projekat komandom **go mod init _jedinstven_naziv_projekta_**. Naziv projekta je najčešće u formatu URL-a, npr. _prvi_primer.xws.com_.

Nakon inicijalizacije projekta primetićete da je kreirana datoteka `go.mod` koja sadrži za sada samo naziv projekta (modula) i verziju go-a koju koristi.

U narednom koraku ćemo promeniti Hello world primer, tako da isto ispisuje ali u nekoj boji. Za to ćemo koristiti paket [fatih/color](https://github.com/fatih/color)

```go
package main

import "github.com/fatih/color"

func main(){
  color.Cyan("Hello world")
}
```

Ako odmah pokušamo da pokrenemo novi program dobićemo grešku jer kompajler ne prepoznaje `fatih/color` paket, koji nije deo standardne biblioteke. Komanda `go mod tidy` proverava koji se sve paketi koriste u našem programu, i ako neki ne prepoznaje, pokušaće da ga skine sa interneta. Nakon što izvršite ovu komandu, videćete da su u `go.mod` datoteku upisani svi paketi koji se koriste: i `color` paket, ali i paketi koji su neophodni za rad `color` paketa. `go.mod` datoteku možemo da poredimo sa `pom.xml`-om u mavenu, samo što je sintaksa znatno jednostavnija, ili `package.json`-om u Node projektima.

Takođe možete primetiti da je kreirana i nova datoteka, `go.sum`, gde se nalaze preciznije informacije o paketima, tačna verzija koja je skinuta, kao i checksum poslednjeg komita u trenutku skidanja paketa. Ovo je ekvivalent `package-lock.json`-u u Node-u.

Ako sada pokrenemo program, trebalo bi da vidimo ispis Hello world u ~svetlo plavoj boji.

Snimak vezan za pakete i dependency management možete pogledati <a href='https://youtu.be/0ymOGFQRch4'>ovde</a>.

