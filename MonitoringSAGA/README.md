## SAGA i monitoring

Materijali:  
- <a href='https://github.com/lukaDoric/SOA/blob/main/S5/asinhrona_komunikacija.md'>Asinhrona komunikacija.</a>  
- <a href='https://github.com/lukaDoric/SOA/blob/main/S5/api_composition_saga.md'>API composition i SAGA.</a>
- <a href='https://github.com/lukaDoric/SOA/blob/main/S5/monitoring.md'>Monitoring/Logging/Tracing</a>

<b>Napomena:</b> Kako komunikacija dve dokerizovane ASP .NET aplikacije zahteva rad sa sertifikatima i dodatnu konfiguraciju projekata, servise (ili gateway) u .NET ne morate dokerizovati (ko želi da se pozabavi ovim može slobodno i to da uradi). Takođe kao i za prethodnu KT ako ne možete da pokrenete sve servise kroz docker (zbog memorije) uradite u lokalu ili dokerizujte one servise nad kojima ste implementirali SAGU ili bilo kakav monitoring. Dakle možete sve uraditi i lokalno bez dokera ali je bitno da pokažete compose na odbrani. Sve navedene tehnologije iz S5 treba da prilagodite tenhologijama koje trenutno koristite u projektu ili istražite sličnu tehnologiju za vaš servis.
