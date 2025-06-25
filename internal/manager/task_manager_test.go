package manager

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/xjncx/taskmanager/internal/manager/mocks"
	"github.com/xjncx/taskmanager/internal/repository"
	"go.uber.org/mock/gomock"
)

func Test_Create(t *testing.T) {

	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockRepository(ctrl)
	tm := NewTaskManager(mockRepo)

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()

		mockRepo.EXPECT().Add(gomock.Any()).Return(nil)

		id, err := tm.Create(ctx)

		assert.NoError(t, err)
		assert.NotEqual(t, uuid.Nil, id)
	})

	t.Run("repo add fails", func(t *testing.T) {
		ctx := context.Background()

		mockRepo.EXPECT().Add(gomock.Any()).Return(repository.ErrTaskExists)

		id, err := tm.Create(ctx)

		assert.Error(t, err)
		assert.Equal(t, uuid.Nil, id)
	})
}
