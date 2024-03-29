package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"
)

const monitoringTimes = 5
const delay = 5 * time.Second
const sitesFileName = "sites.txt"
const logFileName = "log.txt"


func main() {
	// hello()
	// variables1()
	// variables2()
	// testMultipleReturn()
	// whileTrue()
	// arrayVsSlice()
	// readAndPrintFile("sites.txt")
	drawMenu()
}

func hello() {
	fmt.Println("Hello, world!")
}

func variables1() {
	var name string = "John" // explicitly typed
	var lastname = "Doe" // inferred type
	fullname := name + " " + lastname // variable declaration var keyword + strings concatenation
	fullname = name + " " + lastname // re-assigned without declaration
	fmt.Println(name, lastname)
	fmt.Println(fullname)
}

func variables2() {
	var n1 float32 = 0.1
	var n2 float32 = 0.2
	n3 := 0.3
	result := n1 + n2
	fmt.Println(result)
	fmt.Println("O tipo da variavel n3 é", reflect.TypeOf(n3)) // float64
}

func drawMenu() {
	version := 1.0
	fmt.Println("Este programa está na versão", version)

	for {
		fmt.Println("1 - Iniciar monitoramento")
		fmt.Println("2 - Exibir logs")
		fmt.Println("0 - Sair do programa")
	
		var option int
		fmt.Scanf("%d", &option) // ponteiro para variavel
		// fmt.Scan(&option) // Forma curta sem especificar formato
	
		fmt.Println("A opção escolhida foi a", option)
	
		logOption(option)
		initOption(option)
	}
}

func input(message string) string {
	var input string
	fmt.Println(message)
	fmt.Scanf("%s", &input)
	return input
}

func logOption(option int) {
	if option == 1 {
		fmt.Println("Iniciando monitoramento...")
	} else if option == 2 {
		fmt.Println("Exibindo logs...")
	}	else if option == 0 {
		fmt.Println("Saindo...")
	}	else {
		fmt.Println("Opção inválida")
	}
}

func initOption(option int) {
	// Na linguagem go o break não é nessesario no switch case
	// Mas se tiver break o compilador não reclama
	switch option {
		case 0:	exitProgram()
		case 1:	startMonitoring()
		case 2:	showLogs()
		default: println("Opção inválida")
	}
}

func startMonitoring() {
	sites := getSitesFromFile(sitesFileName)
	for i := 0; i < monitoringTimes; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			checkSite(site)
		}
		fmt.Println("")
		time.Sleep(delay)
	}
}

func checkSite(site string) {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro ao acessar o site", site)
	} else if response.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso!")
		registerLog(site, true)
	} else {
		fmt.Println("Site", site, "está com problemas. Status Code:", response.StatusCode)
		registerLog(site, false)
	}
}

func showLogs() {
	readAndPrintFile(logFileName)
}

func exitProgram() {
	os.Exit(0)
}

func multipleReturn() (int, string) {
	// No GO uma função pode retornar mais de uma coisa
	return 1, "Hello"
}

func testMultipleReturn() {
	code, message := multipleReturn()
	fmt.Println(code, message)
}

func whileTrue() {
	// No go não existe a palavra chave while mas podemos usar o for
	counter := 0
	for {
		fmt.Println("Counter:", counter)
		counter++
		if counter == 10 {
			break
		}
	}
}

func getSites() []string {
	sites := []string {
		"https://www.google.com",
		"https://www.youtube.com",
		"https://random-status-code.herokuapp.com/",
	}
	return sites
}

func arrayVsSlice() {
	// No go arrays tem tamanho fixo
	// Slices são como se fossem arrays mas podem ter tamanho variável
	slice := []int { 1, 2, 3 }
	array := [3]int { 1, 2, 3 }

	slice = append(slice, 4)

	fmt.Println(reflect.TypeOf(slice), slice)
	fmt.Println(reflect.TypeOf(array), array)

	fmt.Println("Tamanho do slice:", len(slice))
	fmt.Println("Tamanho do array:", len(array))

	fmt.Println("Capacidade do slice:", cap(slice), "(Dobro do tamanho original)")
	fmt.Println("Capacidade do array:", cap(array))

	sliceFromArray := array[0:2]

	fmt.Println(sliceFromArray)
}

func getSitesFromFile(fileName string) []string {
	sites := []string {}

	file, err := os.Open(fileName)
	
	if err != nil {
		fmt.Println("Ocorreu um erro ao abrir o arquivo", fileName)
		return sites
	}
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sites = append(sites, scanner.Text())
	}

	file.Close()

	return sites
}

func readAndPrintFile(fileName string) {
	fileBytesArray, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Ocorreu um erro ao abrir o arquivo", fileName)
		return
	}
	fileContent := string(fileBytesArray)

	fmt.Println(fileContent)
}

func readFileLinesWithReader(fileName string) []string {
	lines := []string {}
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Ocorreu um erro ao abrir o arquivo", fileName)
		file.Close()
		return lines
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		lines = append(lines, line)
		if err == io.EOF {
			file.Close()
			break
		}
	}
	return lines
}

// os.Open e os.OpenFile são funções que abrem arquivos, a segunda permite criar o arquivo caso não exista
func registerLog(site string, online bool) {
	const readWriteChMod os.FileMode = 0666
	const readCreateAppendFlag = os.O_RDWR | os.O_CREATE | os.O_APPEND

	file, err := os.OpenFile(logFileName, readCreateAppendFlag, readWriteChMod)

	if err != nil {
		fmt.Println("Ocorreu um erro ao abrir o arquivo", logFileName)
		return
	}

	status := ternaryString(online, "online", "offline")
	timeFormat := "02/01/2006 15:04:05"
	// O pacote time possui funções para manipular datas e horas
	// A função Now() retorna a data e hora atual e permite formata-la com o método .Format() que recebe um formato em string específico do Go
	date := time.Now().Format(timeFormat)
	logLine := date + " - " + status + " - " + site + "\n"

	file.WriteString(logLine)
	file.Close()
}

func ternaryString(condition bool, trueResult string, falseResult string) string {
	if condition {
		return trueResult
	} else {
		return falseResult
	}
}
