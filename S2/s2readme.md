## Drugi sprint (18.03 - 1.04)

Na osnovu arhitekture sa prve KT (početna monolitna aplikacija i dva izolovana servisa) potrebno je:  
1. Dokerizovati čitav bekend deo (frontend nije potrebno dokerizovati).  
1.1 Potrebno je napisati Dockerfile za svaki servis (monolitnu i dva izolovana servisa).  
1.2 Potrebno je napisati docker compose fajl koji uključuje u sebe podizanje sva tri servisa i odgovarajuće baze podataka.  
1.3 Potrebno je omogućiti komunikaciju između izolovanih kontejnera i od frontenda do kontejnera monolitne aplikacije.  

<a href='https://github.com/lukaDoric/SOA/blob/main/S2/docker.md'>Teorijske osnove Docker-a.</a>   
<a href='https://github.com/lukaDoric/SOA/blob/main/S2/docker-golang.md'>Primer rada sa Docker-om (Golang).</a>  
<a href='https://github.com/lukaDoric/SOA/blob/main/S2/docker-compose-NET.md'>Primer rada sa Docker-om (.NET).</a>   

Napomena: U materijalima ima i primer za dokerizaciju servisa pisanih u Javi.  
Ukoliko ste krenuli razvoj projekta od modularnog monolita (.NET) savet je da kada pogledate sve materijale samo proširite poslednji (.NET) primer i njegov compose sa dodatna dva servisa.
