### gRPC Gateway

Pošto gRPC ne koristi JSON, XML ili neki drugi tekstualni tip poruka već binarni, ponekad može biti nezgodno testirati vašu aplikaciju ili može biti nezgodno da se vaša aplikacija otvori za veći krug klijenata. Jedna opcija je da oni koriste gRPC takođe, međutim ako to nije opcija mora postojati način da se uradi translacija JSON u protobuf i obrnuto. Tu operaciju radi **gRPC GatewayPlugin**.  

<img src="https://i.ibb.co/PTv2gn7/grpc-rest-gateway.png" alt="grpc-rest-gateway" border="0">  

Kao i kod standardnog gRPC-a, isto koristimo .proto datoteku za specifikaciju našeg gateway-a. Pre upotrbe ovog alata, potrebno je instalirati par biblioteka koje će sve to omogućiti. Primer možete videti na sledećem linku - https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/introduction/.  

Potrebno je proširiti *.proto file dodatnom sintaksom koja će biti korišćena za translaciju JSON-a u proto poruke kao na primeru ispod.  

```
service Greeter {
// Sends a greeting
rpc SayHello (HelloRequest) returns (HelloReply) {
    option (google.api.http) = {
        post: "/v1/example/echo"
        body: "*"
        };
    }
}
```  

Ovde možemo da vidimo da je prethodno definisani servis proširen za par elemenata gde vidmo da za HTTP POST zahtev se koristi **/v1/example/eho** putanja i da prihvata sve poruke u **body**. Isto tako mozemo da specificiramo i ostale HTTP metode.  

Druga razlika je u tome, što moramo malo proširiti naš Golang kod, da pokrenemo web server koji će prihvatiti HTTP zahteve spoljnog sveta i prebaciti ih u gRPC i protobuf zahtev i propagirati dalje u naš sistem. Primer možete videti na linku - https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/adding_annotations/.  

gRPC ima mnoštvo korisnih plugina koje možete da iskoristite i postoji dosta dostupnog middleware-a za bezbednost i proveru raznih stvaru u vašim zahtevima. U početku, oni mogu da deluju malo konfuzno i komplikovano ali kada se naviknete oni će vam znatno olakšati posao. Svi funkcionišu na sličan način, a to je presretanje zahteva, implementacija nekakve logike i propagacija zahteva dalje ili ne u zavisnosti od tipa middleware-a.  

Primere za razne biblioteke možete naći na linku - https://github.com/grpc-ecosystem/go-grpc-middleware. A uvek možete implementirati i vaš, ako za time imate potrebe.  

## Primer postavke gateway-a (Golang)  
### (Obavezno pogledati primere iako koristite .NET Gateway jer su osnovni koncepti objašnjeni u videima)
  
U narednom <a href='https://www.youtube.com/watch?v=sIkZAWDgWkI'>videu</a> možeš pogledati postavku gRPC gateway-a (Golang).   
Primer iz videa je dostupan na sledećem <a href='https://drive.google.com/file/d/1h3o1FMtmmUbCZBGK4_IFcPiZDxpahow9/view?usp=drive_link'>link-u</a>. 

protoc komanda koja je korišćena u primeru:  

protoc -I ./proto \
--go_out ./proto --go_opt paths=source_relative \
--go-grpc_out ./proto --go-grpc_opt paths=source_relative \
--grpc-gateway_out ./proto --grpc-gateway_opt paths=source_relative \
./proto/greeter/greeter-service.proto

U narednom <a href='https://youtu.be/RRG3LVRBuco'>videu</a> možeš pogledati postavku gRPC gateway-a (Golang) koji je sposoban da samostalno obradi zahteve (bez daljeg prosleđivanja servisima).   
Primer iz videa je dostupan na sledećem <a href='https://drive.google.com/file/d/1KeoXdCDcLFM_0IPtCy_ZNfIHpWAAaCwd/view?usp=sharing'>link-u</a>.    

Primer sa složenijim Protobuf porukama (Napomena primer u sebi ima uključen tracer za praćenje zahteva i key-value bazu koju nismo posebno obrađivali ali vam može biti značajan da vidite neke složenije poruke, primer nije obavezno za sada prolaziti). TODO: Dodati primer. 

## Primer postavke gateway-a (.NET)
### (Kako se koriste koncepti iz prethodnih Golang videa, prvo je potrebno njih pogledati, ovo je samo primer kako bi izgledala implementacija u .NET tehnologiji)

Po uzoru na Golang primer, slično ćemo primeniti i na ASP .NET Core aplikaciju. Cilj je da između mikroservisa i između API Gateway-a i mikroservisa imamo RPC komunikaciju, dok se svetu tj. bilo kom klijentu moramo prilagoditi odnosno pružiti im mogućnost da koriste naše usluge po REST specifikaciji.  

Krećemo od početnog projekta gde ćemo fokus staviti na API sloj (cilj je da ga izolujemo). Najpravilnije bi bilo da pored API sloja
nemamo ni jedan modul više prisutan, jer ovo treba da predstavlja samo ulaznu tačku u naš sistem. Dakle želimo da odvojimo Stakeholders modul od API projekta tj. da API projekat postane fizički izolovan i da predstavlja API Gateway (jedinu ulaznu tačku koja stoji ispred mikroservisne arhitekture) pri čemu će još imati i funkciju prevođenja RESTful zahteva u gRPC. 

<b>Napomena:</b> Nakon što izolujete Stakeholders modul moguće je da će vam zbog prethodne arhitekture ostati još neki od modula prikačeni na API projekat ali ih za sada možete zanemariti.

### Kreiranje API Gateway-a

Preporuka je preuzeti <a href='https://drive.google.com/file/d/1zRoRaByk2OVCymbOJTdnKxdNqqVkemJO/view?usp=drive_link'>primer</a> i kroz njega pratiti naredne korake.    

1. Prvi korak je da dodamo potrebne dependency-je u projekat kako bi radili sa gRPCem.  

U okviru Explorer.API, u okviru ItemGroup elementa dodajemo sledeće:  

```code
<ItemGroup>
    <PackageReference Include="Grpc.AspNetCore" Version="2.49.0" />
    <PackageReference Include="Google.Protobuf" Version="3.26.1" />
    <PackageReference Include="Grpc.Net.Client" Version="2.52.0" />
    <PackageReference Include="Grpc.Tools" Version="2.62.0">
      <IncludeAssets>runtime; build; native; contentfiles; analyzers; buildtransitive</IncludeAssets>
      <PrivateAssets>all</PrivateAssets>
    </PackageReference>
    <PackageReference Include="Microsoft.AspNetCore.Grpc.JsonTranscoding" Version="7.0.17" />
</ItemGroup>
```
U okviru PropertyGroup elementa, dodajemo sledeće:  

    <IncludeHttpRuleProtos>true</IncludeHttpRuleProtos>  

Nakon toga možemo napisati .proto specifikaciju za jedan kontroler koji će prihvatati kredencijale od korisnika po REST specifikaciji prevoditi taj zahtev u RPC poruku koju će proslediti izolovanom Stakeholders modulu.  

2. Kreiramo Protos folder u okviru Explorer.API-a.
3. Preuzmite i skinite Google folder (iz okačenog primera) i ubacite u Protos folder.

![image](https://github.com/lukaDoric/SOA/assets/57589408/fb3c9612-ffbd-44d6-8b8e-770578a83ab2)


4. Pišemo .proto specifikaciju u okviru authentication.proto fajla.

```code
syntax = "proto3";

option csharp_namespace = "GrpcServiceTranscoding";
import "Protos/google/api/annotations.proto";

package Authorize;

service Authorize {
  rpc Authorize (Credentials) returns (AuthenticationTokens) {
    option (google.api.http) = {
      post: "/v1/authorize"
      body: "*"
    };
  }
}

// The request message containing the user's name.
message Credentials {

  string Username = 1;
  string Password = 2;
}

// The response message containing the greetings.
message AuthenticationTokens {
  int32 Id = 1;
  string AccessToken = 2;
}
```

5. Stavljamo putanju do .proto specifikacije u okviru ItemGroup elementa - <Protobuf Include="Protos\authentication.proto" />
6. Pokrećemo build komandu (u VisualStudio okruženju). 

Napomena: Nakon ovog koraka možete resetovati okruženje jer nekad nije svesno novogenerisanih fajlova.
Napomena: Sada u okviru src -> Explorer.API -> obj -> Debug -> net7.0 -> Protos treba da vidiš generisane proto specifikacije (Authentication.cs i AuthenticationGrpc.cs).

7. Napravimo kontroler koji će implementirati proto specifikaciju (npr. AuthenticationProtoController, slika ispod).

```code
public class AuthenticationProtoController : Authorize.AuthorizeBase
{
    private readonly ILogger<AuthenticationProtoController> _logger;

    public AuthenticationProtoController(ILogger<AuthenticationProtoController> logger)
    {
        _logger = logger;
    }

    public override async Task<AuthenticationTokens> Authorize(Credentials request,
        ServerCallContext context)
    {
        var httpHandler = new HttpClientHandler();
        httpHandler.ServerCertificateCustomValidationCallback = HttpClientHandler.DangerousAcceptAnyServerCertificateValidator;
        var channel = GrpcChannel.ForAddress("https://localhost:44332", new GrpcChannelOptions { HttpHandler = httpHandler });

        var client = new Authorize.AuthorizeClient(channel);
        var response = await client.AuthorizeAsync(request);

        Console.WriteLine(response.AccessToken);

        return await Task.FromResult(new AuthenticationTokens
        {
            Id = response.Id,
            AccessToken = response.AccessToken
        });
    }
}
````

8. Radimo override generisane metode po proto specifikaciji. Obrati pažnju da kada kontaktiraš drugi mikroservis preko RPCa moraš koristiti sslPort tj. Http 2 protokol jer RPC samo na njemu funkcioniše. Kako bi izbegli rad sa sertifikatima, poturili smo opcije kroz httpHandler.

9. U Program.cs je potrebno dodati:  
  
builder.Services.AddGrpc().AddJsonTranscoding();  
app.MapGrpcService<AuthenticationProtoController>();

### Kreiranje servisa kojeg će kontaktirate Gateway
  
Preporuka je preuzeti <a href='https://drive.google.com/file/d/1bq7snn4MvWhsoJh4aSB_VvS0aLLdIGAB/view?usp=drive_link'>primer</a> i kroz njega pratiti naredne korake.  
  
Pošto moramo sada da kontaktiramo izolovani Stakeholders modul po RPCu, i njemu će biti potrebna (veoma slična) .proto specifikacija kako bi ga mogli kontaktirati i dobiti odgovor.

1. Prvo je potrebno izolovati Stakeholders modul da u jednom nezavisnom servisu budu samo njegov API sloj i sam Stakeholders modul.
2. Dodamo identične dependency-je kao i za Gateway.
3. Kreiramo Protos folder u okviru njegovog API sloja. (slika) Napomena: Nije nam potreban Google folder niti option deo u porukama jer ne radimo konverziju REST -> RPC već samo obrađujemo RPC poruke.

![image](https://github.com/lukaDoric/SOA/assets/57589408/077029ae-939f-4871-b05a-34b19d8523b7)

4. Pišemo .proto specifikaciju u okviru authentication.proto fajla.

```code

syntax = "proto3";

option csharp_namespace = "GrpcServiceTranscoding";

package Authorize;

// The greeting service definition.
service Authorize {
  rpc Authorize (Credentials) returns (AuthenticationTokens) {
  }
}

// The request message containing the user's name.
message Credentials {
  string Username = 1;
  string Password = 2;
}

// The response message containing the greetings.
message AuthenticationTokens {
  int32 Id = 1;
  string AccessToken = 2;
}

```

5. Pokrenemo build komandu (u VisualStudio okruženju).

Napomena: Nakon ovog koraka možete resetovati okruženje jer nekad nije svesno novogenerisanih fajlova.
Napomena: Sada u okviru src -> Explorer.API -> obj -> Debug -> net7.0 -> Protos treba da vidiš generisane proto specifikacije.

6. Napravimo kontroler koji će implementirati proto specifikaciju tj. raditi override generisane metode.

```code

public class AuthenticationProtoController : Authorize.AuthorizeBase
{
    private readonly ILogger<AuthenticationProtoController> _logger;
    private readonly IAuthenticationService _authenticationService;

    public AuthenticationProtoController(ILogger<AuthenticationProtoController> logger, IAuthenticationService authenticationService)
    {
        _logger = logger;
        _authenticationService = authenticationService;
    }

    public override Task<AuthenticationTokens> Authorize(Credentials request,
        ServerCallContext context)
    {
        var credentials = new Stakeholders.API.Dtos.CredentialsDto { Password = request.Password, Username  = request.Username };
        var result = _authenticationService.Login(credentials);
       
        return Task.FromResult(new AuthenticationTokens
        {
            Id = (int)result.Value.Id,
            AccessToken = result.Value.AccessToken,
        });
    }
}

```

7. U Program.cs je potrebno dodati:

builder.Services.AddGrpc();
app.MapGrpcService<AuthenticationProtoController>();

Napomena: Sa obzirom da ima puno detalja (tj. na više mesta je potrebno staviti tačno određenu liniju koda, pogledati kako je urađeno u primerima).
Takođe u okviru API Gateway-a možete sve kontrolere koji implementiraju proto specifikaciju zaštiti kao i do sada određenim policy-em.

Sada možemo da pokrenemo oba servisa (Gateway i Stakeholders), potom kontaktiramo API Gateway, po REST-u sa imenom i šifrom. On će preuzeti zahtev po REST-u, prevesti ga u RPC, kontaktirati Stakeholders mikroservis po RPC gde ćemo dobiti odgovor u vidu JWT-a i prevesti nazad odgovor korisniku po REST specifikaciji.
