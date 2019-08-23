package receiver

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"../props"
)

type Sizer interface {
	Size() int64
}

const htmlTemp=`<html><form enctype="multipart/form-data" action="/upload" method="POST">
<p>文件: <input name="userfile" type="file" /><input type="submit" value="上传" /></p>
</form><p>%s</p></html>`

func CreateHandler(props propsReader.AppConfigProperties) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		UploadHandler(w, r, props)
	}
}

func UploadHandler(w http.ResponseWriter, r *http.Request, props propsReader.AppConfigProperties) {
	if "POST" == r.Method {
		file, handler, err := r.FormFile("userfile")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		defer file.Close()
		filePath := props["homedir"] + "/" + handler.Filename
		f, err := os.Create(filePath)
		defer f.Close()
		io.Copy(f, file)
		fileStat, err := f.Stat()
		var fileSizeStr string
		if fileStat.Size() > (1024*1024){
			fileSizeStr = fmt.Sprintf("%dM", fileStat.Size()/(1024*1024))
		} else if fileStat.Size() > 1024 {
			fileSizeStr = fmt.Sprintf("%dK", fileStat.Size()/1024)
		} else {
			fileSizeStr = fmt.Sprintf("%dBytes", fileStat.Size())
		}
		log.Println("file received and stored in ", filePath)
		msg := fmt.Sprintf(" 文件已上传到%s, 文件大小 %s", filePath, fileSizeStr )
		fmt.Fprintf(w, htmlTemp, msg)
		return
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, fmt.Sprintf(htmlTemp,""))
}



