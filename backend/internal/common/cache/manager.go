package cache

var Manager *_Manager

type _Manager struct {
	Controllers map[string]Controller
}

func init() {
	m := &_Manager{
		Controllers: make(map[string]Controller),
	}

	m.RegisterController(NewInReviewCache(), NewLocalCache())
	Manager = m
}

func (m *_Manager) GetController(name string) Controller {
	if c, ok := m.Controllers[name]; ok {
		return c
	}

	return nil
}

func (m *_Manager) RegisterController(controllers ...Controller) error {
	for _, c := range controllers {
		err := c.Preload(struct{}{})
		if err != nil {
			return err
		}
		m.Controllers[c.ID()] = c
	}
	return nil
}
