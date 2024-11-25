package test

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/johnfercher/go-tree/node"
	"github.com/stretchr/testify/assert"

	"github.com/johnfercher/maroto/v2/pkg/core"
)

var (
	marotoFile              = ".maroto.yml"
	goModFile               = "go.mod"
	configSingleton *Config = nil
)

type Node struct {
	Value   interface{}            `json:"value,omitempty"`
	Type    string                 `json:"type"`
	Details map[string]interface{} `json:"details,omitempty"`
	Nodes   []*Node                `json:"nodes,omitempty"`
}

// MarotoTest is the unit test instance.
type MarotoTest struct {
	t    *testing.T
	node *node.Node[core.Structure]
}

// New creates the MarotoTest instance to unit tests.
func New(t *testing.T) *MarotoTest {
	if configSingleton == nil {
		cfg, err := loadMarotoConfig()
		if err != nil {
			assert.Fail(t, fmt.Sprintf("could not load maroto config: %s", err.Error()))
		}

		configSingleton = cfg
	}

	return &MarotoTest{
		t: t,
	}
}

// Assert validates if the structure is the same as defined by Equals method.
func (m *MarotoTest) Assert(structure *node.Node[core.Structure]) *MarotoTest {
	m.node = structure
	return m
}

// Equals defines which file will be loaded to do the comparison.
func (m *MarotoTest) Equals(file string) *MarotoTest {
	actual := m.buildNode(m.node)
	actualBytes, _ := json.Marshal(actual)
	actualString := string(actualBytes)

	indentedExpectBytes, err := os.ReadFile(configSingleton.getAbsoluteFilePath(file))
	if err != nil {
		assert.Fail(m.t, err.Error())
	}

	savedNode := &Node{}
	_ = json.Unmarshal(indentedExpectBytes, savedNode)
	expectedBytes, _ := json.Marshal(savedNode)

	assert.Equal(m.t, string(expectedBytes), actualString, fmt.Sprintf("json: %s", configSingleton.getAbsoluteFilePath(file)))
	return m
}

// Save is an auxiliary method to update the file to be asserted.
func (m *MarotoTest) Save(file string) *MarotoTest {
	actual := m.buildNode(m.node)
	actualBytes, _ := json.MarshalIndent(actual, "", "\t")

	err := os.WriteFile(configSingleton.getAbsoluteFilePath(file), actualBytes, os.ModePerm)
	if err != nil {
		assert.Fail(m.t, err.Error())
	}

	return m
}

func (m *MarotoTest) buildNode(node *node.Node[core.Structure]) *Node {
	data := node.GetData()
	actual := &Node{
		Type:    data.Type,
		Value:   data.Value,
		Details: data.Details,
	}

	nexts := node.GetNexts()
	for _, next := range nexts {
		actual.Nodes = append(actual.Nodes, m.buildNode(next))
	}

	return actual
}

func loadMarotoConfig() (*Config, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return nil, errors.New("unable to get the current filename")
	}

	return &Config{
		AbsolutePath: filepath.Dir(filepath.Dir(filepath.Dir(filename))),
		TestPath:     "test/maroto",
	}, nil
}

func hasFileInPath(file string, path string) (bool, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return false, err
	}

	for _, entry := range entries {
		if entry.Name() == file {
			return true, nil
		}
	}

	return false, nil
}

func getParentDir(path string) string {
	dirs := strings.Split(path, "/")
	dirs = dirs[:len(dirs)-2]

	var newPath string
	for _, dir := range dirs {
		newPath += dir + "/"
	}

	return newPath
}
