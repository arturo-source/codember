# Reto 2: ¡Atrapa a esos ciber criminales!

## Problema

Un grupo de ciber criminales están usando mensajes encriptados para comunicarse. El FBI nos ha pedido ayuda para descifrarlos.

Los mensajes son cadenas de texto que incluyen números enteros muy largos y espacios en blanco. Aunque los números no parecen tener sentido... una chica llamada Alice ha descubierto que podrían usar el código ASCII de las letras en minúscula.

Con su método ha conseguido descifrar estos mensajes:

```txt
"109105100117" -> midu
"9911110010110998101114" -> codember
"9911110010110998101114 109105100117" -> codember midu
"11210897121 116101116114105115" -> play tetris
```

Pero han interceptado un mensaje más largo que no han podido y nos han dicho que es muy importante que lo descifremos:

```txt
11610497110107115 102111114 11210897121105110103 9911110010110998101114 11210810197115101 11510497114101
```

Ahora que ya sabes esto, <https://codember.dev/encrypted.txt>

## Pistas

Recuerda que los mensajes son cadenas de texto conformadas por números y espacios en blanco.
Parece que los números tienen algo que ver con el código ASCII.
Los espacios en blanco parece que son simplemente espacios...
Cómo enviar la solución
Usa el comando "submit" para enviar tu solución con la frase descifrada, en minúsculas y respetando los espacios en blanco. Por ejemplo:

```sh
submit this is fine
```

## Mi solución

<details>
<summary>Si haces click aquí te harás spoiler de la solución creada por mí. Hazlo bajo tu responsabilidad.</summary>

```go
func ReadFile(path string) (string, error) {
	dat, err := os.ReadFile(path)
	return string(dat), err
}

func IsValidChar(char rune) bool {
	return char >= 'a' && char <= 'z'
}

func main() {
	content, err := ReadFile("data")
	if err != nil {
		panic(err)
	}

	var asciiCode string
	var message string
	for _, char := range content {
		if char == ' ' {
			asciiInt, _ := strconv.Atoi(asciiCode)
			asciiCode = ""

			message += string(rune(asciiInt))
			message += " "

			continue
		}

		asciiCode += string(char)
		asciiInt, _ := strconv.Atoi(asciiCode)
		if IsValidChar(rune(asciiInt)) {
			message += string(rune(asciiInt))
			asciiCode = ""
		}
	}

	fmt.Println("submit", message)
}
```

<details>
