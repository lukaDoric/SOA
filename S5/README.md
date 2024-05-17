## Peti sprint (13.05 - 03.06)

Na osnovu dosadašnjeg izgrađenog projekta potrebno je:  
1. Implementirati SAGA obrazac na barem jednom mestu u projektu. Svaki slučaj SAGE mora uključiti barem dva mikroservisa. Dakle potrebno je osmisliti ili iskoristiti trenutne funkcionalnosti koje uključuju barem 2 mikroservisa kako bi se nad njima implmentirala SAGA.  

2. Tracing:  
Neophodno je implementirati tracing barem negde u okviru mikroservisne aplikacije i prikazati u nekom alatu za vizualizaciju.  

3. Logging:  
Neophodno je implementirati agregraciju logova barem negde u okviru mikroservisne aplikacije i prikazati u nekom alatu za vizualizaciju.  

4. Metrike:
   
4.1 Metrike operativnog sistema host mašine na kojoj će mikroservisna
aplikacija biti podignuta. Minimum treba obezbedtiti informacije o
iskorišćenju procesora, RAM memorije, file sistema i protok mrežnog
saobraćaja.  

4.2 Metrike kontejnera koji se koriste u mikroservisnoj aplikaciji. Minimum
treba obezbedtiti informacije o iskorišćenju procesora, RAM memorije,
file sistema i protok mrežnog saobraćaja. 

Materijali:  
- <a href='https://github.com/lukaDoric/SOA/blob/main/S5/asinhrona_komunikacija.md'>Asinhrona komunikacija.</a>  
- <a href='https://github.com/lukaDoric/SOA/blob/main/S5/api_composition_saga.md'>API composition i SAGA.</a>
- <a href='https://github.com/lukaDoric/SOA/blob/main/S5/monitoring.md'>Monitoring/Logging/Tracing</a>

<b>Napomena:</b> Kako komunikacija dve dokerizovane ASP .NET aplikacije zahteva rad sa sertifikatima i dodatnu konfiguraciju projekata, servise (ili gateway) u .NET ne morate dokerizovati. Takođe kao i za prethodnu KT ako ne možete da pokrenete sve servise kroz docker (zbog memorije) uradite u lokalu ili dokerizujte one servise nad kojima ste implementirali SAGU ili bilo kakav monitoring. Dakle možete sve uraditi i lokalno bez dokera ali je bitno da pokažete compose na odbrani. Sve navedene tehnologije iz S5 treba da prilagodite tenhologijama koje trenutno koristite u projektu ili istražite sličnu tehnologiju za vaš servis.
