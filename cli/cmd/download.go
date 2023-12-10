/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

type downloadConfig struct {
	gatewayAddr string
	outFilePath string
	filename    string
}

var downloadCfg downloadConfig

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "downloads a file",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Создаем запрос
		client := &http.Client{}
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/file", uploadCfg.gatewayAddr), nil)
		if err != nil {
			return fmt.Errorf("cannot make request object: %v", err)
		}

		q := req.URL.Query()
		q.Add("filename", downloadCfg.filename)
		req.URL.RawQuery = q.Encode()

		// Устанавливаем необходимые заголовки
		req.Header.Set("Content-Type", "application/octet-stream")

		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("cannot send request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body := "<failed to read response body or it is empty>"
			if bytes, err := io.ReadAll(resp.Body); err == nil && len(bytes) > 0 {
				body = string(bytes)
			}
			return fmt.Errorf("unexpected response status code %d, data: %s", resp.StatusCode, body)
		}

		hashReader := newHashingReader(resp.Body)

		outPath := downloadCfg.outFilePath
		if outPath == "" {
			outPath = downloadCfg.filename
		}
		outFile, err := os.Create(downloadCfg.outFilePath)
		if err != nil {
			return fmt.Errorf("cannot open output file: %v", err)
		}

		if _, err := io.Copy(outFile, hashReader); err != nil {
			return fmt.Errorf("failed to save file to disk: %v", err)
		}

		cmd.Printf("File successfully download to %s, md5-hash: %s\n", outPath, hashReader.Hash())

		return nil
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	downloadCmd.Flags().StringVarP(&downloadCfg.gatewayAddr,
		"gateway", "g", "http://localhost:8081",
		"Gateway address in the form of <schema>://<host>:<port>",
	)
	downloadCmd.Flags().StringVarP(&downloadCfg.filename,
		"filename", "f", "",
		"Filename to be downloaded",
	)
	downloadCmd.Flags().StringVarP(&downloadCfg.outFilePath,
		"out", "o", "",
		"Path where the downloaded file will be stored. If empty working directory will be used",
	)
}
