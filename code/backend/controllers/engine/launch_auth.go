package engine

type AuthLaunchData struct{}

func (c *Controller) LaunchAuth(data AuthLaunchData) (string, error) {
	return c.launchAuth(data)
}

func (c *Controller) launchAuth(data AuthLaunchData) (string, error) {
	return "", nil
}
