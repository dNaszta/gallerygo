package config

import "testing"

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
