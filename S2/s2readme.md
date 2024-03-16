## Drugi sprint (18.03 - 1.04)

Na osnovu arhitekture sa prve KT (početna monolitna aplikacija i dva izolovana servisa) potrebno je:  
1. Dokerizovati čitav bekend deo (frontend nije potrebno dokerizovati).  
1.1 Potrebno je napisati Dockerfile za svaki servis (monolitnu i dva izolovana servisa).  
1.2 Potrebno je napisati docker compose fajl koji uključuje u sebe podizanje sva tri servisa i odgovarajuće baze podataka.  
1.3 Potrebno je omogućiti komunikaciju između izolovanih kontejnera i od frontenda do kontejnera monolitne aplikacije.  

<a href=''>Teorijske osnove Docker-a.</a>   
<a href=''>Primer rada sa Docker-om (Golang).</a>  
<a href=''>Primer rada sa Docker-om (.NET).</a>   

<h1>Docker</h1>

<h2>Video materijal koji prati ovu lekciju:</h2>

[Docker slika za servis i bazu podataka](https://youtu.be/8cyjjYR6LzY)

[Pravilan redosled komandi prilikom definisanja Docker slike](https://youtu.be/bogpeVjMIKE)

[Multistage build](https://youtu.be/_EYy81VGrJ0)

[docker compose](https://youtu.be/LnP4ibmKhIg)

[Primer docker-compose.yaml sa više servisa i env varijablama](https://youtu.be/gtnnIr3aET4)

<h2>1. Motivacija</h2>

Sa pojavom virtuelnih mašina (VM) omogućeno je izbegavanje situacija gde se fizički serveri koriste na takav način da je iskoristivost resursa vrlo mala što je u prošlosti često bio slučaj (iskoristivost resursa često bude od 10-20%). Virtuelne mašine su apstrakcija fizičkog hardvera koje omogućuju pretvaranje jednog servera u više manjih servera. Svaka VM-a uključuje punu kopiju operativnog sistema, aplikacije, biblioteke pri čemu se ispod njih nalazi hypervisor odnosno softver koji omogućuje kreiranje, pokretanje i izvršavanje više VM-a na jednom fizičkom računaru (type 1 i type 2 hypervisor) i omogućuje deljenje fizičkih resursa (memoriju, procesor) između njih. Dakle, virtuelne mašine (virtuelni serveri) su jeftiniji od fizičkih servera s obzirom da troše deo resursa istog. Pored manje cene omogućuju lakše upravljanje, bolje skaliranje, konzistentno okruženje za izvršavanje aplikacija što ih čini odličnom podlogom za pružanje usluga web servisa.

Sa pojavom virtuelnih mašina, svet je postao bolje mesto, ali i dalje je bilo prostora za napredak. Ono što je negativna strana VM-a jeste da svaka zahteva underlying OS što znači da će se deo resursa koristiti za podizanje i izvršavanje operativnog sistema. Takođe, operativni sistemi uključuju potencijalni overhead u obliku dodatnih potreba za licencama, potreba za administracijom (updates, patches) itd.

Virtuelne mašine vs kontejneri:

![0](https://github.com/lukaDoric/SOA/assets/45179708/2b0a77cf-7431-4558-984e-ec9b95dd8397)

Ovi nedostaci su u priču uključili kontejnere. Za razliku od virtuelnih mašina gde svaka ima sopstveni OS i oslanja se na hypervisor, kontejneri se oslanjaju na jedan host OS i dele njegove funkcije kernela (takođe i binaries, libraries itd.) i samim tim su lakši (lightweight) i u priličnoj meri se smanjuje overhead koji donose VM-e. Kontejnerske tehnologije su bile prisutne duže vremena ali nisu bile previše popularne jer je kreiranje i upravljanje kontejnerima bilo dosta kompleksno ali je Docker uspeo to da promeni.

Na sledećim linkovima možeš pogledati video snimke o virtualnim mašinama, njihovim manama i o Linux c-grupama:  
[Virtualne mašine](https://youtu.be/xxC2yAfamSM)  
[Mane virtuelnih mašina](https://youtu.be/MbVYKEUDab0)  
[Linux cgroups & namespaces](https://youtu.be/Mby8-4twxNY)

<h2>2. Šta je Docker i koje su njegove komponente?</h2>

Docker je open-source platforma koja automatizuje proces deployment-a aplikacija u softverske kontejnere. Dizajniran je tako da omogući lagano i brzo okruženje za izvršavanje naših aplikacija kao i izuzetno lako premeštanje aplikacija iz jednog okruženja u drugo (`test` -> `production`).

Njegove osnovne komponente su:

- Docker Engine
- Docker Images
- Registries (Docker Hub)
- Docker containers
- Docker Swarm

Kada pričamo o Docker Engine-u, govorimo o klasičnoj klijent-server aplikaciji. Docker klijent nam pruža CLI (command line interface) putem kojeg unosimo komande na osnovu kojih se generišu API request-ovi koji se šalju serveru (Docker daemon-u) koji ih obrađuje.

![image-002](https://github.com/lukaDoric/SOA/assets/45179708/c0e4918a-0c78-4abb-9184-3003b54d9f4f)

Sam Docker daemon je nakon refaktorisanja (zbog toga što je narastao u jedan veliki monolit) ostao bez ikakvog koda koji zaista kreira i pokreće kontejnere. On se obraća putem gRPC API-a preko lokalnog Linux socket-a `containerd`-u (long running daemon-u) koji predstavlja "API fasadu" koja omogućuje startovanje containerd-shim-a odnosno roditeljskog procesa za svaki kontejner gde runc (container runtime) vrši kreiranje kontejnera. Sloj ispod containerd-a vrši kompletan rad sa kernelom odnosno koristi njegove funkcije. Iako arhitektura izgleda prilično kompleksno, ovakva podela omogućuje da se pojedine komponente bez ikakvih problema zamenjuju a da to ne utiče na pokrenute kontejnere što sa administratorske tačke gledišta puno olakšava stvari. Na primer, moguće je promeniti verziju Docker-a a da se pri tome ne moraju zaustavljati već pokrenuti kontejneri.

Na sledećim linkovima možeš pogledati video snimke o Docker-u:  
[Uvod u Docker](https://youtu.be/MAivaLjKDaY)  
[Interna arhitektura Docker-a](https://youtu.be/nat6d-kNTGU)

<h2>3. Šta su Docker slike?</h2>

Generalno je poznat koncept slike kada je priča o virtuelnim mašinama. Za sličnu stvar se koriste i Docker slike, odnosno predstavljaju build-time konstrukt od kojih nastaju kontejneri, ali se tu sličnost završava. Docker slike predstavljaju skup read-only layer-a gde svaki sloj predstavlja različitosti u fajlsistemu u odnosu na prethodni sloj, pri čemu uvek postoji jedan bazni (base) sloj. Upotrebom storage driver-a skup svih slojeva čini root filesystem kontejnera, odnosno svi slojevi izgledaju kao jedan unificirani fajlsistem. Pojednostavljeno docker slike čine templejt na osnovu kog se kreira docker kontejner.

![image-003](https://github.com/lukaDoric/SOA/assets/45179708/cbc05507-5818-4642-8c54-850bd7136e81)

![image-004](https://github.com/lukaDoric/SOA/assets/45179708/97093118-f6a0-452d-97da-b8e748057547)

Svi ovi read-only slojevi predstavljaju osnovu za svaki kontejner koji se pokreće i ne mogu se menjati. Prilikom pokretanja svakog kontejnera, Docker dodaje još jedan sloj koji je read-write tipa i u koji se upisuju nove datoteke i sve izmene. Ukoliko želimo da menjamo neki fajl koji se nalazi u nekom read-only sloju, taj fajl će biti kopiran u read-write sloj, biće izmenjen i kao takav dalje korišćen. Originalna verzija će i dalje postojati (nepromenjena), ali nalaziće se "skrivena" ispod nove verzije.

![image-005](https://github.com/lukaDoric/SOA/assets/45179708/f8271035-8fad-40cd-931b-941f81c69d8a)

Ovakav mehanizam se zove Copy-on-write i delom čini Docker zaista moćnim. Koliko god kontejnera da kreiramo, read-only slojevi će uvek biti isti, tj. ostaće nepromenjeni, samo će svaki kontejner dobiti sopstveni read-write sloj. Na ovaj način se štedi jako puno prostora na disku jer kada smo jednom preuzeli/kreirali sliku, koliko god kontejnera da pokrenemo, slika ostaje apsolutno nepromenjena.

Na sledećim linkovima možeš pogledati video snimke o docker slikama:  
[Uvod u Docker slike](https://youtu.be/LsjntJHEe_8)  
[Prva Docker slika](https://youtu.be/xfuSdAcusfw)

<h2>4. Odakle se preuzimaju postojeće slike?</h2>

Docker čuva slike u registrima, pri čemu postoje dva tipa odnosno javni i privatni. Javni registar kojim upravlja Docker Inc. se zove DockerHub i na njemu svako može da napravi nalog i da tamo čuva i deli sopstvene slike. Postoje dva tipa slika a to su oficijelne koje žive na top nivou DockerHub namespace-a (npr. Ubuntu, Redis itd.) i neoficijelne (korisničke). Takođe je moguće napraviti privatni registar u kome se mogu čuvati slike i sve to sakriti iza firewall-a što je ponekad neophodno za pojedine organizacije.

Postojeći Docker registri nude mesto gde korisnici mogu da preuzmu već postojeće slike koje su kreirali drugi korisnici ili organizacije. Ovo omogućava brzo deljenje i razvoj aplikacija, jer korisnici mogu da iskoriste prethodno kreirane slike kao osnovu za svoje aplikacije ili da jednostavno pokrenu servise potrebne za svoje aplikacije bez potrebe da sami kreiraju sve od početka. Oficijelne slike na DockerHub-u su verifikovane i pružaju siguran temelj za izgradnju kontejnerizovanih aplikacija, dok neoficijelne slike pružaju širok spektar alata i aplikacija koje je zajednica razvila.

<h2>5. Šta predstavljaju kontejneri?</h2>

Kako slike predstavljaju build-time konstrukt (templejt), tako su kontejneri run-time konstrukt (pokreću se na osnovu templejta). Gruba analogija odnosa između slike i kontejnera se može posmatrati kao klasa i instanca te klase. Kontejneri predstavljaju lightweight execution environment koji omogućuju izolovanje aplikacije i njenih zavisnosti koristeći `kernel namespaces` i `cgroups` mehanizme.

![image-006](https://github.com/lukaDoric/SOA/assets/45179708/7657bc58-9b52-4e5e-a492-09839f6552e6)

`Namespaces` nam omogućuju izolaciju, odnosno da podelimo naš operativni sistem na manje izolovanih virtuelnih operativnih sistema (kontejnera). Odnosno, kontejneri se ponašaju kao zasebni operativni sistemi (kao kod VM-a) samo što to nisu, jer svi dele isti kernel na host OS-u. Svaki kontejner ima sopstveni skup namespace-ova (kada pričamo o Linux-u to su namespace-ovi sa slike 6) pri čemu je njegov pristup ograničen isključivo na taj prostor imena, odnosno svaki kontejner nije uopšte svestan postojanja drugih kontejnera.

![image-007](https://github.com/lukaDoric/SOA/assets/45179708/22ddf7dc-1f88-4fb6-8c50-09676f6b2bea)

Međutim, iako imamo potpunu izolaciju, to nam nije skroz dovoljno. Kao i svaki multi-tenant sistem, uvek postoji opasnost od noisy neighbors-a, odnosno neophodan nam je mehanizam kojim ćemo ograničiti upotrebu resursa host OS-a od strane svih kontejnera, kako se ne bi desilo da jedan kontejner troši mnogo više resursa od drugih. To nam omogućava control groups (cgroups) kernel mehanizam (slika ispod).

![image-008](https://github.com/lukaDoric/SOA/assets/45179708/a10e8547-435a-4f69-b962-8062c1431ee9)

<h2>6. Kako raditi sa kontejnerima?</h2>

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

```docker
# Comment
INSTRUCTION arguments
```

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

<h2>10. Docker Swarm</h2>

Docker ​Swarm je alat koji omogućava orkestraciju nad kontejnerima. ​Docker Swarm​ 
ima implementiran l​oad balancer ​i ​discovery service, ​servisi koji su neophodni u mikroservisnoj arhitekturi. Za aktiviranje ​Docker Swarm-​a neophodno je podesit ​Docker da radi u ​swarm režimu komandom: `​docker swarm init`.​ Za pokretanje prethodnog primera pomoću ​Docker Swarm-​a, neophodne je dopuniti određene stvari u `docker-compose.yml` fajlu.

U fajlu je dodata `​deploy` ​sekcija koja govori kako treba da se uradi ​deployment​servisa:
- **replicas​** koliko instaci kontejnera treba da bude aktivno
- **parallelism​** i **​delay**​  ukoliko broj aktivnih kontejnera manji od specifiranog, ​Docker pokreće nove instance kontejnera dok ne postigne željeni broj. Ova sekcija definiše kako je to potrebno postići. Na primer ukoliko definišemo *r​eplicas*​: 10, *parallelism*: 2 i *delay*: 10s, u slučaju da su pali svih 10 kontejnera, ​Docker​će istovremeno podizati 2 kontejnera pri čemu će nakog uspešnog podizanja oba kontejnera sačekati 10 sekundi i opet pokušavati da podigne istovremeno 2 kontejnera. Ovaj proces se ponavlja sve dok se ne postigne željeni broj instanci kontejnera.
- **restart_policy**​ definiše pod kojim okolnostima je neophodno podizati nove kontejnere.
  
Pokretanje se vrši pomoću naredne komande (​*demo​* predstavlja naziv ​stack-​a koji može biti proizvoljan, dok -​c f​lagom definišemo putanju do ​yml ​fajla):
`docker stack deploy -c stack-file.yml demo`
Za praćenje stanja servisa neophodne je izvršiti narednu komandu:
`docker stack services demo`
