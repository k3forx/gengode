package service

// go:generate mockgen -source=./user.go -destination=./mock/user.go -package=mock -mock_names=User=User
type User interface{}
