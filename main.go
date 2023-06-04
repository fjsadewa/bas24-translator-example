package main

import (
	"fmt"

	gt "github.com/bas24/googletranslatefree"
)

func main() {
    text := "Musisi Fiersa Besari buka suara usai mobil yang ditumpanginya menabrak batu dan nyaris masuk jurang di Jalan Poros Bone-Makassar, Sulawesi Selatan (Sulsel). Dia mengatakan insiden tersebut terjadi karena sopir travel yang mereka tumpangi mengantuk."
    result, err := gt.Translate(text, "id", "en")
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(result)
}