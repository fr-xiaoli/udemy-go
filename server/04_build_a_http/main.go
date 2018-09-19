package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	request(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			mux(conn, ln)
		}
		if ln == "" {
			break
		}
		i++
	}
}

func mux(conn net.Conn, ln string) {
	// request Line
	m := strings.Fields(ln)[0]
	u := strings.Fields(ln)[1]
	fmt.Println("*** METHOD:", m)
	fmt.Println("*** URI:", u)

	// multiplexer
	if m == "GET" && u == "/" {
		index(u, conn)
	}
	if m == "GET" && u == "/about" {
		about(u, conn)
	}
	if m == "GET" && u == "/contact" {
		contact(u, conn)
	}
	if m == "GET" && u == "/apply" {
		apply(u, conn)
	}
	if m == "POST" && u == "/apply" {
		applyProcess(u, conn)
	}
}

func generateBody(uri string, hasApplyForm bool) string {
	uri = strings.Replace(uri, "/", "", -1)
	if uri == "" {
		uri = "index"
	}
	title := strings.Title(uri)
	formPlaceholder := ""
	if hasApplyForm {
		formPlaceholder = `<div>
				<form method="post" action="/apply">
					<input type="submit" value="apply">
				</form>
			</div>`
	}
	body := fmt.Sprintf(`<!DOCTYPE html>
	<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Index</title>
		</head>
		<body>
			<h1>%s</h1>
			<a href="/">index</a><br/>
			<a href="/about">about</a><br/>
			<a href="/contact">contact</a><br/>
			<a href="/apply">apply</a>
			%s
		</body>
	</html>`, title, formPlaceholder)
	return body
}

func index(uri string, conn net.Conn) {
	body := generateBody(uri, false)
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func about(uri string, conn net.Conn) {
	body := generateBody(uri, false)
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func contact(uri string, conn net.Conn) {
	body := generateBody(uri, false)
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func apply(uri string, conn net.Conn) {
	body := generateBody(uri, true)
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func applyProcess(uri string, conn net.Conn) {
	body := generateBody(uri, false)
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
