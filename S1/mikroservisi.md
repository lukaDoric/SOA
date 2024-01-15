# Osnovni koncepti mikroservisne arhitekture

Kroz prvi sprint ćeš već upoznati osnovne koncepte mikroservisne arhitekture.

Mikroservisna arhitektura spada u tipove servisno orijentisanih arhitektura. Mikroservisna aplikacija je kolekcija međusobno slabo zavisnih servisa modelovanih oko poslovnog domena.
Svaki servis enkapsulira određenu funkcionalnost i čini je dostupnom drugim servisima preko mreže. Na slici ispod možemo videti primer dijagrama arhitekture mikroservisne aplikacije.
Za početak mikroservisnu aplikaciju možeš zamisliti kao modularni monolit (koji smo već gradili) samo što ne možemo direktno da pozivamo servise jednog modula iz drugog već to moramo učiniti kroz http (kasnije rpc) zahteve.

<img src="https://th.bing.com/th/id/R.981b89570cb0eac2a5e6694468c09b4b?rik=RuRxeRr68LOrKw&riu=http%3a%2f%2fkhoadinh.github.io%2fassets%2fmedia%2fmicroservices_architecture_diagram.png&ehk=WhJl0eYtq9%2b%2fUxwTZxlGym7p8GRcp01fWrANAr%2bSuFU%3d&risl=&pid=ImgRaw&r=0" alt="diagram" border="0" width=70% />

Servis skriva detalje implementacije poslovne logike i strukturu podataka koje čuva, a spoljašnjem svetu prikazuje API (REST, RPC, asinhroni) preko kog klijenti mogu interagovati sa servisom.
Iz perspektive klijenata, svaki servis je black box koji ima nepromenljiv ili retko promenljiv interfejs za komunikaciju. Na ovaj način stvaramo jasnu granicu između servisa i njegovog okruženja.
Izmene u implementaciji ne treba da utiču na klijente servisa - interfejs skriva detalje koji se često menjaju.
Iz ovog razloga, deljenje na primer skladišta podataka između servisa se snažno obeshrabruje zbog toga što je format podataka često promenljiva kategorija, a deljenjem  i taj format implicitno postaje deo interfejsa preko kog servisi komuniciraju.
Kada su servisi jako međusobno zavisni, gubi se velika većina benefita ovakve arhitekture.

## Ključni koncepti mikroservisa

### Nezavisan deployment

Osnovna ideja je da izmenu u nekom servisu možemo napraviti i učiniti dostupnom korisnicima, a da ona ne utiče na rad drugih servisa. Kako bismo to postigli, neophodno je da servisi budu slabo uvezani. Stabilan i jasno definisan API je preduslov za postizanje ovog cilja i predstavlja osnovnu heuristiku pri postavljanju granica između servisa. Ako konstantno menjamo interfejs - dogovor koji imamo prema drugim servisima, timovi koji rade na tim servisima će biti u obavezi da menjaju implementaciju svojih servisa pri svakoj izmeni koju napravimo.

### Modelovanje oko poslovnog domena

Kada se vodimo ciljem kreiranja slabo zavisnih servisa, granice vođene poslovnim domenom umesto tehničkim funkcijama (perzistencija, poslovna logika, prezentacija) su znatno bolji izbor. Izmena ili dodavanje nove poslovne funkcionalnosti najčešće obuhvata izmene u svim tehničkim slojevima, što znači da želimo da svi ti slojevi budu deo jednog servisa. Sa druge strane, nepovezane poslovne funkcionalnosti mogu se deploy-ovati nezavisno i iz toga sledi da ćemo ih smestiti u različite servise. Ukratko, preferiramo koheziju poslovnog domena iznad tehničke kohezije.

Dobre smernice za određivanje granice mikroservisa nudi nam domain-driven design i pojam poddomena i bounded konteksta. Poddomen može biti neka poslovna celina dok bounded kontekst možemo posmatrati kao pogled na domen iz određene perspektive. Recimo, u aplikaciji za upravljanje porudžbinama treba da rukujemo proizvodima. Iz perspektive upravljanja zalihama, bitno nam je da znamo koliko imamo dostupnih proizvoda (određene boje, veličine, proizvođača itd). Sa druge strane, modul za naplatu porudžbina treba da bude svestan cene proizvoda, popusta itd. Različiti bounded konteksti uključuju različite aspekte (atribute) entiteta. Često se može napraviti mapiranje bounded konteksta na granice mikroservisa. To mapiranje ne mora uvek biti 1 - 1, uvek postoje izuzeci jer je broj faktora koji utiče na pristup dekompoziciji veliki. Na slici ispod možemo videti kako entitet proizvoda može da se "podeli" na domenske modele različitih poddomena, a samim tim često i mikroservisa. U slučaju vašeg modularnog monolita aplikacija je već logički izdeljena na servise samo ih treba fizički razdvojiti.

<img src="https://i.ibb.co/PQ6w2KC/bc1.png" alt="diagram" border="0" width=40% />
<img src="https://i.ibb.co/7JJstKw/bc2.png" alt="diagram" border="0" width=40% />

### Skrivanje informacija

Jedno od osnovnih pravila koje mikroservisi treba da prate je skrivanje informacija - enkapsulacija. Sa pojmom enkapsulacije ste se upoznali još dok ste savladavali objektno orijentisani programiranje. Princip se poprilično lepo mapira i na mikroservise. Svaki mikroservis koristi bazu koja mu najviše odgovara, napisan je u bilo kojoj tehnologiji itd. i klijenti toga ne treba da budu svesni. Sve što njih treba da zanima su operacije koje servis nudi preko svog interfejsa. Ako su nam potrebni podaci iz nekog servisa, pitamo taj servis da ih dobavi, ne idemo direktno do baze podataka. Isto tako, interfejs treba da bude nezavisan od neke konkretne tehnologija, kako bi se integracija i komunikacija što lakše ostvarile.

### Veličina

Jedna od najčešćih nedoumica vezana za mikroservisnu arhitekturu je koliko mali treba biti mikroservis i kakve sve funkcionalnosti treba da obuhvati. Na ovo postoji mnoštvo odgovora koji su vrlo zavisni od konteksta. Broj linija koda ni broj funkcionalnosti nisu preterano korisne metrike. Manjim servisima mogu upravljati manji timovi, ali koordinacija između timova i servisa postaje složenija jer je neizbežno da će broj funkcionalnosti koje se protežu kroz više servisa porasti, što između ostalog uvodi i dodatan latency u sistem. Sa druge strane, preobimni servisi polako degradiraju ka monolitu i potreban je veliki broj ljudi za održavanje codebase-a. Kao zaključak, sama veličina servisa ne treba da bude glavna vodilja pri dizajnu, već implikacije koje neka podela ima na razvoj i performanse aplikacije.

## Prednosti mikroservisa (u odnosu na monolit)

Pre nego što počnemo da pričamo o prednostima mikroservisnih u odnosu na monolitne aplikacije, vrlo je bitno napomenuti da ove prednosti nisu garantovane ukoliko se prilikom dizajna ne ispoštuju koncepti mikroservisne arhitekture. Zbog toga je taj korak jedan od najvažnijih, ako ne i najvažniji, prilikom razvoja ovakvih aplikacija.

### Heterogenost tehnologija

Odabir tehnološkog steka u slučaju monolita utiče na sve koji su uključeni u rad projekta. Izmena programskog jezika ili radnog okvira je u zrelim projektima retkost zbog toga što se čitava aplikacija mora prepisati i rizik je izuzetno visok. Mikroservisi nas oslobađaju ovog problema jer je osnovna stvar na koju se oslanjamo tehnološki agnostičan interfejs. To nam omogućava da svaki servis implementiramo u jeziku koji zadovoljava naše prioritete (performanse, lakoća razvoja itd). Pored toga, svaki servis u potpunosti upravlja svojim podacima, tako da se skladište podataka može birati na nivou servisa, tako da ono bude najpogodnije za format podataka i upite koje će servis obavljati. Na slici ispod možemo videti koliko različitih tehnologija može biti uključeno u implementaciju samo tri servisa.

<img src="https://i.ibb.co/LNsFpdm/image-VWA211.png" alt="diagram" border="0" width=50% />

### Skalabilnost

Horizontalnom skaliranju sistema u slučaju monolitne aplikacije može se pristupiti na samo jedan način - pokretanjem novih instanci čitave aplikacije. Iako pristup nije optimalan jer su možda samo jedan ili dva modula preopterećena, nemamo drugih opcija. U slučaju mikroservisne aplikacije, situacija je dosta fleksibilnija - možemo podići nove instance samo onih servisa koji su preopterećeni. Na taj način, servisi se mogu pokretati i na resursno ograničenijem hardveru. Recimo, veliki broj cloud servisa bazira se na ovakvoj vrsti elastičnosti, podižemo nove instance samo delova sistema i gasimo ih čim više nemamo potrebe za njima. Kada radimo sa više instanci istog servisa, izuzetno je poželjno da oni budu stateless, jer na taj način izbegavamo razne probleme vezane za koordinaciju i sinhronizaciju podataka.

<img src="https://i.ibb.co/KzKxDXn/image-MQ9-U11.png" alt="diagram" border="0" width=70% />

### Olakšan deployment

Svaka izmena u monolitu zahteva ponovni deployment čitave aplikacije. Takva operacija je vrlo rizična i samim tim se često čeka da se izmene nakupe da bi se obavila. Međutim, u tom slučaju verovatnoća neispravno unetih izmena raste, što je drugi problem. Kada koristimo mikroservise, izmenu možemo izolovati na jedan servis i brzo ga ponovno deploy-ovati bez uticaja na ostatak servisa. Čak i kada unesemo bug u sistem nekom izmenom, ona se jednostavnije može izolovati na taj servis i nad njim se može primeniti rollback mehanizam.

## Problemi u mikroservisnim aplikacijama

### Latency

Komunikacija preko mreže znatno je sporija od komunikacije unutar jednog procesa ili više procesa na jednoj mašini. Kako smo rekli da servisi međusobno komuniciraju pomoću mrežnih poziva, jako je važno da i ovaj aspekt uzmemo u obzir prilikom formiranja naših sistema. Ako aplikaciju dekomponujemo na jako male servise, latency koji nastaje zbog potrebe za njihovom komunikacijom može postati previsok. Pored prirodno sporije komunikacije, mreža je nepouzdana i vrlo lako možemo doći u situaciju da ne možemo da kontaktiramo neki servis. Ovo u dizajn servisa unosi dodatna razmatranja (retry, timeout itd) koja ranije nisu bila u prvom planu, barem ne u tolikoj meri.

### Otkazi

U sistemu sa velikim brojem komponenti, neminovno je da će neka od njih u nekom trenutku prestati da radi. Moramo primeniti mehanizme koji će biti sposobni da detektuju i otklone ovakve kvarove pre nego što oni počnu da utiču na celokupan sistem i izazovu kaskadni otkaz svih komponenti. Česta situacija u kojoj se otkaz sistema kaskadno propagira je lanac sinhronih poziva između komponenti. Ako servis A poziva servis B, a servis B poziva servis C, pri čemu je servis C trenutno nedosupan, polako će početi da se troše resursi i servisa A i servisa B zbog čekanja koje smo uveli, tako da i oni mogu da postanu neresponsivni u nekom trenutku. Zbog ovakvih situacija u mikroservisnim aplikacijama prefereira se asinhrona neblokirajuća komunikacija između servisa. Kada ona nije moguća, postoje druge tehnike koje se mogu primeniti, kao što je Circuit breaker šablon.

### Konzistentnost podataka

Prelaskom na mikroservisnu arhitekturu odričemo se mogućnosti da na jednostavan način skladištimo podatke u jednoj bazi podataka i njoj prepustimo brigu o svim transakcijama i konzistentnosti. Recimo da imamo dva servisa kojima je potrebno da čuvaju neki podskup atributa korisnika, a u preseku ta dva skupa nalazi se njegov username. Kada izdamo komandu za izmenu username-a servisu koji je za to zadužen, ta izmena mora se propagirati i do svih ostalih servisa koji čuvaju ovaj podatak. Od trenutka izmene u jednoj bazi do trenutka kada je ta izmena napravljena u svim bazama, naš sistem nalazi se u nekonzistentnom stanju. Korisnici koji šalju upite mogu dobiti različite rezultate, validacije se mogu obavljati nad zastarelim podacima itd. Ovim problemom više ćemo se pozabaviti na nekom narednom terminu vežbi, ali je nešto čega moramo biti svesni.

### Monitoring

Na slici ispod možemo videti dijagram komunikacije mikroservisa u Netflix aplikaciji. Pozivi se često protežu kroz više servisa i nije jednostavno identifikovati gde je tačno uzrok nekog problema kada do njega dođe. Iz tog razloga poseban akcenat moramo staviti na observability naših aplikacija - logove, metrike i distribuirani tracing. O ovome ćete čuti dosta više na narednim terminima vežbi.

<img src="https://ibb.co/KKPFXKm" alt="diagram" border="0" width=50% />
