# Reto 3: La zebra de colores

## Problema

TMChein ya se está preparando para las fiestas y quiere empezar a decorar la casa con las luces de navidad.

Quiere comprar una pero sus favoritas son las que tienen dos colores que se van alternando. Como una zebra de dos colores.

Ha hecho que las luces sean Arrays y cada posición un color. Y quiere saber qué luces tienen las zebras más largas y cuál es el último color de esa sucesión de colores. Por ejemplo:

```json
['red', 'blue', 'red', 'blue', 'green'] -> 4, blue
['green', 'red', 'blue', 'gray'] -> 2, gray
['blue', 'blue', 'blue', 'blue'] -> 1, blue
['red', 'green', 'red', 'green', 'red', 'green'] -> 6, green
['blue', 'red', 'blue', 'red', 'gray'] -> 4, red
['red', 'red', 'blue', 'red', 'red', 'red', 'green'] -> 3, red
['red', 'blue', 'red', 'green', 'red', 'green', 'red', 'green'] -> 6, green
```

Fíjate que sólo quiere saber la longitud de cuando dos colores se van alternando. Una vez que se rompe la alternancia de los dos colores, deja de contar.

Ahora que ya sabes esto, <https://codember.dev/colors.txt>

Recuerda que una zebra de colores es cuando dos colores se alternan una y otra vez. Si se repite un color en la posición siguiente o es un tercer color, entonces se deja de contar.
Lo que queremos calcular es la tira de colores más larga en forma de zebra y el último color de esa tira de colores.

## Mi solución

<details>
<summary>Si haces click aquí te harás spoiler de la solución creada por mí. Hazlo bajo tu responsabilidad.</summary>

Este ejercicio es sencillo de entender pero difícil de programar. La idea para resolverlo es tener dos variables que son el último color que hemso visto `lastColor`, y el siguiente color que esperamos `nextColor`. Entonces recorreremos el array de colores buscando dos colores que se alternen.

Si el color que esperamos `nextColor`, es distinto al color actual `currColor`, significa que no se están alternando. Además si el color se repite (`lastColor == currColor`), tampoco se están alternando.

Entonces, ahora sabemos lo que tiene que pasar para que los colores dejen de considerarse "alternos".

En cualquier caso, calculamos para la siguiente iteración cuál es el `lastColor`, obviamente es `currColor`, ¿y el `nextColor`? Efectivamente, `lastColor`. Esto implica que se están "alternando". Esto en Go podemos hacerlo en una sola línea como vemos con `lastColor, nextColor = currColor, lastColor`.

```go
func ReadJson(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var colors []string
	err = json.NewDecoder(f).Decode(&colors)

	return colors, err
}

func main() {
	colors, err := ReadJson("data")
	if err != nil {
		panic(err)
	}

	var maxZebraPoints = 0
	var maxZebraLastColor = ""

	var lastColor = ""
	var nextColor = colors[0]
	var currMaxPoints = 1

	for _, currColor := range colors {
		if currColor != nextColor || lastColor == currColor {
			currMaxPoints = 1
		}

		currMaxPoints++
		lastColor, nextColor = currColor, lastColor

		if currMaxPoints > maxZebraPoints {
			maxZebraPoints = currMaxPoints
			maxZebraLastColor = lastColor
		}
	}

	fmt.Printf("submit %d@%s\n", maxZebraPoints, maxZebraLastColor)
}
```

</details>
