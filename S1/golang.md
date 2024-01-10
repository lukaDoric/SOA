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

### Strukture podataka
Video o strukturama podataka možete pogledati <a href='https://youtu.be/Crb5irYjq60'>ovde</a>.  
Video o interfejsu možete pogledati <a href='https://youtu.be/h6Npp4Bfymc'>ovde</a>.

## Nizovi, slice-ovi i mape

## Nizovi

Go podržava rad sa statičkim nizovima. Prilikom deklaracije ili inicijalizacije ovih nizova, neophodno je navesti i tip podataka koji će sadržati, i broj elemenata. U primeru ispod možete videti kako se statički nizovi deklarišu, inicijalizuju i kako se pristupa određenim elementima.

```go
var klubovi [2]string                       // prazan niz koji sadrži "nulte" vrednosti, što su u ovom slučaju dva prazna stringa
klubovi := [2]string{"Javor", "Cukaricki"}   // deklaracija i inicijalizovanje niza
klubovi[1] = "Spartak"                      // promena drugog elementa
fmt.Println(klubovi[2])                     // index out of bounds greška
```

## Slices

Statički nizovi se, međutim, veoma retko pojavljuju u go-u, već se umesto njih uglavnom koriste `slice`-ovi. Slice je apstrakcija nad statičkim nizom. Sastoji se od 3 podatka:
1. Pokazivač na statički niz nad kojim je napravljen
2. Informaciju o dužini slice-a (Statički niz može da sadrži 5 elemenata, ali da su samo prva 3 deo slice-a. Zato nam je, pored pokazivača na startnu poziciju statičkog niza, potrebna i dužina slice-a)
3. Kapacitet. Sve elemente koje upisujemo u slice zapravo moramo upisati u neku memorijsku lokaciju. To će uvek biti na nekoj poziciji unutar statičkog niza (ako hoćemo da upišemo broj na početku slice-a, u pozadini će se broj upisati na početak statičkog niza). Zbog ove činjenice i dalje ne smemo da prekoračimo dužinu statičkog niza, i da pokušamo da upišemo na 6. poziciju u slice-u neki element, ako je statički niz dužine 5: dobićemo index out of bounds grešku. Da bi ovo izbegli često nas interesuje koliko imamo mesta u slice-u, to jest koliki je kapacitet slice-a.

Primer definisanja slice-a nad statičkim nizom:

```go
var slice = []int{1,2,3,4,5}  // Sintaksa je ista kao za inicijalizaciju statičkog niza, samo nema broj 
                             // u uglastim zagradama koji precizno definiše dužinu
fmt.Println(slice) // 1,2,3,4,5
```
U primeru gore smo kreirali statički niz od 5 elemenata (1,2,3,4,5), podesili pokazivač slice-a na prvi element tog niza, i podesili dužinu i kapacitet slice na 5

Slice možemo takođe da definišemo nad već postojećim nizom tako što ćemo "preseći" (iliti slice-ovati) deo tog niza
```go
var staticki_niz = [5]int{1,2,3,4,5} 
slice := staticki_niz[1:3]
fmt.Println(slice) // 2,3
```
Ovde smo kreirali slice tako što smo namestili pokazivač na drugi element statičkog niza (ne početak niza!), podesili dužinu na 2, a kapacitet na 4
Ako pokušamo da pristupimo trećem elementu slice-a, dobićemo index out of bounds grešku, iako statički niz ispod ima još 2 elementa ekstra. Kako onda da dodamo element na slice - uz pomoć funkcije `append`

```go
// Kod se nastavlja na prethodni primer
slice = append(slice, 9)
fmt.Println(slice) // 2,3,9
```

Funkcija append radi tako što upisuje novi element u statički niz koji se nalazi "ispod" slice-a sve dok u tom statičkom nizu ima mesta. U trenutku kad izažemo izvan opsega statičkog niza, go kreira novi statički niz koji je duplo duži i prepiše sve vrednosti iz starog niza, + novu vrednost koju želimo da append-ujemo.

Ne moramo da dodajemo jedan po jedan element, možemo više elemenata odjedno ili kompletan slice na drugi slice

```go
// Kod se nastavlja na prethodni primer
slice = append(slice, 7,6,5)
fmt.Println(slice) // 2,3,9,7,6,5
slice2 := []int{11,10,9} 
slice = append(slice, slice2)
fmt.Println(slice) // 2,3,9,7,6,5,11,10,9
```

Postoji i funkcija `make` pomoću koje možemo da, izmežu ostalog, kreiramo i slice-ove:

```go
slice = make([]string, 5) //kreira slice nad nizom od 5 praznih stringova
```
ili
```go
slice := make([]string, 0, 5) //kreira slice sa kapacitetom od 5 elemenata, i sa dužinom od 0 elemenata
// Ovo je ekvivalentno kao da smo uradili:
staticki_niz := [5]string{}
slice := staticki_niz[:0]
```

Ako kasnije hoćemo da prođemo kroz slice, možemo koristiti `range`:

```go
for index, value := range slice {
  fmt.Printf("Na mestu broj %d se nalazi broj %d", index, value)
}
```

Sliceovi služe da nam malo olakšaju upravljanje statičkim nizovima, jer go radi za nas u pozadini kreiranje novog statičkog niza sa većim kapacitetom kada je to neophodno. 

Sve što je ovde napisano je sa još više detalja opisano u [ovom članku](https://go.dev/blog/slices-intro) na zvaničnoj go stranici.

Snimak o slice-ovima možete pogledati <a href='https://youtu.be/ZupwE3TCgdY'>ovde</a>.  

## Mape

Key-value struktura u go-u naziva se mapa. Prilikom kreiranja mapa treba da definišemo tip podataka koji ćemo koristiti kao ključ, i tip koji ćemo koristiti kao vrednost.

```go
mapa := make(map[int]string)
// Ili možemo odma da uradimo i inicijalizaciju
mapa := map[int]string{1: "Ivan", 32451: "Jovan"}
// Pristup i modifikacija mape
mapa[2]="Milana"
// Brisanje elemenata
delete(mapa, 32451)

// Iteracija kroz mapu
for key, value := range mapa {
  fmt.Printf("Key: %d-value %s\n", key, value)
}
```
