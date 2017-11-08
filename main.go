package main

import (
	"time"
	"fmt"
	"path/filepath"
	"flag"
	"io/ioutil"
	"os"
	"path"
	"io"
)

var (
	oPath = flag.String("o","c:/","originFilePath")
	dPath = flag.String("d","c:/","destnationFilePath")
	ext = flag.String("e",".exe","fileExtension")
)
func main(){
	flag.Parse()

	t := time.NewTicker(60 * time.Second)
	for {
		select {
		case <-t.C:
			 fmt.Println(dirwalk(*oPath,*dPath,*ext))
		}
	}
	t.Stop()
}

func dirwalk(pOPath string,pDPath string,pExt string) []string {

	files, err := ioutil.ReadDir(pOPath)  //file列挙
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if path.Ext(file.Name()) ==  pExt{
			ofile := filepath.Join(pOPath,file.Name())
			dfile := filepath.Join(pDPath,file.Name())

			FileCopy(ofile,dfile)

			paths = append(paths, filepath.Join(pOPath, file.Name()))
		}
	}
	return paths
}

//ファイルコピー
func FileCopy(pOfile string,pDfile string){
	oldFile,err := os.Open(pOfile)
	if err != nil{
		panic(err)
	}
	newFile,err := os.Create(pDfile)
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	_,err = io.Copy(newFile,oldFile)
	if err != nil {
		panic(err)
	}
	oldFile.Close()
	os.Remove(pOfile)
}