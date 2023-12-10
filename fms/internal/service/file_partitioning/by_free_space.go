package file_partitioning

import (
	"fms/internal/domain"
	"log"
	"sort"

	"github.com/pkg/errors"
)

// ByFreeSpaceStrategy распределяет файл по хранилищам по убыванию свободного места пропорционально кол-ву свободного места в них
type ByFreeSpaceStrategy struct{}

func (b ByFreeSpaceStrategy) Partition(file domain.FullFileInfo, parts uint, storages []domain.Storage) ([]domain.FilePart, error) {
	if parts == 0 {
		return nil, errors.New("parts must be >= 0")
	}

	if len(storages) < int(parts) {
		log.Printf("currently there are %d/%d storages", len(storages), parts)
		return nil, domain.ErrNotEnoughStorages
	}

	// Сортировка хранилищ по убыванию свободного места
	sort.Slice(storages, func(i, j int) bool {
		return storages[i].SpaceAvailableBytes > storages[j].SpaceAvailableBytes
	})

	// Берем parts хранилищ с самым большим свободным местом
	totalSpace := uint64(0)
	for i := 0; i < int(parts); i++ {
		totalSpace += storages[i].SpaceAvailableBytes
	}

	if file.Size > totalSpace {
		return nil, domain.ErrNotEnoughSpace
	}

	fileParts := make([]domain.FilePart, 0, parts)
	remainingSize := file.Size
	for i := 0; i < int(parts); i++ {
		// Распределяем размер частей пропорционально свободному месту в хранилищах
		proportion := float64(storages[i].SpaceAvailableBytes) / float64(totalSpace)
		partSize := uint64(float64(file.Size) * proportion)
		if partSize > remainingSize {
			partSize = remainingSize
		}

		fileParts = append(fileParts, domain.FilePart{
			PartId:  i + 1,
			Storage: storages[i],
			Size:    partSize,
		})
		remainingSize -= partSize
	}

	// Из-за округлений могло остаться < N байтов, где N - кол-во частей. Попробуем впихнуть куда сможем, возможно, в несколько хранилищ
	if remainingSize > 0 {
		for i, p := range fileParts {
			availableSpace := p.Storage.SpaceAvailableBytes - p.Size
			if availableSpace > 0 {
				additionalSize := min(availableSpace, remainingSize)
				fileParts[i].Size += additionalSize
				remainingSize -= additionalSize
			}
			if remainingSize == 0 {
				break
			}
		}
	}
	// если никуда не смогли впихнуть, значит не хватило места
	if remainingSize > 0 {
		return nil, errors.New("cannot distribute remaining part")
	}

	return fileParts, nil
}
