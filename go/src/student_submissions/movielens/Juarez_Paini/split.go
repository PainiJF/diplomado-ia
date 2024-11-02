package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func SplitBigFile(file_name string, number_of_chunks int, directory string) []string {

	file, err := os.Open(directory + file_name + ".csv") //
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("Error extracting data from file %v: %s", file_name, err)
	}

	rowsPerFile := len(data) / number_of_chunks

	var filesCreated []string

	for i := 0; i < number_of_chunks; i++ {
		var tempName string
		tempName = file_name + "_" + fmt.Sprintf("%02d", i+1)
		WriteCsv(data[i*rowsPerFile:(i+1)*rowsPerFile], tempName, directory) // Escribe el archivo
		filesCreated = append(filesCreated, tempName)                        // Agrega el nombre del archivo a l
	}

	return filesCreated
}

func WriteCsv(data [][]string, name string, path string) {
	csvFile, err := os.Create(path + "/" + name + ".csv")
	if err != nil {
		log.Fatalf("Error creating new csv file %v: %s", name, err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	err = writer.WriteAll(data)
	if err != nil {
		log.Fatalf("Error writing new csv file %v: %s", name, err)
	}

	fmt.Printf("Archivo %s ha sido creado con %v filas\n", name, len(data))
}

// En obtaining Data se leeran los archivos ratings0i.cvs y movies.cvs para obtener los datos necesarios
// posteriormente esta función se trabajará con 10 workers simultaneos
func Obtaining_Data(file_name string) {

	ratings_file, err := os.Open("/Users/paini/Desktop/DiplomadoAI24-25/ml-25m/" + file_name + ".csv") //
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer ratings_file.Close()

	Reader_ratings := csv.NewReader(ratings_file)
	data_ratings, err := Reader_ratings.ReadAll()
	if err != nil {
		log.Fatalf("Error extracting data from file %v: %s", ratings_file, err)
	}

	movies_file, err := os.Open("/Users/paini/Desktop/DiplomadoAI24-25/ml-25m/movies.csv") //
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer movies_file.Close()

	Reader_movies := csv.NewReader(movies_file)
	data_movies, err := Reader_movies.ReadAll()
	if err != nil {
		log.Fatalf("Error extracting data from file %v: %s", movies_file, err)
	}

	// En estas lineas definimos todos los géneros y las listas ra
	kg := []string{"Action", "Adventure", "Animation", "Children", "Comedy", "Crime", "Documentary",
		"Drama", "Fantasy", "Film-Noir", "Horror", "IMAX", "Musical", "Mystery", "Romance",
		"Sci-Fi", "Thriller", "War", "Western", "(no genres listed)"}
	number_genres := len(kg) // number of known genres
	ra := make([][]float64, number_genres)
	ca := make([][]int, number_genres)
	for i := 0; i < number_genres; i++ {
		ra[i] = make([]float64, 1)
		ca[i] = make([]int, 1)
	}

	if strings.Contains(data_movies[1][2], kg[19]) {
		println("SI")
	}

	for i := 0; i < len(data_ratings)-2500005; i++ {
		for j := 0; j < len(data_movies); j++ {
			if strings.Contains(data_ratings[i][1], data_movies[j][0]) {
				fmt.Println(data_ratings[i][1], data_movies[j][0]) // prueba

			}

		}
		// Accede a cada elemento usando el índice i
		//movie_ID := data_ratings[i][1]         // ID de la película
		//rating_for_movie := data_ratings[i][2] // Rating de la película

		//fmt.Println(movie_ID, rating_for_movie)
	}

	println(len(data_ratings))

	println(number_genres)
	fmt.Println(ra)
	fmt.Println(ca)
	print("--------")
	println(data_movies[1][2])
	println(data_ratings[0][2])

}

func main() {
	start := time.Now()
	//SplitBigFile("ratings", 10, "/Users/paini/Desktop/DiplomadoAI24-25/ml-25m/")
	elapsed := time.Since(start) // Calcula el tiempo transcurrido
	fmt.Printf("Tiempo total transcurrido: %s\n", elapsed)
	Obtaining_Data("ratings_01")

}
