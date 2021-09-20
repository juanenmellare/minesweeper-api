package services

import (
	"github.com/google/uuid"
	"math/rand"
	"minesweeper-api/errors"
	"minesweeper-api/models"
	"minesweeper-api/repositories"
	"time"
)

type GamesService interface {
	Create(settings *models.Settings) (*models.Game, *errors.ApiError)
	FindById(uuid *uuid.UUID, hasToPreload bool) (*models.Game, *errors.ApiError)
	ExecuteFieldAction(gameUuid *uuid.UUID, fieldUuid *uuid.UUID, fieldStatus models.FieldStatus) *errors.ApiError
}

type gamesServiceImpl struct {
	gamesRepository    repositories.GamesRepository
	fieldsRepository   repositories.FieldsRepository
}

func NewGamesService(gamesRepository repositories.GamesRepository,
	fieldsRepository repositories.FieldsRepository) GamesService {
	return &gamesServiceImpl{
		gamesRepository:    gamesRepository,
		fieldsRepository:   fieldsRepository,
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
			field.SetInitialValue()
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
		StartedAt: time.Now(), Settings: *settings, Minefield: minefield, Status: models.GameStatusInProgress,
	}

	if err := g.gamesRepository.Create(game); err != nil {
		return nil, err
	}

	return game, nil
}

func (g gamesServiceImpl) FindById(uuid *uuid.UUID, hasToPreload bool) (*models.Game, *errors.ApiError) {
	game, err := g.gamesRepository.FindById(uuid, hasToPreload)
	if err != nil {
		return nil, err
	}

	if game.EndedAt == nil {
		game.Duration = int(time.Now().Sub(game.StartedAt).Seconds())
	}

	return game, nil
}

func hasLost(field *models.Field, _ *models.Game, _ repositories.FieldsRepository) (
	*models.GameStatus, *errors.ApiError) {
	if field.IsMine() {
		gameStatus := models.GameStatusLost
		return &gameStatus, nil
	}

	return nil, nil
}

func hasWon(field *models.Field, game *models.Game,
	fieldsRepository repositories.FieldsRepository) (*models.GameStatus, *errors.ApiError) {
	flaggedMines, err := fieldsRepository.FindMineFieldsFlaggedByGame(&field.GameId)
	if err != nil {
		return nil, err
	}

	if game.Settings.MinesQuantity == len(*flaggedMines) {
		gameStatus := models.GameStatusWon
		return &gameStatus, nil
	}

	return nil, nil
}

var isGameFinishedStrategyMap = map[models.FieldStatus]func(
	field *models.Field, game *models.Game,
	fieldsRepository repositories.FieldsRepository) (*models.GameStatus, *errors.ApiError){
	models.FieldStatusShown:   hasLost,
	models.FieldStatusFlagged: hasWon,
}

func (g gamesServiceImpl) validateIfIsFinished(fieldStatus models.FieldStatus, field *models.Field, game *models.Game) *errors.ApiError {
	if hasGameFinished, ok := isGameFinishedStrategyMap[fieldStatus]; ok {
		gameStatus, err := hasGameFinished(field, game, g.fieldsRepository)
		if err != nil {
			return err
		}
		if gameStatus != nil {
			game.Status = *gameStatus
			now := time.Now()
			game.EndedAt = &now
			game.Duration = int(game.EndedAt.Sub(game.StartedAt).Seconds())
			if err != g.gamesRepository.Update(game) {
				return err
			}
		}
	}

	return nil
}

func (g gamesServiceImpl) ExecuteFieldAction(gameUuid *uuid.UUID, fieldUuid *uuid.UUID,
	fieldStatus models.FieldStatus) *errors.ApiError {
	field, err := g.fieldsRepository.FindByIdAndGameId(fieldUuid, gameUuid)
	if err != nil {
		return err
	}

	game, err := g.gamesRepository.FindById(&field.GameId, false)
	if err != nil {
		return err
	}

	if game.Status != models.GameStatusInProgress {
		return errors.NewBadRequestApiError(errors.NewError("game " + game.ID.String() + " is finished"))
	}

	if err := field.SetStatus(fieldStatus); err != nil {
		return errors.NewBadRequestApiError(err)
	}

	if err = g.fieldsRepository.Update(field); err != nil {
		return err
	}

	if err = g.validateIfIsFinished(fieldStatus, field, game); err != nil {
		return err
	}

	return nil
}
