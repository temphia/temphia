package pagedash

func (pd *PageDash) actionLoad(req LoadRequest) (*LoadResponse, error) {

	return &LoadResponse{
		Name:     pd.model.Name,
		Data:     pd.model.StaticData,
		Sections: pd.model.Sections,
	}, nil
}

func (pd *PageDash) actionBuild(req BuildRequest) (*BuildRespone, error) {

	return nil, nil
}
