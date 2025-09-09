package repository

import "github.com/swethabj/movies-api/config"

func RawMovieJoin() ([]MovieJoinRow, error) {
	var rows []MovieJoinRow
	sql := `
	SELECT m.id as movie_id, m.title, m.release_year, g.name as genre_name, a.name as actor_name
	FROM movies m
	LEFT JOIN genres g ON m.genre_id = g.id
	LEFT JOIN movie_actors ma ON m.id = ma.movie_id
	LEFT JOIN actors a ON ma.actor_id = a.id
	ORDER BY m.id;
	`
	if err := config.DB.Raw(sql).Scan(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}

type GMovieJoinRow struct {
	MovieNames string `json:"movie_name"`
}

func GetMoviesByGenreName(name string) ([]GMovieJoinRow, error) {
	var rows []GMovieJoinRow
	sql := `SELECT m.title as MovieNames
			FROM movies m
			INNER JOIN genres g ON m.genre_id = g.id
			WHERE g.name = ?
			ORDER BY m.title;`

	if err := config.DB.Raw(sql, name).Scan(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}
