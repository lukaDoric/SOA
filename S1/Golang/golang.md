# Uvod u Golang

Go (Golang) je programski jezik razvijen od strane Google-a krajem 2000-ih godina. U početku se koristio kao interni alat samo unutar kompanije, da bi krajem 2009. jezik objavili u vidu _open-source_ projekta. Među kreatorima jezika nalazi se i **Ken Thompson**, tvorac _Unix_ operativnog sistema. 

### Osnovne karakteristike Go-a

- **Kompajliran** jezik, nalik na C i C++. Za pokretanje programa napisanih u Go-u nije potrebno ništa osim binarne datoteke koja se kreira prilikom kompajliranja. Nema potrebe za izvornim kodom i dodatnim alatima poput virtualnih mašina unutar kojih se kod izvršava (Java) ili interpretera (Python).
- Koristi “**garbage collector**” kao i Python i Java, te korisnici ne moraju da vode računa o alociranju i dealociranju memorije prilikom pisanja koda.
- Go je razvijen u trenutku kada su multiprocesorske mašine i procesori sa više jezgara bili već uveliko u širokoj upotrebi. Kreatorima jezika je jedan od prioriteta bio omogućavanje jednostavnog **pisanja konkurentnog koda**, i njegovo efikasno izvršavanje. Go ima podršku za rad sa nitima (go-rutinama) ne u vidu eksternih biblioteka, već u vidu funkcionalnosti koje dolaze zajedno sa samim jezikom, poput kreiranja funkcija, struktura ili promenjivih.
- Jednostavno **testiranje koda**, te ni ovde nema potrebe za korišćenje eksternih biblioteka i/ili alata.

### Zašto Go?

Cilj predmetnog projekta je izrada _mikroservisne_ veb aplikacije. O ovoj temi ćete detaljno čuti u narednim terminima vežbi, ali za sada možemo da kažemo da ćemo pisati nekoliko "jednostavnijih" programa umesto jednog kompleksnog, od kojih će svaki obavljati deo funkcionalnosti naše veb aplikacije. 

Kada pišemo kompletnu veb aplikaciju u vidu jednog programa koji koristi gigabajte RAM memorije tokom svog izvršavanja, i koji će raditi na tradicionalnim veb serverima, _overhead_ koji nose dodatni alati neophodni za rad same aplikacije (poput virtualne mašine ili interpretera) ne predstavlja veliki problem (procentualno troše malo resursa u odnosu na samu aplikaciju). Međutim, ako hoćemo da tu istu veb aplikaciju razdelimo na 10 ili 15 jednostavnijih programa, od kojih će svaki trošiti svega nekoliko stotina MB RAM memorije, odjednom _overhead_ koji nose propratni alati postaje primetan. Ovaj problem dolazi još do većeg izražaja ako uzmemo u obzir da ćemo te iste servise verovatno pokretati na mašinama _cloud provider_-a poput AWS-a, koji svoje usluge naplaćuju, delom, na osnovu resursa koje naši servisi troše. Zbog ovoga je prilikom razvijanja mikroservisnih aplikacija poželjno korišćenje kompajliranih jezika, poput Go-a.

Jedna od glavnih mana kompajliranih jezika jeste distribucija programa. Moramo posebno kompajlirati program za svaki tip mašine na kojoj će on biti pokrenut. Međutim, ako uzmemo u obzir da se naša mikroservisna aplikacija neće pokretati na korisničkim računarima, već na serverima gde mi imamo kontrolu nad hardverom, ovo nam neće predstavljati problem.

Pored toga Go ima _out-of-the-box_ podršku za rad sa mrežnim protokolima poput **HTTP**-a i **gRPC**-a, što je još jedan plus s obzirom da pišemo veb servise.

Na sledećem <a href='https://github.com/lukaDoric/SOA/blob/main/S1/Golang/prviProgram.md'>link-u</a> možeš videti kako da pokreneš prvi Golang program.  
Na sledećem <a href='https://github.com/lukaDoric/SOA/blob/main/S1/Golang/strukturePodataka.md'>link-u</a> možeš videti strukture podataka i osnovnu sintaksu u Golang-u.  
Na sledećem <a href='https://github.com/lukaDoric/SOA/blob/main/S1/Golang/slice.md'>link-u</a> možeš videti mape i slice-ove u Golang-u.  
Na sledećem <a href='https://github.com/lukaDoric/SOA/blob/main/S1/Golang/konkurentni.md'>link-u</a> možeš videti pisanje konkurentnog pograma u Golang-u.  
Na kraju <a href='https://github.com/lukaDoric/SOA/blob/main/S1/Golang/rest.md'>ovde</a> možeš videti kako da napišeš čitav REST servis u Golang-u.  
