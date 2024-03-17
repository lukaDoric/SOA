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

### 2.1 Docker engine

Kada pričamo o Docker Engine-u, govorimo o klasičnoj klijent-server aplikaciji. Docker klijent nam pruža CLI (command line interface) putem kojeg unosimo komande na osnovu kojih se generišu API request-ovi koji se šalju serveru (Docker daemon-u) koji ih obrađuje.

![image-002](https://github.com/lukaDoric/SOA/assets/45179708/c0e4918a-0c78-4abb-9184-3003b54d9f4f)

Sam Docker daemon je nakon refaktorisanja (zbog toga što je narastao u jedan veliki monolit) ostao bez ikakvog koda koji zaista kreira i pokreće kontejnere. On se obraća putem gRPC API-a preko lokalnog Linux socket-a `containerd`-u (long running daemon-u) koji predstavlja "API fasadu" koja omogućuje startovanje containerd-shim-a odnosno roditeljskog procesa za svaki kontejner gde runc (container runtime) vrši kreiranje kontejnera. Sloj ispod containerd-a vrši kompletan rad sa kernelom odnosno koristi njegove funkcije.  
Iako arhitektura izgleda prilično kompleksno, ovakva podela omogućuje da se pojedine komponente bez ikakvih problema zamenjuju a da to ne utiče na pokrenute kontejnere što sa administratorske tačke gledišta puno olakšava stvari. Na primer, moguće je promeniti verziju Docker-a a da se pri tome ne moraju zaustavljati već pokrenuti kontejneri.

Na sledećim linkovima možeš pogledati video snimke o Docker-u:  
[Uvod u Docker](https://youtu.be/MAivaLjKDaY)  
[Interna arhitektura Docker-a](https://youtu.be/nat6d-kNTGU)

### 2.2 Docker slike

Generalno je poznat koncept slike kada je priča o virtuelnim mašinama. Za sličnu stvar se koriste i Docker slike, odnosno predstavljaju build-time konstrukt od kojih nastaju kontejneri, ali se tu sličnost završava. Pojednostavljeno docker slike čine templejt na osnovu kog se kreira docker kontejner. Docker slike predstavljaju skup read-only layer-a gde svaki sloj predstavlja različitosti u fajlsistemu u odnosu na prethodni sloj, pri čemu uvek postoji jedan bazni (base) sloj. Upotrebom storage driver-a skup svih slojeva čini root filesystem kontejnera, odnosno svi slojevi izgledaju kao jedan unificirani fajlsistem.

![image-003](https://github.com/lukaDoric/SOA/assets/45179708/cbc05507-5818-4642-8c54-850bd7136e81)

![image-004](https://github.com/lukaDoric/SOA/assets/45179708/97093118-f6a0-452d-97da-b8e748057547)

Svi ovi read-only slojevi predstavljaju osnovu za svaki kontejner koji se pokreće i ne mogu se menjati. Prilikom pokretanja svakog kontejnera, Docker dodaje još jedan sloj koji je read-write tipa i u koji se upisuju nove datoteke i sve izmene. Ukoliko želimo da menjamo neki fajl koji se nalazi u nekom read-only sloju, taj fajl će biti kopiran u read-write sloj, biće izmenjen i kao takav dalje korišćen. Originalna verzija će i dalje postojati (nepromenjena), ali nalaziće se "skrivena" ispod nove verzije.

![image-005](https://github.com/lukaDoric/SOA/assets/45179708/f8271035-8fad-40cd-931b-941f81c69d8a)

Ovakav mehanizam se zove Copy-on-write i delom čini Docker zaista moćnim. Koliko god kontejnera da kreiramo, read-only slojevi će uvek biti isti, tj. ostaće nepromenjeni, samo će svaki kontejner dobiti sopstveni read-write sloj. Na ovaj način se štedi jako puno prostora na disku jer kada smo jednom preuzeli/kreirali sliku, koliko god kontejnera da pokrenemo, slika ostaje apsolutno nepromenjena.

Na sledećem linku možeš pogledati video snimak o Docker slikama:  
[Uvod u Docker slike](https://youtu.be/LsjntJHEe_8)  

### 2.3 Docker registri

Docker čuva slike u registrima, pri čemu postoje dva tipa odnosno javni i privatni. Javni registar kojim upravlja Docker Inc. se zove DockerHub i na njemu svako može da napravi nalog i da tamo čuva i deli sopstvene slike. Postoje dva tipa slika a to su oficijelne koje žive na top nivou DockerHub namespace-a (npr. Ubuntu, Redis itd.) i neoficijelne (korisničke). Takođe je moguće napraviti privatni registar u kome se mogu čuvati slike i sve to sakriti iza firewall-a što je ponekad neophodno za pojedine organizacije.

Postojeći Docker registri nude mesto gde korisnici mogu da preuzmu već postojeće slike koje su kreirali drugi korisnici ili organizacije. Ovo omogućava brzo deljenje i razvoj aplikacija, jer korisnici mogu da iskoriste prethodno kreirane slike kao osnovu za svoje aplikacije ili da jednostavno pokrenu servise potrebne za svoje aplikacije bez potrebe da sami kreiraju sve od početka. Oficijelne slike na DockerHub-u su verifikovane i pružaju siguran temelj za izgradnju kontejnerizovanih aplikacija, dok neoficijelne slike pružaju širok spektar alata i aplikacija koje je zajednica razvila.

### 2.4 Docker kontejneri

Kako slike predstavljaju build-time konstrukt (templejt), tako su kontejneri run-time konstrukt (pokreću se na osnovu templejta). Gruba analogija odnosa između slike i kontejnera se može posmatrati kao klasa i instanca te klase. Kontejneri predstavljaju lightweight execution environment koji omogućuju izolovanje aplikacije i njenih zavisnosti koristeći `kernel namespaces` i `cgroups` mehanizme.

![image-006](https://github.com/lukaDoric/SOA/assets/45179708/7657bc58-9b52-4e5e-a492-09839f6552e6)

`Namespaces` nam omogućuju izolaciju, odnosno da podelimo naš operativni sistem na manje izolovanih virtuelnih operativnih sistema (kontejnera). Odnosno, kontejneri se ponašaju kao zasebni operativni sistemi (kao kod VM-a) samo što to nisu, jer svi dele isti kernel na host OS-u. Svaki kontejner ima sopstveni skup namespace-ova (kada pričamo o Linux-u to su namespace-ovi sa slike 6) pri čemu je njegov pristup ograničen isključivo na taj prostor imena, odnosno svaki kontejner nije uopšte svestan postojanja drugih kontejnera.

![image-007](https://github.com/lukaDoric/SOA/assets/45179708/22ddf7dc-1f88-4fb6-8c50-09676f6b2bea)

Međutim, iako imamo potpunu izolaciju, to nam nije skroz dovoljno. Kao i svaki multi-tenant sistem, uvek postoji opasnost od noisy neighbors-a, odnosno neophodan nam je mehanizam kojim ćemo ograničiti upotrebu resursa host OS-a od strane svih kontejnera, kako se ne bi desilo da jedan kontejner troši mnogo više resursa od drugih. To nam omogućava control groups (cgroups) kernel mehanizam (slika ispod).

![image-008](https://github.com/lukaDoric/SOA/assets/45179708/a10e8547-435a-4f69-b962-8062c1431ee9)

