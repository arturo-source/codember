# Reto 1: ¡Arregla Twitter!

> Esta plantilla ha sido creada por `wayaba`, puedes verla en el repositorio <https://github.com/wayaba/codember-2022>

## Problema

Twitter ha sido comprado y quieren eliminar los bots. Te han pedido ayuda para detectar el número de usuarios en su base de datos que tienen datos corruptos.

La base de datos es muy antigua y está en un formato extraño. Los perfiles requieren tener los siguientes datos:

```json
usr: nombre de usuario
eme: email
psw: contraseña
age: edad
loc: ubicación
fll: número de seguidores
```

Todo está en un fichero donde los datos de usuario son una secuencia de pares `key:value`, que pueden estar en la misma línea o separado por líneas, y cada usuario está separado por un salto de línea. ¡Ojo porque puede estar todo desordenado!

## Ejemplo de input

```json
usr:@midudev eme:mi@gmail.com psw:123456 age:22 loc:bcn fll:82

fll:111 eme:yrfa@gmail.com usr:@codember psw:123456 age:21 loc:World

psw:11133 loc:Canary fll:333 usr:@pheralb eme:pheralb@gmail.com

usr:@itziar age:19 loc:isle psw:aaa fll:222 eme:itzi@gmail.com
```

El primer usuario SÍ es válido. Tiene todos los campos.
El segundo usuario SÍ es válido. Tiene todos los campos.
El tercer usuario NO es válido. Le falta el campo `age`.
El cuarto usuario SÍ es válido. Tiene todos los campos..

## Cómo enviar la solución

Usa el comando "submit" para enviar tu solución con el número de usuarios correctos + el nombre del último usuario válido. Por ejemplo:

```sh
submit 482@midudev
```

## Mi solución

<details>
<summary>Si haces click aquí te harás spoiler de la solución creada por mí. Hazlo bajo tu responsabilidad.</summary>

La idea para resolver este problema es tener un array con los campos necesasrios (en este caso `DATA_TO_BE_PRESENT`), y obtener de cada una de los usuarios, las claves.

El problema tiene dos temas que resolver. El primero es el de obtener un usuario, que puede estar en una línea o en varias. Esto lo solucionamos leyendo el fichero como un array de lineas. Recorremos este array y cada vez que encuentramos una linea en blanco, añadimos ese usuario al array de usuarios.

La segunda parte trata de obtener las claves, como todas las clave-valor están separadas por un espacio `" "`, es sencillo, se divide el usuario por espacios. Esta operación devuelve un array que recorreremos también, y dividiremos también, en este caso por dos puntos  `":"`. Ahora tendremos en la primera posición del array resultante la clave.

Guardando el array de claves del usuario ya tendríamos el ejercicio casi acabado. Recorremos el array inicialmente mencionado `DATA_TO_BE_PRESENT`. Y si alguno de los valores de eset array no está presente en el usuario, el usuario es inválido.

Esto lo hacemos con todos los usuarios y ya tendremos el resultado final.

```go
var DATA_TO_BE_PRESENT = []string{"usr", "eme", "psw", "age", "loc", "fll"}

func ReadAllLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func LinesToData(lines []string) []string {
	var data []string
	var lineBuffer string
	for _, line := range lines {
		if line == "" {
			data = append(data, lineBuffer)
			lineBuffer = ""
			continue
		}
		lineBuffer += line + " "
	}
	if lineBuffer != "" {
		data = append(data, lineBuffer)
	}

	return data
}

func Contains(slice []string, element string) bool {
	for _, a := range slice {
		if a == element {
			return true
		}
	}
	return false
}

func DataIsValid(data string) bool {
	keys := []string{}
	keyValuePairs := strings.Split(data, " ")

	for _, pair := range keyValuePairs {
		keyValue := strings.Split(pair, ":")
		keys = append(keys, keyValue[0])
	}

	for _, key := range DATA_TO_BE_PRESENT {
		if !Contains(keys, key) {
			return false
		}
	}

	return true
}

func GetUsername(data string) string {
	keyValuePairs := strings.Split(data, " ")
	for _, pair := range keyValuePairs {
		keyValue := strings.Split(pair, ":")
		if keyValue[0] == "usr" {
			return keyValue[1]
		}
	}
	return ""
}

func main() {
	lines, err := ReadAllLines("data")
	if err != nil {
		panic(err)
	}
	data := LinesToData(lines)

	invalidUsers := 0
	lastValidUsername := ""
	for _, dataLine := range data {
		if DataIsValid(dataLine) {
			invalidUsers++
			lastValidUsername = GetUsername(dataLine)
		}
	}

	fmt.Print("submit ", invalidUsers, lastValidUsername)
}
```

</details>
