package loaders

// import vikstrous/dataloadgen with your other imports
import (
	"time"

	"github.com/bisal-dd/go-backend-boilerplate/graph/loaders/userLoader"
	"github.com/bisal-dd/go-backend-boilerplate/graph/model"
	"github.com/vikstrous/dataloadgen"
	"gorm.io/gorm"
)


type Loaders struct {
	UserLoader *dataloadgen.Loader[string, *model.User]
}

func NewLoaders(conn *gorm.DB) *Loaders {
	return &Loaders{
		UserLoader: dataloadgen.NewLoader(userLoader.NewUserReader(conn).GetUsers, dataloadgen.WithWait(time.Millisecond)),
	}
}

