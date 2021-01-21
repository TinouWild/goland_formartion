# Golang Formation

## Généralités

go run main.go == go build + ./executable

### GOPATH

```
/go
  |__ /bin
  |__ /pkg
  |__ /src
```
Contient 3 dossiers. Le /bin correspond aux binaires générés via le go install. Le /pkg correspond aux paquets nécessaires à l'execution du code situé dans /src. Le /src contient tout le code GO qui peut être hiérarchisé en projets, sous-projets, modules, etc.

## Variables

- Possible de déclarer une variable sans valeur d'initialisation
- Pour une déclaration en dehors d'une fonction, utilisation obligatoire de la syntaxe var toto = X (et non toto := X)
- Accessibilité d'une variable en fonction de la casse de la première lettre

## Fonctions
```
func sumNamedReturn(x, y, z int) int { }
func sumNamedReturn(x, y, z int) (sum int) { }
```
Dans le deuxième cas, sum est déclaré au moment de l'execution et on peut ensuite l'utiliser pour lui affecter une valeur dans la méthode

