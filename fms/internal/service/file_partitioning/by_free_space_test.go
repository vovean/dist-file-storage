package file_partitioning

import (
	"fms/internal/domain"
	"reflect"
	"testing"
)

func TestByFreeSpaceStrategy_Partition(t *testing.T) {
	type args struct {
		file     domain.FullFileInfo
		parts    uint
		storages []domain.Storage
	}
	tests := []struct {
		name    string
		args    args
		want    []domain.FilePart
		wantErr bool
	}{
		{
			name: "basic",
			args: args{
				file: domain.FullFileInfo{
					Size: 10_000,
				},
				parts: 3,
				storages: []domain.Storage{
					{
						Id:                  1,
						SpaceAvailableBytes: 6000,
					},
					{
						Id:                  2,
						SpaceAvailableBytes: 3000,
					},
					{
						Id:                  3,
						SpaceAvailableBytes: 1000,
					},
				},
			},
			want: []domain.FilePart{
				{
					PartId: 1,
					Storage: domain.Storage{
						Id:                  1,
						SpaceAvailableBytes: 6000,
					},
					Size: 6000,
				},
				{
					PartId: 2,
					Storage: domain.Storage{
						Id:                  2,
						SpaceAvailableBytes: 3000,
					},
					Size: 3000,
				},
				{
					PartId: 3,
					Storage: domain.Storage{
						Id:                  3,
						SpaceAvailableBytes: 1000,
					},
					Size: 1000,
				},
			},
			wantErr: false,
		},
		{
			name: "double storage free space",
			args: args{
				file: domain.FullFileInfo{
					Size: 10_000,
				},
				parts: 3,
				storages: []domain.Storage{
					{
						Id:                  1,
						SpaceAvailableBytes: 12000,
					},
					{
						Id:                  2,
						SpaceAvailableBytes: 6000,
					},
					{
						Id:                  3,
						SpaceAvailableBytes: 2000,
					},
				},
			},
			want: []domain.FilePart{
				{
					PartId: 1,
					Storage: domain.Storage{
						Id:                  1,
						SpaceAvailableBytes: 12000,
					},
					Size: 6000,
				},
				{
					PartId: 2,
					Storage: domain.Storage{
						Id:                  2,
						SpaceAvailableBytes: 6000,
					},
					Size: 3000,
				},
				{
					PartId: 3,
					Storage: domain.Storage{
						Id:                  3,
						SpaceAvailableBytes: 2000,
					},
					Size: 1000,
				},
			},
			wantErr: false,
		},
		{
			name: "rounding extra bytes",
			args: args{
				file: domain.FullFileInfo{
					Size: 10_001,
				},
				parts: 3,
				storages: []domain.Storage{
					{
						Id:                  1,
						SpaceAvailableBytes: 12000,
					},
					{
						Id:                  2,
						SpaceAvailableBytes: 6000,
					},
					{
						Id:                  3,
						SpaceAvailableBytes: 2000,
					},
				},
			},
			want: []domain.FilePart{
				{
					PartId: 1,
					Storage: domain.Storage{
						Id:                  1,
						SpaceAvailableBytes: 12000,
					},
					Size: 6001,
				},
				{
					PartId: 2,
					Storage: domain.Storage{
						Id:                  2,
						SpaceAvailableBytes: 6000,
					},
					Size: 3000,
				},
				{
					PartId: 3,
					Storage: domain.Storage{
						Id:                  3,
						SpaceAvailableBytes: 2000,
					},
					Size: 1000,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := ByFreeSpaceStrategy{}
			got, err := b.Partition(tt.args.file, tt.args.parts, tt.args.storages)
			if (err != nil) != tt.wantErr {
				t.Errorf("Partition() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Partition() got = %v, want %v", got, tt.want)
			}
		})
	}
}
