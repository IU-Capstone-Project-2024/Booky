package booky

import (
	pb "booky-back/api/booky"
	"booky-back/internal/config"
	"booky-back/internal/gpt"
	"booky-back/internal/storage"
	inmemory "booky-back/internal/storage/in-memory"
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
