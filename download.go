package main
 
import (
	"io"
	"os"
	"fmt"
	"bufio"
	"net/http"
	"path/filepath"
)
 

func main() {

	file, err := os.OpenFile("url.txt", os.O_RDWR, 0666)
	//url, filename :="",""
	if err != nil {
		fmt.Println("Fail to open file url.txt !")
		//fmt.Println(err)
	} else {
		buf := bufio.NewReader(file)
		burl, isPrefix, err := buf.ReadLine()
		url := string(burl)
		_ = isPrefix
		if err != nil {
			fmt.Println("Read error!")
		} else {
			filename := filepath.Base(url)
			fmt.Print("Link: ")
			fmt.Println(url)
			fmt.Print("File: ")
			fmt.Println(filename)
			
			if len(url) != 0 {
				res, err := http.Get(url)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Start to download...")
					fmt.Println("")
				}
				f, err := os.Create(filename)
					if err != nil {
						fmt.Println(err)
					}
				io.Copy(f, res.Body)
			}
		}
		
	}
}
