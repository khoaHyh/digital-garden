package projects

type Project struct {
	Title       string
	Description string
	Link        string
}

var projects = []Project{
	{
		Title:       "SonicSight",
		Description: "A web application that transforms audio files into spectrograms and uses deep learning to classify sounds. Currently specialized in distinguishing between cat and dog sounds, it leverages a model I trained on Kaggle, deployed on Hugging Face with a FastHTML frontend.",
		Link:        "https://sonicsight.dev/",
	},
}
