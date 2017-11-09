package config

import (
	"testing"
)

func TestConfigs_ToString(t *testing.T) {
	testConf := Configs{
		Port: ":8080",
		Log: "test",
		Sizes: []SizeConfig{
			SizeConfig{
				Suffix: "a",
				Width: 10,
				Height: 20,
			},
		},
		MongoDB: mongoSettings{
			ConnectionString: "localhost",
			Database: "database",
			GalleryCollection: "collection",
		},
	}

	expect := "{" +
		"\"Port\":\":8080\"," +
		"\"Log\":\"test\"," +
		"\"Sizes\":[" +
			"{" +
				"\"Suffix\":\"a\"," +
				"\"Width\":10," +
				"\"Height\":20" +
			"}" +
		"]," +
		"\"MongoDB\":{" +
			"\"ConnectionString\":\"localhost\"," +
			"\"Database\":\"database\"," +
			"\"GalleryCollection\":\"collection\"}}"
	json := testConf.ToString()

	if json != expect {
		t.Logf("ToString failed, expecting %s, got %s", expect, json)
		t.Fail()
	}
}


func TestConfigs_ToJSON(t *testing.T) {
	testConf := Configs{
		Port: ":8080",
		Log: "test",
		Sizes: []SizeConfig{
			SizeConfig{
				Suffix: "a",
				Width: 10,
				Height: 20,
			},
		},
		MongoDB: mongoSettings{
			ConnectionString: "localhost",
			Database: "database",
			GalleryCollection: "collection",
		},
	}

	expect := "{" +
		"\"Port\":\":8080\"," +
		"\"Log\":\"test\"," +
		"\"Sizes\":[" +
		"{" +
		"\"Suffix\":\"a\"," +
		"\"Width\":10," +
		"\"Height\":20" +
		"}" +
		"]," +
		"\"MongoDB\":{" +
		"\"ConnectionString\":\"localhost\"," +
		"\"Database\":\"database\"," +
		"\"GalleryCollection\":\"collection\"}}"
	json := string(testConf.ToJSON())

	if json != expect {
		t.Logf("ToJSON failed, expecting %s, got %s", expect, json)
		t.Fail()
	}
}

func TestLoad(t *testing.T) {
	testPath := "../config_test.json"
	expect := Configs{
		Port: ":8080",
		Log: "test",
		Sizes: []SizeConfig{
			SizeConfig{
				Suffix: "a",
				Width: 10,
				Height: 20,
			},
		},
		MongoDB: mongoSettings{
			ConnectionString: "localhost",
			Database: "database",
			GalleryCollection: "collection",
		},
	}

	var testConfig Configs
	Load(testPath, &testConfig)
	if testConfig.ToString() != expect.ToString() {
		t.Logf("Load failed, expecting %v, got %v", expect, testConfig)
		t.Fail()
	}
}