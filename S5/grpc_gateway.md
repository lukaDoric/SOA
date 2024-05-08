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

Video materijal:
https://www.youtube.com/watch?v=sIkZAWDgWkI
https://youtu.be/RRG3LVRBuco

gRPC Gateway primer: https://github.com/lukaDoric/SOA/tree/main/S5/gRPC-Gateway
Gateway sa samostalnom obradom zahteva: 
