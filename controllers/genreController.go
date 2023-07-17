package controllers

//func GetGenreIdsByName(name string) []models.Genre {
//	namesGenre := strings.Split(name, ",")
//	var genreIds []models.Genre
//	err := initializers.DB.Table("genres").Select("id").Where("genreName IN ?", namesGenre).Find(&genreIds).Error
//	if err != nil {
//		log.Fatal("Ошибка при выполнении запроса:", err)
//	}
//	return genreIds
//}
