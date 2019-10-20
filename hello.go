/*
Para testar executar o seguinte comando no CMD:
(go run localdoarquivo\arquivo.go)
> go run hello.go
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

/*
  Função inicializadora do sistema em golang.
*/
func main() {
	//testarPrint()
	//testarCalculo()
	//testarPrintf()
	//testarInput()
	//testarArray()
	//testarOperacoes()
	//testarSwitch()
	//testarFor()
	//fmt.Println("Retorno:", testarReturn())
	//testarArgumento("Felipe", 20)
	//testarArrayMultidimensional()
	//testarDicionario()
	//testarEstrutura()
	//testarInterface()
	//testarFuncaoMatematica()
	//testarDataHora()
	//testarFicheiroEscreve()
	//testarFicheiroLe()
	//testarOperadoresBitABit()
	//testarObterTipoVariavel()
	////-------------------------------//
	//fmt.Println(testarDoisRetornos())
	//// ou
	//nome, idade := testarDoisRetornos()
	//fmt.Println("Nome:", nome, "Idade:", idade)
	////-------------------------------//
	//testarSliceAppend()
	testarMenuMultiplaEscolha()
}

/***************************************************************************************************/
func testarPrint() {
	fmt.Println("Olá \"mundo\"!")
	fmt.Println("Felipe", "Sulzbach")
	fmt.Println(10 + 10)
	fmt.Println(10.5 - 5.5 - 5)
	fmt.Println(-10 * -2)
	fmt.Println(10 / 10)
	fmt.Println(9 % 2)

	var descricao string = "Declarando o tipo."
	descricao2 := "Sem declaração."
	fmt.Println(descricao)
	fmt.Println(descricao2)
}

/***************************************************************************************************/
func testarCalculo() {
	x := 2
	x += 3
	fmt.Println("Soma de x:", x)
	x -= 2
	fmt.Println("Subtração de x:", x)
	x *= 2
	fmt.Println("Multiplicação de x:", x)
	x /= 2
	fmt.Println("Divisão de x:", x)
	x %= 3
	fmt.Println("Resto da divisão de x:", x)
}

/***************************************************************************************************/
func testarPrintf() {
	sobrenome, nome := "Sulzbach", "Felipe"
	idade := 20

	qtDigitosNome := len(nome)
	qtDigitosSobrenome := len(sobrenome)

	fmt.Println(nome, sobrenome, idade, qtDigitosNome, qtDigitosSobrenome)
	fmt.Printf("Nome: %s %s\nIdade: %d", strings.ToUpper(nome), strings.ToLower(sobrenome), idade)
}

/***************************************************************************************************/
func testarInput() {
	inputDoUtilizador := bufio.NewReader(os.Stdin)
	fmt.Println("Insira seu nome: ")
	nome_1, _ := inputDoUtilizador.ReadString('\n')

	fmt.Println("Insira seu sobrenome: ")
	sobrenome_1, _ := inputDoUtilizador.ReadString('\n')

	fmt.Println("Insira sua idade: ")
	idade_1, _ := inputDoUtilizador.ReadString('\n')

	fmt.Printf("Nome: %sSobrenome: %sIdade: %s", nome_1, sobrenome_1, idade_1)
}

/***************************************************************************************************/
func testarArray() {
	var cores1 [5]string

	cores1[0] = "Azul"
	cores1[1] = "Verde"
	cores1[2] = "Amarelo"
	cores1[3] = "Vermelho"
	cores1[4] = "Laranja"

	cores2 := [...]string{"Azul", "Verde", "Amarelo", "Vermelho", "Laranja"}

	fmt.Println("Número de cores:", len(cores1), "\nNúmero de cores:", len(cores2))

	fmt.Println("Primeira cor:", cores1[0], "\nPrimeira cor:", cores2[0])
	fmt.Println("Última cor:", cores1[len(cores1)-1], "\nÚltima cor:", cores2[len(cores2)-1])
}

/***************************************************************************************************/
func testarOperacoes() {
	x := 30
	y := 20

	if x == 10 {
		fmt.Println("O valor de x é 10.")
	} else if x == 20 {
		fmt.Println("O valor de x é 20.")
	} else if x != 10 && x != 20 {
		fmt.Println("O valor de x é diferente de 10 e de 20.")
	} else if x == y {
		fmt.Println("O valor de x é igual ao valor de y.")
	} else if x > y {
		fmt.Println("O valor de x é maior que o valor de y.")
	} else if x < y {
		fmt.Println("O valor de x é menor que o valor de y.")
	} else if x > y || x == y {
		fmt.Println("O valor de x é maior ou igual ao valor de y.")
	} else if x == 10 && x < y {
		fmt.Println("O valor de x é igual a 10 e menor que o valor de y.")
	}
}

/***************************************************************************************************/
func testarSwitch() {
	caso := "D"

	switch caso {
	case "A":
		fmt.Println("O caso A existe.")
	case "B":
		fmt.Println("O caso B existe.")
	case "C":
		fmt.Println("O caso C existe.")
	default:
		fmt.Println("O caso", caso, "não existe.")
	}
}

/***************************************************************************************************/
func testarFor() {
	lista := [...]string{"registro00", "registro01", "registro02", "registro03", "registro04", "registro05", "registro06", "registro07", "registro08", "registro09", "registro10"}

	for index := 0; index < len(lista); index++ {
		fmt.Printf("lista[%d]: %s\n", index, lista[index])
	}

	fmt.Printf("\n\n")

	for index, registro := range lista {
		fmt.Printf("lista[%d]: %s\n", index, registro)
	}

	fmt.Printf("\n\n")

	contador := 1
	for contador <= 10 {
		fmt.Printf("Contador: %d\n", contador)
		contador++
	}

	fmt.Printf("\n\n")

	contador = 1
	for contador <= 10 {
		fmt.Printf("Contador: %d\n", contador)

		if contador == 5 {
			break
		}
		contador++
	}

	fmt.Printf("\n\n")

	contador = 0
	for contador < 10 {
		contador++

		if contador == 5 {
			continue
		}

		fmt.Printf("Contador: %d\n", contador)
	}
}

/***************************************************************************************************/
func testarReturn() int {
	return 30
}

/***************************************************************************************************/
func testarArgumento(nome string, idade int) {
	fmt.Printf("Nome: %s\nIdade: %d", nome, idade)
}

/***************************************************************************************************/
var arrayMultidimensional = [5][4]int{
	{1, 2, 3, 4},
	{1, 1, 1, 1},
	{2, 2, 2, 2},
	{3, 3, 3, 3},
	{4, 4, 4, 4},
}

func testarArrayMultidimensional() {
	for linha := 0; linha < 5; linha++ {
		for coluna := 0; coluna < 4; coluna++ {
			fmt.Print(arrayMultidimensional[linha][coluna], "\t")
		}

		fmt.Println()
	}
}

/***************************************************************************************************/
var pessoas = map[string]int{
	"Nome A": 22,
	"Nome B": 37,
}

func testarDicionario() {
	fmt.Println("Número de pessoas:", len(pessoas))

	// inserir novos
	pessoas["Nome C"] = 52
	pessoas["Nome D"] = 68

	fmt.Println("Número de pessoas:", len(pessoas))

	// alterar existente
	pessoas["Nome C"] = 100

	// remover um
	delete(pessoas, "Nome B")

	fmt.Println("Número de pessoas:", len(pessoas))

	for chave, valor := range pessoas {
		fmt.Println("Nome:", chave)
		fmt.Println("idade:", valor)
	}

	// limpando o map
	pessoas = make(map[string]int)

	fmt.Println("Número de pessoas:", len(pessoas))
}

/***************************************************************************************************/
type Estrutura struct {
	numero int
}

func (estrutura Estrutura) imprimirNumero() {
	if estrutura.numero >= 0 && estrutura.numero <= 100 {
		fmt.Println("Número:", estrutura.numero)
	} else {
		fmt.Println("O número precisa ser >= 0 e <= a 100.")
	}
}
func testarEstrutura() {
	estruturaA := Estrutura{100}
	estruturaB := Estrutura{}

	estruturaB.numero = 200

	estruturaA.imprimirNumero()
	estruturaB.imprimirNumero()
}

/***************************************************************************************************/
type FiguraGeometrica interface {
	area() float64
	perimetro() float64
}
type Retangulo struct {
	largura, altura float64
}
type Circulo struct {
	raio float64
}

func (retangulo Retangulo) area() float64 {
	return retangulo.largura * retangulo.altura
}
func (retangulo Retangulo) perimetro() float64 {
	return (retangulo.largura * 2) + (retangulo.altura * 2)
}
func (circulo Circulo) area() float64 {
	return math.Pi * circulo.raio * circulo.raio
}
func (circulo Circulo) perimetro() float64 {
	return math.Pi * 2 * circulo.raio
}
func calcularAreaPerimetro(figuraGeometrica FiguraGeometrica) {
	fmt.Println("Área:", figuraGeometrica.area())
	fmt.Println("Perímetro:", figuraGeometrica.perimetro())
}
func testarInterface() {
	retangulo := Retangulo{3, 4}
	fmt.Println("Retângulo")
	calcularAreaPerimetro(retangulo)

	circulo := Circulo{5}
	fmt.Println("Círculo")
	calcularAreaPerimetro(circulo)
}

/***************************************************************************************************/
func testarFuncaoMatematica() {
	fmt.Println("Número utilizado:", 10.5)
	fmt.Println("Resultado do Seno:", math.Sin(10.5))
	fmt.Println("Resultado do Coseno:", math.Cos(10.5))
	fmt.Println("Resultado da Tangente:", math.Tan(10.5))
	fmt.Println("Resultado da função que remove casas decimais:", math.Floor(10.5))
	fmt.Println("Resultado da função que arredonda números:", math.Ceil(10.5))
	fmt.Println("Resultado da função que retorna o valor absoluto:", math.Abs(-10.5))
	fmt.Println("Resultado da função que calcula a raiz quadrada:", math.Sqrt(10.5))

	fmt.Println("Números utilizados:", 2, "e", 5)
	fmt.Println("Resultado da função que retorna um valor elevado a outro:", math.Pow(2, 5))
}

/***************************************************************************************************/
func testarDataHora() {
	tempo := time.Now()
	fmt.Println(tempo)

	fmt.Println("Dia:", tempo.Day())
	fmt.Println("Mês:", tempo.Month())
	fmt.Println("Ano:", tempo.Year())

	fmt.Println("Hora:", tempo.Hour())
	fmt.Println("Minuto:", tempo.Minute())
	fmt.Println("Segundo:", tempo.Second())
}

/***************************************************************************************************/
var caminhoEscrita = "F:\\googledrive\\estudos\\java\\go\\Ficheiro.txt"
var paises = [...]string{"Brasil", "Portugal", "Espanha", "França", "Itália", "Austrália", "Índia"}

func escreverFicheiro() {
	var ficheiro, _ = os.OpenFile(caminhoEscrita, os.O_WRONLY, 0644)
	defer ficheiro.Close()

	for _, pais := range paises {
		ficheiro.WriteString(pais + "\n")
	}

	ficheiro.Sync()
}
func testarFicheiroEscreve() {
	// "_" = blank identifier (Identificador em branco) - Usado para guardar um valor irrelevante ao
	// sistema, de forma temporária. Por exemplo, quando uma função retorna dois valores e apenas um
	// deles é relevante para o sistema e o outro pode ser descartado.
	var _, ficheiro = os.Stat(caminhoEscrita)

	if !os.IsNotExist(ficheiro) {
		escreverFicheiro()
	}
}

/***************************************************************************************************/
var caminhoLeitura = "F:\\googledrive\\estudos\\java\\go\\Ficheiro.txt"

func lerFicheiro() {
	var ficheiro, _ = os.OpenFile(caminhoEscrita, os.O_RDONLY, 0644)
	defer ficheiro.Close()

	var conteudo = make([]byte, 1024)

	ficheiro.Read(conteudo)

	fmt.Println(string(conteudo))
}
func testarFicheiroLe() {
	// "_" = blank identifier (Identificador em branco) - Usado para guardar um valor irrelevante ao
	// sistema, de forma temporária. Por exemplo, quando uma função retorna dois valores e apenas um
	// deles é relevante para o sistema e o outro pode ser descartado.
	var _, ficheiro = os.Stat(caminhoLeitura)

	if !os.IsNotExist(ficheiro) {
		lerFicheiro()
	}
}

/***************************************************************************************************/
func testarObterTipoVariavel() {
	nome := "Felipe"
	idade := 20
	versao := 1.1

	fmt.Println("O tipo da variável nome é", reflect.TypeOf(nome))
	fmt.Println("O tipo da variável idade é", reflect.TypeOf(idade))
	fmt.Println("O tipo da variável versao é", reflect.TypeOf(versao))
}

/***************************************************************************************************/
func testarDoisRetornos() (string, int) {
	nome := "Felipe"
	idade := 20
	return nome, idade
}

/***************************************************************************************************/
func testarSliceAppend() {
	nomes := []string{"Nome01", "Nome02", "Nome03"}
	fmt.Println(nomes)
	fmt.Println("Quantidade:", len(nomes))
	fmt.Println("Capacidade:", cap(nomes))

	nomes = append(nomes, "Nome04")
	fmt.Println(nomes)
	fmt.Println("Quantidade:", len(nomes))
	fmt.Println("Capacidade:", cap(nomes))

	nomes = append(nomes, "Nome05", "Nome06", "Nome07", "Nome08")
	fmt.Println(nomes)
	fmt.Println("Quantidade:", len(nomes))
	fmt.Println("Capacidade:", cap(nomes))
}

/***************************************************************************************************/
const qtdMonitoramento = 3
const tempoMonitoramento = 5

func exibirIntroducao() {
	nome := "Felipe"
	versao := 1.1

	fmt.Println("Olá senhor", nome)
	fmt.Println("Este programa está na versão", versao)
}
func exiibirMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}
func lerComando() int {
	var comando int
	// & - informa ao GO que a variável 'commando' é um input, irá receber informação digitada pelo usuário.
	// o '&' na frente da variável retorna o endereço em que a variável está na memória do computador.
	fmt.Scan(&comando)

	fmt.Println("O comando digitado foi", comando)
	fmt.Println("O endereço de memória, da variável 'comando' é", &comando)
	return comando
}
func iniciarMonitoramento() {
	//sites := []string{"https://www.alura.com.br/", "https://www.alura.com.br/dfsdfsdf", "https://www.google.com.br/"}
	sites := lerArquivoSites()

	for i := 0; i < qtdMonitoramento; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testarSite(site)
		}
		time.Sleep(tempoMonitoramento * time.Second)
	}
}
func testarSite(site string) {
	reposta, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if reposta.StatusCode == 200 {
		fmt.Println("Site carregado com sucesso!")
		registrarLog(site, true)
	} else {
		fmt.Println("Site com problemas!")
		fmt.Println("Código de status:", reposta.StatusCode)
		registrarLog(site, false)
	}
}
func lerArquivoSites() []string {
	var sites []string
	sites = lerArquivoTexto("suport/sites.txt")

	return sites
}
func registrarLog(site string, status bool) {
	// Codigo '0666' para dar permissao de escrita (na criacao) de arquivo novo.
	arquivo, err := os.OpenFile("suport/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
		arquivo.Close()
	}

	fmt.Println(arquivo)

	// Verificar no link 'https://golang.org/src/time/format.go', os formatos possiveis.
	arquivo.WriteString("[" + time.Now().Format("2006/01/02 15:04:05") + "] " + site + "- online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}
func exibirLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}
func testarMenuMultiplaEscolha() {
	exibirIntroducao()
	// GOLAND não possui while, usamos o for para que a aplicação rode em um looping infinito (somente para o looping, quando passado o comando zero).
	for {
		exiibirMenu()
		comando := lerComando()

		switch comando {
		case 1:
			fmt.Println("Monitorando...")
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
			exibirLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			// -1 indica que ocorreu um erro inesperado.
			os.Exit(-1)
		}
	}
}
func lerArquivoTexto(pathF string) []string {
	var linhas []string

	//arquivo, err := ioutil.ReadFile("sites.txt")
	//fmt.Println(string(arquivo))

	arquivo, err := os.Open(pathF)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
		arquivo.Close()
		os.Exit(0)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		linhas = append(linhas, linha)

		if err == io.EOF {
			break
		}
	}
	fmt.Println(linhas)

	arquivo.Close()

	return linhas
}
