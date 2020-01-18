package main

type Retiever interface {
	Get(url string) string
}

func download(r Retiever) string {
	return r.Get("www.immoc.com")
}

func main() {

}
