package handler

import (
	"cram-school-reserve-server/back/infra/rdb"
	"cram-school-reserve-server/front/pkg"

	"log"
	"net/http"

	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
)

type ReservationInfo struct {
	StudentID string
	Time      []string
}

func CreateCSVTemplate(c *gin.Context, ID int) ([]string, int, *rdb.Form, string) {
	user := pkg.CheckLogin(c)
	if user == nil {
		// Handle error here
	}
	if student, err := user.(*rdb.Student); err {
		form, err := rdb.GetForm(c, ID)
		if err != nil {
			http.Error(c.Writer, "Failed to get form", http.StatusInternalServerError)
		}

		dates, err := rdb.CreateReservationPeriodList(*form)
		if err != nil {
			http.Error(c.Writer, "Failed to create reservation period list", http.StatusInternalServerError)

		}

		formattedDates, err := rdb.FormatDates(dates)
		if err != nil {
			http.Error(c.Writer, "Failed to format dates", http.StatusInternalServerError)
			log.Println(err)
		}
		return formattedDates, ID, form, student.ID
	} else if _, err := user.(*rdb.Teacher); err {
		form, err := rdb.GetForm(c, ID)
		if err != nil {
			http.Error(c.Writer, "Failed to get form", http.StatusInternalServerError)

		}

		dates, err := rdb.CreateReservationPeriodList(*form)
		if err != nil {
			http.Error(c.Writer, "Failed to create reservation period list", http.StatusInternalServerError)

		}

		formattedDates, err := rdb.FormatDates(dates)
		if err != nil {
			http.Error(c.Writer, "Failed to format dates", http.StatusInternalServerError)
			log.Println(err)
		}
		return formattedDates, ID, form, ""
	}
	return nil, 0, nil, ""
}

func GetCreateCSVForAllForms(c *gin.Context) {
	formattedDates, _, _, _ := CreateCSVTemplate(c, 1)
	dateReservations := make(map[string][]ReservationInfo)

	for _, date := range formattedDates {
		for i := 1; i <= 3; i++ {
			reservations, err := rdb.GetReservationByFormIDAndDate(c, i, date)

			if err != nil {
				http.Error(c.Writer, "Failed to get reservation", http.StatusInternalServerError)

			}
			for _, reservation := range reservations {
				time := strings.Split(reservation.Time, "-")
				// 予約情報を作成
				info := ReservationInfo{
					StudentID: reservation.StudentID,
					Time:      time,
				}
				// マップに予約情報を追加
				dateReservations[date] = append(dateReservations[date], info)
			}
		}
	}

	CreateExcel(c, "all", dateReservations, formattedDates, "teacher")
	DownloadExcel("all", c)
}

func GetCreateCSVForOneForm(c *gin.Context) {
	ID, err := strconv.Atoi(c.Query("form_id"))
	if err != nil {
		http.Error(c.Writer, "Invalid form ID", http.StatusBadRequest)
		return
	}
	formattedDates, ID, form, _ := CreateCSVTemplate(c, ID)
	dateReservations := make(map[string][]ReservationInfo)

	for _, date := range formattedDates {

		reservations, err := rdb.GetReservationByFormIDAndDate(c, ID, date)

		if err != nil {
			http.Error(c.Writer, "Failed to get reservation", http.StatusInternalServerError)

		}
		for _, reservation := range reservations {
			time := strings.Split(reservation.Time, "-")
			// 予約情報を作成
			info := ReservationInfo{
				StudentID: reservation.StudentID,
				Time:      time,
			}
			// マップに予約情報を追加
			dateReservations[date] = append(dateReservations[date], info)
		}
	}

	CreateExcel(c, form.Name, dateReservations, formattedDates, "teacher")
	DownloadExcel(form.Name, c)
}

func GetCreateCSVForOneStudent(c *gin.Context) {
	formattedDates, _, form, studentID := CreateCSVTemplate(c, 1)
	dateReservations := make(map[string][]ReservationInfo)

	for _, date := range formattedDates {
		for i := 1; i <= 3; i++ {
			reservations, err := rdb.GetReservationByFormIDAndStudentIDAndDate(c, i, studentID, date)

			if err != nil {
				http.Error(c.Writer, "Failed to get reservation", http.StatusInternalServerError)

			}
			for _, reservation := range reservations {
				time := strings.Split(reservation.Time, "-")
				// 予約情報を作成
				info := ReservationInfo{
					StudentID: reservation.StudentID,
					Time:      time,
				}
				// マップに予約情報を追加
				dateReservations[date] = append(dateReservations[date], info)
			}
		}
	}

	CreateExcel(c, form.Name, dateReservations, formattedDates, "stundet")
	DownloadExcel(form.Name, c)
}

func CreateExcel(c *gin.Context, formName string, dateReservations map[string][]ReservationInfo, dates []string, user string) {
	// 新しいExcelファイルを作成
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Reservations")
	if err != nil {
		log.Fatal(err)
	}

	oneDayReservationInfo := make([][]string, 10)
	// 各日付について処理
	for i, date := range dates {

		infos := dateReservations[date]

		for j := range oneDayReservationInfo {
			oneDayReservationInfo[j] = make([]string, 0) // 各要素もスライスとして初期化
		}
		for _, info := range infos {
			for _, time := range info.Time {
				student, err := rdb.GetStudent(c, info.StudentID)
				if err != nil {
					log.Fatal(err)
				}
				studentName := student.Name
				// 時間を整数に変換
				reservedClassNumber, err := strconv.Atoi(time)
				if err != nil {
					log.Fatal(err)
				}
				oneDayReservationInfo[reservedClassNumber] = append(oneDayReservationInfo[reservedClassNumber], studentName)
			}
		}
		if user == "student" {
			InputToExcelForStudent(sheet, date, oneDayReservationInfo, i)
		} else {
			InputToExcel(sheet, date, oneDayReservationInfo, i)
		}
	}

	// Excelファイルを保存
	err = file.Save("xlsx/" + formName + "_reservations.xlsx")
	if err != nil {
		log.Fatal(err)
	}
}

func InputToExcel(sheet *xlsx.Sheet, date string, oneDayReservationInfo [][]string, nowPage int) {
	// 新しいスタイルを作成します
	style := xlsx.NewStyle()

	// 罫線を作成します
	border := *xlsx.NewBorder("thin", "thin", "thin", "thin")

	// スタイルに罫線を追加します
	style.Border = border

	timeAxis := []string{
		"14:35~15:15",
		"15:20~16:00",
		"16:05~16:45",
		"16:50~17:30",
		"17:35~18:15",
		"18:20~19:00",
		"19:05~19:45",
		"19:50~20:30",
		"20:35~21:15",
		"21:20~22:00",
	}

	sheet.Cell(0+66*nowPage, 0).Value = "時間/日付"
	sheet.Cell(0+66*nowPage, 5).Value = date

	boldStyle := xlsx.NewStyle()
	boldBorder := *xlsx.NewBorder("medium", "medium", "medium", "medium")
	boldStyle.Border = boldBorder

	startCol, endCol := 0, 8
	startRow, endRow := 0, 8
	for i, time := range timeAxis {
		if i < 6 {
			for k := 0; k < 6; k++ {
				times := strings.Split(time, "~")
				if k == 2 {
					sheet.Cell((k+6*i+1)+66*nowPage, 0).Value = times[0]
				} else if k == 3 {
					sheet.Cell((k+6*i+1)+66*nowPage, 0).Value = "~"
				} else if k == 4 {
					sheet.Cell((k+6*i+1)+66*nowPage, 0).Value = times[1]
				} else if k == 0 {
					sheet.Cell((k+6*i+1)+66*nowPage, 0).Value = "----"
				} else if k == 5 {
					sheet.Cell((k+6*i+1)+66*nowPage, 0).Value = "----"
				} else {
					sheet.Cell((k+6*i+1)+66*nowPage, 0).Value = ""
				}
				if k == 0 {
					startRow = (k + 6*i + 1) + 66*nowPage
				}
				if k == 5 {
					endRow = (k + 6*i + 1) + 66*nowPage
				}
			}
			//BoldBorder(sheet, startRow, endRow, startCol, endCol, boldStyle)
		} else {
			for k := 0; k < 7; k++ {
				times := strings.Split(time, "~")
				if k == 2 {
					sheet.Cell((36+k+7*(i-6)+1)+66*nowPage, 0).Value = times[0]
				} else if k == 3 {
					sheet.Cell((36+k+7*(i-6)+1)+66*nowPage, 0).Value = "~"
				} else if k == 4 {
					sheet.Cell((36+k+7*(i-6)+1)+66*nowPage, 0).Value = times[1]
				} else if k == 0 {
					sheet.Cell((36+k+7*(i-6)+1)+66*nowPage, 0).Value = "----"
				} else if k == 6 {
					sheet.Cell((36+k+7*(i-6)+1)+66*nowPage, 0).Value = "----"
				} else {
					sheet.Cell((36+k+7*(i-6)+1)+66*nowPage, 0).Value = ""
				}
				if k == 0 {
					startRow = (36 + k + 7*(i-6) + 1) + 66*nowPage
				}
				if k == 6 {
					endRow = (36 + k + 7*(i-6) + 1) + 66*nowPage
				}
			}
			//BoldBorder(sheet, startRow, endRow, startCol, endCol, boldStyle)
		}
	}

	for i, info := range oneDayReservationInfo {
		rowCounter := 1
		lineCounter := 0
		if i < 6 {
			for _, studentInfo := range info {
				sheet.Cell(6*i+1+66*nowPage+lineCounter, rowCounter).Value = studentInfo
				if rowCounter < 9 {
					rowCounter++
				} else {
					rowCounter = 1
					lineCounter++
				}
			}
		} else {
			for _, studentInfo := range info {
				sheet.Cell(36+7*(i-6)+1+66*nowPage, rowCounter).Value = studentInfo
				if rowCounter < 9 {
					rowCounter++
				} else {
					rowCounter = 1
					lineCounter++
				}
			}
		}
	}

	startRow, endRow = 1+66*nowPage, 66+66*nowPage
	startCol, endCol = 1, 8

	// 指定した範囲のセルにスタイルを適用します
	for r := startRow; r <= endRow; r++ {
		for c := startCol; c <= endCol; c++ {
			cell := sheet.Cell(r, c)
			cell.SetStyle(style)
		}
	}
}

func InputToExcelForStudent(sheet *xlsx.Sheet, date string, oneDayReservationInfo [][]string, nowPage int) {
	// 新しいスタイルを作成します
	style := xlsx.NewStyle()

	// 罫線を作成します
	border := *xlsx.NewBorder("thin", "thin", "thin", "thin")

	// スタイルに罫線を追加します
	style.Border = border

	timeAxis := []string{
		"14:35~15:15",
		"15:20~16:00",
		"16:05~16:45",
		"16:50~17:30",
		"17:35~18:15",
		"18:20~19:00",
		"19:05~19:45",
		"19:50~20:30",
		"20:35~21:15",
		"21:20~22:00",
	}

	sheet.Cell(0+66*nowPage, 0).Value = "時間/日付"

	boldStyle := xlsx.NewStyle()
	boldBorder := *xlsx.NewBorder("medium", "medium", "medium", "medium")
	boldStyle.Border = boldBorder

	startCol, endCol := 0, 8
	startRow, endRow := 0, 8
	for i, time := range timeAxis {
		if i < 6 {
			for k := 0; k < 6; k++ {
				times := strings.Split(time, "~")
				if k == 2 {
					sheet.Cell((k+6*i+1)+66*nowPage, 0).Value = times[0]
				} else if k == 3 {
					sheet.Cell((k+6*i+1)+66*nowPage, 0).Value = "~"
				} else if k == 4 {
					sheet.Cell((k+6*i+1)+66*nowPage, 0).Value = times[1]
				} else if k == 0 {
					sheet.Cell((k+6*i+1)+66*nowPage, 0).Value = "----"
				} else if k == 5 {
					sheet.Cell((k+6*i+1)+66*nowPage, 0).Value = "----"
				} else {
					sheet.Cell((k+6*i+1)+66*nowPage, 0).Value = ""
				}
				if k == 0 {
					startRow = (k + 6*i + 1) + 66*nowPage
				}
				if k == 5 {
					endRow = (k + 6*i + 1) + 66*nowPage
				}
			}
			//BoldBorder(sheet, startRow, endRow, startCol, endCol, boldStyle)
		} else {
			for k := 0; k < 7; k++ {
				times := strings.Split(time, "~")
				if k == 2 {
					sheet.Cell((36+k+7*(i-6)+1)+66*nowPage, 0).Value = times[0]
				} else if k == 3 {
					sheet.Cell((36+k+7*(i-6)+1)+66*nowPage, 0).Value = "~"
				} else if k == 4 {
					sheet.Cell((36+k+7*(i-6)+1)+66*nowPage, 0).Value = times[1]
				} else if k == 0 {
					sheet.Cell((36+k+7*(i-6)+1)+66*nowPage, 0).Value = "----"
				} else if k == 6 {
					sheet.Cell((36+k+7*(i-6)+1)+66*nowPage, 0).Value = "----"
				} else {
					sheet.Cell((36+k+7*(i-6)+1)+66*nowPage, 0).Value = ""
				}
				if k == 0 {
					startRow = (36 + k + 7*(i-6) + 1) + 66*nowPage
				}
				if k == 6 {
					endRow = (36 + k + 7*(i-6) + 1) + 66*nowPage
				}
			}
			//BoldBorder(sheet, startRow, endRow, startCol, endCol, boldStyle)
		}
	}

	for i, info := range oneDayReservationInfo {
		rowCounter := 1
		lineCounter := 0
		sheet.Cell(0+66*nowPage, i+1).Value = date
		if i < 6 {
			for _, studentInfo := range info {
				sheet.Cell(6*i+1+66*nowPage+lineCounter, rowCounter).Value = studentInfo
				if rowCounter < 9 {
					rowCounter++
				} else {
					rowCounter = 1
					lineCounter++
				}
			}
		} else {
			for _, studentInfo := range info {
				sheet.Cell(36+7*(i-6)+1+66*nowPage, rowCounter).Value = studentInfo
				if rowCounter < 9 {
					rowCounter++
				} else {
					rowCounter = 1
					lineCounter++
				}
			}
		}
	}

	startRow, endRow = 1+66*nowPage, 66+66*nowPage
	startCol, endCol = 1, 8

	// 指定した範囲のセルにスタイルを適用します
	for r := startRow; r <= endRow; r++ {
		for c := startCol; c <= endCol; c++ {
			cell := sheet.Cell(r, c)
			cell.SetStyle(style)
		}
	}
}

func BoldBorder(sheet *xlsx.Sheet, startRow, endRow, startCol, endCol int, boldStyle *xlsx.Style) {
	for r := startRow; r <= endRow; r++ {
		for c := startCol; c <= endCol; c++ {
			cell := sheet.Cell(r, c)
			cell.SetStyle(boldStyle)
		}
	}
}
