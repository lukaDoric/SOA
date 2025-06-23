# Observability

Observability podrazuvema uvođenja metrika, logova i trace-ova (još se nazivaju **The three pillars of observability**) u sistem koji se želi nadgledati. Ova tri koncepta pomažu da se pruži sveobuhvatno razumevanje kako sistem funkcioniše i mogu se koristiti za identifikaciju i rešavanje problema.

- **Logovi** su zapisi događaja koji se dešavaju unutar sistema, kao što su greške ili radnje korisnika. Oni pružaju detaljan prikaz onoga što se dešava u datom trenutku i mogu se koristiti za otklanjanje grešaka ili revizije.
- **Metrike** su numerička merenja koja prate performanse sistema tokom vremena. Mogu se koristiti za postavljanje upozorenja i identifikovanje trendova ili anomalija. 
- **Tracing** podrazumeva praćenje zahteva kroz sistem da bi se pratio njegov put i identifikovala sva uska grla ili greške. Pruža detaljan pregled toka podataka i može se koristiti za optimizaciju performansi sistema.

Timovi mogu da steknu potpuniju sliku o svojim sistemima i bolje razumeju kako oni rade. Ovo može dovesti do brže dijagnoze i rešavanja problema, kao i do poboljšanja ukupne pouzdanosti i performansi sistema.Uočljivost je posebno važna u složenim distribuiranim sistemima, gde tradicionalni alati za praćenje možda ne pružaju dovoljnu vidljivost. Sa vidljivošću, timovi mogu lakše razumeti interakcije između različitih komponenti i usluga i brzo identifikovati probleme koji mogu da utiču na ukupne performanse sistema. 

Ekosistem je  poprilično veliki kada su pitanju alati koji pružaju prikupljnje logova, metrika i trace-ova. Među najpopularnijima su:

- Monitoring: Prometheus, Datadog, Zabbix.
- Logging: Loki, ElasticSearch.
- Log aggregator: Fluent-bit, Promtail. 
- Tracing: Jaeger, Tempo, Zipkin.
-  Visualisation: Grafana, Kibana. 

Svaki od prethodno navedenih alata ima svoj SDK koji omogućava integraciju između samog alata i aplikacija. Programeri su primorani da koriste SDK i pomoću njega da rade instrumentaciju metrika, logova ili trace-ova (zavisi od konkretnog alata). Ukoliko je neophodno promeniti neki alat iz observability stack-a, neophodno je koristi SDK od tog alata i izmeniti u svim aplikacijama kod koji se odnosio na instrumentaciju. U velikim sistemima ovo predstavlja ozbiljan problem jer je migracija na novi alat jako bolna. Ovaj problem pokušava da reši [OpenTelemetry](https://opentelemetry.io/) koji uvodi standard za metrike, logove i trace-ove. Sve više observability  alata se prilagođava ovom standardu kako bi se omogućilo da se jednom napisan kod za instrumentaciju ne menja ukoliko se promeniti alat iz observability stack-a. OpenTelemetry je još uvek pod razvojom i razvojni status se razlikuje od jezika do jezika. Status razvoja svih jezika je dostupan [ovde](https://opentelemetry.io/status/). 

**Napomena**: Pitanje je vremena kada će indrustrija da usvoji OpenTelemetry  kao standard tako je topla preporuka krenuti što pre sa upoznavanjem sa OpenTelemetry ekosistemom.

U primeru observability stack čine:

- Prometheus
- Loki
- Fluent-bit
- Jaeger
- Grafana 

Primeri koda koji se koriste u nastavnku lekcije su dostupni na:

- Go sa gin-gonic framework-om: https://github.com/lukaDoric/SOA/tree/main/MonitoringSAGA/monitoring_go
- Java sa SpringBoot framework-om: https://github.com/DanijelRadakovic/server

Kompletan observability stack se pokreće pomoću docker-a, tako da je taj deo identičan u oba primera.  




# Logging

Kada je u pitanju prikupljanje logova aplikacija koje se izvršavaju u kontejneru, sve što je neophodno uraditi jeste konfigurisati da aplikacija upisuje logove na standardni izlas što je po default-u podešeno. Odnosno nema nikakvog dodatnog posla da bi se omogućilo prikupljnaje logava aplikacije. Docker prikuplja logove i privremeno ih skladišti lokalno na fajl sistem sve dok kontejner ne prestane sa radom. Pristupanjem logova se radi pomoću komande:

 `docker log <container -name>`

Za trajno skladištenje logova kontejnera neophono je koristiti odgovarajući alate:

- Log baza: baza koja trajno čuva logove. Treba da omogući indeksiranje logova radi brze pretrage logova. 
- Log agregator: prikuplja logove sa nekog izvora, normalizuje ih i prosledjuje ih u odgovarajuću log bazu.

U konkretnom primeru kao log baza se koristi Loki. Fluent-bit je agregator koji prikuplja docker logove svih kontejnera i prosleđuje Loki-u. Takođe, Grafana se koristi za vizualizaciju logova tako što koristi Loki kao datasource.

Video materijal: https://youtu.be/wWw3fRR9NTk



# Tracing

Za tracing se koristi OpenTelemetry biblioteka s obzirom na to da postoji stabilna verzija za go. Manipulacija trace-ovima se radi pomoću TracingProvider.

### Instanciranje TracingProvider 

```go
var tp *trace.TracerProvider

func initTracer() (*trace.TracerProvider, error) {
// Ukoliko je definisana JAEGER_ENDPOINT env var, intanciraj JagerTracer koji šalje trace-ove Jaeger-u,
// u suprotnom instanciraj FileTracer koji upisuje trace-ove u json fajl 
	url := os.Getenv("JAEGER_ENDPOINT")
	if len(url) > 0 {
		return initJaegerTracer(url)
	} else {
		return initFileTracer()
	}
}

func initFileTracer() (*trace.TracerProvider, error) {
	log.Println("Initializing tracing to traces.json")
	f, err := os.Create("traces.json")
	if err != nil {
		return nil, err
	}
	exporter, err := stdouttrace.New(
		stdouttrace.WithWriter(f),
		stdouttrace.WithPrettyPrint(),
	)
	if err != nil {
		return nil, err
	}
	return trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithSampler(trace.AlwaysSample()),
	), nil
}

func initJaegerTracer(url string) (*trace.TracerProvider, error) {
	log.Printf("Initializing tracing to jaeger at %s\n", url)
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	return trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
		)),
	), nil
}

``` 

## Intrumentacija 

```go
func weaponGetHandler(ctx *gin.Context) {
    // kreiraj novi trace
	traceContext, span := tp.Tracer(serviceName).Start(ctx, "weapon-get")
	defer func() { span.End() }()  // na kraju funkcije zatvori trace

   // kreiraj novi subtrace za komunikaciju sa bazom
	span.AddEvent("Establishing connection to the database")
	db := getDB(ctx)
	if db == nil {
		return
	}
	var weapons []Weapon
	err := db.ModelContext(traceContext, &weapons).Select()
	if err != nil {
        // u slučaju greške zatvori trace i zabeleži grešku u samom trace-u
		httpErrorInternalServerError(err, span, ctx)
		return
	}
	ctx.JSON(http.StatusOK, weapons)
}

func httpErrorInternalServerError(err error, span trace.Span, ctx *gin.Context) {
	httpError(err, span, ctx, http.StatusInternalServerError)
}

func httpError(err error, span trace.Span, ctx *gin.Context, status int) {
	log.Println(err.Error())
	span.RecordError(err) // zabeleži grešku u trace-u
	span.SetStatus(codes.Error, err.Error()) // podesi status
	ctx.String(status, err.Error())
}
```

Video materijal: https://youtu.be/XPwZguzQY4o


# Monitoring

Za instrumentaciju metrika za Prometheus alat koristi se [SDK](https://prometheus.io/docs/guides/go-application/) za golang. SDK po default-u expose-uje metrike koje se odnose na samu go aplikaciju (heap usage, garbace collector latency itd.).   S obrzirom na to da se u primeru koristi gin-gonic framework, korististi se [ginprometheus](https://github.com/zsais/go-gin-prometheus) biblioteka koja expose-uje metrike za HTTP zahteve. Biblioteka koristi Prometheus SDK pomoću kojeg su definisane HTTP metrike. Ukoliko je potrebno definsati proizvodnu metriku, koristiti SDK u te svrhe. Metrike su expose-ovane na [/metrics](https://github.com/zsais/go-gin-prometheus/blob/master/middleware.go#L17) endpoint-u koji Prometheus periodično scrape-uje.   

Video Materijal: https://youtu.be/bjq3faDNsio
