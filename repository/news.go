package repository

import (
	"kango-api/database"
	"kango-api/model"
)

func GetNews() (*model.NewsResponse, error) {
	db, err := database.Connect()
	if err != nil {
		//fmt.Println(err.Error())
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("select id, title, summary, photo from cj_berita")
	if err != nil {
		//fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close()
	var result []model.News
	for rows.Next() {
		var each = model.News{}
		var err = rows.Scan(&each.ID, &each.Title, &each.Summary, &each.Photo)
		if err != nil {
			//fmt.Println(err.Error())
			return nil, err
		}
		result = append(result, each)
	}
	if err = rows.Err(); err != nil {
		//fmt.Println(err.Error())
		return nil, err
	}

	return &model.NewsResponse{
		Status:  true,
		Message: "Sukses",
		Data:    result,
	}, nil
}

func GetNewsDetail(id int) (*model.NewsDetail, error) {
	db, err := database.Connect()
	if err != nil {
		//fmt.Println(err.Error())
		return nil, err
	}
	defer db.Close()

	var result = model.NewsDetail{}

	row := db.QueryRow("select id, title, summary, photo, body from cj_berita where id=?", id)
	err = row.Scan(&result.ID, &result.Title, &result.Summary, &result.Photo, &result.Body)
	if err != nil {
		//fmt.Println(err.Error())
		return nil, err
	}

	return &result, nil
}
