package types

import (
	"time"
)

type User struct {
	ID        int       
	FirstName string    
	LastName  string  
	Email     string   
	Password  string  
	CreatedAt time.Time
}

type File struct {
	ID        int
	UserId	  int
	FolderId  int
	FileName  string
	FileSize  int
	FileType  string
	FilePath  string
	Description string
	MimeType  string
	CreatedAt time.Time
}

type Folder struct {
	ID        	int
	UserId	  	int
	FolderName  string
	Description string
	CreatedAt 	time.Time
	UpdateAt 	time.Time
}

type CreateFolderPayload struct {
	FolderName  string `json:"folderName" validate:"required"`
	Description string `json:"description"`
}

type UpdateFolderPayload struct {
	FolderName  string `json:"folderName" validate:"required"`
	Description string `json:"description"`
}

type UploadFilePayload struct {
	FileName  string `json:"fileName" validate:"required"`
	FileSize  int    `json:"fileSize" validate:"required"`
	FileType  string `json:"fileType" validate:"required"`
	FilePath  string `json:"filePath" validate:"required"`
	FolderId  int    `json:"folderId" validate:"required"`
	Description string `json:"description"`
	MimeType  string `json:"mimeType" validate:"required"`
}

type UpdateFilePayload struct {
	FileName  string `json:"fileName" validate:"required"`
	FileSize  int    `json:"fileSize" validate:"required"`
	FileType  string `json:"fileType" validate:"required"`
	FilePath  string `json:"filePath" validate:"required"`
	FolderId  int    `json:"folderId" validate:"required"`
	Description string `json:"description"`
	MimeType  string `json:"mimeType" validate:"required"`
}

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=130"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}

type FileStore interface {
	GetFileByID(id int) (*File, error)
	GetFilesByFolderId(folderId int) ([]File, error)
	CreateFile(File) error
	UpdateFile(File) error
	DeleteFile(id int) error
}

type FolderStore interface {
	GetFolderByID(id int) (*Folder, error)
	GetFoldersByUserId(userId int) ([]Folder, error)
	CreateFolder(Folder) error
	UpdateFolder(Folder) error
	DeleteFolder(id int) error
}
