package dto

type CreateCourseDto struct {
	Title         string       `json:"title" binding:"required"`
	Description   string       `json:"description" binding:"required"`
	Image         *string      `json:"image"`
	TelegramGroup *string      `json:"telegram_group"`
	Requirements  string       `json:"requirements" binding:"required"`
	Level         string       `json:"level" binding:"required"`
	Price         int          `json:"price" binding:"required"`
	Author        string       `json:"author" binding:"required"`
	Chapters      []ChapterDto `json:"chapters" binding:"required"`
	CategoryID    uint         `json:"category_id" binding:"required"`
}

type ChapterDto struct {
	Name    string      `json:"name" binding:"required"`
	Modules []ModuleDto `json:"modules" binding:"required"`
}

type ModuleDto struct {
	Title     string `json:"title" binding:"required"`
	Duration  int    `json:"duration" binding:"required"`
	URL       string `json:"url" binding:"required"`
	IsTrailer bool   `json:"is_trailer" binding:"required"`
}
