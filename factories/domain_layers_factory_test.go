package factories

import (
	"github.com/stretchr/testify/assert"
	"minesweeper-api/controllers"
	"testing"
)

func TestNewDomainLayersFactory(t *testing.T) {
	domainLayersFactory := NewDomainLayersFactory(nil)

	assert.Implements(t, (*DomainLayersFactory)(nil), domainLayersFactory)
}

func Test_domainLayersFactoryImpl_GetGamesController(t *testing.T) {
	domainLayersFactory := NewDomainLayersFactory(nil)

	assert.Implements(t, (*controllers.GamesController)(nil), domainLayersFactory.GetGamesController())
}

func Test_domainLayersFactoryImpl_GetHealthChecksController(t *testing.T) {
	domainLayersFactory := NewDomainLayersFactory(nil)

	assert.Implements(t, (*controllers.HealthChecksController)(nil), domainLayersFactory.GetHealthChecksController())
}
