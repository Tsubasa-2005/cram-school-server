package handler

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func DownloadExcel(formName string, c *gin.Context) {

	// Excelファイルのパスを指定します。
	// この例では、CreateExcel関数で作成したExcelファイルの名前とパスを指定します。
	excelFilePath := "csvFiles/" + formName + "_reservations.xlsx"
	downloadName := formName + "_reservations.xlsx"

	// Excelファイルを開きます。
	file, err := os.Open(excelFilePath)
	if err != nil {
		c.String(http.StatusInternalServerError, "Unable to open the file: %v", err)
		return
	}
	defer file.Close()

	// ファイルの内容を読み込みます。
	content, err := ioutil.ReadAll(file)
	if err != nil {
		c.String(http.StatusInternalServerError, "Unable to read the file: %v", err)
		return
	}

	// ファイルの内容をHTTPレスポンスとして送信します。
	// "Content-Disposition"ヘッダーを設定することで、ブラウザはこのレスポンスをダウンロードとして扱います。
	c.Header("Content-Disposition", "attachment; filename="+downloadName)
	c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", content)

	if _, err := os.Stat(excelFilePath); os.IsNotExist(err) {
		c.String(http.StatusInternalServerError, "File does not exist: %v", err)
		return
	}

	// なぜかフォルダが削除されない。ファイルだけを指定しても削除されなかった。
	//err = os.RemoveAll("csvFiles")
	//if err != nil {
	//	c.String(http.StatusInternalServerError, "Unable to delete the file: %v", err)
	//	return
	//}
	//err = os.MkdirAll("csvFiless", 0755)
	//if err != nil {
	//	log.Fatalf("Failed to create directory: %v", err)
	//}

	c.Redirect(http.StatusSeeOther, "/teacher")
}
