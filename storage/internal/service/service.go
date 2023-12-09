package service

import (
	"context"
	goerrors "errors"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"storage/internal/domain"

	"github.com/pkg/errors"
)

// В данном случае бизнес логика сервиса - это хранение файлов, так что не будет выделять еще и слой репозитория,
// объединим его с сервисом. В рамках тестового, думаю, точно будет достаточно
type Service struct {
	root      string
	sizeBytes uint64
}

// FolderExists проверяет существование папки
func FolderExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// CanWriteToFolder проверяет доступ на запись в папку
func CanWriteToFolder(path string) bool {
	tmpfile, err := os.CreateTemp(path, "test")
	if err != nil {
		return false
	}

	// Удаляем временный файл после создания
	if err := tmpfile.Close(); err != nil {
		log.Printf("cannot close tmp file: %v\n", err)
	}
	if err := os.Remove(tmpfile.Name()); err != nil {
		log.Printf("cannot remove temp file: %v\n", err)
	}

	return true
}

func New(root string, sizeBytes uint64) (*Service, error) {
	exists, err := FolderExists(root)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot check folder %s existance", root)
	}
	if !exists {
		return nil, fmt.Errorf("provided root folder %s doesn't exist", root)
	}

	canWrite := CanWriteToFolder(root)
	if !canWrite {
		return nil, errors.New("provided storage root is not writeable")
	}

	return &Service{root: root, sizeBytes: sizeBytes}, nil
}

// Используею named return ради того, чтобы иметь возможность добавить сообщение к ошибке в defer.
// Naked return все еще строгое табу
func WriteFileFromReader(filePath string, r io.Reader) (err error) {
	// todo в продакшене был бы другой уровень доступа
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return errors.Wrap(err, "create folders")
	}

	file, err := os.Create(filePath)
	if err != nil {
		return errors.Wrapf(err, "create file %s", filePath)
	}
	defer func() {
		if errC := file.Close(); errC != nil {
			err = goerrors.Join(err, errors.Wrap(err, "close file"))
		}
	}()

	_, err = io.Copy(file, r)

	return errors.Wrap(err, "write file from reader")
}

func (s *Service) Store(ctx context.Context, req domain.StoreFileRequest) error {
	r := &contextAwareReader{
		ctx:    ctx,
		reader: req.Content,
	}

	return WriteFileFromReader(path.Join(s.root, req.Path), r)
}

func (s *Service) Serve(ctx context.Context, req domain.ServeFileRequest) (io.ReadCloser, error) {
	file, err := os.Open(path.Join(s.root, req.Path))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, domain.ErrFileNotFound
		}
		return nil, errors.Wrapf(err, "open %s", req.Path)
	}
	return file, err
}

func (s *Service) Info(ctx context.Context) (domain.StorageInfo, error) {
	return domain.StorageInfo{Size: s.sizeBytes}, nil
}

func (s *Service) Delete(filepath string) error {
	err := os.Remove(path.Join(s.root, filepath))
	if err != nil && !os.IsNotExist(err) {
		// Возвращаем ошибку, только если это не ошибка "файл не найден"
		return err
	}
	return nil
}
