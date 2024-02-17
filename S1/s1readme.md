## Prvi sprint (26.02 - 18.03)

Počevši od monolinte aplikacije (koju smo gradili na predmetu "Projektovanje softvera") za prvi sprint je potrebno:
1. <b>Fizički izolovati</b> (neke module) i omogućiti http komunikaciju između njih tj. kreirati početnu mikroservisnu arhitekturu. Više o mikroservisima možeš pročitati <a href='https://github.com/lukaDoric/SOA/blob/main/S1/mikroservisi.md'>ovde </a>.
Kako će biti potrebno uspostaviti http komunikaciju između mikroservisa, možeš pogledati sledeće linkove:  
<a href='https://learn.microsoft.com/en-us/dotnet/fundamentals/networking/http/httpclient'> Make HTTP requests with the HttpClient class - C# </a>  
<a href='https://pkg.go.dev/net/http'> Http package - Golang </a>

2. U podtimovima od po dvoje potrebno je da <b>prevedete izolovani modul</b> u programski jezik koji se razlikuje od onog u kom je trenutno aplikacija, preporuka je da to bude Golang. Više o Golangu možeš pročitati <a href='https://github.com/lukaDoric/SOA/blob/main/S1/Golang/golang.md'>ovde</a>.

Naš predlog je da za prvi sprint izolujete Tours i Encounters module (jer imaju najviše funkcionalnosti), svakako možete odabrati i druge ali bitno je ispuniti dva zahteva:
1. Da postoji http komunikacija sa drugim delovima aplikacije.
2. Da svako prevede barem tri funkcionalnosti tako da može da se demonstrira kroz front aplikaciju da funkcionalnost radi.
