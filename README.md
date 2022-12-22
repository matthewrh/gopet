# GoPet

[![Go Report Card](https://goreportcard.com/badge/github.com/matthewrh/gopet)](https://goreportcard.com/report/github.com/matthewrh/gopet)

GoPet is a simple Tamagotchi-style game that you can run from your command line!

## Installation

1. Make sure you have Go version `>=1.19.1` installed. Verify this by running `go version`. Make sure your `$GOROOT` environment variable is set!
2. Run `go install github.com/matthewrh/gopet@latest` to install the package.
3. To start the application, run `go run gopet`.

## Usage Manual

- Upon starting the application for the first time, you will be prompted to name your pet.
- After entering a name, a new window should open that has your pet and various actions you can perform!
- Use the left and right arrow keys to cycle between different actions. To perform an action, hit the enter key.
- You'll notice that some stats for your GoPet will appear in the upper left hand corner. Stats will decay over time as your GoPet gets hungry, sad, or tired. Once different stats cross certain thresholds, some stats will decay at faster rates. Make sure to pay attention to your GoPet to prevent it from losing life points!
- `Feeding` your GoPet will reduce hunger
- `Playing` with your GoPet will improve happiness
- `Washing` your GoPet will prevent it from getting dirty (and potentially impact happiness!)
- Putting your GoPet to `sleep` will reduce fatigue and regenerate life points

## License

GoPet is licensed under the [MIT license](https://github.com/matthewrh/gopet/blob/main/LICENSE).
