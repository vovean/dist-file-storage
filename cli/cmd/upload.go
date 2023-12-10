package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

type uploadConfig struct {
	gatewayAddr string
	filePath    string
}

var uploadCfg uploadConfig

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "uploads file to storage",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Открываем файл
		file, err := os.Open(uploadCfg.filePath)
		if err != nil {
			return fmt.Errorf("cannot open file: %v", err)
		}
		defer file.Close()

		hashReader := newHashingReader(file)

		// Получаем информацию о размере файла
		fileInfo, err := file.Stat()
		if err != nil {
			log.Fatalf("cannot get file stats: %v", err)
		}

		// Создаем запрос
		client := &http.Client{}
		req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/file", uploadCfg.gatewayAddr), hashReader)
		if err != nil {
			return fmt.Errorf("cannot make request object: %v", err)
		}

		q := req.URL.Query()
		q.Add("filename", fileInfo.Name())
		req.URL.RawQuery = q.Encode()

		// Устанавливаем необходимые заголовки
		req.Header.Set("Content-Type", "application/octet-stream")
		req.ContentLength = fileInfo.Size()

		// Выполнение запроса
		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("cannot send request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated {
			body := "<failed to read response body or it is empty>"
			if bytes, err := io.ReadAll(resp.Body); err == nil && len(bytes) > 0 {
				body = string(bytes)
			}
			return fmt.Errorf("unexpected response status code %d, data: %s", resp.StatusCode, body)
		}

		cmd.Printf("File has been successfully uploaded, md5-hash: %s\n", hashReader.Hash())

		return nil
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)

	uploadCmd.Flags().StringVarP(&uploadCfg.gatewayAddr,
		"gateway", "g", "http://localhost:8081",
		"Gateway address in the form of <schema>://<host>:<port>",
	)
	uploadCmd.Flags().StringVarP(&uploadCfg.filePath,
		"file", "f", "",
		"Path to the file to be uploaded",
	)
}
