package setting

type DatabaseSettings struct {
	Dbtype   string
	Host     string
	Password string
	Username string
}

func (s *Setting) ReadSection(key string, v interface{}) error {
	err := s.vp.UnmarshalKey(key, &v)
	if err != nil {
		return err
	}
	return nil
}
