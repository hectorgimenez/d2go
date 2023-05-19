<h1 align="center">d2go</h1>

---

Tooling for Diablo II: Resurrected written in Go. It provides a library to read game memory and a public SDK to be
imported, along with data structures and some other tools.

### Libraries

- [data](https://github.com/hectorgimenez/d2go/tree/main/pkg/data) - D2R Game data structures
- [memory](https://github.com/hectorgimenez/d2go/tree/main/pkg/memory) - D2R memory reader (it provides the data
  structures)
- [nip](https://github.com/hectorgimenez/d2go/tree/main/pkg/nip) - A very basic NIP file parser
- [itemfilter](https://github.com/hectorgimenez/d2go/tree/main/pkg/itemfilter) - Based on game data, it provides an item
  pickup filtering

### Tools

- [cmd/itemwatcher](https://github.com/hectorgimenez/d2go/tree/main/cmd/itemwatcher) - Small tool that plays a sound
  when an item passing the filtering process is dropped
