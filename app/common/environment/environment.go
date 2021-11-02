package environment

type Environment string

const (
	Test       = "test"
	Developer  = "developer"
	Sandbox    = "sandbox"
	Production = "production"
)

//IsValid check if environment is valid
func IsValid(env string) bool {
	return env == Test || env == Developer || env == Production || env == Sandbox
}

func (e Environment) String() string {
	return string(e)
}

//FromString return environment from string
func FromString(env string) Environment {
	switch env {
	case "test":
		return Test
	case "developer":
		return Developer
	case "production":
		return Production
	case "sandbox":
		return Sandbox
	default:
		return "invalid"
	}
}

//CheckEnvironment check if JUS_ENVIRONMENT is set
func CheckEnvironment(env string) {
	// Compare and validate environments
	if !IsValid(env) {
		panic("JUS_ENVIRONMENT not setted!\n" +
			"To correct this, run this command in your terminal: export JUS_ENVIRONMENT=developer\n" +
			"Available environments: developer, sandbox, production, test")
	}
}
