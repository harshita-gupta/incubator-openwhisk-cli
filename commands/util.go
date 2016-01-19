package commands

import (
	"errors"

	"github.ibm.com/Bluemix/whisk-cli/client"
)

func parseKeyValueArray(args []string) ([]client.KeyValue, error) {
	parsed := []client.KeyValue{}
	if len(args)%2 != 0 {
		err := errors.New("key|value arguments must be submitted in pairs")
		return parsed, err
	}

	for i := 0; i < len(args); i += 2 {
		keyValue := client.KeyValue{
			Key:   args[i],
			Value: args[i+1],
		}
		parsed = append(parsed, keyValue)

	}
	return parsed, nil
}

func parseParameters(args []string) (client.Parameters, error) {
	parameters := client.Parameters{}
	parsedArgs, err := parseKeyValueArray(args)
	if err != nil {
		return parameters, err
	}
	parameters = client.Parameters(parsedArgs)
	return parameters, nil
}

func parseAnnotations(args []string) (client.Annotations, error) {
	annotations := client.Annotations{}
	parsedArgs, err := parseKeyValueArray(args)
	if err != nil {
		return annotations, err
	}
	annotations = client.Annotations(parsedArgs)
	return annotations, nil
}
