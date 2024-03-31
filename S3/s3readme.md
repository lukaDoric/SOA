## Treći sprint (01.04 - 22.04)

Na osnovu dosadašnjeg izgrađenog projekta potrebno je:  
1. Na nekom od bilo koja dva izolovana servisa zameniti SQL bazu <b>dokument NoSQL bazom</b> (preporuka MongoDB).  
2. Napraviti mikroservis koji će voditi računa o međusobnom praćenju korisnika (Follower Microservice - možete sami odbarati tehnologiju u kojoj ćete implementirati ovaj mikroservis).  
2.1 Potrebno je omogućiti korisnicima da zaprate druge korisnike.  
2.2 Korisnici mogu da čitaju blogove samo onih korisnika koje su zapratili.  
2.3 Omogućiti korisnicima da vide preporuke koga bi mogli da zaprate. Na primer osoba koju ja pratim, prati neke druge  
    profile pa mi sistem daje preporuku da ih i ja zapratim.  
2.4 Sve iz tačke dva realizovati pomoću graf NoSQL baze (preporuka Neo4j).  
3. Novi mikroservis i nove baze potrebno je dokerizovati, i kao i za drugu KT pokrenuti sve kroz docker compose.


<b>Napomena:</b> Ukoliko ste u prethodnim KT izolovali Stakeholders servis i razdvajali ga na servis za autentifikaciju (Auth Microservice) i na servis koji vodi računa
o ostalim informacijama korisnika (UserManagement Microservice) možete u okviru UserManagement mikroservisa implementirati sve iz tačke 2. Ukoliko niste izolovali
ništa vezano za Stakeholders modul preporuka je da napravite novi Follower Microservice koji će voditi računa o svemu iz tačke 2.

<a href='https://github.com/lukaDoric/SOA/blob/main/S3/sql-NoSQL.md'>Teorijske osnove NoSQL baza.</a>   
<a href='https://github.com/lukaDoric/SOA/blob/main/S3/mongo-go.md'>Primer rada sa dokument NoSQL bazom (MongoDB).</a>  
<a href='https://github.com/lukaDoric/SOA/blob/main/S3/neo4j-mongo.md'>Primer rada sa graf NoSQL bazom (Neo4j).</a>  
<a href=''>Primer REST servisa sa dokument i graf bazom.</a>   

