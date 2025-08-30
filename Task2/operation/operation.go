package operation

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"task2/models"
	"time"
)

func InsertUser(u *models.User, ch chan models.LogInfo) error {
	bytes, err := json.Marshal(u)
	if err != nil {
		return err
	}
	fw := New("data.data")
	fmt.Fprintln(fw, string(bytes))
	//ch := make(chan models.LogInfo)

	// data := models.LogInfo{
	// 	Create_on: time.Now(),
	// 	Action:    "Insert",
	// }
	// ch <- data
	// InsertLogDetails(ch)
	go func() {
		ch <- models.LogInfo{
			Create_on: time.Now(),
			Action:    "Insert",
		}
	}()

	return nil

}

func InsertLogDetails(ch chan models.LogInfo) {?>1?
	// logInfo := <-ch
	// fw := New("audit.data")
	// fmt.Fprintln(fw, logInfo)
	go func() {
		for logInfo := range ch {
			bytes, _ := json.Marshal(logInfo)
			fw := New("audit.data")
			fmt.Fprintln(fw, string(bytes))
		}
	}()
	time.Sleep(time.Millisecond * 20)
}

func GetUserByEmail(email string, ch chan models.LogInfo) error {
	data, err := os.ReadFile("data.data")
	if err != nil {
		return err
	}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		var user models.User
		json.Unmarshal([]byte(line), &user)

		if user.Email == email {
			fmt.Println("User found:", user)

			ch <- models.LogInfo{
				Create_on: time.Now(),
				Action:    "GET",
			}

			return nil
		}

	}

	fmt.Println("No user found with email:", email)
	return nil

}

type FileWriter struct {
	FileName string
}

func New(filename string) *FileWriter {
	return &FileWriter{FileName: filename}
}

func (fw *FileWriter) Write(p []byte) (n int, err error) {
	f, err := os.OpenFile(fw.FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}
	// n, err = f.Write(p)
	// return n, err
	defer f.Close()
	return f.Write(p)
}
