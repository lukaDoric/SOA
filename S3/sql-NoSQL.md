## SQL baze podataka

Kako ste do sada imali prilike da radite sa SQL bazama podataka (poput PostgreSQL-a) prvi deo lekcije biće posvećen njihovim karakteristikama. Kao motivaciju za uvođenje NoSQL baza podataka u drugom delu biće predstavljene osnovne karateristike NoSQL baza.

Relacione baze podataka (SQL) su alat za skladištenje i preuzimanje **strukturiranih podataka**. U većini slučajeva u okviru SQL-a postoje tabele koje su mapirane na model podataka (klase). SQL baze nazivaju i **relacionim bazama** jer su podaci koje skladištimo (tabele) uglavnom u relacijama (vezama). U okviru relacionih baza podataka podaci su smešteni u okviru tabela, a tabele su povezane jedinstvenim ključevima.

Tabela Users
```
|Id |Username | Password | Role |
|-- |-------- | -------- |----- |
| 0 | pera@12 | pera123  |0     |
| 1 | marko.2 | marko123 |1     |
```

Tabela Instructors
```
|Id  |UserId   | Name     | Surname  |
|--- |-------- | -------- |--------- |
|100 | 0       | Pera     | Peric    |
```

Tabela Students
```
|Id  |UserId  | Name     | Surname  | Index   | IsArchieved
|--- |------- | -------- |--------- |-------- |-----------
|120 | 1      | Marko    |Markovic  |TT-2-2022| False
```

Dakle imamo tabelu sa svim korisnicima (podaci vezani za kredencijale korisnika), potom dve tabele koje se odnose na dve uloge korisnika (instruktore i učenike). Kako bismo izvukli informacije o studentu preko korisničkog imena moramo spojiti tabele korisnika i studenta. Bilo koju informaciju koju skladištimo relativno lako možemo izvući pomoću SQL upita.

SQL baze podataka se obavezno koriste u domenima gde imamo potrebe za **kompleksnim transakcijama** (banke) jer su "stabilnije" od NoSQL baza i pružaju nam **integritet** nad podacima.

Podsetimo se ACID principa vezanih za transkacije:

1. Atomičnost (engl. Atomicity): Svi delovi transakcije se izvršavaju u potpunosti i uspešno, inače cela transakcija propada.
2. Konzistentnost (engl. Consistency): Podaci ostaju konzistentni u čitavoj relacionoj bazi podataka.
3. Izolacija (engl. Isolation): Svaka transakcija je nezavisna od drugih transakcija.
4. Otpornost (engl. Durability): Čak i ako sistem padne, podaci iz završenih transakcija su sigurno uskladišteni.

## NoSQL baze podataka

NoSQL pokret reprezentuje sve tipove baza podatka koji nisu relacione baze podataka. Njihov **model skladištenja, upita i svih ostalih operacija** suštinski se razlikuje od relacionih baza.

Samo čuvanje podataka je drugačijeg oblika od tabela i relacija koje se sreću u relacionim bazama. Za korišćenje SQL baze ćemo se opredeliti kada imamo strukturirane podatke dok ćemo NoSQL iskoristiti kada imamo i strukturirane i **polu strukturirane i nestrukturirane podatke**. Što znači da NoSQL **nema striktnu šemu** za podatke sa kojima barata. Dodatno SQL upiti nam više nisu na raspolaganju već svaka NoSQL baza koristi **specifičan jezik za upite**.

Sledeći primer ilustruje format skladištenja u okviru dokument NoSQL baze podataka o kojima će biti više reči u nastavku lekcije.

```
{  
     "_id": 1,
     "name": pera,
     "surname": peric,
     "address":
          {
               "country": "Serbia",
               "city": "Novi Sad",
               "street": "Bulevar Slobodana Jovanovića"
          }
}
```
```
{
     "_id": 2,
     "name": mika,
     "surname": mikić,
}
```

Primer iznad ilustruje dva **dokumenta** u okviru NoSQL dokument baze. Dokument predstavlja informacije o jednom objektu i obično je u JSON (BSON) formatu. Svako polje objekta koje čuvamo može biti različitog tipa string, number, boolean, drugi objekat itd.

Konkretno u primeru iznad imamo skladištena dva korisnika pri čemu prvi ima i adresu, dok drugi nema adresu. Ovakvo ponašanje je karakteristično za NoSQL baze jer **nemaju fiksnu šemu podataka, tj. ona je fleksibilna**.

Ovo je samo jedan od primera NoSQL baze. Pored dokument baze postoje i Key-Value, Graph, Wide-Column i Search Engine NoSQL baze. U okviru ovog kursa akcenat će biti na dokument i graf NoSQL bazama.

## Zbog čega NoSQL?

Na važnosti počinju da dobijaju još u ranim nastancima cloud computing-a, iz prostog razloga nemogućnosti relacionih baza da se primene u distribuiranim sistemima. CAP teorema definisana od strane dr Eric Brewer-a kaže da ne možemo zadovoljiti sve tri osobine u okviru baze: **Konzistentnost, Dostupnost i Particioniranje** i da uvek možemo da dobijemo najviše dve od tri osobine. Na arhitektama je da odluče šta od tih osobina žele da imaju (particioniranje nije moguće izbeći u distribuiranim sistemima). Slika 1 prikazuje CAP teoremu, sa osobinama i predstavnicima baza podataka (više o ovome: https://www.ibm.com/topics/cap-theorem).

<img src="https://i.ibb.co/b28MhyW/CAP.png" alt="CAP" border="0" />


Koncept NoSQL baza postaje popularan sa gigantima kao što su Google, Facebook, Amazon itd. Sa rastom svojih proizvoda imali su potrebu za skladištenjem velike količine podataka a RDBMS se nije pokazao dobro u
takvoj situaciji. Kako bi rešili ovaj problem mogli su skalirati sistem vertikalno (dodavanje još procesorske snage, memorije itd.) ali je ovo dosta skupo. Alternativa je da distribuiramo naš sistem na vise host-ova odnosno da skaliramo sistem horizontalno (slika ispod). Tu dolazi do isticanja NoSQL baza jer se skaliraju bolje nego RDBMS.

<img src="https://i.ibb.co/vx7gMgZ/SCALE.png" alt="SCALE" border="0">

Ovi tipovi baza (NoSQL) obično **ne podržavaju join operacije niti imaju transakcije kao relacione baze i tako striktnu formu podataka**. Zbog ovih osobina izuzetno se lako koriste u mikroservisima i pružaju mogućnost da vaš domen problema prilagodite bazi a ne da se dovijate fiksnom modelu relacionih baza. To ne znači da transakcije nije moguće implementirati, ali se one obično implementiraju na aplikativnom sloju (npr. SAGA patern o kojem će biti reči u nastavku kursa).
Sa druge strane sa obzirom da nemamo operaciju join (takvu kakvu poznajemo u SQL upitima) moramo da poznajemo svoje upite kako bismo se odlučili koju NoSQL bazu ćemo koristiti. Sa obzirom da su mikroservisi relativno mali i trebalo bi da imaju jednu odgovornost ovo ne predstavlja velik problem.





