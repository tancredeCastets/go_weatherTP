# Weather TP

## Partie A — Comparaison JSON / XML

| Donnée            | JSON                                                           | XML                                                |
|-------------------|----------------------------------------------------------------|----------------------------------------------------|
| Pays              | "Stations"["country"]                                          | attribut country   balise <station>                |
| Coordonnées       | "Stations"["location"{"latitude" , "longitude"}]               | attributs lat, lon   balise <station><coordinates> |
| Altitude          | "Stations"["altitude_m"]                                       | attribut altitude     balise <station><station>             |
| Modèle de capteur | "Stations"["device"{"type"}]                                   | attribut model       balise <station><hardware>             |
| Température       | "Stations"["observations"{"temperature_celsius"}]              | measure type="temperature" balise <station><observation>    |
| Conditions ciel   | "Stations"["observations"{"conditions"}]                       | measure type="sky"       balise <station><observation>      |
| Vent              | "Stations"["observations"{wind("speed_kmh", "direction_deg")}] | speed, "direction"     balise <station><observation<wind>>  |
