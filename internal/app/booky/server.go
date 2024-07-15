package booky

import (
	pb "booky-back/api/booky"
	"booky-back/internal/config"
	"booky-back/internal/gpt"
	yandex "booky-back/internal/gpt/yandex_gpt"
	"booky-back/internal/storage"
	inmemory "booky-back/internal/storage/in-memory"
	"booky-back/internal/storage/s3"
)

type Server struct {
	Config  *config.Config
	Storage *storage.Storage
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

func getStorage(config *config.StorageConfig) *storage.Storage {
	var courseStorage storage.CourseModel
	switch config.CourseStorage {
	default:
		courseStorage = inmemory.NewCourseStorage()
	}

	var noteStorage storage.NoteModel
	switch config.NoteStorage {
	default:
		noteStorage = inmemory.NewNoteStorage()
	}

	var fileStorage storage.FileModel
	switch config.FileStorage {
	case "S3":
		fileStorage = s3.NewFileStorage(config)
	default:
		fileStorage = inmemory.NewFileStorage()
	}

	var userStorage storage.UserModel
	switch config.UserStorage {
	default:
		userStorage = inmemory.NewUserStorage()
	}

	return &storage.Storage{
		CourseStorage: courseStorage,
		NoteStorage:   noteStorage,
		FileStorage:   fileStorage,
		UserStorage:   userStorage,
	}
}
