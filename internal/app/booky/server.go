package booky

import (
	pb "booky-back/api/booky"
	"booky-back/internal/config"
	"booky-back/internal/pkg/gpt"
	"booky-back/internal/pkg/gpt/yandex_gpt"
	storage2 "booky-back/internal/pkg/storage"
	inmemory2 "booky-back/internal/pkg/storage/in-memory"
	"booky-back/internal/pkg/storage/s3"
)

type Server struct {
	Config  *config.Config
	Storage *storage2.Storage
	GPT     *gpt.GPT

	pb.UnimplementedBookyServiceServer
}

func NewServer(config *config.Config) *Server {
	return &Server{
		Config:  config,
		Storage: getStorage(&config.Storage),
		GPT:     &gpt.GPT{AiModel: yandex.NewYandexGPT(&config.Gpt)},
	}
}

func getStorage(config *config.StorageConfig) *storage2.Storage {
	var courseStorage storage2.CourseModel
	switch config.CourseStorage {
	default:
		courseStorage = inmemory2.NewCourseStorage()
	}

	var noteStorage storage2.NoteModel
	switch config.NoteStorage {
	default:
		noteStorage = inmemory2.NewNoteStorage()
	}

	var fileStorage storage2.FileModel
	switch config.FileStorage {
	case "S3":
		fileStorage = s3.NewFileStorage(config)
	default:
		fileStorage = inmemory2.NewFileStorage()
	}

	var userStorage storage2.UserModel
	switch config.UserStorage {
	default:
		userStorage = inmemory2.NewUserStorage()
	}

	return &storage2.Storage{
		CourseStorage: courseStorage,
		NoteStorage:   noteStorage,
		FileStorage:   fileStorage,
		UserStorage:   userStorage,
	}
}
