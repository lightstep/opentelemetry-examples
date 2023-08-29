package nginx_test

import(
	"fmt"
	"os"
	"testing"

	tc "github.com/testcontainers/testcontainers-go"
)

// specifyEnv
// passThroughEnv

func setup() error {
	accessToken := os.Getenv("LS_ACCESS_TOKEN")
	composeFilePaths := []string {"docker-compose.yml"}
	identifier := strings.ToLower(uuid.New().String())

	compose := tc.NewLocalDockerCompose(composeFilePaths, identifier)
	execError := compose.
	    WithCommand([]string{"up", "-d"}).
	    WithEnv(map[string]string {
		"LS_ACCESS_TOKEN": accessToken,
	    }).
	    Invoke()

	if err := execError.Error; err != nil {
		return fmt.Errorf("Could not run compose file: %v - %v", composeFilePaths, err)
	}
	return nil
}

func TestMain(m *testing.M) {
	log.Println("*** SETUP ***")
	if err := setup(); err != nil {
	    os.Exit(1)
	}

	exitVal := m.Run()

	log.Println("*** TEARDOWN ***")
	if execError := compose.Down(); execError.Error != nil {
	    // return fmt.Errorf("Could not run compose file: %v - %v", composeFilePaths, err)
	    os.Exit(1)
	}
	os.Exit(exitVal)
}

func TestA(t *testing.T) {

}

func TestB(t *testing.T) {

}
