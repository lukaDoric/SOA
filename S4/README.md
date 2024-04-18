## Četvrti sprint (22.04 - 13.05)
Na osnovu dosadašnjeg izgrađenog projekta potrebno je:  
1. Izolovati Stakeholders modul (tj. modul zadužen za izdavanje tokena) od API sloja.  
2. API sloj izmeniti tako da prihvata REST zahteve i prevodi ih u RPC (grpc-gateway).  
3. Bilo kakva komunikacija iza API Gateway-a mora da se odvija po RPC-u.  
4. Potrebno je sve dokerizovati (Ukoliko nemate ni jednu mašinu u timu sposobnu da pokrene sve servise u 4. KT možete prezentovati bez dokera ali morate pokazati Dockerfile-ove i docker-compose).

Napomena: Ako ste razvijali projekat od početnog rešenja gateway možete ostaviti u .NET-u ali ima malo slabiju podršku za rad sa RPC tj. nema automatsko preusmeravanje zahteva na određeni servis već ćete morati ručno da instancirate klijenta u kontroleru. Nakon pogledanih materijala i za .NET i za Golang odlučite na nivou tima u kojoj tehnologiji ćete implementirati API Gateway.

<a href='https://github.com/lukaDoric/SOA/blob/main/S4/api-gateway.md'>Teorijske osnove API Gateway-a.</a>   
<a href='https://github.com/lukaDoric/SOA/blob/main/S4/rpc.md'>gRPC osnove.</a>  
<a href='https://github.com/lukaDoric/SOA/blob/main/S4/rpc-gateway.md'>gRPC Gateway.</a>  
