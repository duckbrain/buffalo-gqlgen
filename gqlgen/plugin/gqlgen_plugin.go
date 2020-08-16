package plugin

// TODO plugin that auto implements resolvers using pop

type Plugin struct {
}

func (Plugin) Name() string {
	return "Buffalo"
}

func New() *Plugin {
	return &Plugin{}
}
