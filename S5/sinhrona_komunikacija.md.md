### REST

Do sada ste imali prilike da radite sa **REST** servisima. Fokus REST web servisa je na **resursima**, i kako omogućiti pristup tim resursima. Resurs može biti predstavljen kao objekat u memoriji, datoteka na disku, podatak iz aplikacije ili baze podataka itd. Prilikom dizajniranja sistema prvo je potrebno da identifikujemo resurse, i da ustanovimo kako su međusobno povezani. Ovaj postupak je sličan modelovanju baze podataka. Kada smo identifikovali resurse, sledeći korak je da pronađemo način kako da te resurse reprezentujemo u našem sistemu. Za te potrebe možemo koristiti bilo koji format za reprezentaciju resursa (JSON, XML npr.). Da bi dobili sadržaj sa udaljene lokacije (od nekog servera), klijent mora napraviti HTTP zahtev, i poslati ga web servisu. Nakon svakog HTTP zahteva, sledi i HTTP odgovor od servera ka klijentu tj. onome ko je zahtev poslao. Klijent može biti korisnik, ili može biti neka druga aplikacija (npr. drugi web servis).  

Svaki HTTP zahtev se sastoji od nekoliko elemenata:  
• **VERB** - GET, PUT, POST, DELETE, OPTIONS, itd. odnosno koju operaciju želmo da uradimo.  
• **URL** - Putanja do resursa nad kojim ce operacija biti izvedena.  

Kada klijent dobije odgovor nazad, dobiće sadržaj (ako ga ima), ali i status kod. Ovaj kod nam govori da li je prethodno zahtevana operacija izvršena uspešno, ili ne. Status kod je reprezentovan celim brojem I to: 

• **Success 2xx** sve je prošlo ok  
• **Redirection 3xx** desila se redirekcija na neki drugi servis  
• **Error 4xx, 5xx** javila se greška 

Servisi koje pozivamo preko internet imaju već definisanu strukturu (tačnu putanju metod kojim se pozivaju, podatke koje očekuju, kako se pretražuju). Ako mi implementiramo web servis, onda mi namećemo ova pravila. Slika ispod predstavlja primer klijent-server komunikacije.

<img src="https://i.ibb.co/BcfKX5v/rest.png" alt="rest" border="0">

### JSON

JSON je format za lakšu razmenu podataka, kao i XML nezavistan je od programskog jezika i tehnologije koja se koristi. Podaci se zapisuju kao parovi ključ:vrednost.
Ključ se navodi kao tekst pod duplim navodnicima nakon čega sledi vrednost na primer:  
• “firstName":"John"

JSON vrednosti mogu biti neke od unapred definisanih:  
• A number (integer or floating point)  
• A string (in double quotes)  
• A Boolean (true or false)  
• An array (in square brackets)  
• An object (in curly braces)  
• Null    


JSON object se zapisuje u parovima ključ:vrednost koji se nalaze unutar
vitičastih zagrada:  
{"firstName":"John", "lastName":"Doe"}  

JSON Array sadrži ključ, nakon čega sledi niz elemenata u uglasim
zagradama:  

"employees":[  
{"firstName":"John", "lastName":"Doe"},  
{"firstName":"Anna", “lastName”:”Smith"},  
{"firstName":"Peter","lastName":"Jones"}  
]  

JSON može da kombinuje razne tipove podataka unutar jednog niza, o
tome treba voditi računa kada koristimo strogo tipizirane jezike da ne bi
doslo do problema prilikom konverzije. Mešanje tipova treba izbegavati.

### Remote procedure call

**RPC (Remote Procedure Call)** i REST predstavljaju dva različita pristupa za formiranje API-ja servisa dostupnih preko mreže ali je bitno napomenuti da nemaju nikakve sličnosti. Osnova ideja RPC-a je iniciranje **operacije udaljenih programa** kao da su oni dostupni lokalno. Na taj način, skriva se komunikacija koja se odvija preko mreže. Dok je REST orijentisan na resurse, u osnovi različitih RPC radnih okvira i protokola nalaze se **akcije**. U situacijama kada je potrebno obezbediti efikasnu i brzu komunikaciju i razmenu što manje količine podataka preko mreže, RPC rešenja mogu biti pogodnija opcija od REST-a.  

RPC je poput pozivanja **metode** lokalnog objekta (samo što bi se u tom slučaju zvao LPC - Local Procedure Call). A ako pričamo o pozivanju akcije koja se nalazi u drugom mikroservisu onda je to RPC - Remote Procedure Call. (https://www.techtarget.com/searchapparchitecture/definition/Remote-Procedure-Call-RPC)

### gRPC

**gRPC** je moderan, izuzetno brz, open source framework za izradu servisno orijentisanih aplikacija (mikroservisa) zasnovan na Remote Procedure Call (RPC) principima, odnosno pozivima metoda udaljenog objekta.  
Ovaj framework omogućava slanje izuzetno **malih poruka**, **bidirekcioni streaming**, **http/2 način komunikacije**. Ekosistem oko njega je izuzetno bogat i moguće ga je proširiti koristeći plugin mehanizam.

1. Zbog toga što koristi HTTP/2 i Protobuf
(Protocol Buffers) binarni format poruka, vrlo
je brz i poruke koje se razmenjuju su znatno
manje u odnosu na one koju su tekstualnog, na
primer JSON ili XML formata.
2. Za veliki broj programskih jezika dostupni su
alati i biblioteke potrebne za rad sa gRPC-jem.

<img src="https://i.ibb.co/dpZbX7y/grpc-proto-concept.png" alt="grpc-proto-concept" border="0">

**Napomena:** Komunikaciju između bilo koja dva servisa je moguće implementirati i putem REST-a i putem RPC-a. Kako u okviru projekta gradite mikroservisnu arhitekturu RPC je bolja opcija za "razmenu akcija" između mikroservisa. Svakako komunikacija klijenta (front dela aplikacije) i gateway-a ostaje po REST arhitekturalnom paternu.

### Protobuf

gRPC framework **ne koristi JSON ili XML format** za prenos poruka između servisa, vec svoj binaran tip koji se zove **protobuf** . Definicija protobuf poruka je izuzetno jednostavna i za te potrebe se koristi DSL iz čega se dobija tip poruke za odgovarajući jezik ili jezike. Pored poruka, u istom file-u, možemo definisati i specifikacije servisa odnosno koje poruke naši servisi prihvataju, ali i šta su povratne vrednosti naših servisa. Primer ispod ilustruje definiciju servisa sa svojim porukama.

### Postupak implementacije gRPC servisa

Kako bi implementirali dva servisa koji komuniciraju pomoću gRPC-a trebalo bi izvršiti sledeće korake:
1. Definisanje servisa u okviru .proto fajla
2. Generisanje serverskog i klijentskog koda na osnovu .proto fajla
3. Implementacija procedura na serverskoj strani
4. Pozivanje procedura na klijentskoj strani

Za potrebe definicije servisa sadržaj se čuva u datoteci sa ekstenzijom **.proto**.  

Primer jedne **.proto** datoteke sa definisanom **Protobuf** porukom:

```
// Definicija gRPC servisa
service Greeter {
// Definicija rpc metode
rpc SayHello (HelloRequest) returns (HelloReply) {} }
// Definicija protobuf poruke koja će biti ulazni podatak za servis
message HelloRequest {
string name = 1;
}
// Definicija povratne vrednosti našeg servisa
message HelloReply {
string message = 1;
}
```

- Za svako polje prvo se navodi tip, zatim naziv i nakon toga broj koji mora biti jedinstven na nivou polja u okviru poruke (potreban za serijalizaciju poruke)  
- Ako polje predstavlja listu vrednosti, na početku se navodi ključna reč repeated  
- Detaljno uputstvo za upotrebu jezika možete pronaći na sledećem linku (https://protobuf.dev/programming-guides/proto3/)

U prethodno definisanom primeru, vidimo jednostavan gRPC servis koji prima jednu protobuf poruku koja ima samo jedan atribut tipa string, i poruku koja reprezentuje povratnu vrednost našeg servisa koja takođe ima jedan atribut tipa string. Naravno, protobuf nije ograničen samo na tip string, možemo koristiti i druge tipove podataka, ali isto tako može da ugnjezdimo i druge poruke da bi dobili kompleksniju poruku ili odgovor. Moguće opcije za atribute možete videti u samoj dokumentaciji (https://protobuf.dev/overview/).  

Bitno je napomenuti, da gRPC servisi prihvataju tip **Message** , tako da čak i ako želimo da samo vratimo jedan podatak kao u prethodnom primeru, taj atribut moramo da “obmotamo” u Message tip, dok je naziv poruke proizvoljan. Nakon što smo specificirali naše poruke i servise, potrebno je generisati kod za željeni jezik. To se radi koristeći alat **protoc** .  
Ovaj alat je potrebno instalirati na vašu mašinu.  
- Primer instalacije za windows možete videti na sledećem linku - https://www.youtube.com/watch?v=ES_GI-lmhEU,  
- Za unix like operativne sisteme (max/linux) možete videti na sledećem linku - https://grpc.io/docs/protoc-installation/.  

Kada je protoc instaliran na vašu radnu mašinu (a definisali ste vaše gRPC servise i protobuf poruke u fajlu sa ekstenziom .proto), možete pristupiti generisanju potrebnih elemenata za vaš željeni jezik. Primer generisanja za programski jezik go možete naći na sledećem linku - https://grpc.io/docs/languages/go/quickstart/, kao i na ovom linku - https://grpc.io/docs/languages/go/basics/ (sekcija Generating client and server code). Ako generisanje prođe bez problema, ono što je dalje potrebno jeste da definišete server (golang struktura), koja implementira sve metode definisane u vašem servisu. Kod ispod daje primer definicije servera sa metodom za proto specifikaciju iz gore definisanog primera.

```go
type server struct {
...
}
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
return &pb.HelloReply{Message: "Hello again " + in.GetName()}, nil
}
```

Komanda za generisanje go struktura na osnovu protobuf-a:

```bash
protoc --go_out=./product --go_opt=paths=source_relative \
--go-grpc_out=./product --go-grpc_opt=paths=source_relative \
product_service.proto
```

Jedna napomena jeste da će sadržaj generisanja završiti na mestu gde vi specificirate i pod ekstenzijom **naziv_proto_fajla.pb.go** , ako niste specificirali drugačije. U prethodno definisanom fajlu pb.HelloRequest je poruka koja se nalazi unutar pb biblioteke, gde je pb skraćeni naziv za punu putanju gde se generisani *.pb.go fajl nalazi. Kompletan primer možete videti na sledećem linku - https://github.com/grpc/grpc-go/tree/master/examples/helloworld, kao i primer generisanog klijenta. Pored servera, i poruka protoc generiše i klijent koji možete koristiti da pozivate druge servise. Isto kao i kod REST a, ako imate dva servisa koja treba da komuniciju u tom slučaju jedan je klijent (traži uslugu), a drugi je server (obrađuje sadržaj).

Video primer kreiranja gRPC komunikacije: https://www.youtube.com/watch?v=KHy_gHRePpU
Primer gRPC komunikacije: https://github.com/lukaDoric/SOA/tree/main/S5/RPC_primer
