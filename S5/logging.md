Kada je u pitanju prikupljanje logova aplikacija koje se izvršavaju u kontejneru, sve što je neophodno uraditi jeste konfigurisati da aplikacija upisuje logove na standardni izlas što je po default-u podešeno. Odnosno nema nikakvog dodatnog posla da bi se omogućilo prikupljnaje logava aplikacije. Docker prikuplja logove i privremeno ih skladišti lokalno na fajl sistem sve dok kontejner ne prestane sa radom. Pristupanjem logova se radi pomoću komande:

 `docker log <container -name>`

Za trajno skladištenje logova kontejnera neophono je koristiti odgovarajući alate:

- Log baza: baza koja trajno čuva logove. Treba da omogući indeksiranje logova radi brze pretrage logova. 
- Log agregator: prikuplja logove sa nekog izvora, normalizuje ih i prosledjuje ih u odgovarajuću log bazu.

U konkretnom primeru kao log baza se koristi Loki. Fluent-bit je agregator koji prikuplja docker logove svih kontejnera i prosleđuje Loki-u. Takođe, Grafana se koristi za vizualizaciju logova tako što koristi Loki kao datasource.

Video materijal: https://youtu.be/wWw3fRR9NTk
