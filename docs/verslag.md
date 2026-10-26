# Titel verslag

Versienummer : 1

Auteur : Tim van Loo 

Studentnummer : 0270033

Klas : 25SDBA

Naam Keuzedeel: Software verdieping - 1

Code Keuzedeel: K0505

Datum : 16-09-2026

---

<div style="page-break-after: always;"></div>

## Inhoudsopgave

- [Titel verslag](#titel-verslag)
  - [Inhoudsopgave](#inhoudsopgave)
  - [Inleiding](#inleiding)
  - [Maakt een keuze voor software](#maakt-een-keuze-voor-software)
    - [Ondezoek](#onderzoek)
    - [Vergelijking](#vergelijking)
    - [Keuze](#keuze)
    - [Arbeidsmarkt](#arbeidsmarkt)
    - [Onderbouwing onderzoek en keuze](#onderbouwing-onderzoek-en-keuze)
  - [Bronnen](#bronnen)

<div style="page-break-after: always;"></div>

## Inleiding

Het technologielandschap ontwikkeld zich razendsnel waardoor organisaties continu moeten afwegen welke tools het beste aansluiten bij hun doelen. De keuze voor een specifieke programmeertaal is hierin curciaal, het bepaalt niet alleen de efficiëntie en schaalbaarheid van de software, maar beinvoeldt ook de productiviteit van je werknemers/ontwikkelingsteam. Waar getalsmatige data vaak laat zien welke talen populair zijn, blijft de vraag waarom ontwikkelaars en bedrijven vaak een voorkeur hebben bij een taal boven de andere. Deze vraag blijft vaak onbeantwoord. In dit marktonderzoek beantwoord ik deze vraag door verschillende talen te vergelijken. Vervolgens kies ik er één uit om me te verdiepen in de redenen waarom juist déze taal zo populair is.

## Maakt een keuze voor software

### Onderzoek

- [Softwarepakket Rust]
  - In deze functie wordt Rust gebruikt om complexe, wiskundige software te ontwikkelen voor de reis- en transportindustrie. Het ontwikkelteam zet deze taal specifiek in omdat de applicaties uiterst performant moeten zijn en honderden requests per seconde foutloos moeten kunnen afhandelen.
  - [Zie Link](https://codeblauw.careers/vacature/rust-developer/)
  - [Zie Screenshot](/KeuzeDeel-Verdieping-Software/images/Screenshot%202026-09-25%20at%2012-34-07%20Vacature%20-%20Rust%20Developer%20Eindhoven%20-%20CodeBlauw.png)


- [Softwarepakket Rust]
  - In deze functie wordt Rust gebruikt om simulatiesoftware te ontwikkelen voor houwitser- en mortiersimulatoren. Het team zet deze taal in om ballistische berekeningen uit te voeren, hardware en sensoren aan te sturen en realtime te communiceren met de specialistische apparatuur in de simulator.
  - [Zie Link](https://vanhalteren.com/careers/software-developer-rust/?utm_campaign=google_jobs_apply&utm_source=google_jobs_apply&utm_medium=organic)
  - [Zie Screenshot](/KeuzeDeel-Verdieping-Software/images/Screenshot%202026-09-25%20at%2012-35-15%20Software%20Developer%20Rust%20Van%20Halteren%20Technologies.png)

- [Softwarepakket Rust]
  - In deze functie wordt Rust gebruikt om microservices te ontwikkelen in de AWS Cloud, naast Java, Kotlin en Go. Het team kiest zelf welke taal het beste past bij het project en zet Rust in voor software die honderden requests per seconde moet verwerken.
  - [Zie Link](https://matchmatters.nl/vacature/golang-rust-developer])
  - [Zie Screenshot](/KeuzeDeel-Verdieping-Software/images/Screenshot%202026-09-25%20at%2012-35-31%20Go%20_%20Rust%20Developer%20–%20MatchMatters.png)

- [Softwarepakket Go]
  - In deze functie wordt Go gebruikt om de server-side fundering te bouwen voor Castle Busters, waarmee het spel verandert van een client-based naar een server-validated, competitieve game. Het team gebruikt Go voor serversided validatie van de metagame en voor het ontwerpen van API's waar de Unity-client tegenaan integreert.
  - [Zie Link](https://jobs.ashbyhq.com/voodoo/e51a4c50-ee4a-4505-b438-6211681bbf7d?utm_source=nedy21l72a)
  - [Zie Screenshot](/KeuzeDeel-Verdieping-Software/images/Screenshot%202026-09-25%20at%2012-35-39%20Senior%20Backend%20Engineer%20(Golang)%20-%20Castle%20Busters%20@%20Voodoo.png)
  

- [Softwarepakket Go]
  - In deze functie wordt Go gebruikt om backendservices en API's te ontwikkelen voor het betaalplatform van Payter. Het team zet Go in om schaalbare, veilige en high-performance applicaties te bouwen die wereldwijd contactloze en cashless betaaltransacties mogelijk maken.
  - [Zie Link](https://payter.recruitee.com/o/backend-developer-golang])
  - [Zie Screenshot](/KeuzeDeel-Verdieping-Software/images/Screenshot%202026-09-25%20at%2012-35-45%20Payter%20B.V.%20-%20Backend%20Developer%20Golang.png)
  

- [Softwarepakket Elixir]
  - In deze functie wordt Elixir gebruikt als onderdeel van de backend-stack (naast Ruby on Rails en GraphQL) voor het bouwen van bedrijfskritische systemen bij Mobiel.nl. Het platform verwerkt complexe processen rond producten, abonnementen, leveranciers, logistiek en orderstromen voor een systeem dat dagelijks door duizenden consumenten wordt gebruikt.
  - [Zie Link](https://werkenbij.mobiel.nl/o/senior-backend-engineer)
  - [Zie Screenshot](/KeuzeDeel-Verdieping-Software/images/Screenshot%202026-09-25%20at%2012-35-56%20Mobiel.nl%20-%20Vacature%20Senior%20Backend%20Engineer.png)

- [Aanvulling]
  - Tijdens het onderzoek viel op dat het aanbod van vacatures per taal sterk verschilt. Voor Rust waren relatief veel vacatures te vinden, voor Go een gemiddeld aantal, terwijl vacatures voor Elixir aanzienlijk schaarser waren. Dit verschil in vindbaarheid geeft mogelijk al een eerste indicatie van de relatieve populariteit en marktvraag van deze talen.
  
### Vergelijking

| Criterium | Rust | Go | Elixir |
| :--- | :--- | :--- | :--- |
| **Marktrelevantie (vacatures)** [^1] | Ruim aanbod op de arbeidsmarkt (momenteel 50+ openstaande vacatures) | Middelmatige marktvraag (momenteel circa 10 openstaande vacatures) | Specifieke nichemarkt (beperkt aanbod van circa 4 vacatures) |
| **Snelheid / Performance** | Uitzonderlijk hoge prestaties (volledig zonder vertraging van een Garbage Collector) [^3] | Zeer hoge prestaties (geoptimaliseerd inclusief het gebruik van een Garbage Collector) [^5] | Hoge prestaties gericht op I/O-doorvoer (dankzij de BEAM virtuele machine)[^6] |
| **Tiobe Index**  | Sterk vertegenwoordigd op de 10e positie in de ranglijst [^4] | Goed vertegenwoordigd op de 12e positie in de ranglijst [^4] | Valt buiten de top van de ranglijst (positie 50 en hoger) [^4] |
| **Geheugenveiligheid**  | [Gegarandeerde en bewezen hoge veiligheid](https://blog.google/security/rust-in-android-move-fast-fix-things/) [^2] | Hoge veiligheid (automatisch beheerd door de ingebouwde Garbage Collector) [^8] | Hoge veiligheid door automatische Garbage Collection (GC) gecombineerd met een strikte 'share nothing' architectuur, waardoor processen elkaars geheugen niet kunnen beinvloeden. [^7] |
| **Ontwikkelsnelheid & Leercurve** | Zeer steile leercurve het expliciet geheugenbeheer lijkt mij te hoog gegrepen om in een kleine periode onder de knie te krijgen. [^10] | Zeer lage drempel: GoLang staat er om bekend dat het een minimalistisch programmeertaal is. Hopelijk lukt het binnen een paar dagen tot weken om productief te zijn in mijn eigen leeromgeving. [^11] | Gemiddelde leercuve: Vriendelijke Ruby-like omgeving/syntax. Maar het feit dat ik op veel websites lees dat het moeilijk is om mee te beginnen zal deze taal niet mijn keuze worden. [^12] |


### Keuze

Mijn Keuze is gevallen op Go(GoLang). Waar Rust een steilere leercurve heeft en veel tijd eist om de strikte geheugenregels onder de knie te krijgen, blinkt Go juist uit in eenvoudigheid en leesbaarheid. De taal is specifiek ontworpen om makkelijker te begrijpen en sneller dingen te kunnen schrijven. Dit zorgt voor een sneller begrip van de taal waardoor ik makkelijker kan beginnen aan het werken aan een project.

### Arbeidsmarkt

Hoewel mijn steekproef aangetoont heeft dat er lokaal op dit moment iets minder directe vacatures voor Go zijn in vergelijking met Rust, is de positie van Go in de bredere IT wereld enorm sterk vooral als het gaat over Cloud Computing. Door een focus te leggen op GoLang zorg ik ervoor dat ik een goede fundering leg voor de alsmaar groeiende arbeidsmarkt. [^9]

### Onderbouwing onderzoek en keuze

Voor dit onderzoek heb ik drie moderne programmeertalen onder de loep gelegd: Rust, Go, en Elixir. Ze zijn allemaal veilig en snel, maar pakken het heel anders aan. Rust kiest voor controle, Elixir voor system die niet kunnen/mogen crashen, en Go is speciaal ontworpen om simpel en snel in te kunnen ontwikkelen. Als ik kijk naar alle vergelijking criteria kom ik tot de conclusie dat ik voor GoLang kies vanwege de lage drempel, gebruikersvriendelijke syntax en een vrij actieve deelname van mensen aan Stackoverflow qua nieuwe vragen over GoLang. 

---

## Maakt zich de software eigen

### Reflectie

**S(ituatie)**: 

Binnen mijn studie volg ik het keuzedeel Software verdieping. mijn focus heeft vooral op front-end gelegen waardoor mijn ervaring met back-end development nog vrij laag lag.

**T(aak)**:

Er wordt van mij verwacht dat ik een marktonderzoek  doe en een verslag aanlever met een vergelijkingsonderzoek om een keuze te maken uit verschillende talen.

**A(ctie)**:

Ik ben begonnen met het onderzoeken van relevante informatie via technische artikelen die ik op het internet kon vinden.

Zodra ik tegen dingen aanliep zoals het vinden van vacatures voor Elixir heb ik ervoor gekozen om vrij snel die taal als een mogelijk voor mijn keuze te laten vallen. 

Ik heb uiteindelijk mijn keuze gebaseerd op de vergelijking met de tijd die ik heb voor het maken van een prototype en de moeilijkheid van het leren van een hele nieuwe taal.

**R(esultaat)**:

Door mijn onderzoek en de vergelijking ligt er nu een volwaardig verslag op tafel die ik kan gebruiken als ondersteuning bij mijn prototype voorstel.

Mijn keuze is beland op Go(Lang) vooral door de lagere instapdrempel, eenvoudige syntax, en een sterke positie op de arbeidsmarkt.

**R(eflectie)**:

Ik ben erg tevreden over mijn werkhouding en doorzettingsvermogen. 

Ik moet wel eerlijk zijn over het feit dat ik het erg lastig vind om het gehele theorie-gedeelte te begrijpen op een goed niveau zonder dat ik ook maar een enkele lijn code heb geschreven in GoLang. 

Ik zal dan ook dit hele verslag meenemen naar de volgende fase en het gebruiken als hulpmiddel voor het maken van een prototype en verder onderzoeken van bijvoorbeeld de syntax van Go(Lang) zelf.

**Wat weet ik al**

Mijn huidige kennis ligt vooral aan de kant van front-end development. Ik heb een basis in HTML, CSS, en Javascript en ben inmiddels begonnen aan het leren van frameworks zoals Angular en Vue. Daarnaast heb ik al ervaring opgedaan tijdens teamprojecten, namelijk via SCRUM. Hierdoor begrijp ik beter hoe een applicatie in elkaar zit en hoe een front-end met een server communiceert.

**Wat weet ik nog niet**

Omdat mijn focus op de front-end lag, en ik mij voor mijn toekomstige stage/loopbaan ook op de front-end development richt, mis ik nog praktijk ervaring met echte back-end logica en de bijhorende talen. Verder merk ik op dat de dingen waar ik onderzoek naar heb gedaan mij nog niet eigen zijn. Bijvoorbeeld de werking van een garbage collector, index statistieken en performance matrixen. Ik heb de data gevonden en geanalyseerd voor het onderzoek maar beheers de technische werking van de bovengenoemde dingen nog niet. 

**Wat wil ik nog leren**

Ik wil de overstap maken van theorie en onderzoek doen naar een werkende applicatie maken door daadwerkelijk te beginnen aan een programma schrijven met GoLang. Mijn doel is om het minimalistisch te houden en de lat niet te hoog te leggen. Ik wil hier een balans in vinden door verder onderzoek te doen naar wat de taal kan en wat haalbaar is voor mij. Uiteindelijk is mijn doel ook om basiskennis van GoLang te kunnen opdoen en bewijs materiaal aan te leveren die dit aantoont.

### Leerdoelen formuleren

#### Technische doelen (Go)

**Leerdoel 1: De basis en syntax van Go onder de knie krijgen**

- **Wat wil ik leren:** Omdat ik vooral ervaring heb met front-end, wil ik leren hoe Go werkt met vaste datatypes, variabelen, loops en functies.
- **Hoe laat ik dit zien (SMART):** Ik toon dit aan door een werkende Tic-Tac-Toe CLI game te bouwen waarin het spelen soepel verloopt. Ik wil gaan toelichten in mijn code met uitleg waarom iets zo werkt.
- **Tijdsplanning:** Dit wil ik afgerond hebben in de eerste/tweede week van het project (Ten minste de basiskennis).

- Ik gebruik een tutorial van GoLang om de basis van de taal te gebruiken. [^13] (https://go.dev/tour/welcome/1)

**Leerdoel 2: Snappen en toepassen van pointers**

- **Wat wil ik leren:** Pointers zijn nieuw voor mij. Ik wil leren wat een pointer precies doet in het geheugen en hoe je die gebruikt om data aan te passen.
- **Hoe laat ik dit zien (SMART):** Ik pas pointers toe in mijn Tic-Tac-Toe game om het 3x3 speelbord bij te werken wanneer een speler bijvoorbeeld een zet doet.
- **Tijdsplanning:** Dit wil ik afgerond hebben in de derde week van de start van mijn project.

**Leerdoel 3: Foutafhandeling in Go toepassen**

- **Wat wil ik leren:** Ik wil leren hoe Go errors teruggeeft en hoe je deze netjes afhandelt.
- **Hoe laat ik dit zien (SMART):** Ik zorg dat mijn applicatie niet crasht als een speler een verkeerde invoer geeft. De game geeft dan een duidelijke foutmelding in de terminal en vraagt om een nieuwe zet.
- **Tijdsplanning:** Dit wil ik afgerond hebben in de laatste fase van mijn project. Dus vooral als het gaat om debugging wil ik kunnen laten zien dat de fouten goed gelogd worden en ik efficient de dingen kan oplossen via een correcte foutafhandeling.

#### Persoonlijke doelen & werkwijze

**Leerdoel 4: Plannen en overzicht houden met GitHub Projects**

- **Wat wil ik leren:** Ik wil tijdens dit project gestructureerd werken via een Kanban-systeem, zodat ik het overzicht bewaar over wat er nog moet gebeuren.
- **Hoe laat ik dit zien (SMART):** Ik maak vanaf het begin een Kanban-bord aan in GitHub Projects met kolommen (zoals To Do, In Progress, Done). Ik evalueer regelmatig [zelf invullen: bijv. wekelijks / per werksessie] of ik nog op schema lig en pas de planning daarop aan.
- **Tijdsplanning:** Dit pas ik toe gedurende het hele project.

**Leerdoel 5: Een duidelijk logboek bijhouden**

- **Wat wil ik leren:** Ik wil leren om mijn proces, keuzes en problemen beter vast te leggen terwijl ik bezig ben, in plaats van achteraf pas te reflecteren.
- **Hoe laat ik dit zien (SMART):** Na elke werksessie noteer ik hoeveel tijd ik erin heb gestoken, tegen welke fouten of problemen ik aanliep en hoe ik die heb opgelost. Dit werk ik uit in het logboek van dit verslag.
- **Tijdsplanning:** Dit houd ik bij tijdens het hele project.

### Prototype voorstel

Voor mijn prototype ga ik een CLI applicatie bouwen van het spel Tic Tac Toe. Dit project biedt een omgeving om te focusen op een nieuwe taal en een volledig nieuwe syntax. Het spel is qua regels simpel, waardoor ik al mijn aandacht kan richten op de technische implementatie en het schrijven van mooie en schone code.

**Functionaliteiten**: 

- Visuele weergave: Een 3x3 spelbord dat dynamisch in de terminal wordt geprint en na elke zet wordt geupdate.

- Spelersinvoer: Twee spelers die om de beurt een zet kunnen doen door een getal in te voeren gebaseerd op een grid. 

- Validatie en foutafhandeling: Het programma controleert of een invoer uberhaupt geldig is en zal dit ook weergeven in de terminal.

- Spelregels toepassen: Het programma controleert na elke beurt automatisch of er sprake is van een winnaar of een gelijkspel.

**Technische uitwerking**: 

- Variabelen, Datatypes en Loops: Voor het bijhouden vande beurten en het draaien van het spel zelf. Dit moet blijven draaien tot het spel voorbij is.

- Arrays: Voor het opslaan van de staat van het spelbord de 3x3 grid.

- Functies: De logica wordt opgedeeld in kleine, herbruikbare functies.

- Error handling: Ik zal Go's ingebouwde manier van foutafhandeling actief toepassen bij het verwerken van de CLI invoer.


### Planning

**Week 1-2**: Setup & Syntax leren
 - Go installeren, VS Code configureren, leren werken met `go run main.go` en `go build`.
 - De officiële interactieve gids doen: A Tour of Go (tour.golang.org) [^13].
 - Begrijpen wat Go anders maakt dan JS: striktere datatypes (`int, string, bool`), geen `let/const` maar `var` en `:=`.

**Week 3-4**: Basis logica & structuren
 - Oefenen met loops (for is de enige loop in GoLang) en conditionals (if/else, switch).
 - Slices vs Arrays (hoe Go omgaat met lijsten van data).
 - Structs ontdekken (hoe Go data structureert zonder classes).

**Week 5**: Eerste mini-script & opzet
 - Een mini-console script schrijven: vraag de gebruiker om input in de terminal en print iets uit.
 - GitHub Projects bord inrichten voor het bijhouden van de taken.

**Week 6**: Speelbord & Spelloop
 - 3x3 grid definiëren in Go (`[3][3]string`).
 - Functie maken die het bord netjes print in de terminal met rasterlijnen.
 - De spelloop maken: beurt afwisselen tussen speler "X" en speler "O".

**Week 7**: Winst- en gelijkspelcontrole
 - Logica schrijven om rijen, kolommen en diagonalen te checken voor drie op een rij.
 - Checken of het bord vol is zonder winnaar (gelijkspel).

**Week 8**: Pointers & Foutafhandeling
 - Pointers toepassen om het speelbord direct bij te werken in functies.
 - Foutafhandeling inbouwen: zorgen dat het spel niet crasht bij verkeerde invoer (zoals letters of bezette vakjes).

**Week 9**: Video & Documentatie
 - Korte schermopname / video maken waarin de werkende CLI game wordt getoond.
 - Code netjes maken en toelichting toevoegen met commentaar.
 - `product/README.md` invullen met uitleg hoe het spel gestart moet worden.

**Week 10**: Afronding & Inleveren
 - Logboek compleet maken.
 - Reflectie en eindconclusie in het verslag schrijven.
 - Project controleren en inleveren.

**Leerdoel 1** (Basis Go): Week 1 t/m 6.  
**Leerdoel 2** (Pointers): Week 8.  
**Leerdoel 3** (Foutafhandeling): Week 8.  
**Leerdoel 4 & 5** (Planning en logboek): Gedurende het gehele project (Week 1 t/m 10 via GitHub Projects en dit document).

### Logboek

Hou een logboek bij van de tijd die je besteed aan het maken van je prototype. Noteer hierin ook de stappen die je hebt gezet om je leerdoelen te behalen.

## Prototype

Leg uit wat je prototype is en hoe je dit hebt gemaakt. Voeg ook een link naar je video van je prototype.

## Eindconclusie

Leerdoelen behaald? Wat heb je geleerd? Wat zou je een volgende keer anders doen? Wat zijn je vervolgstappen?

<div style="page-break-after: always;"></div>

## Bronnen

[^1] Gebaseerd op actuele zoektocht op meerdere vacature websites (September 2026).

[^2] [Google Security Blog: Rust in Android](https://blog.google/security/rust-in-android-move-fast-fix-things/) - Beschrijft hoe Rust geheugenveiligheid garandeert zonder runtime overhead. (Geraadpleegd op 24 September 2026)

[^3] [What does rust have instead of a garbage collector + Benchmark](https://stackoverflow.com/questions/32677420/what-does-rust-have-instead-of-a-garbage-collector) - Beschrijft hoe Rust geen GC heeft. (Geraadpleegd op 24 September 2026)

[^3] [Rust presteert superieur in CPU-intensieve operaties ten opzichte van talen met een Garbage Collector.](https://programming-language-benchmarks.vercel.app/rust) (Geraadpleegd op 24 September 2026)

[^4] [TIOBE index](https://www.tiobe.com/tiobe-index/) (Geraadpleegd op 24 September 2026)

[^5] [Go benchmark](https://programming-language-benchmarks.vercel.app/go) (Geraadpleegd op 24 September 2026)

[^6] [Beam virtual machine](https://medium.com/flatiron-labs/elixir-and-the-beam-how-concurrency-really-works-3cc151cddd61) (Geraadpleegd op 28 September 2026)

[^7] [Elixir does not share processes their data](https://dockyard.com/blog/2021/03/30/elixir-is-safe)(Geraadpleegd op September 2026)

[^8] [Garbage collection (GoLang)](https://dev.to/shrsv/gos-garbage-collector-how-it-keeps-your-code-clean-48nn) (Geraadpleegd op September 28 2026)

[^9][GoLang relevancy](https://www.avidclan.com/blog/Golang-the-Go-To-Programming-Language/) (Geraadpleegd op September 28 2026)

[^10][Rust learning curve](https://rustify.rs/articles/is-rust-hard-to-learn) (Geraadpleegd op September 28 2026)

[^11][Golang Easy to Learn?](https://dev.to/adrianfinantyo/is-golang-easy-to-learn-540l) (Geraadpleegd op September 28 2026)

[^12][Elixir Hard to study because there isn't much out there to learn from?](https://news.ycombinator.com/item?id=40923362) (Geraadpleegd op September 28 2026)

[^13][A tour of GoLang] (https://go.dev/tour/welcome/1) (Geraadpleegd op October 8 2026)