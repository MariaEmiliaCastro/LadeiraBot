package main

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Variavel utilizada para os parametros de linha de comando
var (
	Token string
)

func init() {

	flag.StringVar(&Token, "t", "Nzc4NDEyNTEyODc4MDAyMTc3.X7RnJQ.uy6A-mhavH9U83xmuB_MDYEvDUo", "Bot Token")
	flag.Parse()
}

//baseado nisso aqui https://golangcode.com/check-if-element-exists-in-slice/
func Find(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)

	if arr.Kind() != reflect.Array {
		return false
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

func main() {

	// Cria nova sessão do bot no discord utilizando o token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Registra messageCreate func como um callback para eventos gerados por MessageCreate.
	dg.AddHandler(messageCreate)

	// Evento de recebimento de mensagens
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	// Abre uma conexão websocket com o Discord e inicia a escuta.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Fica aqui até o usuário digitar CTRL-C no console para encerrar a sessão.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Fecha a sessão do Discord de maneira limpa.
	dg.Close()
}

// Esta função será chamada sempre que um novo evento de mensagem (addHandler acima) ocorre nos channel que o bot tem acesso
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	comandos := [...]string{"world","country","state","help"}
	//estados := [...]string{"AC", "AL", "AM", "AP", "BA", "CE", "DF", "ES", "GO", "MA", "MT", "MS", "MG", "PA", "PB", "PR", "PE", "PI", "RJ", "RN", "RO", "RS", "RR", "SC", "SE", "SP", "TO"}

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}



	if m.Content == "!cruzeiro" {
		s.ChannelMessageSend(m.ChannelID, "***É o melhor time de todos! :blue_heart:***")
	}


	if strings.HasPrefix(m.Content, "!corona"){

		query:=strings.Split(m.Content, " ")

		if len(query) >1 { 
			if(!Find(comandos, query[1])){
				s.ChannelMessageSend(m.ChannelID, "Comando não encontrado! Digite `!corona help` para verificar os comandos disponíveis.")
			}else if len(query)==2{

				if (query[1]=="world") {
					worldInfo, _ := allCountriesCorona()
					s.ChannelMessageSend(m.ChannelID, "```diff\nInformacao sobre corona no Mundo\n- Casos:\t\t"+
					strconv.Itoa(worldInfo.Cases)+
					"\n- Mortes:\t\t"+strconv.Itoa(worldInfo.Deaths)+
					"\n+ Recuperados:\t"+strconv.Itoa(worldInfo.Recovered)+
					"\n```")
				}else{
					s.ChannelMessageSend(m.ChannelID,"**Comandos atuais:**\n\n***`corona`*** - Retorna informações sobre o corona no Brasil.\n***`corona country <countryName>`*** - Retorna informações sobre o corona ao redor do mundo.\n***`corona state <uf>`*** - Retorna informações sobre o corona de um estado específico.\n***`corona country <countryName>`*** - Retorna informações sobre o corona de um país específico. O nome do país deve ser digitado em inglês.\n***`corona help`*** - Lista todos os comandos do bot.\n\neg: `!corona world`")
					
				}

			}else if len(query)==3{

				if (query[1]=="country"){
					queryData, _ := getCountry(query[2])
					s.ChannelMessageSend(m.ChannelID, "```diff\n"+queryData.Country+" - Informações sobre corona no país.\n- Casos:\t\t"+
					strconv.Itoa(queryData.Cases)+
					"\n- Mortes:\t\t"+strconv.Itoa(queryData.Deaths)+
					"\n+ Recuperados:\t"+strconv.Itoa(queryData.Recovered)+
					"\n```")
				}else if (query[1]=="state"){
					queryData, _ := getState(query[2])
					s.ChannelMessageSend(m.ChannelID, "```diff\n"+queryData.State+" - Informações sobre corona no estado.\n- Casos:\t\t"+
					strconv.Itoa(queryData.Cases)+
					"\n- Mortes:\t\t"+strconv.Itoa(queryData.Deaths)+
					"\n+ Suspeitos:\t"+strconv.Itoa(queryData.Suspects)+
					"\n```")

				}




			}

		}else {

			queryData, _ := getCountry("brazil")
			s.ChannelMessageSend(m.ChannelID, "```diff\nInformacao sobre corona no Brasil\n- Casos:\t\t"+
			strconv.Itoa(queryData.Cases)+
			"\n- Mortes:\t\t"+strconv.Itoa(queryData.Deaths)+
			"\n+ Recuperados:\t"+strconv.Itoa(queryData.Recovered)+
			"\n```")
			
		}


	// Se a mensagem for "!ladeira", responder com "abaixo!"
	if m.Content == "!ladeira" {
		s.ChannelMessageSend(m.ChannelID, "abaixo!")
	}

	if strings.HasPrefix(m.Content, "!corona world") {
		worldInfo, _ := allCountriesCorona()

		s.ChannelMessageSend(m.ChannelID, "```diff\nInformacao sobre corona no Mundo\n- Casos:\t\t"+
			strconv.Itoa(worldInfo.Cases)+
			"\n- Mortes:\t\t"+strconv.Itoa(worldInfo.Deaths)+
			"\n+ Recuperados:\t"+strconv.Itoa(worldInfo.Recovered)+
			"\n```")
	}
}
