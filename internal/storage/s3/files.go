package s3

import (
	"booky-back/internal/config"
	"booky-back/internal/models"
	"bytes"
	"context"
	"fmt"
	"io"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	metaCourseID    = "course_id"
	metaFilename    = "filename"
	metaPublisherID = "publisher_id"
	metaCreatedAt   = "created_at"
)

type FileStorage struct {
	client *s3.Client
	bucket string
}

func NewFileStorage(c *config.StorageConfig) *FileStorage {
	return &FileStorage{
		client: getClient(c),
		bucket: c.S3.Bucket,
	}
}

func (s *FileStorage) CreateFile(file *models.File) (*models.File, error) {
	file.ID = uuid.New().String()
	file.CreatedAt = timestamppb.Now()

	metadata := map[string]string{
		metaCourseID:    file.CourseID,
		metaFilename:    file.Filename,
		metaPublisherID: file.Publisher.ID,
		metaCreatedAt:   file.CreatedAt.AsTime().Format("2006-01-02T15:04:05Z"),
	}

	input := &s3.PutObjectInput{
		Bucket:        aws.String(s.bucket),
		Key:           aws.String(file.ID),
		Body:          bytes.NewReader(file.Content),
		ContentLength: lo.ToPtr(int64(len(file.Content))),
		ContentType:   aws.String("application/octet-stream"),
		Metadata:      metadata,
	}

	_, err := s.client.PutObject(context.Background(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to upload file to S3: %w", err)
	}

	return file, nil
}

func (s *FileStorage) GetFile(id string) (*models.File, error) {
	input := &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(id),
	}

	resp, err := s.client.GetObject(context.Background(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to download file from S3: %w", err)
	}
	defer resp.Body.Close()

	fileData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read file data: %w", err)
	}

	createdAt, err := time.Parse("2006-01-02T15:04:05Z", resp.Metadata[metaCreatedAt])
	if err != nil {
		return nil, fmt.Errorf("failed to parse CreatedAt metadata: %w", err)
	}

	file := &models.File{
		ID:        id,
		CourseID:  resp.Metadata[metaCourseID],
		Content:   fileData,
		Filename:  resp.Metadata[metaFilename],
		Publisher: models.User{ID: resp.Metadata[metaPublisherID]},
		CreatedAt: timestamppb.New(createdAt),
	}

	return file, nil
}

func (s *FileStorage) DeleteFile(id string) error {
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(id),
	}

	_, err := s.client.DeleteObject(context.Background(), input)
	if err != nil {
		return fmt.Errorf("failed to delete file from S3: %w", err)
	}

	return nil
}

func (s *FileStorage) ListFiles(courseID string) ([]*models.File, error) {
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(s.bucket),
	}

	resp, err := s.client.ListObjectsV2(context.Background(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to list files in S3: %w", err)
	}

	var files []*models.File
	for _, item := range resp.Contents {
		headInput := &s3.HeadObjectInput{
			Bucket: aws.String(s.bucket),
			Key:    item.Key,
		}

		headResp, err := s.client.HeadObject(context.Background(), headInput)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve file metadata: %w", err)
		}

		if headResp.Metadata[metaCourseID] == courseID {
			createdAt, err := time.Parse("2006-01-02T15:04:05Z", headResp.Metadata[metaCreatedAt])
			if err != nil {
				return nil, fmt.Errorf("failed to parse CreatedAt metadata: %w", err)
			}

			file := &models.File{
				ID:        *item.Key,
				CourseID:  headResp.Metadata[metaCourseID],
				Filename:  headResp.Metadata[metaFilename],
				Publisher: models.User{ID: headResp.Metadata[metaPublisherID]},
				CreatedAt: timestamppb.New(createdAt),
			}
			files = append(files, file)
		}
	}

	return files, nil
}
