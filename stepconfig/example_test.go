package stepconfig_test

import (
	"fmt"
	"log"
	"os"

	"github.com/bitrise-tools/go-steputils/stepconfig"
)

type Configuration struct {

	// Env vars specified in the struct tags are converted to the respective basic data types.
	Name        string `env:"name"`
	BuildNumber int    `env:"build_number"`
	IsUpdate    bool   `env:"is_update"`

	// List items have to be separated by pipe '|', like: "item1|item2"
	Items []string `env:"items"`

	// Secrets are not shown in the output.
	Password stepconfig.Secret `env:"password"`

	// If the env var is not set, the field will be set to the type's default value.
	Empty string `env:"empty"`

	// Env vars marked as 'required' have to be set.
	Mandatory string `env:"mandatory,required"`

	// File validation is checks if the file is exist in the specified path.
	TempFile string `env:"tmpfile,file"`

	// Dir checks if the file exist and it is a directory.
	TempDir string `env:"tmpdir,dir"`

	// Value options can be listed using the notation "opt[opt1,opt2,opt3]"
	// The value of the env var should be one of the options.
	ExportMethod string `env:"export_method,opt[dev,qa,prod]"`
}

var envs = map[string]string{
	"name":          "Example",
	"build_number":  "11",
	"is_update":     "yes",
	"items":         "item1|item2|item3",
	"password":      "pass1234",
	"empty":         "",
	"mandatory":     "present",
	"tmpfile":       "/etc/hosts",
	"tmpdir":        "/tmp",
	"export_method": "dev",
}

func Example() {
	var c Configuration
	os.Clearenv()

	// Set env vars for the example.
	for env, value := range envs {
		err := os.Setenv(env, value)
		if err != nil {
			log.Fatalf("Couldn't set env vars: %v\n", err)
		}
	}

	if err := stepconfig.Parse(&c); err != nil {
		log.Fatalf("Couldn't create config: %v\n", err)
	}

	fmt.Println(c)
	// Output: {Example 11 true [item1 item2 item3] *****  present /etc/hosts /tmp dev}
}