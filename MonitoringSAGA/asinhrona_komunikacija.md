### Komunikacija mikroservisa

Jedan od najvećih izazova prilikom prelaska sa monolitne na mikroservisnu arhitekturu aplikacije jeste izmena u načinu komunikacije između komponenti
(lokalni vs udaljeni pozivi procedura). Sinhroni pozivi zahtevaju istovremenu dostupnost svih učesnika (problem znatno izraženiji kada imamo ulančane pozive).
Težimo da komunikaciju između servisa smanjimo na minimum, kao i da ti pozivi, kad god je moguće, ne budu blokirajući (**RPC** vs **messaging**).

Tipove komunikacije možemo podeliti na dva načina:
- Prema prirodi komunikacionog protokola:
   - Sinhroni - Klijent šalje zahtev i blokiran je sve dok ne dobije odgovor od servera
   - Asinhroni - Pošiljalac šalje poruku ne čekajući na odgovor klijenta (ukoliko ga uopšte očekuje)
 - Prema broju primalaca poruke/zahteva:
   - Jedan - Zahtev dobija i obrađuje tačno jedan primalac
   - Više - Klijent šalje poruku preko nekog kanala, iz kog poruke preuzima i obrađuje potencijalno više primalaca

### Messaging

Asinhrona komunikacija, gde servisi razmenjuju poruke upotrebom kanala
  - Možemo ostvariti različite stilove komunikacije:
     - **Request/response** - šaljemo poruku jednom primaocu i čekamo odgovor
     - **Notifcations** - šaljemo poruku jednom primaocu ne očekivajući nikakav odgovor
     - **Request/asynchronous response** - šaljemo poruku jednom primaocu i očekujemo da ćemo nekada u budućnosti dobiti odgovor
     - **Publish/subscribe** - šaljemo poruku koju može preuzeti više primalaca
     - **Publish/asynchronous response** - šaljemo poruku koju može preuzeti više primalaca, koji nam nakon toga mogu vratiti odgovor

### NATS

NATS je posrednik za razmenu poruka između servisa

Razlikujemo dva tipa učesnika u komunikaciji:
- Publisher - Objavljuje poruku na neku temu (eng. subject)  
- Subscriber - Prijavljuje se da prima poruke iz određenog subject-a

Svaka poruka sadrži:  
 - subject
 - payload (niz bajtova)
 - headers
 - reply subject (opciono)

<img src="https://i.ibb.co/z4jtrL5/nats1.png" alt="nats1" border="0">


#### Queue grupa

Dodatna funkcionalnost koju NATS nudi je queue  
- Svaki subscriber preuzima poruke iz jednog subject-a
- Subscriber-u se može dodeliti queue kom pripada, a svi subscriber-i koji pripadaju istom queue-u čine queue grupu
- Svaka poruka koja se objavi u subject biće isporučena samo jednom članu queue grupe

<img src="https://i.ibb.co/6cQwp6J/nats2.png" alt="nats2" border="0">

**Napomena:** U okviru projekta ideja je da sinhrone pozive između servisa razvijete putem gRPC-a, dok bi asinhrone pozive trebalo realizovati kroz NATS (SAGA pattern o kom će biti više reči kasnije).

Video materijal: https://www.youtube.com/watch?v=JHMekgDkc-8

Primeri: https://github.com/lukaDoric/SOA/tree/main/MonitoringSAGA/NATS_primer
