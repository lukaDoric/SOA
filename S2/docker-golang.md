
Na sledećim linkovima možeš pogledati video snimke o docker slikama:  
[Prva Docker slika](https://youtu.be/xfuSdAcusfw)

<h1>Docker</h1>

<h2>Video materijal koji prati ovu lekciju:</h2>

[Docker slika za servis i bazu podataka](https://youtu.be/8cyjjYR6LzY)

[Pravilan redosled komandi prilikom definisanja Docker slike](https://youtu.be/bogpeVjMIKE)

[Multistage build](https://youtu.be/_EYy81VGrJ0)

[docker compose](https://youtu.be/LnP4ibmKhIg)

[Primer docker-compose.yaml sa više servisa i env varijablama](https://youtu.be/gtnnIr3aET4)

<h2> Kako raditi sa kontejnerima?</h2>

Pre nego što bi mogli bilo šta da radimo sa kontejnerima neophodno je izvršiti instalaciju Docker.

Detaljna uputstva za instaliranje Docker-a:
- [Linux](https://docs.docker.com/engine/install/ubuntu/)
- [Windows](https://docs.docker.com/desktop/install/windows-install/) (Nepohodno je da imate instaliran WSL 2, što bi trebalo da se nalazi na svim verizijama Windows 10 i 11)
- [Mac](https://docs.docker.com/desktop/install/mac-install/) Obratite pažnju ako imate nove računare sa Apple čipovima. Dosta docker slika ne radi kako treba na ARM arhitekturi.

Nakon instalacije neophodno je proveriti da li je instalacija bila uspešna. U terminalu otkucati komandu: sudo docker info (windows mašine: docker info).

Napomena: Ukoliko ne želite da izvršavate Docker naredbe sa povišenim privilegijama (da kucate sudo) onda je neophodno nakon instalacije ispratiti par koraka ispisanih u dokumentaciji: [https://docs.docker.com/install/linux/linux-postinstall/](https://docs.docker.com/install/linux/linux-postinstall/).

<img width="1512" alt="Screenshot 2024-03-15 at 12 37 22" src="https://github.com/lukaDoric/SOA/assets/45179708/669e684d-82a1-4be6-b602-7ecdeadea975">

Rezultat naredbe jesu informacije o broju kontejnera, broju slika, ​storage driver-​u i ostalim bazičnim konfiguracijama.

Ukoliko želimo da pokrenemo neki kontejner kucamo komandu: docker run naziv_slike. U konkretnom slučaju otkucaćemo: `docker run -i -t ubuntu /bin/bash`

<img width="861" alt="Screenshot 2024-03-15 at 12 39 11" src="https://github.com/lukaDoric/SOA/assets/45179708/0dbf8d53-e2e3-4ccb-92a3-ee97b6a11d65">

Dakle, šta se najpre dogodilo? Docker nije uspeo da pronađe sliku sa datim nazivom na lokalnom računaru pa se obratio javnom registru (DockerHub-u) i krenuo da povlači poslednju stable verziju (označena tagom latest) slike. Rekli smo da se slike sastoje iz više layer-a pa je preuzeo svaki sloj (linije koje se završavaju sa Pull complete). Nakon preuzimanja pokrenuo je nov kontejner. Ovde smo dodali i dva flega prilikom pokretanja komande. Fleg -i i -t. Prvi naglašava da je neophodno održati standard input (STDIN) dok drugi fleg dodeljuje pseudo terminal (terminal koji ima funkcije kao i pravi fizički terminal). Nakon naziva slike zadali smo i komandu koja je pokrenula Linux shell pri čemu nam se pokretanje kontejnera prikazuje kao na slici.

Kada pokrenemo top komandu unutar kontejnera vidimo da je to jedini proces koji je zapravo pokrenut u našem kontejneru.

<img width="861" alt="Screenshot 2024-03-15 at 12 39 55" src="https://github.com/lukaDoric/SOA/assets/45179708/340f6757-314e-4002-97d5-ee7184499a94">

Sa komandom `exit` napuštamo kontejner i vraćamo se na glavni terminal. Ono što je bitno razumeti je da smo sa ovom komandom ugasili glavni proces kontejnera i samim tim smo ugasili i kontejner.

Sa komandom `docker ps` smo zatražili izlistavanje svih pokrenutih kontejnera.

<img width="1500" alt="Screenshot 2024-03-15 at 12 41 32" src="https://github.com/lukaDoric/SOA/assets/45179708/eb6990a4-3b8d-4a41-adce-177f4fb188ff">

S obzirom da smo sa exit ugasili glavni proces našeg kontejnera (samim tim i njega), prilikom izvršenja gorepomenute komande neće biti izlistane informacije o kontejneru. Dodavanjem flega -a izlistavamo i pokrenute i zaustavljene kontejnere dok sa flegom -l izlistavamo informacije o poslednjem kontejneru koji je bio pokrenut bez obzira da li je i dalje pokrenut ili je zaustavljen. Sa flegom -n x slična priča kao i sa -l, s tim što ovde eksplicitno naglašavamo za koliko kontejnera želimo da vidimo informacije. Konkretne stvari koje nam se prikazuju jesu:

- **ID** - Identifikator kontejnera.
- **IMAGE** - Slika od koje je kreiran kontejner.
- **COMMAND** - Izvršena komanda.
- **STATUS** - Status našeg kontejnera (koliko je dugo pokrenut/ugašen).
- **PORTS** - Izloženi portovi.
- **NAMES** - Naziv kontejnera (Ako nije eksplicitno zadat putem flega biće generisano ime).

Sa komandom docker images izlistavamo informacije o svim preuzetim i kreiranim slikama.

<img width="1500" alt="Screenshot 2024-03-15 at 12 42 15" src="https://github.com/lukaDoric/SOA/assets/45179708/8eb29568-9242-478e-961b-73adb0dbd48c">

Informacije koje nam se prikazuju su:

- **REPOSITORY** - Repozitorijum sa koje je slika preuzeta.
- **TAG** - Oznaka koja najčešće ima za ulogu da prikaže verziju slike (npr. za Ubuntu je to 18.04/18.10 itd.). Ukoliko ne naglasimo koji tag želimo, biće preuzeta poslednja stable verzija slike.
- **IMAGE ID** - Identifikator slike.
- **CREATED** - Kada je slika kreirana.
- **SIZE** - Veličina slike.

Sa komandom `docker run` smo istovremeno preuzeli sliku i odmah pokrenuli kontejner od nje. Možemo izvršiti i samo preuzimanje slike bez naknadnog pokretanja putem komande `docker pull naziv_slike:tag`.

<img width="818" alt="Screenshot 2024-03-15 at 12 44 26" src="https://github.com/lukaDoric/SOA/assets/45179708/4f59593e-2b4d-438e-bfd9-2a2e529a45a9">

U konkretnom slučaju preuzeli smo Fedora sliku gde smo sa tagom naznačili verziju 20.

Neke od vrlo korisnih komandi:
- `docker rm naziv_kontejnera` (dodatno fleg `-f` za brisanje kontejnera koji je pokrenut. Umesto naziva se može koristiti i id).
- `docker start naziv_kontejnera` (pokretanje kontejnera sa zadatim nazivom; može se koristiti i id).
- `docker stop naziv_kontejnera` (zaustavljanje kontejnera sa zadatim nazivom; može se koristiti i id).
- `docker exec` (omogućuje izvršavanje komandi unutar kontejnera).
- `docker rmi naziv_slike` (omogućuje brisanje slike po nazivu).

Postoji naravno još komandi i puno dodatnih flegova za svaku komandu, i dodatne informacije o svakoj se mogu naći u odličnoj zvaničnoj dokumentaciji: [https://docs.docker.com/engine/reference/commandline/docker/](https://docs.docker.com/engine/reference/commandline/docker/)

<h2>7. Kako kreirati sopstvene slike?</h2>

Videli smo kako da pokrenemo kontejnere na osnovu već postojećih slika, ali ono što nas konkretno interesuje jeste kako da kreiramo sopstvene slike i da pomoću njih pokrenemo naše kontejnere u kojima će se izvršavati neki konkretan mikroservis (u primeru neka Spring-Boot aplikacija). Za potrebe kreiranja naše slike neophodno je da kreiramo Dockerfile (sa tim nazivom) odnosno tekstualnu datoteku (najbolja praksa je da se ona nalazi u root direktorijumu projekta) koja koristi bazični DSL sa instrukcijama za kreiranje slika. Kada kreiramo taj fajl, komandom docker image build ćemo kreirati našu sliku izvršavanjem instrukcija koje smo napisali, i zatim ćemo od te slike startovati kontejner.

Format je relativno jednostavan, i instrukcije koje postoje su:

- **FROM** Pomoću ove instrukcije definišemo koja je bazna slika za predstojeće instrukcije koje će biti izvršene. Svaki fajl mora početi FROM instrukcijom, s tim što je moguće imati više FROM instrukcija u istom Dockerfile-u. Bazična slika bi trebala da bude oficijelna i po potrebi sa latest tagom jer su te slike proverene.
- **ADD** Ova instrukcija kopira fajlove sa zadate destinacije u fajlsistem slike na odredišnoj destinaciji (biće dodat novi sloj u slici).
RUN - Omogućuje izvršavanje komande pri čemu će rezultat biti novi sloj (layer) u samoj slici.
- **COPY** Slično kao i ADD instrukcija, s tim što ADD omogućuje da source bude i URL, dok COPY zahteva fizičku putanju na disku (biće dodat novi sloj u slici).
- **WORKDIR** Postavlja putanju odakle će pojedine komande biti izvršene.
- **EXPOSE** Definišemo port kako bi mogli da odradimo mapiranje portova da bi kontejneri mogli da komuniciraju sa spoljašnjim svetom.
- **ENTRYPOINT** Postavljamo executable koji će biti pokrenut sa pokretanjem kontejnera.
- **ENV** Podešavanje environment varijabli.
- **LABEL** Dodaje metapodatke slike poput verzije, maintainer itd.

Postoji još instrukcija koje se mogu definisati u Dockerfile-u i više informacija o njima kao i o formatu argumenata instrukcija možete naći u zvaničnoj dokumentaciji: [https://docs.docker.com/engine/reference/builder/](https://docs.docker.com/engine/reference/builder/)

Takođe, preporuke koje diktira najbolja praksa možete naći u zvaničnoj dokumentaciji: [https://docs.docker.com/develop/develop-images/dockerfile_best-practices/](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/)

Dakle, tok koji treba ispoštovati je pisanje koda, pa kad s tim završimo, kreiramo Dockerfile, pokrećemo docker image build kako bi kreirali sliku, i zatim startujemo kontejner na osnovu slike koju smo kreirali.

-----------------------------------------------
Objašnjenje par termina vezanih za kreiranje Docker slika:

`Host mašina`: Računar na kom su pokrenuti docker kontejneri. U vašem slučaju će to biti vaš laptop ili desktop.

`docker build` : Komanda koja uzima `Dockerfile` i na osnovu njega kreira sliku. Kao argument uzima putanju do direktorijuma gde se nalazi Dockerfile. Pošto ćete najčešće biti pozicionirani u direktorijumu gde se nalazi Dockerfile, ovaj argument će biti "`.`" Potrebno je i da damo naziv slike koju kreiramo. Za to koristimo falg `-t`. Naziv slike ima i _tag_, ili verziju, koji se nalazi nakon "`:`". Primer naziva slike može biti `mysql:5.1` ili `mysql:beta`, gde je "mysql" naziv slike, a "beta" ili "5.1" tagovi. Ako ne navedemo tag, već samo naziv, Docker automatski dodeljuje tag "latest".

Kompletan primer ove komande je: `docker build -t static_web_server .`

<img width="1510" alt="Screenshot 2024-03-15 at 12 52 25" src="https://github.com/lukaDoric/SOA/assets/45179708/e25deb70-f161-429a-bc99-1c0a7ff2ce81">

`docer run`: Komanda koja na osnovu slike kreira i pokreće kontejner (izvršno okruženje). Kao argument uzima naziv slike. Prvo proverava da li se slika sa tim imenom nalazi na računaru, a zatim, ako je ne pronađe, pokušava da vidi da li može tu sliku da skine sa [`Docker hub-a`](https://hub.docker.com/). Ova komanda prima mnogo opcionih parametara koji detaljnije definišu ponašanje našeg kontejnera. Ovde ću nabrojati samo nekoliko najosnovnijih:
- -p lista port binding-a. Možemo da navedemo koje portove naše host mašine spajamo sa portovima docker containera (npr -p 8080:8080)
- -e lista environment variabli koje će se upisati unutar kontejnera prilikom pokretanja (npr. -e DATABASE_USERNAME='student')
- -d kad budete pokretali prvi kontejner, videćete da on preuzme kontrolu nad vašim terminalom. Ako ne želite to da se desi, na kraju ove komande dodate lag -d, što je skraćenica za detached.

Primer komande: `docker run -p 8080:8080 static_web_server -d`

<img width="1144" alt="Screenshot 2024-03-15 at 13 05 20" src="https://github.com/lukaDoric/SOA/assets/45179708/b65e4b7d-a467-4c29-af5a-ed052b6e03df">

Ukoliko ukucamo komandu `docker images`, kreirana slika će nam biti prikazana kao i sve ostale preuzete slike.

Ako hoćemo preko dockera da pokrenemo kompleksniju go aplikaciju koja komunicira sa bazom podataka, moramo promeniti nekoliko stvari. Za početak menjamo Dockerfile.

Bazna slika više nije `FROM alpine` već `FROM golag:alpine`. Ova druga slika u sebi sadrži go kompajler, te ćemo moći od go source koda da napravimo executable datoteku. Samim tim će slika koju dobijemo koristeći novu verziju Dockerfile-a zauzimati više memorija, zbog dodatnih alata koji će biti prisutni. Pored ove promene, takođe bi morali da podesimo komunikaciju između naše golang aplikacije i mysql baze podataka. Više ne možemo da se povezujemo na "localhost:3306" iz aplikacije, jer u okviru virtualne mašine (kontejnera) gde je pokrenut veb server ne postoji baza podataka. Bazu želimo da podignemo u zasebnom kontejneru, kako bi izolovali rad aove dve nezavisne aplikacije. Da bi omogućili komunikaciju između 2 kontejnera možemo koristiti Docker mreže, što ćemo videti u 9. poglavlju zajedno za `docker compose` alatom.

Napredne stvari:  

● [Multi-stage build](https://docs.docker.com/develop/develop-images/multistage-build/)  
● [Kreiranje base slike](https://docs.docker.com/develop/develop-images/baseimages/)  

<h2>8. Docker volumes</h2>

U poglavlju u kome su opisivane slike, bilo je reči o r​ead-only s​lojevima i read-write sloju koji se dodaju iznad prethodnih slojeva za svaki kontejner koji je pokrenut. Sve promene i sav sadržaj se upisuju u taj sloj. Problem sa tim jeste da kada se kontejner obrise, promene će biti potpuno izgubljene.
Zato je ​Docker​ uveo koncept pod nazivom `​volumes`. D​a bi mogli da čuvamo konkretan sadržaj (​persist)​, i po potrebi ga delimo između različitih kontejnera, kreiramo poseban ​`volume` koji je, prosto rečeno, ništa drugo do skup direktorijuma/fajlova koji se nalaze izvan ​default-​nog docker  image file sistema i čiji sadržaj se čuva unutar fajl sistema host mašine, kako se ne bi obrisao nakon gašenja kontejnera.

Kreiranje ​volume-​a je moguće odraditi sa komandom `​docker volume create naziv`.​​Mount​-ovanje se radi prilikom pokretanja sa flegom -​-volume i​li -​v. ​Primer: ​`docker run -i -t -v primer1:/nekiPodaci ubuntu /bin/bash`. ​Dakle najobičnija komanda (koju smo već videli), proširena flegom -​v ​gde smo zadali naziv ​volume-​a i gde će biti izvršeno mount-​ovanje u okviru samog kontejnera.

<img width="1141" alt="Screenshot 2024-03-15 at 15 33 14" src="https://github.com/lukaDoric/SOA/assets/45179708/cb2f4ac0-33ac-40c2-a8a2-0bb87847d806">

Na slici je prikazano najpre kreiranje ​volume-​a, a zatim je pokrenut kontejner kome smo ​mount​-ovali prethodno kreirani ​volume na putanji ​nekiPodaci.​U okviru prvog kontejnera smo i kreirali običan tekstualni fajl. Zatim smo izvršili ​exit (​ugasili glavni proces /​bin/bash ​i samim tim i ugasili kontejner) i pokrenuli nov kontejner kome smo takođe ​mount​-ovali isti ​volume na istoj putanji (apsolutno ne mora biti ista) i kada smo ušli u sam folder, datoteka koju smo prethodno kreirali iz totalno drugog kontejnera i dalje postoji.

<h2>9. Šta raditi sa ostalim mikroservisima?</h2>

U prethodnom poglavljima je objašnjena manipulacija ​volume-​a, kako kreirati sopstvenu sliku i kako od nje kreirati kontejner. Međutim, postavlja se pitanje šta raditi ukoliko imamo više aplikacija, od kojih je neke neophodno pokrenuti u više instanci (kontejnera), koji moraju da komuniciraju međusobno. Tada pojedinačno kreiranja slika i pokretanja kontejnera nije praktično rešenje. Zato se koristi alat `docker compose` ​koji nam omogućuje pokretanje i zaustavljanje ​više aplikacija koristeći jednu komandu, kao i zejdnički ispis logova svih aplikacija na jedan terminal.

Sve što je neophodno jeste da kreiramo fajl pod nazivom `docker-compose.yml`​ U folderu `go-primeri/WebServerWithDB` i `nginx-example` možete videti kako treba ovaj fajl da izgleda

```yml

version: "3.7"
services:
  servers:
    build:
      context: ./
      dockerfile: Dockerfile
    image: students_web_server
    container_name: student_server
    restart: always
    networks:
      - servers
    ports:
      - 8080:8080
    depends_on:
      - database
  
  database:
    image: mysql
    container_name: mysql
    restart: always
    networks:
      - servers
    ports:
      - 4000:3306
    environment:
      MYSQL_ROOT_PASSWORD: root
      MYSQL_DATABASE: students
    volumes:
      - database-data:/var/lib/mysql

volumes:
  database-data:
    name: server-database

networks:
  servers:
    name: servers
    driver: bridge
```

U fajlu za konkretan primer je definisano više direktiva:
- **version** Ovde naglašavamo koju verziju formata želimo da koristimo. Ovo polje
je uvek neophodno i dovoljno je navesti verziju 3 (poslednja verzija formata).
- **services** U ovoj sekciji se definiše niz objekata gde svaki predstavlja servis, odnosno kontejner i takođe ova sekcija je obavezna. Dalje unutar servisa
definišemo:
    - **build** -​ ​Ova direktiva ako je definisana, govori da je neophodno kreirati slike pri čemu se definišu odnosno putanja do direktorijuma na kojoj se
nalazi Dockerfile.​
    - **image​**​ Definiše naziv slika koja će nastati prilikom ​build​-ovanja.
    - **container-name​**​ Definiše naziv kontejnera koji će biti pokrenut.
    - **restart​** Definiše pod kojim okolnostima kontejner treba restartovati
    - **networks​** Definiše mrežu (mreže) u kojoj kontejner treba da se nalazi.
    - **ports​**​ Vrši se mapiranje portova.
    - **environments** Postavlja vrednost environment varijable koje se nalaze u
kontejneru.
    - **volumes​** Definiše volume za koje se kontejner kači.
    - **depends_on** Govori prilikom pokretanja servisa koje su zavisnosti
između njih, odnosno koji servisi moraju biti pokrenuti pre nego što se pokrene konkretan servis.

Za dodatne direktive i njihove vrednosti možete pogledati u zvaničnoj [dokumentaciji ​https://docs.docker.com/compose/](dokumentaciji ​https://docs.docker.com/compose/)
​
Još jedan bitan detalj na koji bi skrenuli pažnju: Ako odete u `main.go` fajl u primeru `WebServerWithDB`, videćete da baza više nije na localhost adresi, već database. Servisi koji su u istoj mreži unutar docker-a mogu komunicirati preko svog imena umesto ip adrese.

Pozicioniramo se na putanju do direktorijuma u kojem se nalazi `docker-compose.yml` i pozovemo naredbu: `​docker compose up --build`​ Sa ovim pokrećemo sve naše servise (kontejnere).

<img width="1501" alt="Screenshot 2024-03-15 at 18 28 59" src="https://github.com/lukaDoric/SOA/assets/45179708/2f9530fd-db6f-4cea-94f9-704e5985ca5e">
