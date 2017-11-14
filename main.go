package main

import (
	"time"
	"fmt"
	"path/filepath"
	"flag"
	"io/ioutil"
	"os"
	"io"
	"src/github.com/jinzhu/gorm"
	_"src/github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	oPath = flag.String("o","c:/","originFilePath")
	dPath = flag.String("d","c:/","destnationFilePath")
	ext = flag.String("e",".exe","fileExtension")
)

type Filelist struct {
	gorm.Model
	FileName string
	Date string
}

func main(){
	flag.Parse()
	fmt.Println("元dir " + *oPath)
	fmt.Println("先dir " + *dPath)
	fmt.Println("拡張子 " + *ext)


	t := time.NewTicker(60 * time.Second)
	for {
		select {
		case <-t.C:
			fileList := dirwalk(*oPath,*dPath,*ext)
			if len(fileList) != 0 {
				//fmt.Println(fileList)
				writeDB(fileList)
			}
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
		fmt.Println(file.Name() + " : " + filepath.Ext(file.Name()))
		if filepath.Ext(file.Name()) ==  pExt{
			ofile := filepath.Join(pOPath,file.Name())
			dfile := filepath.Join(pDPath,file.Name())

			fileCopy(ofile,dfile)

			paths = append(paths, filepath.Join(pOPath, file.Name()))
		}
	}
	return paths
}

//ファイルコピー
func fileCopy(pOfile string,pDfile string){
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

func writeDB(pFileList []string){
	db,err := gorm.Open("sqlite3","List.sqlite3")
	if err != nil{
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&Filelist{})

	for _,file := range pFileList{
		fmt.Println(file)
		db.Create(&Filelist{FileName:file,Date:time.Now().Format("2006/01/02")})
	}
}

func ReadDB(keyword string){
	db, err := gorm.Open("sqlite3","List.sqlite3")
	if err != nil {
		panic("failed to connect database")
	}

	//read
	var list Filelist
	db.Find(&list,"Date = ?",keyword)


}