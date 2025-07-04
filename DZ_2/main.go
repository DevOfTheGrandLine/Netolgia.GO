package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func ReadProcessWrite(inputPath, outputPath string, processFunc func(string) (string, error)) error {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("не удалось открыть входной файл %q: %w", inputPath, err)
	}
	defer inputFile.Close()

	inputData, err := io.ReadAll(inputFile)
	if err != nil {
		return fmt.Errorf("не удалось прочитать файл %q: %w", inputPath, err)
	}

	processedData, err := processFunc(string(inputData))
	if err != nil {
		return fmt.Errorf("ошибка при обработке данных: %w", err)
	}

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("не удалось создать файл %q: %w", outputPath, err)
	}
	defer outputFile.Close()

	_, err = outputFile.WriteString(processedData)
	if err != nil {
		return fmt.Errorf("не удалось записать данные в файл %q: %w", outputPath, err)
	}

	err = outputFile.Sync()
	if err != nil {
		return fmt.Errorf("ошибка синхронизации файла %q: %w", outputPath, err)
	}

	return nil
}

func main() {
	err := ReadProcessWrite(
		"input.txt",
		"output.txt",
		func(text string) (string, error) {
			return strings.ToUpper(text), nil
		},
	)

	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	} else {
		fmt.Println("Файл успешно обработан.")
	}
}
