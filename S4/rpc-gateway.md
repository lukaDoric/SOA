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

## Primer postavke gateway-a (Golang) - (Obavezno pogledati primere iako koristite .NET Gateway jer su osnovni koncepti objašnjeni u videima)
  
U narednom <a href='https://www.youtube.com/watch?v=sIkZAWDgWkI'>videu</a> možeš pogledati postavku gRPC gateway-a (Golang).   
Primer iz videa je dostupan na sledećem <a href='#'>link-u</a>. TODO: Dodati primer. 

protoc komanda koja je korišćena u primeru:  

protoc -I ./proto \
--go_out ./proto --go_opt paths=source_relative \
--go-grpc_out ./proto --go-grpc_opt paths=source_relative \
--grpc-gateway_out ./proto --grpc-gateway_opt paths=source_relative \
./proto/greeter/greeter-service.proto

U narednom <a href='https://youtu.be/RRG3LVRBuco'>videu</a> možeš pogledati postavku gRPC gateway-a (Golang) koji je sposoban da samostalno obradi zahteve (bez daljeg prosleđivanja servisima).   
Primer iz videa je dostupan na sledećem <a href='#'>link-u</a>. TODO: Dodati primer.  

Primer sa složenijim Protobuf porukama (Napomena primer u sebi ima uključen tracer za praćenje zahteva i key-value bazu koju nismo posebno obrađivali ali vam može biti značajan da vidite neke složenije poruke, primer nije obavezno za sada prolaziti). TODO: Dodati primer. 
