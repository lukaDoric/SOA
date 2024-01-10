# Strukture podataka i osnovna sintaksa

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

## Strukture podataka
Video o strukturama podataka možete pogledati <a href='https://youtu.be/Crb5irYjq60'>ovde</a>.  
Video o interfejsima možete pogledati <a href='https://youtu.be/h6Npp4Bfymc'>ovde</a>.
