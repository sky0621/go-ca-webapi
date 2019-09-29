package outputport

import (
	usecasemodel "go-ca-webapi/02_cleanarchitecture/usecase/model"
)

type ItemOutputPort interface {
	RenderSaveResult(target *usecasemodel.SaveItem) error
	RenderListResult(targets []*usecasemodel.ListItem) error
	RenderFailure(err error) error
}
