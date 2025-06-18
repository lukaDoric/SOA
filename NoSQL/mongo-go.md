## Dokument NoSQL

U prethodnom delu lekcije smo se kratko osvrnuli na dokument NoSQL baze podataka. 

Kao glavne karakteristike jedne dokument baze može se izdvojiti:
-  Intuitivan rad sa dokumentima i kolekcijama. 
-  Fleksibilna šema podataka koja nam omogućava da menjamo dokumente kako razvijamo našu aplikaciju.
 - Horizontalno skaliranje.

Dokument baze su NoSQL baze koje umesto da skladište podatke u fiksiranim redovima i kolonama tj. tabelama koriste fleksibilne dokumente (najčešće JSON tj. BSON format - https://www.mongodb.com/json-and-bson). Na primeru ispod se može videti jedan JSON dokument. Ako bismo poredili dokument NoSQL baze sa SQL bazama moglo bi se reći da je dokument jedan red iz tabele a kolekcija sličnih dokumenata cela tabela, svakako tabele i kolekcije dokumenata se i dalje suštinski razlikuju.

### MongoDB

Kao konkretnu dokument NoSQL bazu u okviru vežbi koristićemo MongoDB NoSQL.

Uporedni pregled pojmova u okviru MongoDB i SQL baze podataka:

Baza podataka  -  Baza podataka  
Kolekcija            -  Tabela  
Dokument          - Torka  
Polje                   -  Kolona

### Dokumenti

Dokument obično sadrži informacije o jednom objektu i svim njegovim povezanim metapodacima. Dokumenti čuvaju podatke u parovima polje-vrednost. Vrednosti mogu biti različitih tipova i struktura, uključujući nizove, objekte, stringove, brojeve i datume. Dokumenti mogu biti u formama kao što su JSON, BSON i XML.

Na primeru ispod možemo videti jedan dokument u JSON formatu

```
{
     "_id": 1,
     "first_name": "Tom",
     "email": "tom@example.com",
     "cell": "765-555-5555",
     "likes": [
        "fashion",
        "spas",
        "shopping"
     ],
     "businesses": [
        {
           "name": "Entertainment 1080",
           "partner": "Jean",
           "status": "Bankrupt",
           "date_founded": {
              "$date": "2012-05-19T04:00:00Z"
           }
        },
        {
           "name": "Swag for Tweens",
           "date_founded": {
              "$date": "2012-11-01T04:00:00Z"
           }
        }
     ]
  }
```

### Kolekcije

Kolekcija predstavlja grupu dokumenata. Kolekcije obično skladište dokumente koji imaju sličan sadržaj. Dokumenti u okviru kolekcije ne moraju imati ista polja, jer dokumenti NoSQL baze imaju fleksibilnu šemu.

Oslanjajući se na primer iznad bilo bi logično da korisnika smestimo u users kolekciju. Dodatno možemo ubaciti još korisnika u users kolekciju ali oni ne moraju imati ista polja (primer ispod).

```
 {
     "_id": 2,
     "first_name": "Donna",
     "email": "donna@example.com",
     "spouse": "Joe",
     "likes": [
        "spas",
        "shopping",
        "live tweeting"
     ],
     "businesses": [
        {
           "name": "Castle Realty",
           "status": "Thriving",
           "date_founded": {
              "$date": "2013-11-21T04:00:00Z"
           }
        }
     ]
  }
```


Više možete pročitati na https://www.mongodb.com/document-databases.

U narednom <a href='https://www.youtube.com/watch?v=aL7mN_l2jLM'>videu</a> možeš pogledati osnovne komande nad mongo bazom.  
U narednom <a href='https://www.youtube.com/watch?v=06ehbvxflL4'>videu</a> možeš pogledati integraciju mongo baze sa Golang-om.  
Primer je dostupan na sledećem <a href='https://drive.google.com/file/d/1mCErpAIp7M9aArSQ_3guk0VteRvXr-KN/view?usp=sharing'>link-u</a>.  
