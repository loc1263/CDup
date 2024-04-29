package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Uso: Cdup <nombre_del_archivo> <inicio_campo> <final_campo>")
		return
	}

	nombreArchivo, inicioCampo, finalCampo := os.Args[1], 0, 0

	fmt.Sscanf(os.Args[2], "%d", &inicioCampo)
	fmt.Sscanf(os.Args[3], "%d", &finalCampo)

	file, err := os.Open(nombreArchivo)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	duplicados := make(map[string][]int)
	startTime := time.Now()

	scanner := bufio.NewScanner(file)
	lineaNumero := 1

	for scanner.Scan() {
		linea := strings.TrimSpace(scanner.Text())
		if len(linea) >= finalCampo && inicioCampo < finalCampo {
			campo := linea[inicioCampo:finalCampo]
			if indices, ok := duplicados[campo]; ok {
				fmt.Printf("Valor duplicado: %s (líneas %v y %v)\n", campo, indices, lineaNumero)
				indices = append(indices, lineaNumero)
				duplicados[campo] = indices
			} else {
				duplicados[campo] = []int{lineaNumero}
			}
		}
		lineaNumero++
	}

	elapsedTime := time.Since(startTime).Round(10 * time.Millisecond)
	fmt.Println("Tiempo de ejecución:", elapsedTime)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}
}
