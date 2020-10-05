## ROADMAP
New ideas, thought about needed features will be store in this file.

### Done
* Initial core
    * Import file of specific format **Done as CSV**
    * Export file of specific format **Done as CSV**
    * Turn Restricted Shortest Path extension for CH-algorithm **Propably not modify algorithm, but graph. UPD: done with osm2ch - https://github.com/LdDl/osm2ch#osm2ch**
    * Thoughts and discussions about OSM graph and extensions **Need some ideas about parsing and preparing. UPD: osm2ch - https://github.com/LdDl/osm2ch#osm2ch**
    * Map matcher as another project **WIP (now it is in local git storage)** Here is the one https://github.com/LdDl/horizon
    * Bring OSM parser **WIP It's done in poor way. PRs are welcome**
    * Bring OSM restrictions **WIP It's done in poor way. PRs are welcome**
    * OneTwoMany function (contraction hierarchies) **Done, may be some bench comparisons**
    * Replace int with int64 (OSM purposes) **Done**
    * Separate benchmarks to BENCHMARK.md **Done**

### Planned
* Better CSV format or another format (JSON / binary)