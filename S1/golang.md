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

## Strukture podataka i osnovna sintaksa

## Osnovna sintaksa

U ovoj lekciji ćemo ukratko proći osnovnu sintaksu i koncepte Go jezika. Nećemo zalaziti previše u detalje jer imate sve detaljno objašnjeno uz interaktivne primere na sajtu [gobyexample.com](https://gobyexample.com/). Pre svega pogledajte sekcije za for petlje, if-else i funkcije.


### Deklaracija promenjivih

Go je, poput Jave ili C-a _strongly-typed_ programski jezik, što znači da prilikom deklarisanja promenjive treba naglasiti koji tip podataka će čuvati, i tip promenjive se ne može menjati u toku izvršavanja programa, za razliku od jezika poput JS-a i Pythona. 

```go
 var a int
```

Moguće je deklarisati više promenjivih u jednoj liniji, a ako prilikom deklaracije i inicijalizujemo promenjivu, tip promenjive ne moramo eksplicitno naglasiti. Takođe postoji skraćena sintaksa za deklarisanje i inicijalizovanje promenjive, koja je prikazana u trećoj liniji:

```go
var a, b int
var c = "String"
d := true
```
Da ponovimo još jednom, iako u drugoj i trećoj liniji nije nigde eksplicitno naglašeno koji je tip promenjive, on je implicitno definisan, i ne možemo kasnije u kodu da dodelimo `integer` promenjivoj **_d_** koju smo implicitno deklarisali kao `bool`.

### if-else

Jedina razlika u odnosu na druge jezike jeste što, pored boolean uslova, možemo da definišemo i promenjivu koja će se koristiti isključivo u if/else bloku.

```go
if b%2==0 {
  ...
}else{
  ...
}

// Primer if petlje gde u zaglavlju i definišemo novu promenjivu

if a:=2, b%2==0 {
  fmt.Println(a)
}else{
  fmt.Println(b)
}

fmt.Println(a) // Ovo će izbaciti grešku, jer promenjiva a nije vidljiva izvan if/else bloka u okviru kog je definisana
```

### for petlje

Za for petlju možemo da koristimo različite sintakse, u zavisnosti kako želimo da se ponaša:
1. Ako ne navedemo ni jedan argument nakon ključne reči `for`, onda imamo beskonačnu for pelju iz koje jedino možemo da izađemo koristeći `break` ili `return`
2. Ako postavimo samo uslov (npr. i != 10 ), onda imamo standardnu while petlju
3. Možemo da postavimo samo inicijalno stanje i uslov, a "korak" da preskočimo, samo se onda naravno podrazumeva da će se promenjiva od koje zavisi for petlja menjati unutar for bloka, jer u suprotnom imamo opet beskonačnu petlju
4. Da navedemo sve tri vrednosti, to je klasična for petlja
5. Da koristimo `range` funkciju (primer ćete videti u sledećoj lekciji)

```go
// 1)
for {... bilo kakav kod}

// 2) Ovde se očekuje da je promenjiva i inicijalizovana ranije u kodu
for i<10 { i = i*2 }
// 3)
for i:=0; i<10 {....}
// 4) 
for i:=0; i<10; i++ {...}
```

### funkcije

U go-u se funkcije deklarišu na sledeći način:

```go
func test_funkcija(a, b int, c bool) (int, string) {
}
```

Primećujemo da prilikom definisanja parametara funkcije tip promenjive pišemo nakon naziva, i da, ako imamo više parametara istog tipa možemo samo jednom da naglasimo koji tip je u pitanju  (u našem slučaju su i _a_ i _b_ tipa int). Povratnu vrednost funkcije (to jest tip povratne vrednosti) definišemo nakon liste parametara. Funkcije mogu vratiti više promenjivih, a ne samo jednu. Ako vraćamo samo jednu promenjivu, nije potrebno koristiti zagrade prilikom definisanja povratne vrednosti.

Kada pozivamo funkcije, bitno je da "prihvatimo" sve vrednosti koje funkcija vraća. Sledeći primer nije ispravan jer ne možemo da upišemo 2 vrednosti u jednu promenjivu.

```go
a := test_funkcija(1, 2, "test")
```

Ako nas ipak interesuje samo prva promenjiva, a druga ne, koristićemo sledeću sintaksu:

```go
a, _ := test_funkcija(1, 2, "test")
```

Video o strukturama podataka možete pogledati <a href='https://youtu.be/Crb5irYjq60'>ovde</a>.
Video o interfejsu možete pogledati <a href='https://youtu.be/h6Npp4Bfymc'>ovde</a>.
