## NoSQL graf baze podataka

Za razliku od relacionih i dokument NoSQL baza, graf NoSQL baze skladište **čvorove** (eng. nodes) i **veze između čvorova** (eng. relationships). Opet, podaci se smeštaju **bez striktne šeme** što nam pruža fleksibilnost.

Osnovne karakteristike graf NoSQL baza:
1. U skladištu se čuvaju grafovi u
izvornom obliku - skup povezanih
čvorova.
2. Čvorovi čuvaju informaciju o svojim
susedima - brz pristup na osnovu veza.
3. Ne zahteva dodatne strukture za nove
načine prolaska kroz graf.
4. Brže izvršavanje, manje zauzeće
memorije, pravljano namenski za
grafove.
5. Lako skaliranje sa povećanjem
podataka.

## neo4j

Konkretno u okviru vežbi biće korišćena neo4j graf baza (https://neo4j.com/docs/getting-started/current/languages-guides/neo4j-go/, https://github.com/neo4j/neo4j-go-driver). Neke osnovne karakteristike neo4j baze su sledeće: 

1. Native graf baza podataka (https://neo4j.com/blog/native-vs-non-native-graph-technology/).
2. White-board friendly - logički dijagram se jednostavno preslikava na samu bazu.
3. Skalabilna.
4. Podržan ACID.
5. Jednostavna i raznovsna podrška za upite.
6. Schemeless model podataka - ne zahteva defnisanje šeme.
7. Mogućnost proširenja i dodatka ograničenja nad modelom podataka.
8. Čvorovi i grane grafa su top-level entiteti.
9. Svojstven upitni jezik - Cypher.

### Primer jednog grafa

<img src="https://i.ibb.co/PzmjCJS/GRAPH.png" alt="GRAPH" border="0"><br /><br />


### Terminologija

1. Graf = povezana struktura čvorova  
2. Node = čvor u grafu  
    - predstavlja enitet  
    - odgovara torki u relacionoj bazi  
    - ima id  
    - može imati labele i svojstva  
3. Relationship = grana u grafu  
    - predstavlja usmerenu vezu između 2 čvora  
    - ima id  
    - može imati labele i svojstva  
4. Label = labela; identifikuje tip čvora ili veze  
5. Properties = svojstva; ključ-vrednost parovi koji dodatno opisuju čvorove i veze

### Cypher - jezik za upite

1. Upitni jezik optimizovan za rad sa grafovima.  
2. Jednostavna slikovita sintaksa.  
3. Velika raznovrsnost upita i njihovih kombinacija.  
4. Cheat-sheet sintakse - https://neo4j.com/docs/cypher-cheat-sheet/current/.  
5. Naredbe su case-insensitive dok su labele i svojstva case-sensitive!  

<img src="https://i.ibb.co/VJk9vTr/GRAPH-2.png" alt="GRAPH-2" border="0"><br /><br />

1. () - zagrade predstavljaju čvorove
2. -[]-> - usmerene linije predstavljaju grane
3. {} - vitičaste zagrade predstavljaju svojstva
4. :Naziv - labele se označavaju sa početne :

### Create
Kreiranje čvora:
```
CREATE (pera:Person {name: "Pera", surname: "Peric", age: 23})
```

Kreiranje grane između 2 čvora:
```
CREATE (pera) -[:IS_FRIENDS {since: "2007-09-03"}]-> (mika)
```

### Read

Pronalazak čvora na osnovu polja uz projekciju svojstava:

```
MATCH (n {name: "Pera", surname: "Peric"})
RETURN n.name, n.surname, n.age
```
Pronalazak čvora na osnovu ID-a:
```
MATCH (n)
WHERE ID(n) = 123
RETURN n
```

### Update

Izmena vrednosti svojstva čvora:
```
MATCH (p:Person)
WHERE p.name = "Pera"
SET p.age = 28
```
Dodavanje labele čvoru:
```
MATCH (n)
WHERE ID(n) = 123
SET n:Person
```
### Delete
Brisanje čvora koji nema povezanih grana:
```
MATCH (p:Person)
WHERE p.name = "Pera"
DELETE p
```
Brisanje čvora i povezanih grana:
```
MATCH (n)
WHERE ID(n) = 123
DETACH DELETE n
```  

U narednom <a href='https://www.youtube.com/watch?v=aL7mN_l2jLM'>videu</a> možeš pogledati osnovne komande nad mongo bazom.  
U narednom <a href='https://www.youtube.com/watch?v=06ehbvxflL4'>videu</a> možeš pogledati integraciju mongo baze sa Golang-om.  
Primer je dostupan na sledećem <a href=''>link-u</a>.  
