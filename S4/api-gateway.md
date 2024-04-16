## API gateway

Mikroservisna arhitektura podrazumeva postojanje velikog broja međusobno nezavisnih servisa koji rade kao celina i opslužuju zahteve klijenata aplikacije. Neki od osnovnih problema sa takvom organizacijom su da:

- Klijent u sebi mora sadržati logiku za otkrivanje i adresiranje svakog od servisa koji treba da kontaktira
- Dalja dekompozicija servisa se komplikuje zbog toga što klijent direktno zavisi od unutrašnje organizacije sistema
- Protokoli koje servisi koriste nisu uvek pogodni

Kako bi se navedeni nedostaci prevazišli, uvodi se komponenta koja se naziva API Gateway. API Gateway je servis koji predstavlja jedinstvenu ulaznu tačku u sistem, enkapsulira unutrašnju arhitekturu i nudi API koji klijenti mogu koristiti.

<img src="https://media.licdn.com/dms/image/D4E12AQGvHr7xO9uDgQ/article-cover_image-shrink_423_752/0/1664942913778?e=1684972800&v=beta&t=HD9hAyUV4BkpR4Wb5UNO66UguCY1B0of6aHbj48w-Ao" />

Glavna zaduženja API Gateway-a su:

- Rutiranje pristiglih zahteva
- API Composition (detaljnije o ovoj temi na narednim vežbama)
- Translacija protokola

Svi zahtevi koje eksterni klijenti šalju prvo stižu do API Gateway-a koji u zavisnosti od samog zahteva vrši prosleđivanje zahteva odgovarajućem servisu. Ukoliko zahtev podrazumeva dobavljanje podataka koje poseduju različiti servisi, API Gateway može uputiti odgovarajuće zahteve i agregirati dobijene rezultate, a zatim ih proslediti klijentu. Što se tiče translacije protokola, API Gateway se često implementira tako da nudi RESTful API zbog rasprostranjenosti ovog arhitektonskog stila i raznovrsnosti klijenata koji ga podržavaju, dok aplikativni servisi interno mogu koristiti neke druge protokole, kao što je na primer gRPC. Kada postoje neslaganja u protokolima, API Gateway je zadužen da obavi neophodnu translaciju između eksternog i internog API-ja. 

Nije redak slučaj da API Gateway ima i dodatna zaduđenja, takozvane edge funkcionalnosti. Neke od njih su:

- Ograničenje stope pristupa
- Keširanje
- Prikupljanje metrika
- Logovanje
- Kontrola pristupa
- Terminacija SSL-a
- ...

Kada uvedemo API gateway komponentu, moramo uzeti u obzir i neželjene pojave kao što su:

- Postojanje još jedne komponente koje se mora razvijati i održavati, a pritom mora biti visoko dostupna
- Neophodno je redovno ažurirati API Gateway kako bi API-ji koje servisi nude bili vidljvii spoljašnjem svetu
- Kako predstavlja dodatnu tačka između klijenta i mikroservisa, uvodi se dodatno kašnjenje u obradu svakog zahteva
