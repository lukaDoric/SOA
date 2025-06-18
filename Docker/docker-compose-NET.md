# Docker - ASP .NET Core

Naredni primer predstavlja pokretanje jedne kompleksnije ASP .NET Core  aplikacije. Primer uključuje u sebe jedan <b>Dockerfile</b> za .NET aplikaciju, <b>docker-compose.yml</b> koji pokreće Dockerfile zajedno sa bazom (kontejnerom za postgresql), i na kraju <b>docker-compose-migration.yml</b> zadužen za migracije (jer u okviru .NET projekta ne radimo migracije automatski na pokretanje aplikacije).  

Možeš preuzeti sledeći <a href='https://drive.google.com/file/d/1dsc7rp_bD9Ti4UGQkCfNTKpshlVFNnuP/view?usp=sharing'>primer</a> i analizirati sledeće fajlove - Dockerfile, docker-compose.yml, docker-compose-migration.yml.  

## Dockerfile

U narednih nekoliko linija Dockerfile-a povlačimo slike za izgradnju ASP .NET Core aplikacije i pokrećemo dotnet komande za restore dependency-a i build-a.  

```code
FROM mcr.microsoft.com/dotnet/aspnet:7.0 AS base
WORKDIR /app

FROM mcr.microsoft.com/dotnet/sdk:7.0 as build
WORKDIR /src
COPY . .
WORKDIR /src/src
RUN dotnet restore Explorer.API/Explorer.API.csproj
RUN dotnet build Explorer.API/Explorer.API.csproj -c Release

FROM build as publish
RUN dotnet publish Explorer.API/Explorer.API.csproj -c Release -o /app/publish

ENV ASPNETCORE_URLS=http://+:80
FROM base AS final
COPY --from=publish /app .
WORKDIR /app/publish
CMD ["dotnet", "Explorer.API.dll"]
```
1. Prvo postavljamo baznu sliku za sve naredne stage-ove. Ova komanda puluje ASP.NET sliku verzije 7.0 sa majkrosoftovog repozitorijuma za kontejnere.  
2. Potom se pozicioniramo se u okviru kontejnera u folder /app.  
3. Započinjemo novi stage koji je baziran na .NET SDK slici verzije 7.0. Ovaj stage koristimo za build aplikacije.  
4. Menjamo radni direktorijum na /src u okviru kontejnera.  
5. Kopiramo sve (iz našeg foldera gde je komanda za pokretanje dockerfile-a pokrenuta) u /src folder u okviru kontejnera.  
6. Menjamo radni direktorijum u /src/src.  
7. Radimo restore dependency-a u projektu (unutar kontejnera).  
8. Radimo build projekta unutar kontejnera.
9. Zatim kreiramo novi stage 'publish' koji se zasniva na build stage-u.  
10. Publish-ujemo .NET projekat u direktorijum /app/publish u okviru kontejnera.
11. Kroz ENV varijablu postavljamo port na kojoj bi .NET aplikacija trebala da sluša u okviru kontejnera.  
12. Kreiramo finalni stage.  
13. Kopiramo sve iz publish stage-a u /app direktorijum u okviru kontejnera.  
14. Najzad pokrećemo komandu za pokretanje .NET aplikacije.

Drugi deo Dockerfile-a ima stage za migraciju baze podataka i još uvek ga ne treba pokretati. Dockerfile koji je opisan iznad se koristi za podizanje kontejnera u kom će se pokretati naša monolitna aplikacija (.NET), ovaj Dockerfile biće korišćen od strane docker-compose.yml fajla objašnjenog u nastavku.  

## docker-compose

```code
version: '3.9'

services:
  explorer:
    build:
      dockerfile: Dockerfile
      context: .
      target: final
    restart: on-failure
    networks:
      - application
      - database
    ports:
      - "8080:80"
    environment:
      DATABASE_HOST: database
      DATABASE_PORT: 5432
      DATABASE_PASSWORD: super
      DATABASE_USER: postgres
      DATABASE_SCHEMA: explorer
    depends_on:
      - database
```

1. Od linije services: počinjemo da nabrajamo servise koje ćemo imati u okviru mreže. 
2. Prvi servis je explorer (monolitna aplikacija).  
3. Specificiramo koji Dockerfile da koristi za izgradnju explorera, i koji stage da targetira.  
4. Definišemo kojim mrežama će servis pripadati.  
5. Definišemo mapiranje portova. U konkretnom slučaju monolit će biti dostupan na našem računaru na portu 8080, a zauzeće port 80 u kontejneru.   
6. Potom podešavamo ENV varijable koje će koristiti monolitna aplikacija. Konkretno ovo su ENV varijable vezane za postgre bazu tako da i ona sama (njen kontejner) mora imati isti port, user, password itd.

Drugi deo docker-compose.yml fajla je zadužen za pokretanje kontejnera baze.

```code

   database:
    image: postgres:13
    restart: always
    networks:
      - database
    environment:
      POSTGRES_PASSWORD: super
      POSTGRES_USER: postgres
      POSTGRES_DB: explorer
    volumes:
      - type: volume
        source: database-data
        target: /var/lib/postgresql/data
      - type: bind
        source: explorer-init-data.sql
        target: /tmp/explorer-init.sql
    ports:
      - "5432:5432"
 ```

1. Definišemo sliku koja će se koristiti za bazu.  
2. Definišemo u kojoj mreži će baza biti (bitno je da bude u istoj mreži kao aplikacija koja je koristi).  
3. Dfinišemo ENV varijable za bazu (bitno je da budu isti kao one koje smo prosledili bekend aplikaciji radi konekcije).
4. Definišemo volum-e za bazu kako bi mogli da kopiramo inicijalnu SQL skriptu za bazu koju ćemo kasnije pokrenuti.
5. Definišemo mapiranje portova.  
   
Do kraja se definišu volume-i za bazu i konfiguriše mreža.

## Pokretanje monolitne aplikacije

Komandom docker-compose up u okviru tourism-be foldera ćemo podići svu potrebnu infrastrukturu dakle bekend (monolit) i bazu. Nakon pokretanja komandi potrebno je sačekati neko vreme da svi servisi budu dostupni. Komandom docker ps u terminalu možemo videti pokrenute kontejnere, a možemo ih videti i na docker desktop-u.

![image](https://github.com/lukaDoric/SOA/assets/57589408/5fed68d8-e5c6-4f27-9dd1-61242ce1b5c7)
  
Potom je potrebno pokrenuti komandu docker-compose -f docker-compose-migration.yml up kako bi migrirali podatke (ispratiti u konzoli da su sve migracije prošle). Nakon toga možemo se zakačiti na postgres kontejner da vidimo da li su tu sve tabele koje bi trebale da budu. Komanda je docker exec -it ID-KONTEJNERA-ZA-BAZU bash.  

![image](https://github.com/lukaDoric/SOA/assets/57589408/4c3deff8-1d15-420b-91c8-573dba8982fc)  

Potom se možemo ulogovati u bazu. Komanda je psql -U postgres. Nakon toga možemo izlistati sve baze koje imamo, komanda je \l.  

![image](https://github.com/lukaDoric/SOA/assets/57589408/d796672e-1dc8-4b81-bf38-e480d218ec1b)  

Potom se možemo nakačiti na posebnu bazu pa i na posebnu šemu i proveriti da li su tu tabele. Komanda za kačenje na bazu je \c <IME-BAZE>  

![image](https://github.com/lukaDoric/SOA/assets/57589408/6c42e006-0b60-4893-a9fb-6b8a72241db7)  

Sada možemo izlistati sve šeme komandom \dn  

![image](https://github.com/lukaDoric/SOA/assets/57589408/19d9b722-bccf-41f1-96f0-84c3f129a7a9)  

Na kraju možemo videti sve tabele u okviru šeme komandom \dt stakeholders.*  

![image](https://github.com/lukaDoric/SOA/assets/57589408/ebb46bcd-27a6-4c0f-b3a7-cc22d76fceea)  

Kada smo utvrdili da se sve uspešno migriralo možemo probati da pristupimo aplikaciji tako što u browser-u pristupimo http://localhost:8080/swagger/index.html. Obrati pažnju da je u ovom primeru .NET aplikacije obrisan deo iz konfiguracije projekta tj. sledeći if
```code
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}
```
kako bi nam se swagger pokrenuo i u produkcionom okruženju.  

Na kraju možemo popuniti bazu podacima komandom docker exec -it ID-KONTEJNERA-ZA-BAZU psql -U postgres -d explorer -f /tmp/explorer-init.sql  

![image](https://github.com/lukaDoric/SOA/assets/57589408/d0856479-6c94-4711-abdb-cedf487ff09d)  

Aplikacija je sada dokerizovana, uvezana na bazu i baza je popunjena osnovnim informacijama. Sada je moguće kontaktirati aplikaciju i sa frontend-a, obrati pažnju na putanju koju gaćaš i da je http umesto https (jer nismo u ovom pirmeru podešavali https).

## Zaustavljanje aplikacije

Na kraju možeš sve zaustaviti komandom <b>docker compose down</b>, ukoliko želiš da obrišeš sve zaustavljene kontejnere možeš pokrenuti <b>docker system prune -a</b>, a za čišćenje volume-a možeš pokrenuti <b>docker volume prune</b>














