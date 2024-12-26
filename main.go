package main

func main() {
	ImageConvert(
		ImageConvertConfig{
			//API_KEY: "api_key",
			Dir: Dir{
				//INPUT_PATH:  "./public/input2",
				//OUTPUT_PATH: "./public/output2",
			},
			AppWrite: AppWrite{
				ENDPOINT:   "https://appwrite.autoges.eu/v1",
				API_SECRET: "standard_fbb9d7e7b4d519e932e4bdada4f6df67d96daecbd9092f20d475cdcca3dab74654a78e68c142817bc24a1aafa712d1fbcfe9a53437ed61090e6c77e6a8ab4daacda2529230b3f0882a84de9a60ddbb39622f01466b685ce59c31b18374417ef3d22ac7d0db703699c7ee69018b33e9643b4afb674329e6cd3d43e41e35fc4288",
			},
			Server: ServerConfig{
				//PORT:            "5858",
				//HOST:            "localhost",
				//ENDPOINT_EXPONE: "/api/v1/image",
			},
		},
	)

}
