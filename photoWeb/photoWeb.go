package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	//UploadDir 是一个路径
	UploadDir = "./uploads"
	//TemplateDir html模板
	TemplateDir = "./views"
)

var templates map[string]*template.Template

func init() {
	templates = make(map[string]*template.Template)
	fileInfoArr, err := ioutil.ReadDir(TemplateDir)
	if err != nil {
		panic(err)
	}
	var templateName, templatePath string
	for _, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()
		if ext := filepath.Ext(templateName); ext != ".html" {
			continue
		}
		templatePath = TemplateDir + "/" + templateName
		log.Println("Loading template:", templatePath)
		t := template.Must(template.ParseFiles(templatePath))
		tmpl := strings.Split(templateName, ".")[0]
		templates[tmpl] = t
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// io.WriteString(w, "<html><form method=\"POST\" action=\"/upload\""+" enctype=\"multipart/form-data\">"+
		// 	"Choose an image to upload:<input name=\"image\" type=\"file\" />"+"<input type=\"submit\" value=\"Upload\" />"+
		// 	"</form></html>")
		// t, err := template.ParseFiles("upload.html")
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
		// t.Execute(w, nil)
		if err := renderHTML(w, "upload", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		filename := h.Filename
		defer func() {
			errFile := f.Close()
			if errFile != nil {
				log.Println("文件关闭时发生错误！")
			}
		}()
		t, err := os.Create(UploadDir + "/" + filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//defer t.Close()
		defer func() {
			errFile := t.Close()
			if errFile != nil {
				log.Println("文件关闭时发生错误！")
			}
		}()
		if _, err := io.Copy(t, f); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageID := r.FormValue("id")
	imagePath := UploadDir + "/" + imageID
	if exists := isExists(imagePath); !exists {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir(UploadDir)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	locals := make(map[string]interface{})
	images := []string{}
	for _, fileInfo := range fileInfoArr {
		images = append(images, fileInfo.Name())
	}
	locals["images"] = images
	if err := renderHTML(w, "list", locals); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// t, err := template.ParseFiles("list.html")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// t.Execute(w, locals)

	// var listHTML string
	// for _, fileInfo := range fileInfoArr {
	// 	imgid := fileInfo.Name()
	// 	listHTML += "<li><a href=\"/view?id=" + imgid + "\">" + imgid + "</a></li>"
	// }
	// io.WriteString(w, "<html><ol>"+listHTML+"</ol></html>")
}

func renderHTML(w io.Writer, tmpl string, locals map[string]interface{}) error {
	err := templates[tmpl].Execute(w, locals)
	return err
}

func main() {
	http.HandleFunc("/", listHandler)
	http.HandleFunc("/view", viewHandler)
	http.HandleFunc("/upload", uploadHandler)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}
