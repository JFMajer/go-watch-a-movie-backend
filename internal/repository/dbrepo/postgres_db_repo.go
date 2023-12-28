package dbrepo

import (
	"backend/internal/models"
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type PostgresDBRepo struct {
	DB     *pgxpool.Pool
	Logger zerolog.Logger
}

func (pg *PostgresDBRepo) AllMovies() ([]*models.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, title, release_date, runtime, mpaa_rating, description, coalesce(image, ''), created_at, updated_at from movies order by title`
	rows, err := pg.DB.Query(ctx, query)
	if err != nil {
		pg.Logger.Error().Err(err).Str("query", query).Msg("Error executing query in AllMovies")
		return nil, err
	}
	defer rows.Close()

	var movies []*models.Movie
	for rows.Next() {
		var movie models.Movie
		err := rows.Scan(&movie.ID, &movie.Title, &movie.Year, &movie.Runtime, &movie.MPAARating, &movie.Description, &movie.Image, &movie.CreatedAt, &movie.UpdatedAt)
		if err != nil {
			pg.Logger.Error().Err(err).Msg("Error scanning row into Movie struct")
			continue
		}
		movies = append(movies, &movie)
	}

	if err := rows.Err(); err != nil {
		pg.Logger.Error().Err(err).Msg("Error iterating over rows in AllMovies")
		return nil, err
	}

	return movies, nil
}
