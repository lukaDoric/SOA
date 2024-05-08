## Četvrti sprint (13.05 - 13.05)

Na osnovu dosadašnjeg izgrađenog projekta potrebno je:  
1. Izolovati Stakeholders modul (tj. modul zadužen za izdavanje tokena) od API sloja.  
2. API sloj izmeniti tako da prihvata REST zahteve i prevodi ih u RPC (grpc-gateway) i dalje šalje ka servisima.  
3. Bilo kakva komunikacija iza API Gateway-a mora da se odvija po RPC-u.  
4. U projektu mora postojati validacija JWT (u API sloju ili odvojenom servisu).  
4. Potrebno je sve dokerizovati (Ukoliko nemate ni jednu mašinu u timu sposobnu da pokrene sve servise u 4. KT možete prezentovati bez dokera ali morate pokazati Dockerfile-ove i docker-compose).

Napomena: Ako ste razvijali projekat od početnog rešenja (sa PSW) gateway možete ostaviti u .NET-u ali ima malo slabiju podršku za rad sa RPC tj. nema automatsko preusmeravanje zahteva na određeni servis već ćete morati ručno da instancirate klijenta u kontroleru. Nakon pogledanih materijala i za .NET i za Golang odlučite na nivou tima u kojoj tehnologiji ćete implementirati API Gateway.

- Komunikacija izmedju servisa:
  - <a href='https://github.com/lukaDoric/SOA/blob/main/S5/sinhrona_komunikacija.md'>Sinhrona komunikacija.</a>   
  - <a href='https://github.com/lukaDoric/SOA/blob/main/S5/asinhrona_komunikacija.md'>Asinhrona komunikacija.</a>  
  - <a href='https://github.com/lukaDoric/SOA/blob/main/S5/grpc_gateway.md'>Grpc Gateway.</a>
- <a href='https://github.com/lukaDoric/SOA/blob/main/S5/api_composition_saga.md'>API composition i SAGA.</a>
- <a href='https://github.com/lukaDoric/SOA/blob/main/S5/monitoring.md'>Monitoring/Logging/Tracing</a>

