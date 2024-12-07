package click_handler_service_test

import (
	"context"
	"testing"

	"github.com/raphoester/clickplanet.lol-backend/internal/clicks/adapters/secondary/in_memory_country_checker"
	"github.com/raphoester/clickplanet.lol-backend/internal/clicks/adapters/secondary/in_memory_tile_checker"
	"github.com/raphoester/clickplanet.lol-backend/internal/clicks/adapters/secondary/in_memory_tile_storage"
	"github.com/raphoester/clickplanet.lol-backend/internal/clicks/domain/click_handler_service"
	"github.com/stretchr/testify/suite"
)

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(testSuite))
}

type testSuite struct {
	suite.Suite

	storage *in_memory_tile_storage.Storage
	service *click_handler_service.Service
}

func (s *testSuite) SetupSuite() {
	s.storage = in_memory_tile_storage.New()
	tileChecker := in_memory_tile_checker.New(250_000)
	countryChecker := in_memory_country_checker.New()
	s.service = click_handler_service.New(tileChecker, s.storage, countryChecker)
}

func (s *testSuite) TestNominalCase() {
	err := s.service.HandleClick(context.Background(), 1, "fr")
	s.Assert().NoError(err)
}

func (s *testSuite) TestTileOnZeroIndex() {
	err := s.service.HandleClick(context.Background(), 0, "fr")
	s.Assert().Error(err)
}

func (s *testSuite) TestInvalidCountry() {
	err := s.service.HandleClick(context.Background(), 10, "invalid")
	s.Assert().Error(err)
}

func (s *testSuite) TestInvalidTile() {
	err := s.service.HandleClick(context.Background(), 250_001, "fr")
	s.Assert().Error(err)
}

func (s *testSuite) TestTileOnLimit() { // tiles are 1 indexed from the frontend perspective
	err := s.service.HandleClick(context.Background(), 250_000, "fr")
	s.Assert().NoError(err)
}
