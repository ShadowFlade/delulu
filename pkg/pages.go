package pkg

type IPages struct {
	INDEX     string
	RESULT    string
	ABOUT     string
	RESOURCES string
}

var Pages = IPages{
	INDEX:     "",
	RESULT:    "result",
	ABOUT:     "about",
	RESOURCES: "resources",
}
