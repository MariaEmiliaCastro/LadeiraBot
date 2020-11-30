package main

import (
	"flag"
	"fmt"
	"os"
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

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// Se a mensagem for "!ladeira", responder com "abaixo!"
	if m.Content == "ladeira" {
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
