# Nizovi, slice-ovi i mape

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
