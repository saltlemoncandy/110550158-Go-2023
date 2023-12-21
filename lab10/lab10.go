package main

import (
	"bufio"
	"log"
	"net/http"
	"strings"
	"os"
	"unicode/utf8"
	"regexp"
	"context"
	"github.com/gorilla/websocket"
	"github.com/reactivex/rxgo/v2"
)

type client chan<- string // an outgoing message channel

var (
	entering      = make(chan client)
	leaving       = make(chan client)
	messages      = make(chan rxgo.Item) // all incoming client messages
	ObservableMsg = rxgo.FromChannel(messages)
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	MessageBroadcast := ObservableMsg.Observe()
	for {
		select {
		case msg := <-MessageBroadcast:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli <- msg.V.(string)
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func clientWriter(conn *websocket.Conn, ch <-chan string) {
	for msg := range ch {
		conn.WriteMessage(1, []byte(msg))
	}
}

func wshandle(w http.ResponseWriter, r *http.Request) {
	upgrader := &websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}

	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "你是 " + who + "\n"
	messages <- rxgo.Of(who + " 來到了現場" + "\n")
	entering <- ch

	defer func() {
		log.Println("disconnect !!")
		leaving <- ch
		messages <- rxgo.Of(who + " 離開了" + "\n")
		conn.Close()
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		messages <- rxgo.Of(who + " 表示: " + string(msg))
	}
}

var (
    swearWords     = loadWordsFromFile("swear_word.txt")
    sensitiveNames = loadWordsFromFile("sensitive_name.txt")
)

func loadWordsFromFile(filename string) []string {
    file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var words []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        words = append(words, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    return words
}

func filterSwearWords(msg string) bool {
    for _, word := range swearWords {
        if strings.Contains(msg, word) {
            return false // Filter out message containing swear words
        }
    }
    return true // Allow message if no swear words found
}

func mapSensitiveNames(msg string) string {
    for _, name := range sensitiveNames {
        pattern := regexp.QuoteMeta(name)
        re := regexp.MustCompile(pattern)

        replacer := func(match string) string {
            if utf8.RuneCountInString(name) == 3 {
                runes := []rune(match)
                runes[1] = []rune("*")[0] 
                return string(runes)
            } else if utf8.RuneCountInString(name) == 2 {
                runes := []rune(match)
                runes[1] = []rune("*")[0]
                return string(runes)
            }
            return match 
        }

        msg = re.ReplaceAllStringFunc(msg, replacer)
    }
    return msg
}

func InitObservable() {
    ObservableMsg = ObservableMsg.Filter(func(item interface{}) bool {
        msg, ok := item.(string)
        if !ok {
            return false
        }
        return filterSwearWords(msg)
    }).Map(func(ctx context.Context, item interface{}) (interface{}, error) {
        msg, ok := item.(string)
        if !ok {
            return "", nil
        }
        return mapSensitiveNames(msg), nil
    })
}

func main() {
	InitObservable()
	go broadcaster()
	http.HandleFunc("/wschatroom", wshandle)

	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Println("server start at :8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
