package services

import (
	"math/rand"
	"minesweeper-api/errors"
	"minesweeper-api/models"
	"minesweeper-api/repositories"
	"time"
)

type GamesService interface {
	Create(settings *models.Settings) (*models.Game, *errors.ApiError)
}

type gamesServiceImpl struct {
	gamesRepository repositories.GamesRepository
}

func NewGamesService(gamesRepository repositories.GamesRepository) GamesService {
	return &gamesServiceImpl{
		gamesRepository: gamesRepository,
	}
}

func fillMinefieldWithMines(minefield *[][]models.Field, settings *models.Settings) []models.Position {
	bombsCounter := settings.MinesQuantity

	minesPositions := make([]models.Position, bombsCounter)

	for ok := true; ok; ok = bombsCounter != 0 {
		yPosition := rand.Intn(settings.Height)
		xPosition := rand.Intn(settings.Width)

		if field := &(*minefield)[yPosition][xPosition]; field.IsNil() {
			field.SetMine()
			minesPositions[bombsCounter-1] = models.Position{Y: yPosition, X: xPosition}
			bombsCounter--
		}
	}

	return minesPositions
}

func fillMinefieldWithHints(minefield *[][]models.Field, settings *models.Settings, minesPositions []models.Position) {
	var borderingPositions = []models.Position{
		{X: 0, Y: 1}, {X: 1, Y: 1}, {X: 1, Y: 0}, {X: 1, Y: -1},
		{X: 0, Y: -1}, {X: -1, Y: -1}, {X: -1, Y: 0}, {X: -1, Y: 1},
	}

	for _, position := range minesPositions {
		for _, borderingPosition := range borderingPositions {
			candidatePositionY := position.Y + borderingPosition.Y
			candidatePositionX := position.X + borderingPosition.X

			isCandidateInsideMinefield :=
				candidatePositionY < settings.Height && candidatePositionY != -1 &&
					candidatePositionX < settings.Width && candidatePositionX != -1

			if isCandidateInsideMinefield {
				field := &(*minefield)[candidatePositionY][candidatePositionX]
				if field.IsNil() {
					field.SetInitialHintValue()
				} else if !field.IsMine() {
					field.IncrementHintValue()
				}
			}
		}
	}
}

func fillFieldsPositionsAndFlat(settings *models.Settings, minefield *[][]models.Field) *[]models.Field {
	minefieldSlice := make([]models.Field, settings.Height*settings.Width)

	index := 0
	for y := 0; y < settings.Height; y++ {
		for x := 0; x < settings.Width; x++ {
			field := &(*minefield)[y][x]
			field.SetPosition(y, x)
			minefieldSlice[index] = *field
			index++
		}
	}

	return &minefieldSlice
}

func createMinefield(settings *models.Settings) *[]models.Field {
	minefield := make([][]models.Field, settings.Height)
	for index := range minefield {
		minefield[index] = make([]models.Field, settings.Width)
	}

	minesPositions := fillMinefieldWithMines(&minefield, settings)
	fillMinefieldWithHints(&minefield, settings, minesPositions)
	minefieldFlatted := fillFieldsPositionsAndFlat(settings, &minefield)

	return minefieldFlatted
}

func (g gamesServiceImpl) Create(settings *models.Settings) (*models.Game, *errors.ApiError) {
	if err := settings.Validate(); err != nil {
		return nil, err
	}

	minefield := createMinefield(settings)
	game := &models.Game{
		StartedAt: time.Now(), Settings: *settings, Minefield: minefield, Status: models.StatusInProgress,
	}

	if err := g.gamesRepository.Create(game); err != nil {
		return nil, err
	}

	return game, nil
}
