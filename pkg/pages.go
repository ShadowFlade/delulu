package pkg

type IPages struct {
	INDEX     string
	RESULT    string
	ABOUT     string
	RESOURCES string
	FEEDBACK string
}

var Pages = IPages{
	INDEX:     "",
	RESULT:    "result",
	ABOUT:     "about",
	RESOURCES: "resources",
	FEEDBACK: "feedback",
}
