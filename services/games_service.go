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
	ExecuteFieldAction(gameUuid *uuid.UUID, fieldUuid *uuid.UUID,
		fieldStatus models.FieldStatus) (*models.GameStatus, *errors.ApiError)
}

type gamesServiceImpl struct {
	gamesRepository  repositories.GamesRepository
	fieldsRepository repositories.FieldsRepository
}

func NewGamesService(gamesRepository repositories.GamesRepository,
	fieldsRepository repositories.FieldsRepository) GamesService {
	return &gamesServiceImpl{
		gamesRepository:  gamesRepository,
		fieldsRepository: fieldsRepository,
	}
}

func fillMinefieldWithMines(minefield *[][]models.Field, settings *models.Settings) []models.Position {
	bombsCounter := settings.MinesQuantity
	minesPositions := make([]models.Position, bombsCounter)
	for ok := true; ok; ok = bombsCounter != 0 {
		positionY := rand.Intn(settings.Height)
		positionX := rand.Intn(settings.Width)

		if field := &(*minefield)[positionY][positionX]; field.IsNil() {
			field.SetMine()
			minesPositions[bombsCounter-1] = models.Position{Y: positionY, X: positionX}
			bombsCounter--
		}
	}

	return minesPositions
}

func calculateHintValue(field *models.Field) {
	if field.IsNil() {
		field.SetInitialHintValue()
	} else if !field.IsMine() {
		field.IncrementHintValue()
	}
}

func isPositionInsideMinefield(settings *models.Settings, candidatePositionY, candidatePositionX int) bool {
	return candidatePositionY < settings.Height && candidatePositionY != -1 &&
		candidatePositionX < settings.Width && candidatePositionX != -1
}

var borderingPositions = []models.Position{
	{X: 0, Y: 1}, {X: 1, Y: 1}, {X: 1, Y: 0}, {X: 1, Y: -1},
	{X: 0, Y: -1}, {X: -1, Y: -1}, {X: -1, Y: 0}, {X: -1, Y: 1},
}

func fillMinefieldWithHints(minefield *[][]models.Field, settings *models.Settings, minesPositions []models.Position) {
	for _, position := range minesPositions {
		for _, borderingPosition := range borderingPositions {
			candidatePositionY := position.Y + borderingPosition.Y
			candidatePositionX := position.X + borderingPosition.X

			if isPositionInsideMinefield(settings, candidatePositionY, candidatePositionX) {
				field := &(*minefield)[candidatePositionY][candidatePositionX]
				calculateHintValue(field)
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
			field.SetInitialStatus()
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

	hideMinefieldValues(game)

	return game, nil
}

func hideMinefieldValues(game *models.Game) {
	minefieldLength := len(*game.Minefield)
	hiddenMinefield := make([]models.Field, minefieldLength)

	if game.Status == models.GameStatusInProgress {
		for index := 0; index < minefieldLength; index++ {
			field := (*game.Minefield)[index]
			if field.Status != models.FieldStatusShown {
				field.HideValue()
			}
			hiddenMinefield[index] = field
		}

	}

	game.Minefield = &hiddenMinefield
}

func (g gamesServiceImpl) FindById(uuid *uuid.UUID, hasToPreload bool) (*models.Game, *errors.ApiError) {
	game, err := g.gamesRepository.FindById(uuid, hasToPreload)
	if err != nil {
		return nil, err
	}

	if game.EndedAt == nil {
		game.Duration = int(time.Now().Sub(game.StartedAt).Seconds())
	}

	hideMinefieldValues(game)

	return game, nil
}

func fillMinefield(fields *[]models.Field, settings *models.Settings) *[][]models.Field {
	minefield := make([][]models.Field, settings.Height)
	for index := range minefield {
		minefield[index] = make([]models.Field, settings.Width)
	}

	for _, field := range *fields {
		minefield[field.PositionY][field.PositionX] = field
	}

	return &minefield
}

func showAdjacentFields(g gamesServiceImpl, field *models.Field, minefield *[][]models.Field,
	settings *models.Settings) {
	for _, position := range borderingPositions {
		candidatePositionY := position.Y + field.PositionY
		candidatePositionX := position.X + field.PositionX

		if isPositionInsideMinefield(settings, candidatePositionY, candidatePositionX) {
			candidateField := &(*minefield)[candidatePositionY][candidatePositionX]
			if !candidateField.IsMine() && candidateField.Status == models.FieldStatusHidden {
				candidateField.Show()
				g.fieldsRepository.Update(field)
				showAdjacentFields(g, candidateField, minefield, settings)
			}
		}
	}

}

func executeAfterShow(g gamesServiceImpl, field *models.Field, game *models.Game) (*models.GameStatus, *errors.ApiError) {
	if field.IsMine() {
		gameStatus := models.GameStatusLost
		return &gameStatus, nil
	}

	if field.IsNil() {
		fields, err := g.fieldsRepository.FindByGameId(&game.ID)
		if err != nil {
			return nil, err
		}

		minefield := fillMinefield(fields, &game.Settings)
		showAdjacentFields(g, field, minefield, &game.Settings)
	}

	return nil, nil
}

func executeAfterFlag(g gamesServiceImpl, _ *models.Field, game *models.Game) (*models.GameStatus, *errors.ApiError) {
	flaggedMines, err := g.fieldsRepository.FindMineFieldsFlaggedByGame(&game.ID)
	if err != nil {
		return nil, err
	}

	if game.Settings.MinesQuantity == len(*flaggedMines) {
		gameStatus := models.GameStatusWon
		return &gameStatus, nil
	}

	return nil, nil
}

var candidateFinishActionStrategyMap = map[models.FieldStatus]func(g gamesServiceImpl, field *models.Field,
	game *models.Game) (*models.GameStatus, *errors.ApiError){
	models.FieldStatusShown:   executeAfterShow,
	models.FieldStatusFlagged: executeAfterFlag,
}

func (g gamesServiceImpl) validateIfHasFinished(fieldStatus models.FieldStatus, field *models.Field,
	game *models.Game) *errors.ApiError {
	if afterAction, ok := candidateFinishActionStrategyMap[fieldStatus]; ok {
		gameStatus, err := afterAction(g, field, game)
		if err != nil {
			return err
		}
		if gameStatus != nil {
			game.Status = *gameStatus
			now := time.Now()
			game.EndedAt = &now
			game.Duration = int(game.EndedAt.Sub(game.StartedAt).Seconds())
			if err = g.gamesRepository.Update(game); err != nil {
				return err
			}
		}
	}

	return nil
}

func (g gamesServiceImpl) ExecuteFieldAction(gameUuid *uuid.UUID, fieldUuid *uuid.UUID,
	fieldStatus models.FieldStatus) (*models.GameStatus, *errors.ApiError) {
	game, err := g.gamesRepository.FindById(gameUuid, false)
	if err != nil {
		return nil, err
	}

	if game.Status != models.GameStatusInProgress {
		return nil, errors.NewBadRequestApiError(errors.NewError("game " + game.ID.String() + " is finished"))
	}

	field, err := g.fieldsRepository.FindByIdAndGameId(fieldUuid, gameUuid)
	if err != nil {
		return nil, err
	}

	if err := field.SetStatus(fieldStatus); err != nil {
		return nil, errors.NewBadRequestApiError(err)
	}

	if err = g.fieldsRepository.Update(field); err != nil {
		return nil, err
	}

	if err = g.validateIfHasFinished(fieldStatus, field, game); err != nil {
		return nil, err
	}

	return &game.Status, nil
}
