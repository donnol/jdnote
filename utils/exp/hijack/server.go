package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func StartServer(ctx context.Context, port string, targetPort string) {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if targetPort == "" {
			fmt.Printf("timeout targetPort: %s\n", targetPort)

			// 延迟
			time.Sleep(10 * time.Second)

			// 超时后将连接劫持
			hijacker, ok := w.(http.Hijacker)
			if !ok {
				fmt.Printf("can't hijack\n")
				return
			}
			fmt.Printf("hijack \n")
			conn, bufrw, err := hijacker.Hijack()
			if err != nil {
				panic(err)
			}
			defer conn.Close()

			err = conn.SetWriteDeadline(time.Now().Add(20 * time.Second))
			if err != nil {
				panic(err)
			}

			// _, err = bufrw.WriteString("Now we're speaking raw TCP. Say hi: ")
			// if err != nil {
			// 	panic(err)
			// }
			// bufrw.Flush()
			s, err := bufrw.ReadString('\n')
			if err != nil {
				panic(err)
			}
			fmt.Fprintf(bufrw, "You said: %q\nBye.\n", s)
			bufrw.Flush()
		} else {
			fmt.Printf("hijack targetPort: %s\n", targetPort)

			if err := do(ctx, 5*time.Second, "http://127.0.0.1"+targetPort); err != nil {
				if strings.Contains(err.Error(), "context deadline exceeded") {
					// 超时后将连接劫持
					hijacker, ok := w.(http.Hijacker)
					if !ok {
						fmt.Printf("can't hijack\n")
						return
					}
					fmt.Printf("hijack successfully\n")
					conn, bufrw, err := hijacker.Hijack()
					if err != nil {
						panic(err)
					}
					defer conn.Close()

					err = conn.SetWriteDeadline(time.Now().Add(20 * time.Second))
					if err != nil {
						panic(err)
					}

					// _, err = bufrw.WriteString("Now we're speaking raw TCP. Say hi: ")
					// if err != nil {
					// 	panic(err)
					// }
					// bufrw.Flush()
					// s, err := bufrw.ReadString('\n')
					// if err != nil {
					// 	panic(err)
					// }
					fmt.Fprintf(bufrw, "You said: %q\nBye.\n", "hello")
					bufrw.Flush()
				} else {
					panic(err)
				}
				return
			}
		}

		n, err := w.Write([]byte("hello"))
		if err != nil {
			panic(err)
		}
		fmt.Printf("n: %d, tergetPort: %s\n", n, targetPort)
	}))
	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
