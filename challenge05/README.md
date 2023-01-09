# Reto 5: Battle Royale de frameworks y bibliotecas

## Problema

Hay tanto framework y biblioteca que ya no sabemos qué usar. Así que un comité ha decidido hacer una especie de Los Juegos del Hambre para decidir qué tecnología se queda.

Ha puesto todas las tecnologías en círculo de forma aleatoria. La tecnología en el índice 0 empieza matando a la que tiene justo a la derecha (índice + 1).

El siguiente turno es para la tecnología que esté viva que queda a la derecha de la que se acaba de morir. Y así sucesivamente hasta que sólo quede una. Mira este ejemplo de un grupo de 10 tecnologías, paso a paso:

### Primer paso

```js
         5
      6     4
   7           3
   8           2
      9     1
         0
```

- 0 mata a 1
- 2 mata a 3
- 4 mata a 5
- 6 mata a 7
- 8 mata a 9

### Segundo paso

```js
         X
     6      4
   X           X
   8           2
      X     X
         0
```

- 0 mata a 2
- 4 mata a 6
- 8 mata a 0

### Tercer paso

```js

         X
     X      4
   X           X
   8           X
      X     X
         X
```

- 4 mata a 8

### Resultado final

```js
         X
     X      4
   X           X
   X           X
      X     X
         X
```

La tecnología en el índice 4 es la que ha sobrevivido.

Ahora, para probar que somos capaces de crear un algoritmo que funcione, tenemos la lista de mecenas de la comunidad de midudev: <https://codember.dev/mecenas.json> o en el fichero `data` de esta misma carpeta.

Tienes que crear un algoritmo que nos diga qué usuario sobreviviría usando el mismo sistema.

## Mi solución

<details>
<summary>Si haces click aquí te harás spoiler de la solución creada por mí. Hazlo bajo tu responsabilidad.</summary>

Generalmente resuelvo los problemas de la manera más obvia, y no me paro a optimizarlos, a no ser que sea necesario. A diferencia de los anteriores challenge, en este he querido hacerlo de una manera más eficiente. Ya que la manera obvia, que sería almacenando los índices de los mecenas que sobreviven en un array, y luego eliminarlos del array era confusa con tanto bucle.

He tomado prestada la función `filter` de JS y la he llamado `ArrayFilter`. Para el que no la conozca, esta función recibe un array y una función que recibe el índice y el elemento del array. La función lambda devuelve un booleano que indica si el elemento debe permanecer en el array o no. La función `ArrayFilter` devuelve un nuevo array con los elementos que no han sido eliminados.

Y después de eso he examinado el problema, de tal forma que he encontrado dos casos:

1. Si el número de mecenas es par, mueren todos los mecenas pares.
2. Si el número de mecenas es impar, mueren todos los mecenas pares, y después, el primer mecenas.

Es muy importante el orden en el segundo caso, ya que si se ejecuta al revés, primero moriría el primero (0), y después los impares.

```go
func ReadJson(path string) ([]string, error) {
 f, err := os.Open(path)
 if err != nil {
  return nil, err
 }

 var patrons []string
 err = json.NewDecoder(f).Decode(&patrons)

 return patrons, err
}

func ArrayFilter(arr []int, fn func(i int, el int) bool) []int {
 var n = 0
 for i, el := range arr {
  if fn(i, el) {
   arr[n] = el
   n++
  }
 }

 return arr[:n]
}

func ArrayIndexes(arr []string) []int {
 arrIndexes := make([]int, len(arr))

 for i := range arr {
  arrIndexes[i] = i
 }

 return arrIndexes
}

func main() {
 patrons, err := ReadJson("data")
 if err != nil {
  panic(err)
 }
 patronsIndexes := ArrayIndexes(patrons)

 for len(patronsIndexes) > 1 {
  isEven := len(patronsIndexes)%2 == 0

  patronsIndexes = ArrayFilter(patronsIndexes, func(i int, el int) bool {
   return i%2 == 0
  })

  if !isEven {
   patronsIndexes = patronsIndexes[1:]
  }
 }

 fmt.Printf("submit %s-%d\n", patrons[patronsIndexes[0]], patronsIndexes[0])
}
```

</details>
