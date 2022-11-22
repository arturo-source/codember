# Reto 4: Encuentra la contraseña de tu amigo

## Problema

Un amigo compró 5 BitCoins en 2008. El problema es que lo tenía en un monedero digital... ¡y no se acuerda de la contraseña!

Nos ha pedido ayuda. Y nos ha dado algunas pistas:

- Es una contraseña de 5 dígitos.
- La contraseña tenía el número 5 repetido dos veces.
- El número a la derecha siempre es mayor o igual que el que tiene a la izquierda.

Nos ha puesto algunas ejemplos:

```txt
55678 es correcto lo cumple todo
12555 es correcto, lo cumple todo
55555 es correcto, lo cumple todo
12345 es incorrecto, no tiene el 5 repetido.
57775 es incorrecto, los números no van de forma creciente
```

Dice que el password está entre los números 11098 y 98123. ¿Le podemos decir cuantos números cumplen esas reglas dentro de ese rango?

## Mi solución

<details>
<summary>Si haces click aquí te harás spoiler de la solución creada por mí. Hazlo bajo tu responsabilidad.</summary>

La dificultad de este ejercicio reside en cómo dividir un número entero en dígitos. Yo lo hago con la función `SplitNumberByDigit` que me los devuelve en un array de enteros, aunque están al revés, pero sabiéndolo no es ningún problema.

Para saber si un número es válido hago lo que dice el enunciado. Verifico que los dígitos estén en orden creciente (o decreciente, porque ya he dicho que la función devuelve los dígitos al revés) y que el número 5 esté repetido dos veces.

```go
func SplitNumberByDigit(number int) []int {
	var digits []int
	for number > 0 {
		digits = append(digits, number%10)
		number /= 10
	}
	return digits
}

func IsValidNumber(number int) bool {
	digits := SplitNumberByDigit(number)
	for i := 0; i < len(digits)-1; i++ {
		if digits[i] < digits[i+1] {
			return false
		}
	}

	numOfFives := 0
	for _, digit := range digits {
		if digit == 5 {
			numOfFives++
		}
	}

	return numOfFives >= 2
}

func main() {
	var validNumber []int
	const start = 11098
	const end = 98123

	for i := start; i <= end; i++ {
		if IsValidNumber(i) {
			validNumber = append(validNumber, i)
		}
	}

	fmt.Printf("submit %d-%d\n", len(validNumber), validNumber[55])
}
```

</details>
