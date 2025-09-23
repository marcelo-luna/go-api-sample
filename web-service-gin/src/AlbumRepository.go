package main

import "fmt"

func GetAllAlbuns() []album {

	rows, err := Db.Query("SELECT id, title, artist, price FROM public.albums")
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()

	var albunsResult []album

	for rows.Next() {
		var albumRow album
		err = rows.Scan(&albumRow.ID, &albumRow.Title, &albumRow.Artist, &albumRow.Price)
		if err != nil {
			// handle this error
			panic(err)
		}

		albunsResult = append(albunsResult, albumRow)
		fmt.Println(albumRow)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return albunsResult
}

func InsertAlbum(album album) int {

	var id int
	err := Db.QueryRow(
		"INSERT INTO public.albums (title, artist, price) VALUES ($1, $2, $3) RETURNING id",
		album.Title, album.Artist, album.Price,
	).Scan(&id)
	if err != nil {
		panic(err)
	}
	return id
}
