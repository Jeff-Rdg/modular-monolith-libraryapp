package usecase

import (
	"context"
	"errors"
	"modular-monolith-libraryApp/modules/author/domain"
	"modular-monolith-libraryApp/modules/author/port"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateAuthor_Success(t *testing.T) {
	// arrange
	repoMock := port.NewMockAuthorRepository(t)
	ctx := context.Background()

	input := CreateAuthorInput{
		Name: "Machado de Assis",
	}

	repoMock.EXPECT().
		Save(ctx, domain.Author{Name: input.Name}).
		Return(nil)

	uc := NewCreateAuthor(repoMock)

	// act
	err := uc.Execute(ctx, input)

	// assert
	assert.NoError(t, err)
}

func TestCreateAuthor_RepositoryError(t *testing.T) {
	// arrange
	repoMock := port.NewMockAuthorRepository(t)
	ctx := context.Background()

	input := CreateAuthorInput{
		Name: "Machado de Assis",
	}

	repoMock.EXPECT().
		Save(ctx, domain.Author{Name: input.Name}).
		Return(errors.New("db error"))

	uc := NewCreateAuthor(repoMock)

	// act
	err := uc.Execute(ctx, input)

	// assert
	assert.Error(t, err)
}

func TestCreateAuthor_TableDriven(t *testing.T) {
	cases := []struct {
		name    string
		input   CreateAuthorInput
		repoErr error
		wantErr bool
	}{
		{
			name:    "sucesso",
			input:   CreateAuthorInput{Name: "Machado de Assis"},
			repoErr: nil,
			wantErr: false,
		},
		{
			name:    "erro no repositorio",
			input:   CreateAuthorInput{Name: "José de Alencar"},
			repoErr: errors.New("db connection failed"),
			wantErr: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// arrange
			repoMock := port.NewMockAuthorRepository(t)
			ctx := context.Background()

			repoMock.EXPECT().
				Save(ctx, domain.Author{Name: tc.input.Name}).
				Return(tc.repoErr)

			uc := NewCreateAuthor(repoMock)

			// act
			err := uc.Execute(ctx, tc.input)

			// assert
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// garante que o use case não chama Save com ID ou CreatedAt preenchidos
// já que o domain.Author é montado sem construtor — só Name é passado
func TestCreateAuthor_SaveReceivesOnlyName(t *testing.T) {
	repoMock := port.NewMockAuthorRepository(t)
	ctx := context.Background()

	repoMock.EXPECT().
		Save(ctx, domain.Author{Name: "Clarice Lispector"}).
		RunAndReturn(func(ctx context.Context, a domain.Author) error {
			assert.Equal(t, "Clarice Lispector", a.Name)
			assert.Equal(t, uuid.Nil, a.ID)      // ID não é gerado no use case
			assert.True(t, a.CreatedAt.IsZero()) // CreatedAt não é gerado no use case
			return nil
		})

	uc := NewCreateAuthor(repoMock)
	err := uc.Execute(ctx, CreateAuthorInput{Name: "Clarice Lispector"})

	assert.NoError(t, err)
}
