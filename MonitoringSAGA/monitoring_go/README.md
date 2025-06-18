# Dojo

U ovom primeru imamo CRUD aplikaciju (dojo) goja prima GET i POST zahteve. Koristićemo je za demonstraciju monitoringa i prikupljanja metrika

Prvo treba pokrenuti docker compose gde se nalaze svi alati za metrike:

```shell
cd monitoring
docker compose up -d
```

Zatim treba da pronađemo IP adresu fluent-bit kontejnera (servis za prikupljanje docker logova). Za ovo treba da iskoristite komandu docker ps, i pronađete naziv kontejnera za fluent-bit. U mom primeru je to `monitoring-fluent-bit-1`. 

```shell
docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' monitoring-fluent-bit-1
```

IP adresu koju dobijete treba da upišete u env.conf datoteku, umesto trenutne IP adrese koja stoji pored FLUENT_BIT_ADDRESS promenjive.
Nakon toga možete pokrenuti i mikroservisnu aplikaciju:
```shell
docker compose --env-file env.conf up -d
```

Testirajte aplikaciju uz sledeće HTTP zahteve:
```shell
curl -X POST "http://localhost:8080/weapon?id=0&weapon=katana"
curl -X POST "http://localhost:8080/weapon?id=1&weapon=ninjaStar"
curl -X POST "http://localhost:8080/weapon?id=2&weapon=ninjaSword"

curl "http://localhost:8080/weapon"

```

Na adresi localhost:3000 možete pristupiti Grafana alatu za vizualizaciju metrika. Ako vam traži kredencijale, i username i password su admin. Na osnovu slika koje se nalaze ispod možete kreirati dashboard-ove koji će prikazivati metrike iz go aplikacije, logove, tracing poziva između servisa, itd.

### Tracing

![tracing](./docs/tracing.png)

### Logs

![logs](./docs/logs.png)

### Metrics

![metrics](./docs/metrics.png)
